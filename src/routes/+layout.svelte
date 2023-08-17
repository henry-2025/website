<script>
    import Header from "./Header.svelte";
    import "../app.css";
    import Gradient from "./Gradient.svelte";
</script>

<div class="app">
    <Header />
    <slot />
    <Gradient />
</div>
