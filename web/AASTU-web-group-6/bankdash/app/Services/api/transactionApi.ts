import axios from "axios";

const API_URL = "https://bank-dashboard-6acc.onrender.com/transactions"; // Adjust this to match your actual API base URL
const extension = "/balance-history";
const API_URL_expense =
  "https://bank-dashboard-6acc.onrender.com/transactions/expenses";
const API_URL_income =
  "https://bank-dashboard-6acc.onrender.com/transactions/incomes";
interface TransactionType {
  transactionId: string;
  type: string;
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string | null;
}

const handleRequest = async (
  method: string,
  endpoint: string,
  data?: TransactionType[],
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
    return response.data.data;
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
  public static getTransactions(
    accessToken?: string
  ): Promise<TransactionType[]> {
    return handleRequest("GET", `${API_URL}`, undefined, accessToken);
  }

  public static balanceHistory(
    accessToken?: string
  ): Promise<TransactionType[]> {
    const extension = "/random-balance-history?monthsBeforeFirstTransaction=12";
    return handleRequest(
      "GET",
      `${API_URL}${extension}`,
      undefined,
      accessToken
    );
  }
  public static getExpenseData(
    accessToken?: string
  ): Promise<TransactionType[]> {
    return handleRequest(
      "GET",
      `${API_URL_expense}?page=0&size=1000`,
      undefined,
      accessToken
    );
  }
  public static getIncomeData(
    accessToken?: string
  ): Promise<TransactionType[]> {
    return handleRequest(
      "GET",
      `${API_URL_income}?page=0&size=10`,
      undefined,
      accessToken
    );
  }
}

export default TransactionService;
