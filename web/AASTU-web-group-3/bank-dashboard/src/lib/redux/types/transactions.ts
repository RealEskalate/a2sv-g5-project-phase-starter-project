export interface Transaction {
  transactionId: string;
  type: string;
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string | null;
}

export interface TransactionsResponse {
  success: boolean;
  message: string;
  data: Transaction[];
}

export interface TransactionResponse {
  success: boolean;
  message: string;
  data: Transaction;
}

export interface TransactionRequest {
  type: string;
  description: string;
  amount: number;
  receiverUserName: string;
}

export interface TransactionDepositRequest {
  description: string;
  amount: number;
}

interface Transfer {
  id: string;
  name: string;
  username: string;
  city: string;
  country: string;
  profilePicture: string;
}

export interface LatestTransferResponse {
  success: true;
  message: string;
  data: Transfer[];
}

export interface BalanceData {
  time: string;
  value: number;
}
export interface BalanceHistoryResponse {
  success: true;
  message: string;
  data: BalanceData[];
}

export interface MyExpenseResponse {
  success: boolean;
  message: string;
  data: {
    content: Transaction[];
    totalPages: number;
  };
}

export interface IncomeResponse {
  success: boolean;
  message: string;
  data: {
    content: [
      {
        transactionId: string;
        type: string;
        senderUserName: string;
        description: string;
        date: string;
        amount: number;
        receiverUserName: string;
      }
    ];
    totalPages: number;
  };
}

export interface getQuickTransfersResponse {
  success: boolean;
  message: string;
  data: {
    id: string;
    name: string;
    username: string;
    city: string;
    country: string;
    profilePicture: string;
  }[];
}
