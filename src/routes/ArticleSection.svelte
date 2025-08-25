<script lang="ts">
    import type { Category, Post } from "$lib/blog/posts";

    let props: {
        tag: Category;
        displayTitle: string;
        sortedPosts: Array<Post>;
    } = $props();
</script>

<h3>{props.displayTitle}</h3>
<hr>
<div class="project-list">
    {#each props.sortedPosts as post}
        {#if props.tag == post.category}
            <a href="/blog/{post.slug}">
                <div class="article-summary">
                    <h3>
                        {post.title}
                    </h3>
                    <time>
                        {`${post.date.getMonth() + 1}.${post.date.getDate() + 1}.${post.date.getFullYear()}`}
                    </time>
                    <p>
                        {post.description}
                    </p>
                </div>
            </a>
        {/if}
    {/each}
</div>

<style>

   hr {
    background-color: black;
    height: 0.1px;
    width: 100%;
   }
    .project-list {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        align-items: center;
    }

    .article-summary > h3 {
        font-size: 1em;
        font-weight: 800;
        min-height: 30%;
    }

    .article-summary > p {
        font-size: 0.75em;
    }

    .article-summary > time {
        font-size: 0.75em;
        background-color: rgba(175, 194, 255, 100);
        border-radius: 1rem;
        padding: 0.5em;
    }

    .article-summary {
        width: 10em;
        background-color: rgba(150, 168, 227, 100);
        border-radius: 1em;
        padding: 1rem;
        transition: 0.2s;
        align-content: center;
        margin: 1em;
    }

    a {
        color: inherit;
        text-decoration: inherit;
    }
    .article-summary:hover {
        scale: 1.05;
        cursor: pointer;
    }

    @media (max-width: 800px) {
        .project-list {
            display: flex;
            flex-direction: column;
            align-items: center;
        }

        .article-summary {
            background-color: rgba(150, 168, 227, 100);
            border-radius: 1em;
            padding: 1rem;
            transition: 0.2s;
            align-content: center;
            margin: 1em;
            width: fit-content;
        }
    }
</style>
