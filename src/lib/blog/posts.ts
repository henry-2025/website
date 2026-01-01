
export interface Post {
    title: string;
    slug: string;
    date: Date;
    description: string;
    category: Category;
}

export enum Category {
    Engineering,
    Music,
    Running,
}

export const posts: Array<Post> = [
    {
        title: "ASL Fingerspell",
        slug: "asl-fingerspell",
        date: new Date("2023-08-27"),
        description:
            "designing, training, and deploying a seq2seq model that estimates sign language from pose estimates",
        category: Category.Engineering,
    },
    {
        title: "Can I Finally Build This Site?",
        slug: "can-i-finally-build-this-site",
        date: new Date("2024-03-20"),
        description: "What taking two years to put this site up taught me about software engineering",
        category: Category.Engineering,
    },
    {
        title: "Building a Reactive Audio Visualizer",
        slug: "reactive-audio-visualizer",
        date: new Date("2025-04-20"),
        description: "Processing, streaming, and visualizing sounds on some hardware I put together",
        category: Category.Engineering,
    },
    {
        title: "Drinking the Kool-Aid",
        slug: "10-years-running",
        date: new Date("2026-01-01"),
        description: "The truncated version of my 10-years’ journey from high-school newbie to washed-up hobby racer, some takeaways, and where to now",
        category: Category.Running,
    },
]