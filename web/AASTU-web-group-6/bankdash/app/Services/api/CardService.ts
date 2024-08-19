// src/services/cardService.ts
import axios from "axios";

const API_URL = "https://bank-dashboard-6acc.onrender.com/cards"; // Adjust this to match your actual API base URL

interface Card {
  id?: string;
  balance?: number;
  cardHolder?: string;
  expiryDate?: string;
  passcode?: string;
  cardType?: string;
  cardNumber?: string;
  userId?: string;
}

const handleRequest = async (
  method: string,
  endpoint: string,
  data?: Card,
  accessToken?: string
) => {
  try {
    const response = await axios({
      method,
      url: endpoint,
      data,
      headers: {
        Authorization: accessToken ? `Bearer ${accessToken}` : undefined,
        "Content-Type": "application/json",
      },
    });
    return response.data;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.error("Axios error:", error.message);
    } else {
      console.error("Unexpected error:", error);
    }
    throw error;
  }
};

class CardService {
  public static getAllCards(accessToken?: string): Promise<Card[]> {
    return handleRequest("GET", API_URL, undefined, accessToken);
  }

  public static addCard(card: Card, accessToken: string): Promise<Card> {
    return handleRequest("POST", API_URL, card, accessToken);
  }

  public static getCardById(id: string, accessToken?: string): Promise<Card> {
    return handleRequest("GET", `${API_URL}/${id}`, undefined, accessToken);
  }

  public static deleteCardById(
    id: string,
    accessToken?: string
  ): Promise<void> {
    return handleRequest("DELETE", `${API_URL}/${id}`, undefined, accessToken);
  }
}

export default CardService;
