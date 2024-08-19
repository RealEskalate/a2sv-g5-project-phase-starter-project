import axios from "axios";

const deleteData = async (endpoint: string, accessToken: string, data?: any) => {
  try {
    const response = await axios.delete(endpoint, {
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
      },
      data: data, // Optional data to send with the DELETE request
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

export default deleteData;
