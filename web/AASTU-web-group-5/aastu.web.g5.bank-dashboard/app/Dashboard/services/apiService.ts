// apiService.ts
// to fetch data from "/transactions" endpoint
import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'https://bank-dashboard-6acc.onrender.com/api', // Base URL for the API
  headers: {
    'Content-Type': 'application/json',
  },
});

export const fetchRecentTransactions = async () => {
  try {
    const response = await apiClient.get('/transactions');
    return response.data;
  } catch (error) {
    console.error('Error fetching recent transactions:', error);
    throw error;
  }
};

