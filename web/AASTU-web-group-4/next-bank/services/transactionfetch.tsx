import Cookie from "js-cookie"
const API_BASE_URL = "https://bank-dashboard-6acc.onrender.com";
const token = Cookie.get("accessToken")
  
// GET /transactions
export const getAllTransactions = async () => {
  try{
  const response = await fetch(`${API_BASE_URL}/transactions?page=${0}&size=${5}`, {
    method: 'GET',
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  if (!response.ok)
     {
    console.log("getalltransaction fetch :", response)
    throw new Error("Failed to fetch user details");
  } 
  const data = await response.json();
  console.log("succesful transaction response:", data)
  return data;
}catch (error) {
    console.error("Error:", error);
    throw error;
  }

};

// POST /transactions
export const createTransaction = async (transactionData:any) => {
  const response = await fetch(`${API_BASE_URL}/transactions`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(transactionData),
  });
  return response.json();
};

// POST /transactions/deposit
export const createDeposit = async (depositData:any) => {
  const response = await fetch(`${API_BASE_URL}/transactions/deposit`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(depositData),
  });
  return response.json();
};

// GET /transactions/{id}
export const getTransactionById = async (id:any) => {
  const response = await fetch(`${API_BASE_URL}/transactions/${id}`, {
    method: 'GET',
  });
  return response.json();
};

// GET /transactions/random-balance-history
export const getRandomBalanceHistory = async () => {
  const response = await fetch(`${API_BASE_URL}/transactions/random-balance-history`, {
    method: 'GET',
  });
  return response.json();
};

// GET /transactions/latest-transfers
export const getLatestTransfers = async () => {
  const response = await fetch(`${API_BASE_URL}/transactions/latest-transfers`, {
    method: 'GET',
  });
  return response.json();
};

// GET /transactions/incomes
export const getIncomes = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/transactions/incomes?page=${0}&size=${5}`, {
    method: 'GET',
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  // return response.json();
  if (!response.ok) {
    throw new Error(`Failed to fetch expenses: ${response.statusText}`);
  }

  const data = await response.json();
  console.log('Fetched Expenses:', data);
  return data;
} catch (error) {
  console.error('Error fetching expenses:', error);
  return null;
}
// };
};

// GET /transactions/expenses
export const getExpenses = async () => {
  const response = await fetch(`${API_BASE_URL}/transactions/expenses?page=${0}&size=${5}`, {
    method: 'GET',
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return response.json();
};

// GET /transactions/balance-history
export const getBalanceHistory = async () => {
  const response = await fetch(`${API_BASE_URL}/transactions/balance-history`, {
    method: 'GET',
  });
  return response.json();
};
