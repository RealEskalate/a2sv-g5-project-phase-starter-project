export interface CardContentType {
  id: string;
  cardHolder: string;
  semiCardNumber: string;
  cardType: string;
  expiryDate: string;
  balance: number;
}
export interface CardResponseType {
  content: CardContentType[];
  totalPages: number;
}

export interface SingleCardResponseType {
  id: string;
  balance: number;
  cardHolder: string;
  expiryDate: string;
  cardNumber: string;
  passcode: string;
  cardType: string;
  userId: string;
}
