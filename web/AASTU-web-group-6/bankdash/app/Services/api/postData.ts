import axios from "axios";

const postData = async (
  endpoint: string,
  data: {
    balance: number;
    cardHolder: string;
    expiryDate: string;
    passcode: string;
    cardType: string;
  },
  accessToken: string
) => {
  try {
    const response = await axios.post(endpoint, data, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
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

export default postData;
