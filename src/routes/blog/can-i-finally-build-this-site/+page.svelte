<p>
    <strong>A</strong> lot of people looked at me with concern when I told them I spent an
    afternoon writing a backend using the go standard library to run this blog
    site. Why? There are tens of existing blog libraries for every which ui
    framework, so I could be saving myself not only an afternoon of redundant
    coding but also the work of one coder is almost never as good as a group of
    open source developers? Unless I can create a better experience than
    existing frameworks, then it makes no sense to try to recreate their
    functionality. I acknowledge all of these points and at every step of the
    way, it made me uncomfortable going against what seemed to be conventional
    engineering advice. With the project having reached a presentable state,
    however, I stand confidently on the side of my decisinion to do all this.
</p>

<p>
    There are a number of reasons for this, but number one is recognizing the
    short-term needs of the project. A blog site is by no means a complex piece
    of software and it does not need a team of engineers or a lot of thought to
    be designed and maintained. It is also a highly opinionated piece of
    software. It might be asked to do any number of things including photo
    hosting, dynamic routing, page generation, etc. With such a shallow but
    diffuse set of engineering needs, having a framework to solve some
    boilerplate work starts to seem mismatched with the task at hand.
</p>

<p>
    Say you decide to look beyond this slight conundrum and start with a
    framework anyway. About five months back, I did with Svelte and a markdown
    rendering plugin called mdsvex. Everything was going just fine until I
    wanted to start serving images in one of my project posts. Sounds like a
    simple task, no? To my frustration, there was a solution to this problem but
    it required me to convert my markdown documents (the language of the blog
    source) into a marked-up version of the source that contained embedded
    javascript. Alongside this, I had to learn some Svelte fundamentals about
    file preprocessing to understand why exactly this was necessary.
</p>

<p>
    I ended up accepting the ugly solution and pushing forward towards the
    presentable but simple site that I was envisioning. It was not long before I
    began running into a new issue when trying to render python code snippets
    inside a project writeup. Some of my snippets used list comprehension to
    create dictionaries more succinctly but the mdsvex preprocessor consumes all
    unescaped brackets, within a code block or not. Reimplementing the markdown
    parsing was far beyond my patience and expertise at this point and my
    solution this time was, embarassingly, to write all of my list
    comprehensions as for loops. It was at this point that I finished the first
    iteration of the article and decided to throw the site up in production.
</p>

<p>
    I had just sacrificed some of the very basic qualities of what I was trying
    to deliver for the supposed benefits of a frameworked solution. The time
    cost of this was also enormous; it took _days_ to gain even the most basic
    understanding of all the components I ended up using in this project.
</p>
