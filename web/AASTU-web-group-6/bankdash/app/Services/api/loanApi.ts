import axios from "axios";
import { LoanType, ApiResponse } from '@/types/LoanValue';

const API_URL = "https://bank-dashboard-6acc.onrender.com/active-loans/my-loans"; // Adjust this to match your actual API base URL

const handleRequest = async (
    method: string,
    endpoint: string,
    data?: LoanType[],
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
    public static getLoan(accessToken: string): Promise<LoanType[]> {
        return handleRequest("GET", `${API_URL}`, undefined, accessToken);
    }
}

export default TransactionService;
