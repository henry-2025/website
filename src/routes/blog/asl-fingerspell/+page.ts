import { posts } from '$lib/blog/posts';
import type { Post } from '$lib/blog/posts';

export const load = () => {
    posts.find((post: Post) => post.slug === 'asl-fingerspell');
}