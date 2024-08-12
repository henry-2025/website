<script>
  export let data;
</script>

<link rel="stylesheet" href="/style/prism.css" />

<div>{@html data.content}</div>

<style>
  :global(article) {
    text-align: start;
    margin: 5em 2em;
    padding: 2em 3em;
    background-color: rgba(150, 168, 227, 100);
    border-radius: 1em;
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

  @media (max-width: 800px) {
    :global(article) {
      margin: 2em 0;
      padding: 2em 3em;
      border-radius: 1em;
    }
  }
</style>
