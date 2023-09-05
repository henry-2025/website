<script>
  /** @type {import('./$types').PageData} */
  export let data;
</script>

<h2>Projects</h2>
<div
    class="mx-auto mt-10 grid max-w-2xl grid-cols-1 gap-x-8 gap-y-16 border-t border-gray-400 pt-10 sm:mt-16 sm:pt-16 lg:mx-0 lg:max-w-none lg:grid-cols-3"
>
        {#each data.meta as meta}
    <article class="flex max-w-xl flex-col items-start justify-between">
        <div class="flex items-center gap-x-4 text-xs">
            <time datetime={meta.date} class="text-gray-500">{meta.date.toDateString()}</time>
            <div class="relative z-10 rounded-full text-center bg-gray-300/50 pr-3 py-1.5 font-medium text-gray-600">
                    {meta.category}
                </div>
        </div>
        <div class="group mt-4 relative px-4 rounded-md transition ease-in-out delay-150 bg-gray-200 hover:-translate-y-1 hover:scale-110 hover:bg-indigo-500 duration-300">
            <h3
                class="mt-3 text-lg font-semibold leading-6 text-gray-900 group-hover:text-gray-600"
            >
                <a href={meta.path}>
                    <span class="absolute inset-0" />
                        {meta.title}
                </a>
            </h3>
            <p class="mt-5 line-clamp-3 indent-2 text-sm leading-6">
                    {meta.summary}
            </p>
        </div>
    </article>
        {/each}
</div>

<style>
    a {
        color: #000e47;
    }
</style>
