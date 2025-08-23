<script lang="ts">
  let { children } = $props();
</script>

<article>
  {@render children()}
</article>

<link rel="stylesheet" href="/style/prism.css" />

<svelte:head>
  <style>
    @import url("https://fonts.googleapis.com/css2?family=Roboto+Condensed:ital,wght@0,100..900;1,100..900&display=swap");
  </style>
</svelte:head>

<style>
  article {
    text-align: start;
    margin: 5em 2em;
    padding: 2em 3em;
    background-color: rgba(150, 168, 227, 100);
    border-radius: 1em;
    font-family: "Roboto Condensed";
  }
  :global(article > p) {
    text-indent: 2rem;
  }
  :global(article > :is(h1, h2, h3)) {
    text-align: center;
  }
  :global(article > h1) {
    margin: 1em 0 2em 0;
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

  @media (max-width: 800px) {
    :global(article) {
      margin: 1em 0;
      padding: 2em 3em;
      border-radius: 1em;
    }
  }
</style>
