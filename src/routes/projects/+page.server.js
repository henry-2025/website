import { read } from "$app/server";
import Markdoc from "@markdoc/markdoc";
import yaml from "js-yaml";

const articles = Object.values(
  import.meta.glob("./pages/*.md", {
    query: "?url",
    eager: true,
  }),
).map((x) => x.default);

/** @type {import('./$types').PageServerLoad} */
export async function load({ params }) {
  /** @type Array<Object> */
  let data = [];
  for (let i = 0; i < articles.length; i++) {
    const a = articles[i];
    const postSource = await read(a).text();
    const ast = Markdoc.parse(postSource);
    const frontMatter = yaml.load(ast.attributes.frontmatter);
    const content = Markdoc.transform(ast /* config */);
    console.log(a.match(`/\/.*.md/g`));
    data.push({
      title: frontMatter.title,
      description: frontMatter.description,
      date: frontMatter.date,
      path: a.match("/.*.md"),
    });
  }
  return { articles: data };
}
