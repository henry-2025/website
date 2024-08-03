import { read } from "$app/server";
import Markdoc from "@markdoc/markdoc";
import render from "./renderer";

const articles = import.meta.glob("../pages/*.md", {
  query: "?url",
  eager: true,
});
const pathFromArticleName = new Map(
  Object.entries(articles).map((e) => {
    let file = e[0].replace(/^.*[\\/]/, "");
    file = file.substr(0, file.length - 3);
    return [file, e[1].default];
  }),
);

/** @type {import('./$types').PageServerLoad} */
export async function load({ params }) {
  const file = pathFromArticleName.get(params.slug);
  const postSource = await read(file).text();
  const ast = Markdoc.parse(postSource);
  const content = Markdoc.transform(ast /* config */);
  const html = render(content);
  return {
    content: html,
  };
}
