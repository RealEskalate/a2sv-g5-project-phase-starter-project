import axios from 'axios';
import { authOptions } from '../auth/[...nextauth]/options';
import { getServerSession } from 'next-auth';

const apiClient = axios.create({
  baseURL: 'https://akil-backend.onrender.com',
  headers: {
    'Content-Type': 'application/json',
  },
});

export const setAuthHeader = async () => {
  apiClient.interceptors.request.use(async (config) => {
    const session = await getServerSession(authOptions);
    if (session?.accessToken) {
      config.headers.Authorization = `Bearer ${session.accessToken}`;
      return config;
    }
    return config;
  });
};

setAuthHeader();
export default apiClient;
