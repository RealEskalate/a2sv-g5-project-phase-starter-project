// // src/features/cardSlice.ts
// import { createSlice, PayloadAction } from "@reduxjs/toolkit";

// interface YearlyInvestment {
//   time: string;
//   value: number;
// }

// interface MonthlyRevenue {
//   time: string;
//   value: number;
// }

// interface InvestmentData {
//   totalInvestment: number;
//   rateOfReturn: number;
//   yearlyTotalInvestment: YearlyInvestment[];
//   monthlyRevenue: MonthlyRevenue[];
// }

// interface accountState {
//   cards: Card[];
//   status: "idle" | "loading" | "succeeded" | "failed";
//   error: string | null;
// }

// const initialState: CardState = {
//   cards: [],
//   status: "idle",
//   error: null,
// };

// const accountSlice = createSlice({
//   name: "cards",
//   initialState,
//   reducers: {
//     setCards(state, action: PayloadAction<Card[]>) {
//       state.cards = action.payload;
//     },
//     addCard(state, action: PayloadAction<Card>) {
//       state.cards.push(action.payload);
//     },
//     removeCard(state, action: PayloadAction<string>) {
//       state.cards = state.cards.filter((card) => card.id !== action.payload);
//     },
//     setStatus(state, action: PayloadAction<CardState["status"]>) {
//       state.status = action.payload;
//     },
//     setError(state, action: PayloadAction<string | null>) {
//       state.error = action.payload;
//     },
//   },
// });

// export const { setCards, addCard, removeCard, setStatus, setError } =
//   accountSlice.actions;

// export default accountSlice.reducer;
