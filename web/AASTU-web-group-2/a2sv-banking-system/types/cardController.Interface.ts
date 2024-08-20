export interface Card {
    id: string;
    balance: number;
    cardHolder: string;
    expiryDate: string;
    cardType: string;
  }
  

  export interface GetCardsResponse {
    cards: Card[];
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