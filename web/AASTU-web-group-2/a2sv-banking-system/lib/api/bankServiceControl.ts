import axios from 'axios';
const BASE_URL = 'https://bank-dashboard-1tst.onrender.com';
interface BankServiceData {
  id: string;
  name: string;
  details: string;
  numberOfUsers: number;
  status: string;
  type: string;
  icon: string;
}

interface ApiResponse<T> {
  success: boolean;
  message: string;
  data: T;
}

export const getBankServices = async (page: number, size: number) => {
  try {
    const response = await axios.get<ApiResponse<BankServiceData[]>>(
      `${BASE_URL}/bank-services`,
      { params: { page, size } }
    );
    return response.data;
  } catch (error) {
    console.error('Error fetching bank services', error);
    throw error;
  }
};

export const searchBankServices = async (query: string) => {
  try {
    const response = await axios.get<ApiResponse<BankServiceData[]>>(
      `${BASE_URL}/bank-services/search`,
      { params: { query } }
    );
    return response.data;
  } catch (error) {
    console.error('Error searching bank services', error);
    throw error;
  }
};

export const getBankServiceById = async (id: string) => {
  try {
    const response = await axios.get<ApiResponse<BankServiceData>>(
      `${BASE_URL}/bank-services/${id}`
    );
    return response.data;
  } catch (error) {
    console.error(`Error fetching bank service with ID ${id}`, error);
    throw error;
  }
};

export const createBankService = async (data: Omit<BankServiceData, 'id'>) => {
  try {
    const response = await axios.post<ApiResponse<BankServiceData>>(
      `${BASE_URL}/bank-services`,
      data
    );
    return response.data;
  } catch (error) {
    console.error('Error creating bank service', error);
    throw error;
  }
};

export const updateBankService = async (id: string, data: Partial<Omit<BankServiceData, 'id'>>) => {
  try {
    const response = await axios.put<ApiResponse<BankServiceData>>(
      `${BASE_URL}/bank-services/${id}`,
      data
    );
    return response.data;
  } catch (error) {
    console.error(`Error updating bank service with ID ${id}`, error);
    throw error;
  }
};

export const deleteBankService = async (id: string) => {
  try {
    const response = await axios.delete<ApiResponse<BankServiceData>>(
      `${BASE_URL}/bank-services/${id}`
    );
    return response.data;
  } catch (error) {
    console.error(`Error deleting bank service with ID ${id}`, error);
    throw error;
  }
};
