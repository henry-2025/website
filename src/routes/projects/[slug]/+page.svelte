<script>
    export let data;
</script>

<article>
    <h1>{data.title}</h1>
    <p style="font-style:italic">Published: {data.date}</p>
    <svelte:component this={data.content} />
</article>