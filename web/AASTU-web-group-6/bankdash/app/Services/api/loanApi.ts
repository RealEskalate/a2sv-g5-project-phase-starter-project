import axios from "axios";
import { LoanType, LoanDetail, ApiResponse } from "@/types/LoanValue";

const API_URL = "https://bank-dashboard-rsf1.onrender.com/active-loans"; // Adjust this to match your actual API base URL

const handleRequest = async (
  method: string,
  endpoint: string,
  data?: LoanType[],
  accessToken?: string,
  returnContentOnly?: boolean
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
    return returnContentOnly ? response.data.data.content : response.data.data;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.error("Axios error:", error.message);
    } else {
      console.error("Unexpected error:", error);
    }
    throw error;
  }
};

class TransactionService {
  public static getLoan(
    accessToken: string,
    page: number
  ): Promise<LoanType[]> {
    return handleRequest(
      "GET",
      `${API_URL}/my-loans?page=${page}&size=5`,
      undefined,
      accessToken,
      true
    );
  }
  public static detailData(accessToken: string): Promise<ApiResponse> {
    return handleRequest(
      "GET",
      `${API_URL}/detail-data`,
      undefined,
      accessToken,
      false
    );
  }
}

export default TransactionService;
