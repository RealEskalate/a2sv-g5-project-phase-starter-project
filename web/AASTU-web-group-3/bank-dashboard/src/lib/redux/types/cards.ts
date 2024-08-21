// Card Types
export interface Card {
    id: string;
    cardHolder: string;
    semiCardNumber: string;
    cardType: string;
    balance: number;
    expiryDate: string;
  }
  
  export interface PaginatedCardsResponse {
    content: Card[];
    totalPages: number;
  }
  
  export interface CreateCardRequest {
    balance: number;
    cardHolder: string;
    expiryDate: string;
    passcode: string;
    cardType: string;
  }
  
  export interface CreateCardResponse {
    id: string;
    balance: number;
    cardHolder: string;
    expiryDate: string;
    cardNumber: string;
    passcode: string;
    cardType: string;
    userId: string;
  }
  
  export interface CardDetail {
    id: string;
    balance: number;
    cardHolder: string;
    expiryDate: string;
    cardNumber: string;
    passcode: string;
    cardType: string;
    userId: string;
  }
  