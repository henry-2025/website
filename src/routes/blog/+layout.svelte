<script lang="ts">
  import type { Post } from "$lib/blog/posts";
  import { page } from "$app/state";
  let { data, children }: { data: { posts: Array<Post> }; children: any } =
    $props();

  let activePost: Post | undefined = data.posts
    .filter((post) => page.url.pathname.endsWith(post.slug))
    .at(0);
  if (activePost === undefined) {
    throw new Error(
      `unable to find post metadata for the slug ${page.url.pathname}`,
    );
  }
</script>

<article>
  <h1>{activePost.title}</h1>
  <h2>{activePost.description}</h2>
  <div>
    <time>
      {`${activePost.date.getMonth() + 1}.${activePost.date.getDate() + 1}.${activePost.date.getFullYear()}`}
    </time>
  </div>
  {@render children()}
</article>

<link rel="stylesheet" href="/style/prism.css" />

<svelte:head>
  <style>
    @import url("https://fonts.googleapis.com/css2?family=Playfair+Display:ital,wght@0,400..900;1,400..900&display=swap");
  </style>
</svelte:head>

<style>
  article {
    text-align: start;
    margin: 5em 2em;
    padding: 2em 3em;
    font-family: "Fairplay Display";
  }
  :global(article > p) {
    text-indent: 2rem;
  }
  :global(article > :is(h1, h2, h3, .subtitle)) {
    text-align: center;
  }
  :global(article > h1) {
    margin: 1em 0 2em 0;
  }
  :global(article > .subtitle) {
    font-size: large;
  }
  :global(article > pre) {
    background-color: white;
    padding: 1em;
    border-radius: 1em;
    white-space: pre-wrap;
  }
  :global(article > table) {
    font-family: "Inconsolata", monospace;
  }
  :global(article > p > strong) {
    font-size: 3em;
    color: rgba(150, 168, 227, 100);
  }

  :global(figure > img) {
    width: 100%;
  }

  :global(img) {
    max-width: 100%;
  }

  :global(article > figure) {
    float: inline-start;
    max-width: 50%;
    margin-inline: 1em;
    margin-block: 0;
  }
  :global(figure>figcaption) {
    text-align: center;
    font-size: small;
  }

  :global(article > .footnotes) {
    word-wrap: break-word;
    font-size: small;
    background-color: white;
    padding: 2em 3em;
  }

  @media (max-width: 800px) {
    article {
      margin: 0;
      padding: 0;
      border-radius: 1em;
    }
  }
</style>
