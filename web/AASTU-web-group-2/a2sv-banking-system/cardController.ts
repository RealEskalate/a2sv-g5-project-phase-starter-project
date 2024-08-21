import axios from 'axios';

const API_URL = 'https://bank-dashboard-6acc.onrender.com'; 

export const getAllCards = async () => {
  try {
    const response = await axios.get(`${API_URL}/cards`);
    return response.data;  
  } catch (error) {
    console.error('Error fetching cards:', error);
    throw new Error('Unable to fetch cards');
  }
};

export const createCard = async (cardData: any) => {
  try {
    const response = await axios.post(`${API_URL}/cards`, cardData);
    return response.data;  
  } catch (error) {
    console.error('Error creating card:', error);
    throw new Error('Unable to create card');
  }
};

export const getCardById = async (id: string) => {
  try {
    const response = await axios.get(`${API_URL}/cards/${id}`);
    return response.data;  
  } catch (error) {
    console.error(`Error fetching card with ID ${id}:`, error);
    throw new Error('Unable to fetch card');
  }
};

export const deleteCard = async (id: string) => {
  try {
    const response = await axios.delete(`${API_URL}/cards/${id}`);
    return response.data;  
  } catch (error) {
    console.error(`Error deleting card with ID ${id}:`, error);
    throw new Error('Unable to delete card');
  }
};
