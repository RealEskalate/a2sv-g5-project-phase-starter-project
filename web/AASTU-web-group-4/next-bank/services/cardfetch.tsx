import Cookies from "js-cookie";

// Get All Cards - GET Request
export const getAllCards = async (token: string) => {
  try {
    const response = await fetch(
      "https://next-bank.onrender.com/cards?page=0&size=10",
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    if (!response.ok) {
      console.log(response);
      throw new Error("Failed to fetch cards");
    }

    const data = await response.json();
    return data.content;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Get Card by ID - GET Request
export const getCardById = async (id: string, token: string) => {
  try {
    const response = await fetch(`https://next-bank.onrender.com/cards/${id}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      throw new Error(`Failed to fetch card with ID: ${id}`);
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Create a New Card - POST Request
export const createCard = async (cardData: any, token: string) => {
  try {
    const response = await fetch("https://next-bank.onrender.com/cards", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(cardData),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || "Failed to create a new card");
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Delete Card by ID - DELETE Request
export const deleteCardById = async (id: string, token: string) => {
  try {
    const response = await fetch(`https://next-bank.onrender.com/cards/${id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      throw new Error(`Failed to delete card with ID: ${id}`);
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error:", error);
    throw error;
  }
};

// Export all functions from this file
