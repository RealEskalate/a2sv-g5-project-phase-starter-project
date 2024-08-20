import {PostTransactionRequest, PaginatedTransactionsResponse, QuickTransferData, QuickTransferResponse, BalanceHistoryData, BalanceHistoryResponse, GetTransactionByID, TransactionResponse, TransactionData, GetTransactionsResponse, PostTransactionResponse, PostDepositTransactionRequest} from "@/types/transactionController.interface"

const BASE_URL = 'https://bank-dashboard-6acc.onrender.com';

const getTransactions = async (page: number, size: number): Promise<GetTransactionsResponse> => {
    try {
      const response = await fetch(`${BASE_URL}/transactions?page=${page}&size=${size}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
  
      if (response.status === 200) {
        const data: TransactionResponse = await response.json();
        return { transactions: [data] }; // Wrap the response in the expected format
      } else {
        throw new Error(`Request failed with status code: ${response.status}`);
      }
    } catch (error) {
      console.error('Error fetching transactions:', error);
      throw error;
    }
  };

  const postTransaction = async (transactionDetails: PostTransactionRequest): Promise<PostTransactionResponse> => {
    try {
      const response = await fetch(`${BASE_URL}/transactions`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(transactionDetails),
      });
  
      if (response.status === 200) {
        const data: PostTransactionResponse = await response.json();
        return data;
      } else {
        throw new Error(`Request failed with status code: ${response.status}`);
      }
    } catch (error) {
      console.error('Error posting transaction:', error);
      throw error;
    }
  };
  
  const postTransactionsDeposit = async (transactionDetails: PostDepositTransactionRequest): Promise<TransactionResponse> => {
    try {
      const response = await fetch(`${BASE_URL}/transactions/deposit`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(transactionDetails),
      });
  
      if (response.status === 200) {
        const data: TransactionResponse = await response.json();
        return data;
      } else {
        throw new Error(`Request failed with status code: ${response.status}`);
      }
    } catch (error) {
      console.error('Error posting transaction:', error);
      throw error;
    }
  };
  const getTransactionById = async (transactionId: string): Promise<GetTransactionByID> => {
    try {
      const response = await fetch(`${BASE_URL}/transactions/${transactionId}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
  
      if (response.status === 200) {
        const data: TransactionResponse = await response.json();
        return data;
      } else {
        throw new Error(`Request failed with status code: ${response.status}`);
      }
    } catch (error) {
      console.error('Error fetching transaction:', error);
      throw error;
    }
  };
  const getRandomBalanceHistory = async (monthsBeforeFirstTransaction: number): Promise<BalanceHistoryResponse> => {
    try {
      const response = await fetch(`${BASE_URL}/transactions/random-balance-history?monthsBeforeFirstTransaction=${monthsBeforeFirstTransaction}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
  
      if (response.status === 200) {
        const data: BalanceHistoryResponse = await response.json();
        return data;
      } else {
        throw new Error(`Request failed with status code: ${response.status}`);
      }
    } catch (error) {
      console.error('Error fetching balance history:', error);
      throw error;
    }
  };
  const getQuickTransfers = async (inputInteger: number): Promise<QuickTransferResponse> => {
    try {
      const response = await fetch(`${BASE_URL}/transactions/quick-transfers?input=${inputInteger}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
  
      if (response.status === 200) {
        const data: QuickTransferResponse = await response.json();
        return data;
      } else {
        throw new Error(`Request failed with status code: ${response.status}`);
      }
    } catch (error) {
      console.error('Error fetching quick transfers:', error);
      throw error;
    }
  };
  const getTransactionIncomes = async (page: number, size: number): Promise<PaginatedTransactionsResponse> => {
    try {
      const response = await fetch(`${BASE_URL}/transactions/incomes?page=${page}&size=${size}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
  
      if (response.status === 200) {
        const data: PaginatedTransactionsResponse = await response.json();
        return data;
      } else {
        throw new Error(`Request failed with status code: ${response.status}`);
      }
    } catch (error) {
      console.error('Error fetching paginated transactions:', error);
      throw error;
    }
  };
  const getTransactionsExpenses = async (page: number, size: number): Promise<PaginatedTransactionsResponse> => {
    try {
      const response = await fetch(`${BASE_URL}/transactions/expenses?page=${page}&size=${size}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
  
      if (response.status === 200) {
        const data: PaginatedTransactionsResponse = await response.json();
        return data;
      } else {
        throw new Error(`Request failed with status code: ${response.status}`);
      }
    } catch (error) {
      console.error('Error fetching paginated transactions:', error);
      throw error;
    }
  };
  const getBalanceHistory = async (monthsBeforeFirstTransaction: number): Promise<BalanceHistoryResponse> => {
    try {
      const response = await fetch(`${BASE_URL}/transactions/balance-history`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
  
      if (response.status === 200) {
        const data: BalanceHistoryResponse = await response.json();
        return data;
      } else {
        throw new Error(`Request failed with status code: ${response.status}`);
      }
    } catch (error) {
      console.error('Error fetching balance history:', error);
      throw error;
    }
  };

  export {getBalanceHistory, getRandomBalanceHistory, getQuickTransfers, getTransactionById, getTransactionIncomes, getTransactions, getTransactionsExpenses, postTransactionsDeposit, postTransaction}