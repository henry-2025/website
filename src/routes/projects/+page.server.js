import { read } from "$app/server";
import Markdoc from "@markdoc/markdoc";
import yaml from "js-yaml";

const articlePaths = Object.entries(
  import.meta.glob("./pages/*.md", {
    query: "?url",
    eager: true,
  }),
).map((entry) => {
  let file = entry[0].replace(/^.*[\\/]/, "");
  file = file.substr(0, file.length - 3);
  return { file: file, path: entry[1].default };
});

/** @type {import('./$types').PageServerLoad} */
export async function load({ params }) {
  /** @type Array<Object> */
  let data = [];
  for (let i = 0; i < articlePaths.length; i++) {
    const postSource = await read(articlePaths[i].path).text();
    const ast = Markdoc.parse(postSource);
    const frontMatter = yaml.load(ast.attributes.frontmatter);
    //const content = Markdoc.transform(ast /* config */);
    //const html = Markdoc.renderers.html(content);
    data.push({
      title: frontMatter.title,
      description: frontMatter.description,
      date: frontMatter.date,
      path: articlePaths[i].file,
    });
  }
  return { articles: data };
}
