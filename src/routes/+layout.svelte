<script>
    import Header from './Header.svelte';
    import '../app.css';
    import Gradient from './Gradient.svelte';
</script>

<div class="app">
    <Header />
    <Gradient />
    <main>
        <slot />
    </main>
</div>
