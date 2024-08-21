export interface UserPreference {
  currency: string;
  sentOrReceiveDigitalCurrency: boolean;
  receiveMerchantOrder: boolean;
  accountRecommendations: boolean;
  timeZone: string;
  twoFactorAuthentication: boolean;
}

export interface UserData {
  id: string;
  name: string;
  email: string;
  dateOfBirth: string;
  permanentAddress: string;
  postalCode: string;
  username: string;
  presentAddress: string;
  city: string;
  country: string;
  profilePicture: string;
  accountBalance: number;
  role: string;
  preference: UserPreference;
}

export interface UserResponse {
  success: boolean;
  message: string;
  data: UserData;
}
export interface TransactionContent{
  transactionId: string;
  type: "shopping" | "transfer" | "service" | "deposit";
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string;
}
export interface TransactionData {
  content:TransactionContent[]
  totalPages:number
  
}



export interface TransactionProps {
  success: boolean;
  message : string;
  data : TransactionData;
}

export interface CreditCardProps{
    content : CardDetails[]
    totalPages:number
}
export interface CardDetails {
  id: string;
  cardHolder: string;
  semiCardNumber: string;
  cardType: string;
  balance: number;
  expiryDate: string; 
}

export interface RandomBalanceHistory {
  success: boolean;
  message: string;
  data: BalanceData[];
}

export interface BalanceData {
  time: string;
  value: number;
}

export interface QuickTransferProps {
  success: boolean;
  message : string;
  data : QuickTransferData[];
}

export interface QuickTransferData {
  id: string;
  name: string;
  username: string;
  city: string;
  country: string;
  profilePicture: string; 
}
