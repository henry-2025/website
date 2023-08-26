<script>
	import Header from "./Header.svelte";
	import "../app.css";
	import Gradient from "./Gradient.svelte";
	import Footer from "./Footer.svelte";
	import "$lib/styles/style.scss";
	import { page } from "$app/stores";
</script>

<div class="app">
	<Header />
	{#if $page.url.pathname == "/"}
        <div class="mx-auto max-w-7xl px-6 lg:px-8 py-24, sm:py-20 mb-[10vh]">
		<slot />
    </div>
	{:else}
		<div
			class="mx-auto max-w-7xl px-6 lg:px-8 py-24 sm:py-32 bg-gradient-to-tr from-white min-h-screen"
		>
			<slot />
		</div>
	{/if}
	<Gradient />
	<Footer />
</div>
