
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

export interface TransactionResponse{
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

export interface TransactionDepositRequest{
  description:string;
  amount:number;
}

export interface TransactionDepositResponse{

}