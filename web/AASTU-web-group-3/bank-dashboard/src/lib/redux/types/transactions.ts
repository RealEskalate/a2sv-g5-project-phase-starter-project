export interface Transaction {
    transactionId: string;
    type: string;
    senderUserName: string;
    description: string;
    date: string;
    amount: number;
    receiverUserName: string;
  }
  
  export interface TransactionsResponse {
    success: boolean;
    message: string;
    data: Transaction[];
  }
  