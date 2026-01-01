<p>
    <strong>B</strong>eing able to host a dance party in my apartment might be a
    thing of a past, more collegiate era. I nonetheless embarked on this project
    to engineer the ultimate rave-style light show beneath the occasional
    footsteps of my unknowing landlords. There were ultimately more hours sunk
    into this than I would like to admit and this work is still very much in a
    project state, but I am quite proud of what I was able to get out of my
    hobby coding ambitions on this one.
</p>

<p>
    To start off, what exactly did I build here and isn't it just those cheap,
    blinding LED strips that can be found hanging around the average male
    college dorm room? The short answer is yes, I build a controller to create
    visuals on an addressable LED strip. The longer answer is that I created
    extensible software to be transform system audio into a byte protocol that
    can be streamed wirelessly to this network-enabled red flag <a
        href="#footnote-1">[1]</a
    >.
</p>

<img
    src="/blog/reactive-audio-visualizer/audio-reactive-system-design.png"
    alt="high-level system layout"
/>

<h3>Esp LED Driver</h3>

<p>
    The simplest part of this system is the LED driver, which takes a
    dumb-simple packet format that I came up with and reads a UDP stream
    containing nothing but 4-byte chunks in the following format:
</p>
<pre><code>
    struct packet_bytes {"{"}
        char led_index;
        char r_led;
        char g_led;
        char b_led;
    {"}"};

    union packet_union {"{"}
        packet_bytes packet_rep;
        int int_rep;
    {"}"};
</code></pre>

<p>
    As I said earlier, this runs on an Esp8266 (actually an esp-01 that I had laying around from )
    I decided some time ago that systems was the lowest level I could go while
    still maintaining my sanity writing code, so I ended up using the Esp8266
    and NeoPixel libraries to do the heavy lifting with networking, addressing
    the LED strip, and whatever ESP wrangling needs to be done. The arduino sketch for this was thankfully incredibly small and is probably not going to see many code changes going forward
</p>

<div class="footnotes">
    <div id="footnote-1">
        [1] <a
            href="https://www.reddit.com/r/dating_advice/comments/10u9qkg/tldr_are_led_strip_lights_a_dating_red_flag_if_so/"
            >https://www.reddit.com/r/dating_advice/comments/10u9qkg/tldr_are_led_strip_lights_a_dating_red_flag_if_so/</a
        >
    </div>
</div>
