export interface Card {
    id: string;
    balance: number;
    cardHolder: string;
    expiryDate: string;
    cardType: string;
    semiCardNumber: string 
  }
  

  export interface GetCardsResponse {
    content: Card[];
    totalPages: number
  }
  
  export interface PostCardRequest {
    balance: number;
    cardHolder: string;
    expiryDate: string;
    passcode: string;
    cardType: string;
  }
  
  export interface PostCardResponse extends Card {}
  export interface GetCardByIdResponse extends Card {}  