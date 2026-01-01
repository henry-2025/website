export interface Photo {
    title: string;
    file: string;
    date: Date;
    category: Category;
}

export enum Category {
    Nature,
    Architecture,
    Subjects,
}

export const photos: Array<Photo> = [
    {
        title: "Joy",
        file: "20240826_005728_D355D2E2.heic",
        date: new Date("2025-08-26"),
        category: Category.Subjects,
    },
    {
        title: "Utah sunset",
        file: "20220626_024923_7F2E9FDD.heic",
        date: new Date("2022-06-26"),
        category: Category.Nature,
    },

    {
        title: "Fall Trails",
        file: "Scan204487.jpg",
        date: new Date("10-18-2025"),
        category: Category.Nature,
    },

    {
        title: "Night sunset over dunes",
        file: "IMG_3964.jpg",
        date: new Date("12-26-2025"),
        category: Category.Nature,
    },
    {
        title: "Brimstone",
        file: "20250609_073624_53044B20.jpg",
        date: new Date("06-09-2025"),
        category: Category.Nature,
    },
    {
        title: "Blurry beach sunset",
        file: "20250724_214252_19B78C76.jpg",
        date: new Date("07-24-2025"),
        category: Category.Nature,
    },
    {
        title: "Sprout on wet rocks",
        file: "IMG_3765.jpg",
        date: new Date("09-27-2025"),
        category: Category.Nature,
    },
    {
        title: "99",
        file: "20250810_121338_91CB5229.jpg",
        date: new Date("08-10-2025"),
        category: Category.Architecture,
    },
    {
        title: "Backpackers in snow",
        file: "20250525_084144_EA1D8099.jpg",
        date: new Date("05-25-2025"),
        category: Category.Nature,
    },
    {
        title: "Crater",
        file: "20250609_073635_8E2F706C.jpg",
        date: new Date("06-09-2025"),
        category: Category.Nature,
    },

]