<script>
    import Header from "./Header.svelte";
    import "../app.css";
    import Gradient from "./Gradient.svelte";
    import Footer from "./Footer.svelte";
</script>

<div class="app">
    <Header />
    <slot />
    <Gradient />
    <Footer />
</div>
