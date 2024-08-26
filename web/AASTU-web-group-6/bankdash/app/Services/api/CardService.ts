import axios from "axios";
import { jwtDecode } from "jwt-decode"; // Assuming you're using the jwt-decode library
import { checkAndRefreshToken } from "./hooks/useRefresh";

const API_URL = "https://bank-dashboard-irse.onrender.com/cards";

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
  private static async ensureAccessToken(accessToken?: string): Promise<string> {
    if (accessToken) {
      try {
        const decodedToken = jwtDecode<any>(accessToken);

        const currentTime = Date.now() / 1000;

        if (decodedToken.exp > currentTime) {
          return accessToken;
        }
      } catch (error) {
        console.error("Failed to decode token:", error);
      }
    }

    return await checkAndRefreshToken() as string;
  }

  public static async getAllCards(accessToken?: string): Promise<Card[]> {
    const token = await this.ensureAccessToken(accessToken);
    return handleRequest("GET", `${API_URL}?page=0&size=3`, undefined, token);
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
