import { read } from "$app/server";
import Markdoc from "@markdoc/markdoc";
import render from "./renderer";

/** @type {import('./$types').PageServerLoad} */
export async function load({ params }) {
  const postSource = await read(`../pages/${params.slug}.md`).text();
  const ast = Markdoc.parse(postSource);
  const content = Markdoc.transform(ast /* config */);
  const html = render(content);
  return {
    content: html,
  };
}
