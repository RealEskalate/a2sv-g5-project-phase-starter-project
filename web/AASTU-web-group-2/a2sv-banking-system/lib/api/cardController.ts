import {
  GetCardsResponse,
  PostCardRequest,
  PostCardResponse,
  GetCardByIdResponse,
} from "@/types/cardController.Interface";
const BASE_URL = 'https://a2svwallets.onrender.com';

// Extend the user type to include accessToken

const getCards = async (token:string, page=0, size=1): Promise<GetCardsResponse> => {
  try {
    const response = await fetch(`${BASE_URL}/cards?page=${page}&size=${size}`, {
      headers: {
        'Authorization': `Bearer ${token}`, // Add the token to the headers
      },
    });
    if (response.status === 200) {
      const data: GetCardsResponse = await response.json();
      return data ;
    } else {
      throw new Error(`Request failed with status code: ${response.status}`);
    }
  } catch (error) {
    console.error('Error fetching cards:', error);
    throw error;
  }
};

const postCard = async (cardDetails: PostCardRequest, token:string): Promise<PostCardResponse> => {
  try {
    const response = await fetch(`${BASE_URL}/cards`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`, // Add the token to the headers
      },
      body: JSON.stringify(cardDetails),
    });

    if (response.status === 200) {
      const data: PostCardResponse = await response.json();
      return data;
    } else {
      throw new Error(`Request failed with status code: ${response.status}`);
    }
  } catch (error) {
    console.error('Error posting card:', error);
    throw error;
  }
};

const getCardById = async (id: string, token:string): Promise<GetCardByIdResponse> => {
  try {
    const response = await fetch(`${BASE_URL}/cards/${id}`, {
      headers: {
        'Authorization': `Bearer ${token}`, // Add the token to the headers
      },
    });

    if (response.status === 200) {
      const data: GetCardByIdResponse = await response.json();
      return data;
    } else {
      throw new Error(`Request failed with status code: ${response.status}`);
    }
  } catch (error) {
    console.error(`Error fetching card with ID ${id}:`, error);
    throw error;
  }
};

const deleteCardById = async (id: string, token:string): Promise<void> => {
  try {
    const response = await fetch(`${BASE_URL}/cards/${id}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${token}`, // Add the token to the headers
      },
    });

    if (response.status === 200) {
      console.log(`Card with ID ${id} deleted successfully.`);
    } else {
      throw new Error(`Failed to delete card with ID ${id}. Status code: ${response.status}`);
    }
  } catch (error) {
    console.error(`Error deleting card with ID ${id}:`, error);
    throw error;
  }
};

// Named exports
export { getCards, postCard, getCardById, deleteCardById };