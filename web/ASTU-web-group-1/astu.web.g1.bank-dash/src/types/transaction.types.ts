export interface TransactionDataType {
  transactionId: string;
  type: string;
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string;
}

export interface TransactionResponseType {
  success: boolean;
  message: string;
  data: TransactionDataType[];
}

export interface singleTransactionResponseType {
  success: boolean;
  message: string;
  data: TransactionDataType;
}

// Random Balance History Types
export interface RandomBalanceDataType {
  time: string;
  value: number;
}

export interface RandomBalanceResponseType {
  success: boolean;
  message: string;
  data: RandomBalanceDataType[];
}

// latest Transfers history
export interface LatestTransfersDataType {
  id: string;
  name: string;
  username: string;
  city: string;
  country: string;
  profilePicture: string;
}

export interface LatestTransfersResponseType {
  success: boolean;
  message: string;
  data: LatestTransfersDataType[];
}
