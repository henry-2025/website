import MarkdownIt from "markdown-it";
import Tag from "@markdoc/markdoc/src/tag";
import Prism from "prismjs";
import "prismjs/components/prism-python";
import type { RenderableTreeNodes } from "@markdoc/markdoc/src/types";
const { escapeHtml } = MarkdownIt().utils;

// HTML elements that do not have a matching close tag
// Defined in the HTML standard: https://html.spec.whatwg.org/#void-elements
const voidElements = new Set([
  "area",
  "base",
  "br",
  "col",
  "embed",
  "hr",
  "img",
  "input",
  "link",
  "meta",
  "param",
  "source",
  "track",
  "wbr",
]);

export default function render(node: RenderableTreeNodes): string {
  if (typeof node === "string" || typeof node === "number")
    return escapeHtml(String(node));

  if (Array.isArray(node)) return node.map(render).join("");

  if (node === null || typeof node !== "object" || !Tag.isTag(node)) return "";

  const { name, attributes, children = [] } = node;

  if (!name) return render(children);

  let output = `<${name}`;
  for (const [k, v] of Object.entries(attributes ?? {}))
    output += ` ${k.toLowerCase()}="${escapeHtml(String(v))}"`;
  output += ">";

  if (voidElements.has(name)) return output;

  // render code with prisjs
  if (name === "pre") {
    var autoloader = Prism.plugins.autoloader;
    let renderer = Prism.languages.python;
    if (renderer != null) {
      if (children.length != 1) {
        console.error(
          `expected children with length 1, but got length ${children.length}`,
        );
      } else {
        if (typeof children[0] != "string") {
          throw TypeError("expected string for arg 0 of children");
        }
        output += Prism.highlight(
          children[0],
          renderer,
          attributes["data-language"],
        );
      }
    } else {
      console.error(
        `no renderer found for language ${attributes["data-language"]}`,
      );
    }
  } else if (children.length) output += render(children);
  output += `</${name}>`;

  return output;
}
