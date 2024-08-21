import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Card, CardDetail } from '../types/cards';

interface CardsState {
  cards: Card[];
  cardDetail: CardDetail | null;
  totalPages: number;
}

const initialState: CardsState = {
  cards: [],
  cardDetail: null,
  totalPages: 0,
};

const cardsSlice = createSlice({
  name: 'cards',
  initialState,
  reducers: {
    setCards: (state, action: PayloadAction<Card[]>) => {
      state.cards = action.payload;
    },
    setCardDetail: (state, action: PayloadAction<CardDetail | null>) => {
      state.cardDetail = action.payload;
    },
    setTotalPages: (state, action: PayloadAction<number>) => {
      state.totalPages = action.payload;
    },
  },
});

export const { setCards, setCardDetail, setTotalPages } = cardsSlice.actions;

export default cardsSlice.reducer;
