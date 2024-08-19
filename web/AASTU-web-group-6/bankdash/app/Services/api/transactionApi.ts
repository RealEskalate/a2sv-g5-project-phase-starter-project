import axios from "axios";

const API_URL = "https://bank-dashboard-6acc.onrender.com/transactions"; // Adjust this to match your actual API base URL

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
    public static getTransactions(accessToken?: string): Promise<TransactionType[]> {
        return handleRequest("GET", `${API_URL}`, undefined, accessToken);
    }

    public static balanceHistory(accessToken?: string): Promise<TransactionType[]> {
        const extension = '/balance-history'
        return handleRequest("GET", `${API_URL}${extension}`, undefined, accessToken);
    }
}

export default TransactionService;
