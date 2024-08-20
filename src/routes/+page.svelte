<script>
  /**
   * @typedef Article
   * @type{object}
   * @property {string} title
   * @property {string} path
   * @property {Date} date
   * @property {string} description
   */

  /**
   * @typedef Data
   * @property{Array<Article>} articles
   */

  /** @type {Data} */
  export let data;
</script>

<div class="project-content">
  <div class="project-list">
    {#each data.articles as article}
      <a href="/projects/{article.path}">
        <div class="article-summary">
          <h3>
            {article.title}
          </h3>
          <time>
            {`${article.date.getMonth()}.${article.date.getDay()}.${article.date.getFullYear()}`}
          </time>
          <p>
            {article.description}
          </p>
        </div>
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
