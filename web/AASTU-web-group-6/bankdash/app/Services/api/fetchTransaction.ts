import axios from "axios";

const getTransaction = async (page: number, accessToken: string) => {
  try {
    const response = await axios({
      method: "GET",
      url: `https://bank-dashboard-irse.onrender.com/transactions?page=${page}&size=5`,
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
      },
    });
    return response.data.data.content;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.error("Axios error:", error.message);
    } else {
      console.error("Unexpected error:", error);
    }
    throw error;
  }
};

const getExpense = async (page: number, accessToken: string) => {
  try {
    const response = await axios({
      method: "GET",
      url: `https://bank-dashboard-irse.onrender.com/transactions/expenses?page=${page}&size=5`,
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
      },
    });
    return response.data.data.content;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.error("Axios error:", error.message);
    } else {
      console.error("Unexpected error:", error);
    }
    throw error;
  }
};

const getIncome = async (page: number, accessToken: string) => {
  try {
    const response = await axios({
      method: "GET",
      url: `https://bank-dashboard-irse.onrender.com/transactions/incomes?page=${page}&size=5`,
      headers: {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
      },
    });
    return response.data.data.content;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.error("Axios error:", error.message);
    } else {
      console.error("Unexpected error:", error);
    }
    throw error;
  }
};

export { getExpense, getTransaction, getIncome };
