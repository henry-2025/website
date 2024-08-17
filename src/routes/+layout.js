/** @type {import('./$types').LayoutLoad} */
export const load = ({ url }) => {
  const { pathname } = url;

  return {
    pathname,
  };
};
