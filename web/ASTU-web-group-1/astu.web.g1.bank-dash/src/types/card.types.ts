export interface CardResponseType {
  id: string;
  cardHolder: string;
  semiCardNumber: string;
  cardType: string;
  expiryDate: string;
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
