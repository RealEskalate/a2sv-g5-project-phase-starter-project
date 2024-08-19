import axios from "axios";
import UserValue from "@/types/UserValue";
import UserPreferenceValue from "@/types/UserPreferenceValue";

const API_URL = "https://bank-dashboard-6acc.onrender.com/user";

interface UserResponseValue{

    success: boolean,
    message: string,
    data: UserValue
}

interface UserPreferenceResponseValue{
    
    success: boolean,
    message: string,
    data: UserPreferenceValue
    
}

interface YearlyInvestmentValue {
    time: string;
    value: number;
  }
  
  interface MonthlyRevenueValue {
    time: string;
    value: number;
  }
  
  interface InvestmentDataValue {
    totalInvestment: number;
    rateOfReturn: number;
    yearlyTotalInvestment: YearlyInvestmentValue[];
    monthlyRevenue: MonthlyRevenueValue[];
  }
  
  interface InvestmentResponseValue {
    success: boolean;
    message: string;
    data: InvestmentDataValue;
  }
  

const handleRequest = async (
  method: string,
  endpoint: string,
  data?: any,
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

class UserService {
  public static update(formData: UserValue, accessToken: string): Promise<UserResponseValue> {
    return handleRequest("PUT", `${API_URL}/update`, formData, accessToken);
  }

  public static updatePreference(
    formData: UserPreferenceValue,
    accessToken: string
  ): Promise<UserPreferenceResponseValue> {
    return handleRequest("PUT", `${API_URL}/update-preference`, formData, accessToken);
  }

  public static searchUser(query: string, accessToken: string): Promise<any> {
    return handleRequest("GET", `${API_URL}/${query}`, undefined, accessToken);
  }

  public static randomInvestmentData(years:number,months:number,accessToken: string): Promise<InvestmentResponseValue> {
    return handleRequest("GET", `${API_URL}/random-investment-data?years=${years}&months=${months}`, undefined, accessToken);
  }

  public static current(accessToken: string): Promise<any> {
    return handleRequest("GET", `${API_URL}/current`, undefined, accessToken);
  }
}

export default UserService;
