// src/services/cardService.ts
import axios from "axios";
import { checkAndRefreshToken } from "./hooks/useRefresh";

const API_URL = "https://bank-dashboard-rsf1.onrender.com/cards"; // Adjust this to match your actual API base URL

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

interface DecodedToken {
  exp: number; 
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
    console.log("Response Status",response.status);
    return response.data;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      if (error.response.status === 401)
      console.error("Axios error:", error.message);
    } else {
      console.error("Unexpected error:", error);
    }
    throw error;
  }
};


class CardService {

  private static async ensureAccessToken(accessToken?: string): Promise<string> {
    console.log('checking')
    const decodedToken = jwtDecode<DecodedToken>(accessToken);
    console.log("DECODED TOKEN",decodedToken)
    if (accessToken) {
      return accessToken;
    }
    return await checkAndRefreshToken() as string;
  }

  public static async getAllCards(accessToken?: string): Promise<Card[]> {
    const token = await this.ensureAccessToken(accessToken);
    return handleRequest(
      "GET",
      `${API_URL}?page=0&size=3`,
      undefined,
      token
    );
  }

  public static async addCard(card: Card, accessToken?: string): Promise<Card> {
    const token = await this.ensureAccessToken(accessToken);
    return handleRequest("POST", API_URL, card, token);
  }

  public static async getCardById(id: string, accessToken?: string): Promise<Card> {
    const token = await this.ensureAccessToken(accessToken);
    return handleRequest("GET", `${API_URL}/${id}`, undefined, token);
  }

  public static async deleteCardById(id: string, accessToken?: string): Promise<void> {
    const token = await this.ensureAccessToken(accessToken);
    return handleRequest("DELETE", `${API_URL}/${id}`, undefined, token);
  }
}

export default CardService;
function jwtDecode<T>(accessToken: string | undefined) {
  throw new Error("Function not implemented.");
}

