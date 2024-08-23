const BASE_URL = 'https://web-team-g4.onrender.com/cards';
const HEADERS = {
  "Content-Type": "application/json",
};

// Get All Cards - GET Request
export const getAllCards = async () => {
  try {
    const response = await fetch(" https://web-team-g4.onrender.com/cards", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error("Failed to fetch cards");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Get Card by ID - GET Request
export const getCardById = async (id: string) => {
  try {
    const response = await fetch(
      ` https://web-team-g4.onrender.com/cards/${id}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      });
  
      if (!response.ok) {
        throw new Error(`Failed to fetch card with ID: ${id}`);
      }
  
      const data = await response.json();
      return data;
    } catch (error) {
      console.error('Error:', error);
      throw error;
    }

};

// Create a New Card - POST Request
export const createCard = async (cardData: any) => {
  try {
    const response = await fetch(" https://web-team-g4.onrender.com/cards", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(cardData),
    });

    if (!response.ok) {
      throw new Error("Failed to create a new card");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Delete Card by ID - DELETE Request
export const deleteCardById = async (id: string) => {
  try {
    const response = await fetch(
      ` https://web-team-g4.onrender.com/cards/${id}`,
      {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    if (!response.ok) {
      throw new Error(`Failed to delete card with ID: ${id}`);
    }

    return await response.json();
  } catch (error) {
    console.error(`Error deleting card with ID ${id}:`, error);
    throw error;
  }
};
