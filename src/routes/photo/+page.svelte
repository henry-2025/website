<script lang="ts">
    import type { Photo } from "$lib/photo/photos.ts";
    import { fade } from "svelte/transition";

    let props = $props();
    let { data, children }: { data: { photos: Array<Photo> }; children: any } =
        props;

    let selectedPhoto: Photo | null = $state(null);

    function openLightbox(photo: Photo) {
        selectedPhoto = photo;
        document.body.style.overflow = "hidden";
    }

    function closeLightbox() {
        selectedPhoto = null;
        document.body.style.overflow = "";
    }

    function handleKeyDown(event: KeyboardEvent) {
         if (event.key === "Escape") {
            closeLightbox();
        }
    }
</script>

<svelte:window on:keydown={handleKeyDown}></svelte:window>

<h2>Photo</h2>

<div class="random-layout">
    {#each data.photos as photo}
        <button class="masonry-item" onclick={() => openLightbox(photo)}>
            <img src="https://static.jhpick.com/{photo.file}" alt={photo.title} />
        </button>
    {/each}
</div>

{#if selectedPhoto}
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <div
        transition:fade={{ duration: 200 }}
        class="lightbox"
        onclick={closeLightbox}
    >
        <button class="close-btn" onclick={closeLightbox}>[x]</button>
        <div class="lightbox-content" onclick={(e) => e.stopPropagation()}>
            <img
                src="https://static.jhpick.com/{selectedPhoto.file}"
                alt={selectedPhoto.title}
            />
            <p class="lightbox-caption">{selectedPhoto.title}</p>
        </div>
    </div>
{/if}

<style>
    .random-layout {
        display: flex;
        flex-wrap: wrap;
        gap: 1rem;
        padding: 1rem;
        align-items: flex-start;
    }

    .masonry-item {
        transition: transform 0.2s;
        background: none;
        border: none;
        padding: 0;
        cursor: pointer;
    }

    .masonry-item img {
        width: 100%;
        height: auto;
        display: block;
        border-radius: 4px;
        object-fit: cover;
    }

    .masonry-item:hover {
        transform: scale(1.05);
        z-index: 10;
    }

    /* Lightbox styles */
    .lightbox {
        position: fixed;
        top: 0;
        left: 0;
        width: 100vw;
        height: 100vh;
        background: rgba(0, 0, 0, 0.9);
        display: flex;
        align-items: center;
        justify-content: center;
        z-index: 1000;
        cursor: pointer;
    }

    .lightbox-content {
        max-width: 90vw;
        max-height: 90vh;
        display: flex;
        flex-direction: column;
        align-items: center;
        cursor: default;
    }

    .lightbox-content img {
        max-width: 100%;
        max-height: 85vh;
        object-fit: contain;
        border-radius: 4px;
    }

    .lightbox-caption {
        color: white;
        margin-top: 1rem;
        font-size: 1rem;
        text-align: center;
    }

    .close-btn {
        position: absolute;
        top: 1rem;
        background: transparent;
        right: 1rem;
        border: none;
        color: white;
        font-size: 1rem;
        width: 3rem;
        height: 3rem;
        border-radius: 50%;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
        line-height: 1;
        transition: 0.2s;
        z-index: 1001;
    }

    /* Random scaling and vertical offset pattern */
    .masonry-item:nth-child(13n + 1) {
        flex: 0 0 calc(43% - 1rem);
        margin-top: 0;
        margin-left: 5rem;
    }

    .masonry-item:nth-child(13n + 2) {
        flex: 0 0 calc(29%);
        margin-top: 0.5rem;
    }

    .masonry-item:nth-child(13n + 3) {
        flex: 0 0 calc(75% - 1rem);
        margin-top: 2rem;
        margin-left: 2rem;
    }

    .masonry-item:nth-child(13n + 4) {
        flex: 0 0 calc(43%);
        margin-top: 2rem;
    }

    .masonry-item:nth-child(13n + 5) {
        flex: 0 0 calc(81% + 5em);
        margin-top: 1rem;
        margin-left: 2rem;
    }

    .masonry-item:nth-child(13n + 6) {
        flex: 0 0 calc(72%);
        margin-top: 3rem;
        margin-left: 0;
    }

    .masonry-item:nth-child(13n + 7) {
        flex: 0 0 calc(52% - 1rem);
        margin-top: 3rem;
        margin-left: 3rem;
    }

    .masonry-item:nth-child(13n + 8) {
        flex: 0 0 calc(38%);
        margin-top: 1rem;
    }

    .masonry-item:nth-child(13n + 9) {
        flex: 0 0 calc(79% - 1rem);
        margin-top: 5rem;
        margin-left: 6rem;
    }

    .masonry-item:nth-child(13n + 10) {
        flex: 0 0 calc(42%);
        margin-top: 4rem;
    }

    .masonry-item:nth-child(13n + 11) {
        flex: 0 0 calc(35%);
        margin-top: 4rem;
        margin-left: 2rem;
    }

    .masonry-item:nth-child(13n + 12) {
        flex: 0 0 calc(33%);
        margin-top: 6rem;
    }

    .masonry-item:nth-child(13n) {
        flex: 0 0 calc(36%);
        margin-top: 2.5rem;
    }

    @media (max-width: 768px) {
        .masonry-item:nth-child(n) {
            flex: 0 0 calc(48% - 1rem) !important;
            margin-top: 0 !important;
        }
    }

    @media (max-width: 480px) {
        .masonry-item:nth-child(n) {
            flex: 0 0 100% !important;
            margin-top: 0 !important;
        }
    }
</style>
