import axios from "axios";
const API_URL =
  "https://bank-dashboard-irse.onrender.com/user/random-investment-data";
interface YearlyInvestment {
  time: string;
  value: number;
}

interface MonthlyRevenue {
  time: string;
  value: number;
}

interface InvestmentData {
  totalInvestment: number;
  rateOfReturn: number;
  yearlyTotalInvestment: YearlyInvestment[];
  monthlyRevenue: MonthlyRevenue[];
}
const handleRequest = async (
  method: string,
  endpoint: string,
  data?: InvestmentData,
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
class InvestmentService {
  public static getInvestmentData(
    accessToken?: string
  ): Promise<InvestmentData> {
    return handleRequest(
      "GET",
      `${API_URL}?years=5&months=6`,
      undefined,
      accessToken
    );
  }
}

export default InvestmentService;
