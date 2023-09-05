import path from 'path';

// a list of pages we want to render
let pages = ['asl-fingerspell'];


// parse the metadata from each post markdown file
/** @type {import('./$types').PageServerLoad} */
export async function load({ }) {
    let posts = pages.map(async (page) => {
        return {
            content: await import(/*@vite-ignore*/`./${page}.svelte.md`),
            path: path.join("/projects/", page),
        }
    });
    posts = await Promise.all(posts);

    let meta = posts.map((post) => {
        return {
            title: post.content.metadata.title,
            summary: post.content.metadata.summary,
            category: post.content.metadata.category,
            date: new Date([post.content.metadata.date, '00:00']),
            path: post.path
        }
    }
    ).reverse();
    return {
        meta: meta
    };
}
