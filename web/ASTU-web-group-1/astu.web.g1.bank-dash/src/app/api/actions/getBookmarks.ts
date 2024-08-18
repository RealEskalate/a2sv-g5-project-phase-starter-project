import apiClient from './apiClient';

async function getBookmarks() {
  try {
    const data = await apiClient.get('/bookmarks');
    return data.data;
  } catch (error) {
    throw new Error('error');
  }
}
