<script>
  export let data;
</script>

<link rel="stylesheet" href="/style/prism.css" />

<div id="article">{@html data.content}</div>

<style>
  :global(article) {
    text-align: start;
  }
  :global(article > p) {
    text-indent: 2rem;
  }
  :global(article > :is(h1, h2, h3)) {
    text-align: center;
  }
  :global(article > h1) {
    margin: 4em 0;
  }
  :global(article > pre) {
    background-color: white;
    padding: 1em;
    border-radius: 1em;
    white-space: pre-wrap;
  }

  #article {
    margin: 5em 2em;
    padding: 0 3em 0 3em;
    background-color: rgba(150, 168, 227, 100);
    border-radius: 1em;
  }
</style>
