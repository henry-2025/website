export interface Photo {
    title?: string;
    file: string;
    date: Date;
    caption: string;
    category: Category;
}

export enum Category {
    Engineering,
    Music, 
    Running,
    Nature,
}

export const photos: Array<Photo> = [
    {
        file: "IMG_3964.jpg",
        date: new Date("12-26-2025"),
        caption: "light pollution against dunes",
        category: Category.Nature,
    }
]