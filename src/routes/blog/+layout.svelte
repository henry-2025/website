<script lang="ts">
  let { children } = $props();
</script>

<article>
  {@render children()}
</article>

<link rel="stylesheet" href="/style/prism.css" />

<svelte:head>
 <style>
@import url('https://fonts.googleapis.com/css2?family=Playfair+Display:ital,wght@0,400..900;1,400..900&display=swap');
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
  :global(article>p>strong) {
    font-size: 3em;
    color: rgba(150, 168, 227, 100);
  }

  :global(article>img) {
    max-width: 100%;
  }

  :global(article>.footnotes) {
    word-wrap: break-word;
    font-size: small;
    background-color: white;
    padding: 2em 3em;
  }

  @media (max-width: 800px) {
    :global(article) {
      margin: 1em 0;
      padding: 2em 3em;
      border-radius: 1em;
    }
  }
</style>
