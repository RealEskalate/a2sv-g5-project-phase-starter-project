import axios from 'axios';
const API_URL = "https://bank-dashboard-6acc.onrender.com/user/random-investment-data";
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
    public static getInvestmentData(accessToken?: string): Promise<InvestmentData> {
      return handleRequest("GET", `${API_URL}?years=5&months=6`, undefined, accessToken);
    }
  
  }
  
  export default InvestmentService;
// const investmentApi = async () => {
//     const accessToken = "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJtaWhyZXQiLCJpYXQiOjE3MjM4MDc1MTMsImV4cCI6MTcyMzg5MzkxM30.1_-JRWFuZeesCPQRYnCrpwBTA-2tjJL7yx4H1HCM5Wc0pEDHe6hXlTMh3ivxx9Db";
//     try {
//         const response = await axios.get('https://bank-dashboard-6acc.onrender.com/user/random-investment-data', {
//             params:{
//                 years: 6,
//                 months: 12
//             },
//             headers: {
//                 Authorization: `Bearer ${accessToken}`, 
//             },
//         });
//         console.log("invest" , response.data.data)
//         return response.data.data;
//     } 
//     catch (error) {
//         if (axios.isAxiosError(error)) {
//             console.error('Axios error:', error.message);
//         } else {
//             console.error('Unexpected error:', error);
//         }
//         throw error;
//     }
// };

// export default investmentApi;
