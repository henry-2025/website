const projects = [{
    title: "Kaggle ASL Fingerspell: Pose estimates to text seq2seq model",
    summary: "Methods and some takeaways from my first big deep learning project",
    category: "Deep Learning",
    date: new Date(["2023-08-27", '00:00']),
    path: '/projects/asl-fingerspell'
}];
// For some reason, svelte does not let you load files outside of a parameter call. this is something I am working on
// a list of pages we want to render
//let pages = ['asl-fingerspell'];


// parse the metadata from each post markdown file
/** @type {import('./$types').PageLoad} */
export async function load({ }) {
    //let posts = pages.map(async (page) => {
    //    return {
    //        content: await import(/*@vite-ignore*/`./${page}.svelte.md`),
    //        path: path.join("/projects/", page),
    //    }
    //});
    //posts = await Promise.all(posts);

    //let meta = posts.map((post) => {
    //    return {
    //        title: post.content.metadata.title,
    //        summary: post.content.metadata.summary,
    //        category: post.content.metadata.category,
    //        date: new Date([post.content.metadata.date, '00:00']),
    //        path: post.path
    //    }
    //}
    //).reverse();
    return {
        meta: projects
    };
}
