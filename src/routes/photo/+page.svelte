<script lang="ts">
    import type { Photo } from "$lib/photo/photos.ts";

    let props = $props();
    let { data, children }: { data: { photos: Array<Photo> }; children: any } = props;

    console.log(props)


    // Fisher-Yates implementation
    function shuffle<T>(array: T[]) {
        let currentIndex = array.length,
            randomIndex;

        while (currentIndex !== 0) {
            randomIndex = Math.floor(Math.random() * currentIndex);
            currentIndex--;

            [array[currentIndex], array[randomIndex]] = [
                array[randomIndex],
                array[currentIndex],
            ];
        }

        return array;
    }
    let shuffledPHotos = shuffle(data.photos);
</script>

<h2>Photo</h2>


