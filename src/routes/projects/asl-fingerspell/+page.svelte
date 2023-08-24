<div class="article-container">
    <h2>
        Google ASL Fingerspell Kaggle Challenge: Pose estimates &#8594 text
        seq2seq model
    </h2>
    <p>
        This was my submission to Google's ASL Fingerspell Challenge and
        probably my first real deep learning project beyond the Hello World
        neural nets they have you build on the Pytorch intro page. It's not much
        and I still clearly have a long way to go before I can call myself a
        qualified ML engineer. That said, I am proud of my work here and I don't
        think there have been many other times where I have learned this much in
        such a small amonut of time.
    </p>
    <p>The main tasks in this project were as follows:</p>
    <ol type="1">
        <li>
            <a href="#processing"
                >Processing and loading a ~200G tabular dataset efficiently</a
            >
        </li>
        <li>
            <a href="#cloud-resources"
                >Setting up cloud training and file storage infrastructure</a
            >
        </li>
        <li>
            <a href="#model-design"
                >Designing and training a seq2seq model that converts vector
                time-series data into text</a
            >
        </li>
        <li>
            <a href="#model-deployment"
                >Creating a model deployment that was compliant with the
                competition evaluation standard</a
            >
        </li>
        <li><a href="#model-iteration">Model iteration</a></li>
    </ol>

    <p>
        I go into summary detail on each of these sections, documenting my
        methods and recounting challenges that I encountered during the
        engineering and design processes.
    </p>

    <h3 id="processing">Dataset preprocessing and loading</h3>
    <p>
        The first thing that struck me in this challenge is how hard it is to
        work with tabular data. Intuition would lead you to believe the
        opposite; working with standard format data like audio or image files
        has significantly more library overhead. I mean just try looking up how
        to parse a jpeg using the standard libjpeg methods. No, thank you! While
        that may be true, there are usually only one or a couple different
        formats that are used to store these types of data due to their
        uniformity. There are therefore many nice frameworks like torchaudio,
        tensorflow vision and so forth have been built to abstract almost all of
        that format wrangling away. In the end, you have to do almost no work
        between pointing an image loading call at a file name and having a torch
        tensor ready to go in just milliseconds.
    </p>

    <p>
        Tabular data has a different story <i>because</i> it is such a flexible format.
        There are many ways in which it can be stored natively: CSVs, parquets, binary
        records with another layer of encoding, excel spreadsheets (if that finance
        bro friend of yours has asked you to do some of that &ldquo;coding stuff&rdquo;)... You
        can probably name at least a few more popular ones. At the end of the day,
        there&apos;s a reason why data engineers exist and that is to perform this task
        that no library can do comprehensively.
    </p>

    <p>
        When you take a look at the structure of this competition, it has a
        relational structure. You have a manifest file and a set of parquet
        files. The manifest contains a list of all the training examples, their
        labels and some metadata pointing to their location in the parquet
        files. Each parquet containts some number of vector time series indexed
        by the sequence id given in the manifst. Have a look:
    </p>

    <img src="$lib/assets/asl-fingerspell/dataset.png" alt="" srcset="" />
    <p>
        There are a couple of issues with using this storage format directly
        with your training pipeline. Size is one. The parquet files, although
        they are a binary format, are chunked into files that are very large,
        take a long time to load and leave a large memory footprint during
        training. Furthermore, the number of raw features, 1,629 is intuitively
        high for a model that is just translating fingerspell sequences. Feature
        selection thus becomes an important consideration at this point.
    </p>
    <p>
        At the start of this project, I was inclined to cram the maximum number
        of preprocessing steps I could think of into this stage: coordinate
        centering, standardization, dimensionality reduction... whatever seemed
        like I could use, it was probably in my first preprocessing pipeline. I
        have this bias towards reducing any computational overhead that I see,
        one of the pitfalls of many novice engineers, and my preprocessing
        pipeline was an obvious expression of this. Having completed this part
        of the project, I now know there is a general rule of thumb one can
        follow to avoid this tendency in the data pipeline. Data should be
        preprocessed and stored <i>with as much decoupling from the pipeline as
        possible</i>. If you anticipate wanting to use your full feature set in the
        training stage, then don't do any feature selection or dimensionality
        reduction. The extra three dollars you'll save in monthly S3 storage
        fees is just not worth it lol. Take it from me as I learned during this
        project.
    </p>
    <p>
        Okay, enough talk. Let's get into the nitty-gritty of how to store this
        data in the filesystem. You want to choose a scheme that aligns closely
        with how you plan to load the data during training and validation. If
        you are using a k-folds validation scheme, you'll probably want to
        partition the actual files into folds so that samples can be accessed in
        an independent, shufled manner within thir fold. If you look at
        <a
            href="https://github.com/henry-2025/asl-fingerspell/blob/main/asl_data_preprocessing.ipynb"
            >asl_data_preprocessing.ipynb</a
        >, you'll see that the file names are prefixed with a `fold&#123n&#125`,
        which is how I accomplished this. We then choose a binary format for
        storing the files. If you are training your model in TensorFlow, I
        strongly recommend that you use the library's TFRecord framework. Most
        will underestimate the life-changing impacts of using an integrated binary data
        storage like this baby. But, if you are willing to take my words of
        wisdom as a complete newbie to deep learning, this is absolutely the
        case. When your dataset is stored as TFRecords, you can do parallel
        reads in your dataloader, you can point your file addresses to google
        storage objects (stay tuned for the <a href="#cloud-resources">cloud infrastructure</a> section), you
        can shuffle your otherwise iterable dataset using a variable-length
        shuffle buffer. It's literally all done for you in the the TF library.
        For this reason alone, I went against the advice of all my more
        experienced friends and tried rewriting this entire project for
        TensorFlow. Honestly, it was not a mistake.
    </p>
    <p>
        If you are a bit more sensible and want to write things up in PyTorch,
        that's totally okay as well. Pytorch has a binary tensor storage mechanism that works just fine. 2.
        Best methods for efficient and practical storage 4. Augmentation 5. PCA
    </p>

    <h3 id="cloud-resources">Cloud training and storage resources</h3>

    1. Why cloud tools are necessary and nice, despite being another tooling
    overhead 2. The essentials 3. What I still have to figure out 4. Building a
    pipeline

    <h3 id="model-design">Model design</h3>

    1. Choosing the right mechanism and engineering an efficient, data-accurate
    block 2. seq2seq models for unaligned sequence data 3. More advanced
    mechanisms * ECA for better channel attention * Late dropout * Padding
    mechanisms ## Model deployment 1. how to and how not to do model deployment
    There are like two ways that you can do your engineering: you can either 1.
    Start by building a the simplest system you can that addresses the highest
    priority tasks or... 2. Try to engineer every single feature that comes to
    mind as you build your first iteration ## Model iteration 1. explain why
    this section should go above 2. explain While model training and
    architecture was a big focus of my work on this project, a surprising amount
    of work went into learning the tools and techniques required data loading
    pipelines. Since this was my first real deep learning project beyond the
    "Hello World" neural nets that you build in the Pytorch intro page, Most of
    my for large datasets. It is also my first real deep learning project, which
    I think it My submission for the Google ASL Kaggle challenge. ![training
    example](./fig/example_feature_sequence.gif)

    <h3 id="model-deployment">Model Deployment</h3>
    <h3 id="model-iteration">Model Iteration</h3>
</div>
