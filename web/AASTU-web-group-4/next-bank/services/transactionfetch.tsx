const API_BASE_URL = "https://bank-dashboard-1tst.onrender.com";
// GET /transactions
import Cookies from "js-cookie";

const token = "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJxd2VyIiwiaWF0IjoxNzI0MzE0NzM5LCJleHAiOjE3MjQ0MDExMzl9.B5Avsv1ZUX-DSf7PGwwIRNyAlKk_UAPsy2B9C-geCLtYSOMPOjDYeu9nRkHjT3z7";

export const getAllTransactionsss = async (token:string) => {
  try {
    const response = await fetch(`${API_BASE_URL}/transactions`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.json();
  } catch (error) {
    console.error("Error fetching all transactions:", error);
    throw error;
  }
};

// GET /transactions
export const getAllTransactions = async (page: any, size: any,tokens : string) => {
  console.log("hello:", token);
  const response = await fetch(
    `${API_BASE_URL}/transactions?page=${page}&size=${size}`,
    {
      method: "GET",
      headers: {
        Authorization: `Bearer ${tokens}`,
      },
    }
  );
  return response.json();
};

// POST /transactions
export const createTransaction = async (
  transactionData: any,
  accessToken: string
) => {
  try {
    const response = await fetch(`${API_BASE_URL}/transactions`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(transactionData),
    });
    console.log(response);
    return response.json();
  } catch (error) {
    console.error("Error creating transaction:", error);
    throw error;
  }
};

// POST /transactions/deposit
export const createDeposit = async (depositData: any, accessToken: string) => {
  try {
    const response = await fetch(`${API_BASE_URL}/transactions/deposit`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${accessToken}`,
      },
      body: JSON.stringify(depositData),
    });
    return response.json();
  } catch (error) {
    console.error("Error creating deposit:", error);
    throw error;
  }
};

// GET /transactions/{id}
export const getTransactionById = async (id: any, accessToken: string) => {
  try {
    const response = await fetch(`${API_BASE_URL}/transactions/${id}`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    });
    return response.json();
  } catch (error) {
    console.error("Error fetching transaction by ID:", error);
    throw error;
  }
};

// GET /transactions/random-balance-history
export const getRandomBalanceHistory = async () => {
  try {
    const response = await fetch(
      `${API_BASE_URL}/transactions/random-balance-history?monthsBeforeFirstTransaction=12`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );

    // Check if the response is OK (status code 200-299)
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    // Parse the response JSON
    return response.json();
  } catch (error) {
    console.error("Error fetching random balance history:", error);
    throw error;
  }
};

// GET /transactions/latest-transfers
export const getLatestTransfers = async (
  accessToken: string,
  number: number
) => {
  try {
    const response = await fetch(
      `${API_BASE_URL}/transactions/quick-transfers?number=${number}`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      }
    );
    return response.json();
  } catch (error) {
    console.error("Error fetching latest transfers:", error);
    throw error;
  }
};

// GET /transactions/incomes
export const getIncomes = async (page: any, size: any) => {
  try {
    const response = await fetch(
      `${API_BASE_URL}/transactions/incomes?page=${page}&size=${size}`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    return response.json();
  } catch (error) {
    console.error("Error fetching incomes:", error);
    throw error;
  }
};

// GET /transactions/expenses
export const getExpenses = async (page: any, size: any) => {
  try {
    const response = await fetch(
      `${API_BASE_URL}/transactions/expenses?page=${page}&size=${size}`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    return response.json();
  } catch (error) {
    console.error("Error fetching expenses:", error);
    throw error;
  }
};

// GET /transactions/balance-history
export const getBalanceHistory = async (accessToken: string) => {
  try {
    const response = await fetch(
      `${API_BASE_URL}/transactions/balance-history`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      }
    );
    return response.json();
  } catch (error) {
    console.error("Error fetching balance history:", error);
    throw error;
  }
};
