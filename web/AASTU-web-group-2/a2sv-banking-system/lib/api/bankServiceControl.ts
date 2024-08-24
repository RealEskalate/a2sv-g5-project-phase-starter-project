import axios from 'axios';
import { BankService, BankServiceData, ApiResponse } from '@/types/bankServiceController.interface'; // Importing interfaces

const BASE_URL = 'https://bank-dashboard-mih0.onrender.com';


export const getBankServices = async (token: string, page: number, size: number ): Promise<ApiResponse> => { 
  try {
  console.log("Fetching")
  const response = await fetch(`${BASE_URL}/bank-services?page=${page}&size=${size}`, {
    headers: {
      'Authorization': `Bearer ${token}`, // Add the token to the headers
    },
  });
  if (response.status === 200) {
    const data: ApiResponse = await response.json();
    return data ;
  } else {
    throw new Error(`Request failed with status code: ${response.status}`);
  }
} catch (error) {
  console.error('Error fetching cards:', error);
  throw error;
}
};

// Search bank services by query
export const searchBankServices = async (token: string, query: string ): Promise<BankServiceData[]> => {
  try {
    const response = await fetch(`${BASE_URL}/bank-services/search?query=${query}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (response.status === 200) {
      const data: BankServiceData[] = await response.json();
      return data;
    } else {
      throw new Error(`Request failed with status code: ${response.status}`);
    }
  } catch (error) {
    console.error('Error searching bank services:', error);
    throw error;
  }
};


// Get a specific bank service by ID
export const getBankServiceById = async ( token: string, id: string): Promise<BankServiceData> => {
  try {
    const response = await fetch(`${BASE_URL}/bank-services/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (response.status === 200) {
      const data: BankServiceData = await response.json();
      return data;
    } else {
      throw new Error(`Request failed with status code: ${response.status}`);
    }
  } catch (error) {
    console.error(`Error fetching bank service with ID ${id}:`, error);
    throw error;
  }
};


// Create a new bank service
export const createBankService = async (
  token: string,
  data: Omit<BankServiceData, 'id'>,
  
): Promise<BankServiceData> => {
  try {
    const response = await fetch(`${BASE_URL}/bank-services`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    if (response.status === 200) {
      const responseData: BankServiceData = await response.json();
      return responseData;
    } else {
      throw new Error(`Request failed with status code: ${response.status}`);
    }
  } catch (error) {
    console.error('Error creating bank service:', error);
    throw error;
  }
};


// Update an existing bank service by ID
export const updateBankService = async (
  token: string,
  id: string,
  data: Partial<Omit<BankServiceData, 'id'>>,
  
): Promise<BankServiceData> => {
  try {
    const response = await fetch(`${BASE_URL}/bank-services/${id}`, {
      method: 'PUT',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    if (response.status === 200) {
      const responseData: BankServiceData = await response.json();
      return responseData;
    } else {
      throw new Error(`Request failed with status code: ${response.status}`);
    }
  } catch (error) {
    console.error(`Error updating bank service with ID ${id}:`, error);
    throw error;
  }
};


// Delete a bank service by ID
export const deleteBankService = async (token: string, id: string ): Promise<void> => {
  try {
    const response = await fetch(`${BASE_URL}/bank-services/${id}`, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });


    if (response.status === 200) {
      console.log(`Bank service with ID ${id} deleted successfully.`);
    } else {
      throw new Error(`Request failed with status code: ${response.status}`);
    }
  } catch (error) {
    console.error(`Error deleting bank service with ID ${id}:`, error);
    throw error;
  }
};
