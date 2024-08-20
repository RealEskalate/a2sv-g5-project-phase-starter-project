const API_BASE_URL = 'https://your-api-domain.com';

// GET /transactions
export const getAllTransactions = async () => {
  const response = await fetch(`${API_BASE_URL}/transactions`, {
    method: 'GET',
  });
  return response.json();
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
  const response = await fetch(`${API_BASE_URL}/transactions/incomes`, {
    method: 'GET',
  });
  return response.json();
};

// GET /transactions/expenses
export const getExpenses = async () => {
  const response = await fetch(`${API_BASE_URL}/transactions/expenses`, {
    method: 'GET',
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
