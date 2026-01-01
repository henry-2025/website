export interface Photo {
    title?: string;
    file: string;
    date: Date;
    caption: string;
    category: Category;
}

export enum Category {
    Nature,
    Architecture,
    Subjects,
}

export const photos: Array<Photo> = [
    {
        file: "20240826_005728_D355D2E2.heic",
        date: new Date("2025-08-26"),
        caption: "dog",
        category: Category.Subjects,
    },
    {
        file: "20220626_031848_81127359.heic",
        date: new Date("2022-06-26"),
        caption: "Bonneville sunset",
        category: Category.Nature,
    },

      {
        file: "Scan204487.jpg",
        date: new Date("10-18-2025"),
        caption: "dog on fall trail",
        category: Category.Nature,
    },

    {
        file: "IMG_3964.jpg",
        date: new Date("12-26-2025"),
        caption: "light pollution against dunes",
        category: Category.Nature,
    },
    {
        file: "20250609_073624_53044B20.jpg",
        date: new Date("06-09-2025"),
        caption: "Mt. Baker",
        category: Category.Nature,
    },
    {
        file: "20250724_214252_19B78C76.jpg",
        date: new Date("07-24-2025"),
        caption: "blurry beach sunset",
        category: Category.Nature,
    },
    {
        file: "IMG_3765.jpg",
        date: new Date("09-27-2025"),
        caption: "rock sprout",
        category: Category.Nature,
    },
    {
        file: "20250810_121338_91CB5229.jpg",
        date: new Date("08-10-2025"),
        caption: "underside of 99 bridge, Seattle",
        category: Category.Architecture,
    },
    {
        file: "20250525_084144_EA1D8099.jpg",
        date: new Date("05-25-2025"),
        caption: "backpackers in the snow",
        category: Category.Nature,
    },
    {
        file: "20250609_073635_8E2F706C.jpg",
        date: new Date("06-09-2025"),
        caption: "Mt. Baker volcanic crater",
        category: Category.Nature,
    },

]