const BASE_URL = 'https://web-team-g4.onrender.com/cards';
const HEADERS = {
  "Content-Type": "application/json",
};

// Get All Cards - GET Request
export const getAllCards = async () => {
  try {
    const response = await fetch(BASE_URL, {
      method: 'GET',
      headers: HEADERS,
    });

    if (!response.ok) {
      throw new Error('Failed to fetch cards');
    }

    return await response.json();
  } catch (error) {
    console.error('Error fetching all cards:', error);
    throw error;
  }
};

// Get Card by ID - GET Request
export const getCardById = async (id: string) => {
  try {
    const response = await fetch(`${BASE_URL}/${id}`, {
      method: 'GET',
      headers: HEADERS,
    });

    if (!response.ok) {
      throw new Error(`Failed to fetch card with ID: ${id}`);
    }

    return await response.json();
  } catch (error) {
    console.error(`Error fetching card with ID ${id}:`, error);
    throw error;
  }
};

// Create a New Card - POST Request
export const createCard = async (cardData: any) => {
  try {
    const response = await fetch(BASE_URL, {
      method: 'POST',
      headers: HEADERS,
      body: JSON.stringify(cardData),
    });

    if (!response.ok) {
      throw new Error('Failed to create a new card');
    }

    return await response.json();
  } catch (error) {
    console.error('Error creating a new card:', error);
    throw error;
  }
};

// Delete Card by ID - DELETE Request
export const deleteCardById = async (id: string) => {
  try {
    const response = await fetch(`${BASE_URL}/${id}`, {
      method: 'DELETE',
      headers: HEADERS,
    });

    if (!response.ok) {
      throw new Error(`Failed to delete card with ID: ${id}`);
    }

    return await response.json();
  } catch (error) {
    console.error(`Error deleting card with ID ${id}:`, error);
    throw error;
  }
};
