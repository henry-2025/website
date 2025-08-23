
export interface Post {
    title: string;
    slug: string;
    date: Date;
    description: string;
}

export const posts: Array<Post> = [
    {
        title: "ASL Fingerspell",
        slug: "asl-fingerspell",
        date: new Date("2023-08-27"),
        description:
            "designing, training, and deploying a seq2seq model that estimates sign language from pose estimates",
    },
    {
        title: "Iterations of This Site",
        slug: "this-site",
        date: new Date("2024-03-20"),
        description: "A documentation of the iterations of this site",
    }
]