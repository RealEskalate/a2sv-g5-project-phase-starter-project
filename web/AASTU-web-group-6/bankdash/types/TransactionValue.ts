export interface TransactionType {
  transactionId: string;
  type: string;
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string | null;
}

export interface TransactionResponse {
  success: boolean;
  message: string;
  data: TransactionType[];
}

export interface ChartData {
  day: string;
  amount: number;
};

export interface DailyAmount {
  day: string;
  amount: number;
}