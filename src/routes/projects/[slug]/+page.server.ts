import { read } from "$app/server";
import Markdoc from "@markdoc/markdoc";
import render from "./renderer";
import type { PageServerLoad } from "./$types";

interface FilePath {
  default: string;
}

const articles: Record<string, FilePath> = import.meta.glob(
  "../../../../static/pages/*.md",
  {
    query: "?url",
    eager: true,
  },
);

const pathFromArticleName = new Map(
  Object.entries(articles).map((e) => {
    let file = e[0].replace(/^.*[\\/]/, "");
    file = file.substring(0, file.length - 3);
    return [file, e[1].default];
  }),
);

export const load: PageServerLoad = async ({ params }) => {
  const file = pathFromArticleName.get(params.slug);
  if (typeof file != "string") {
    throw TypeError("expected type string for file");
  }
  const postSource = await read(file).text();
  const ast = Markdoc.parse(postSource);
  const content = Markdoc.transform(ast /* config */);
  const html = render(content);
  return {
    content: html,
  };
};
