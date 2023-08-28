---
title: "Google ASL Fingerspell Kaggle Challenge: Pose estimates to text seq2seq model"
date: "2023-08-27"
---

<script>
  import dataset from '$lib/assets/asl-fingerspell/dataset.png';
  import sequence from '$lib/assets/asl-fingerspell/example_feature_sequence.gif';
</script>

This was my submission to Google's ASL Fingerspell Challenge. It's not much and I still clearly have a long way to go before I can call myself an expert at this stuff. That said, I am proud of my work here and I don't think there have been many other times where I have learned this much in such a small amount of time. Because it is my first real deep learning project, this summary will be more in-depth than just a showcase. I try to recount all of my main challenges and learnings that came out of the engineering/research process. I also shouldn't go any further without thanking the couple of folks who provided their experienced advice when I needed their guidance. Know that your advice went a long way, helping me push through some of the moments of doubt and uncertainty in this.

Without further ado, the writeup has the following structure
1. Primary tasks
  * Processing and loading a ~200G tabular dataset efficiently
  * Setting up cloud training and file storage infrastructure
  * Designing and training a seq2seq model that converts vector time-series data into text
  * Creating a model deployment that was compliant with the competition evaluation standard
2. Guiding engineering principles
  * Liberal feature sets
  * Model iteration

3. Further research
  * Transfer learning with character-level language models

## Primary tasks

### Dataset preprocessing and loading 

The first thing that struck me in this challenge is how hard it is to work with tabular data. Intuition would lead you to believe the opposite; working with standard format data like audio or image files has significantly more library overhead. I mean just try looking up how to parse a jpeg using the standard libjpeg methods. No, thank you! While that may be true, there are usually only one or a couple different formats that are used to store these types of data due to their uniformity. There are therefore many nice frameworks like torchaudio, tensorflow vision and so forth have been built to abstract almost all of that format wrangling away. In the end, you have to do almost no work between pointing an image loading call at a file name and having a torch tensor ready to go in just milliseconds.

Tabular data has a different story *because* it is such a flexible format. There are many ways in which it can be stored natively: CSVs, parquets, binary records with another layer of encoding, excel spreadsheets (if that finance bro friend of yours has asked you to do some of that "coding stuff"), etc. You can probably name at least a few more popular ones. At the end of the day, there's a reason why data engineers exist and that is to perform this task that no library can do comprehensively.
 
When you take a look at the structure of this competition, it has a relational structure. You have a manifest file and a set of parquet files. The manifest contains a list of all the training examples, their labels and some metadata pointing to their location in the parquet files. Each parquet contains some number of vector time series indexed by the sequence id given in the manifest. Have a look:

![Dataset Schema]({dataset})

There are a couple of issues with using this storage format directly with your training pipeline. Size is one. The parquet files, although they are a binary format, are chunked into files that are very large, take a long time to load and leave a large memory footprint during training. Furthermore, the number of raw features, 1,629 is intuitively high for a model that is just translating fingerspell sequences. Feature selection thus becomes an important consideration at this point.


At the start of this project, I was inclined to cram the maximum number of preprocessing steps I could think of into this stage: coordinate centering, standardization, dimensionality reduction... whatever seemed like I could use, it was probably in my first preprocessing pipeline. I have this bias towards reducing any computational overhead that I see, one of the pitfalls of many novice engineers, and my preprocessing pipeline was an obvious expression of this. Having completed this part of the project, I now know there is a general rule of thumb one can follow to avoid this tendency in the data pipeline. Data should be preprocessed and stored *with as much decoupling from the pipeline as possible*. If you anticipate wanting to use your full feature set in the training stage, then don't do any feature selection or dimensionality reduction. The extra three dollars you'll save in monthly S3 storage fees is just not worth it lol. Take it from me as I learned during this project.


Okay, enough talk. Let's get into the nitty-gritty of how to store this data in the filesystem. You want to choose a scheme that aligns closely with how you plan to load the data during training and validation. If you are using a k-folds validation scheme, you'll probably want to partition the actual files into folds so that samples can be accessed in an independent, shuffled manner within their fold. If you look at [asl_data_preprocessing.ipynb](https://github.com/henry-2025/asl-fingerspell/blob/main/asl_data_preprocessing.ipynb), you'll see that the file names are prefixed with a `fold{n}`, which is how I accomplished this. We then choose a binary format for storing the files. If you are training your model in TensorFlow, I strongly recommend that you use the library's TFRecord framework. Most will underestimate the life-changing impacts of using an integrated binary data storage like this baby. But, if you are willing to take my words of wisdom as a complete newbie to deep learning, this is absolutely the case. When your dataset is stored as TFRecords, you can do parallel reads in your dataloader, you can point your file addresses to google storage objects (stay tuned for the [cloud infrastructure](#cloud-resources") section), you can shuffle your otherwise iterable dataset using a variable-length shuffle buffer. It's literally all done for you in the TF library. For this reason alone, I went against the advice of all my more experienced friends and tried rewriting this entire project for TensorFlow. Honestly, it was not a mistake. 

You can see that in the following, these are the only utilities you need to get your dataset written in the preprocessing stage and loaded in the training stage. Very simple, no boilerplate.

```python
def encode_example(sequence: np.ndarray, frame: np.ndarray):
    feature = dict()
    feature['sequence'] = tf.train.Feature(bytes_list=tf.train.BytesList(value=[sequence.tobytes()])),
    feature['frame'] = tf.train.Feature(bytes_list=tf.train.BytesList(value=[frame.tobytes()]))
    return tf.train.Example(features=tf.train.Features(feature=feature)).SerializeToString() 

  def decode_example(b):
      features = dict()
      features['frame'] = tf.io.FixedLenFeature([], dtype=tf.dtypes.string)
      features['sequence'] = tf.io.FixedLenFeature([], dtype=tf.dtypes.string)

      decoded = tf.io.parse_single_example(b, features)
      decoded['frame'] = tf.reshape(tf.io.decode_raw(decoded['frame'], tf.dtypes.float32), (-1, N_FEATURES, 3))
      decoded['sequence'] = tf.io.decode_raw(decoded['sequence'], tf.dtypes.int64)
      return decoded

  # to write a chunk of examples, we use a utility like the following
  def write_chunk(data, write_dir, chunk_num):
          (frames, seqs), fold = data
          chunk_size = len(frames)
          filename = os.path.join(write_dir, f'fold{fold}-{chunk_num}-{chunk_size}.tfrecord')
          options=tf.io.TFRecordOptions(compression_type='GZIP')
          writer = tf.io.TFRecordWriter(filename, options=options)
          for frame, sequence in zip(frames, seqs):
              encoded_bytes = encode_example(sequence, frame)
              writer.write(encoded_bytes)
          writer.close()

  # then when we want to load a dataset, we just have to map our decode function to the dataset
  ds = tf.data.TFRecordDataset(['fold0-0-256.tfrecord', 'fold0-1-256.tfrecord', ...], compression_type='GZIP')
  ds = ds.map(decode_example)
```
    
The way to do something like chunked parallel loading datasets in PyTorch is a little bit more involved, but as is usual with the PyTorch/TensorFlow matchup, you trade off library conformity and a lot of pre-designed features for flexibility. Looking back on this project now, I do prefer the PyTorch method a little more because it encourages you to understand how your project is engineered as opposed to how to get TensorFlow to do what you want. The following is just a dataset of random normal floats and uniform integers that demonstrates how you can store associated features and labels with `torch.save`. Again, if you want to use worker processes to do the loading, you'll have to implement this yourself as PyTorch doesn't have any utilities for that.

```python
n_sequences = int(1e6)
  max_seq_len = 384
  min_seq_len = 50

  max_label_len = 30
  min_label_len = 10

  n_features = 1630
  n_tokens = 59

  sequence_lengths = torch.randint(n_sequences, min_seq_len, max_seq_len)
  label_lengths = torch.randint(n_sequences, max_label_len, min_label_len)

  # sorry about the for loops, this formatter doesn't support list comprehension
  data = dict()
  for i in range(n_sequences):
    example = dict()
    example['sequence'] = torch.randn((sequence_lengths[i], n_features))
    example['label'] = torch.randint(label_lengths[i], 0, n_tokens)
    data[i] = example

  chunk_size = 256
  for chunk in n_sequences // chunk_size:
    chunk = dict()
    for i in range(chunk_size):
      chunk[i] = data[chunk*chunk_size + i]
    torch.save(chunk, f'chunk{chunk}.pt')
```

### Cloud training and storage resources

I had some prior experience with provisioning cloud resources going into this project. The part of this project that arguably would have had the steepest learning curve---figuring out how the heck to get your AWS service accounts to authenticate on remote instances---was already sort of done for me. 

😎

It was still valuable to spend my time on getting familiarized with the various AI training and deployment tools that the big cloud platforms these days have to offer. The obvious three demands that you have for your cloud platform are file storage for your training data, compute for training, and finally a container service for running deployments. In my familiarity with AWS, I was already acquainted with the S3, EC2, and ECR+ECS workflow. I tried sticking with what I knew about these platforms, avoiding having to use the extra bells and whistles that SageMaker. It turns out that this is very hard for a one-off project where you have to setup all environments by yourself and provision the right EC2 instance hardware to match your training needs. I soon realized that this path would easily double the amount of work required by the project so I decided to pivot to SageMaker or an entirely new cloud platform.

Everyone who knows anyone who codes has probably used Google Colab once or twice before. While there are some painful things about the platform, most notably its weird and unpredictable policy of hardware provisioning, I now realize that it is an insanely good free resource for most of these entry-level ML projects. They give you a 12-hour connection, albeit with spotty coverage, to a [TPU V4 instance](https://cloud.google.com/tpu/docs/system-architecture-tpu-vm). For reference, SageMaker's cheapest tier at $0.74/hour runs on an EC2 `ml.g4dn.xlarge` instance, sporting a Nvidia T4 tensor core and 16 GiB memory which you can spec [here](https://www.nvidia.com/en-us/data-center/tesla-t4/). Maybe I'm penny-pinching to prefer not paying several dollars for each training run, but consider that you will be running your training notebook 20-100 times if you're as new to this as I am. Colab is also the platform on which the Kaggle notebooks run, which was an obvious vote in favor of Google's platform. That list made the decision for me and I didn't really have any doubts in other stages of the development process.

I mentioned before that TensorFlow's dataset objects support the gcloud storage protocol natively as file objects. This is awesome because you don't have to manage any of the file downloads in your own code. Cloud storage file access is also quite fast and was never a bottleneck during any of my training pipelines. I assume the fact that all of my resources were in Google Cloud helped with this.

### Model design

At this stage, you can take of your software and data engineering hat and put on your new generative AI hype PhD research transfer learning pants to try and build the next market-transforming neural network architecture. Well, to be a little more realistic, this stage is just about understanding the task and how it can relate to another problem that has been well-explored in research and industry. This should be a sufficient challenge itself. It also makes sense to understand the constraints of our model, considering a factor of increasing significance is [model scaling](https://arxiv.org/abs/2001.08361)---that is, many models have performance that scales with parameter counts. In this respect, we want to focus on a model that can perform optimally within these constraints. For this specific competition, the model must use less than 40Mb of storage in its disk representation. This gives us an upper bound of 40 million f32 parameters, which is intuitively high given the complexity of our application. For reference, AlexNet has 60 million parameters consisting of 5 convolution layers and 3 fully-connected. It was performant on recognition tasks for 256x256 images.

Considering our application, we want to perform a seq2seq task on vector time-series data where the input sizes are somewhere between 100 and 1000 time-steps of 100-200 dimension vectors. Many jump at the mention of seq2seq in language with the thought of an RNN or transformer architecture. However, consider a subtask where we are looking at frozen frames of a fingerspell sequence. Labeling a single character using a single time step or a window around a given time step should be possible. Moreover, it should be the fundamental subtask that we want our model to be good at performing. The lack of any long-term dependence in sequence steps means that the strengths of these two architectures are more or less irrelevant. The 1D CNN, on the other hand, has a stronger .


1. Choosing the right mechanism and engineering an efficient, data-accurate block
2. seq2seq models for unaligned sequence data 3. More advanced mechanisms * ECA
for better channel attention * Late dropout * Padding mechanisms ## Model deployment
1. how to and how not to do model deployment There are like two ways that you can
do your engineering: you can either 1. Start by building a the simplest system
you can that addresses the highest priority tasks or... 2. Try to engineer every
single feature that comes to mind as you build your first iteration ## Model iteration
1. explain why this section should go above 2. explain While model training and
architecture was a big focus of my work on this project, a surprising amount of
work went into learning the tools and techniques required data loading pipelines.
Since this was my first real deep learning project beyond the "Hello World" neural
nets that you build in the Pytorch intro page, Most of my for large datasets. It
is also my first real deep learning project, which I think it My submission for
the Google ASL Kaggle challenge. ![training example]({sequence})

```python
import torch
class Test(torch.nn):
    def __init__(self):
        self.layer1 = torch.nn.Linear(10, 100, dtype=np.float32)
```
### Model Deployment

### Model Iteration
