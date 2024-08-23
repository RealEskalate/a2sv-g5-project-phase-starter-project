import axios from 'axios';
import { BankServiceData, ApiResponse } from '@/types/bankServiceController.interface'; // Importing interfaces

const BASE_URL = 'https://bank-dashboard-o9tl.onrender.com';

// Fetch paginated bank services
export const getBankServices = (page: number, size: number, token: string): Promise<ApiResponse<BankServiceData[]>> => {
  return axios
    .get<ApiResponse<BankServiceData[]>>(`${BASE_URL}/bank-services`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
      params: { page, size },
    })
    .then((response) => response.data)
    .catch((error) => {
      console.error('Error fetching bank services', error);
      throw error;
    });
};

// Search bank services by query
export const searchBankServices = (query: string, token: string): Promise<ApiResponse<BankServiceData[]>> => {
  return axios
    .get<ApiResponse<BankServiceData[]>>(`${BASE_URL}/bank-services/search`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
      params: { query },
    })
    .then((response) => response.data)
    .catch((error) => {
      console.error('Error searching bank services', error);
      throw error;
    });
};

// Get a specific bank service by ID
export const getBankServiceById = (id: string, token: string): Promise<ApiResponse<BankServiceData>> => {
  return axios
    .get<ApiResponse<BankServiceData>>(`${BASE_URL}/bank-services/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
    .then((response) => response.data)
    .catch((error) => {
      console.error(`Error fetching bank service with ID ${id}`, error);
      throw error;
    });
};

// Create a new bank service
export const createBankService = (data: Omit<BankServiceData, 'id'>, token: string): Promise<ApiResponse<BankServiceData>> => {
  return axios
    .post<ApiResponse<BankServiceData>>(`${BASE_URL}/bank-services`, data, {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    })
    .then((response) => response.data)
    .catch((error) => {
      console.error('Error creating bank service', error);
      throw error;
    });
};

// Update an existing bank service by ID
export const updateBankService = (
  id: string,
  data: Partial<Omit<BankServiceData, 'id'>>,
  token: string
): Promise<ApiResponse<BankServiceData>> => {
  return axios
    .put<ApiResponse<BankServiceData>>(`${BASE_URL}/bank-services/${id}`, data, {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    })
    .then((response) => response.data)
    .catch((error) => {
      console.error(`Error updating bank service with ID ${id}`, error);
      throw error;
    });
};

// Delete a bank service by ID
export const deleteBankService = (id: string, token: string): Promise<ApiResponse<BankServiceData>> => {
  return axios
    .delete<ApiResponse<BankServiceData>>(`${BASE_URL}/bank-services/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
    .then((response) => response.data)
    .catch((error) => {
      console.error(`Error deleting bank service with ID ${id}`, error);
      throw error;
    });
};
