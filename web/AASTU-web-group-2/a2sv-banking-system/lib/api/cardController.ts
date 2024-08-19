import { Card, GetCardsResponse, PostCardRequest, PostCardResponse, GetCardByIdResponse } from "@/types/cardController.Interface"


const BASE_URL = 'https://bank-dashboard-6acc.onrender.com';

const getCards = async (): Promise<GetCardsResponse> => {
    try {
      const response = await fetch(`${BASE_URL}/cards`);
  
      if (response.status === 200) {
        const data: Card[] = await response.json();
        return {cards: data};
      } else {
        throw new Error(`Request failed with status code: ${response.status}`);
      }
    } catch (error) {
      console.error('Error fetching cards:', error);
      throw error;
    }
  };

  const postCard = async (cardDetails: PostCardRequest): Promise<PostCardResponse> => {
    try {
      const response = await fetch(`${BASE_URL}/cards`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
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

  const getCardById = async (id: string): Promise<GetCardByIdResponse> => {
    try {
      const response = await fetch(`${BASE_URL}/cards/${id}`);
  
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

  const deleteCardById = async (id: string): Promise<void> => {
    try {
      const response = await fetch(`${BASE_URL}/cards/${id}`, {
        method: 'DELETE',
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
