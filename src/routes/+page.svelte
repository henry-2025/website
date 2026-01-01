<script lang="ts">
  import { Category, type Post } from "$lib/blog/posts.js";
    import ArticleSection from "./ArticleSection.svelte";

  export let data: { posts: Array<Post> };
  let sortedPosts = data.posts.sort(
    (a, b) => b.date.getTime() - a.date.getTime(),
  );
</script>

<div class="project-content">
  <ArticleSection tag={Category.Running} displayTitle="Run" sortedPosts={sortedPosts}/>
  <ArticleSection tag={Category.Engineering} displayTitle="Engineering Projects" sortedPosts={sortedPosts}/>
</div>

<style>
  .project-content {
    display: flex; 
    align-items: center;
    flex-direction: column;
  }
</style>
