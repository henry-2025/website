<script>
    import {slide, fly} from 'svelte/transition';
    let open = false;

</script>

<button type="button" class="-m-2.5 inline-flex items-center justify-center rounded-md p-2.5 text-gray-700">
  <svg class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" aria-hidden="true">
    <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
  </svg>
</button>

<style>
    .dropdown {
        position: absolute;
        top: 58px;
        width: 300px;
        z: 3;
        transform: translateX(-45%);
        background-color: var(--bg);
        border: var(--border);
        border-radius: var(--border-radius);
        padding: 1rem;
        overflow: hidden;
        transition: height var(--speed) ease;
    }
</style>
