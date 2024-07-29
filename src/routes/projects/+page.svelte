<script>
  /** @type {import('./$types').PageServerData} */
  export let data;
</script>

<div class="project-content">
  <h2>Projects</h2>
  <div class="project-list">
    {#each data.articles as article}
      <a href="/projects/{article.path}">
        <article>
          <h3>
            {article.title}
          </h3>
          <time>
            {article.date}
          </time>
          <p>
            {article.description}
          </p>
        </article>
      </a>
    {/each}
  </div>
</div>

<style>
  .project-content {
    display: flex;
    align-items: center;
    flex-direction: column;
  }

  .project-list {
    display: grid;
    gap: 10px;
    grid-template-columns: repeat(3, 1fr);
  }

  article > h3 {
    font-size: 1em;
    font-weight: 800;
    min-height: 30%;
  }

  article > p {
    font-size: 0.75em;
  }

  article > time {
    font-size: 0.75em;
    background-color: rgba(255, 2500);
    border-radius: 1rem;
    padding: 0.5em;
  }

  article {
    width: 10em;
    background-color: rgba(150, 168, 227, 100);
    border-radius: 1em;
    padding: 1rem;
    transition: 0.2s;
    align-content: center;
  }
  a {
    all: none;
  }

  article:hover {
    scale: 1.05;
    cursor: pointer;
  }
</style>
