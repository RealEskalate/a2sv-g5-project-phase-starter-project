import apiClient from './apiClient';

export async function getAllPosts() {
  try {
    const post = await apiClient.get('/opportunities/search');
    return post.data;
  } catch (error) {
    throw new Error('error');
  }
}
