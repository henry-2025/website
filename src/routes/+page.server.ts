import { read } from "$app/server";
import Markdoc from "@markdoc/markdoc";
import yaml from "js-yaml";
import type { PageServerLoad } from "./$types";

interface FrontMatter {
  title: string;
  date: Date;
  description: string;
}

interface FrontMatter {
  title: string;
  description: string;
  date: Date;
}

interface File {
  default: string;
}

interface FilePath {
  file: string;
  path: string;
}
const articleFiles: [string, File][] = Object.entries(
  import.meta.glob("../../static/pages/*.md", {
    query: "?url",
    eager: true,
  }),
);

const articlePaths: FilePath[] = articleFiles.map((entry) => {
  let file = entry[0].replace(/^.*[\\/]/, "");
  file = file.substring(0, file.length - 3);
  return { file: file, path: entry[1].default };
});

export const load: PageServerLoad = async ({ params }) => {
  let data = [];
  for (let i = 0; i < articlePaths.length; i++) {
    const postSource = await read(articlePaths[i].path).text();
    const ast = Markdoc.parse(postSource);

    const frontMatter = yaml.load(ast.attributes.frontmatter) as FrontMatter;
    const date = frontMatter.date;
    data.push({
      title: frontMatter.title,
      description: frontMatter.description,
      date: date,
      path: articlePaths[i].file,
    });
  }
  data.sort((a, b) => {
    return b.date.getTime() - a.date.getTime();
  });
  return { articles: data };
};
