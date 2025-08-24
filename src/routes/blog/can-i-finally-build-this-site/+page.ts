import { posts } from '$lib/blog/posts';
import type { Post } from '$lib/blog/posts';

export const load = () => {
    return { post: posts.find((post: Post) => post.slug === 'can-i-finally-build-this-site') };
}