import axios from 'axios';
const API_URL_expense = "https://bank-dashboard-6acc.onrender.com/transactions/expenses";
const API_URL_income = "https://bank-dashboard-6acc.onrender.com/transactions/incomes";
interface Transaction {
  transactionId: string;
  type: string;
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string;
}
const handleRequest = async (
  method: string,
  endpoint: string,
  data?: Transaction,
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
class LastTransService {
  public static getExpenseData(accessToken?: string): Promise<Transaction[]> {
    return handleRequest("GET", `${API_URL_expense}?page=0&size=3`, undefined, accessToken);
  }
  public static getIncomeData(accessToken?: string): Promise<Transaction[]> {
    return handleRequest("GET", `${API_URL_income}?page=0&size=3`, undefined, accessToken);
  }

}
export default LastTransService;




