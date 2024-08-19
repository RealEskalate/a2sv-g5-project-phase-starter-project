// src/features/cardSlice.ts
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

export interface TransactionType {
  transactionId: string;
  type: string;
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string | null;
}

interface TranState {
  transactions: TransactionType[];
  balanceHist: TransactionType[];
  status: "idle" | "loading" | "succeeded" | "failed";
  error: string | null;
}

const initialState: TranState = {
  transactions: [],
  balanceHist: [],
  status: "idle",
  error: null,
};

const transactionSlice = createSlice({
  name: "transactions",
  initialState,
  reducers: {
    setTran(state, action: PayloadAction<TransactionType[]>) {
      state.transactions = action.payload;
    },
    setBalHist(state, action: PayloadAction<TransactionType[]>) {
      state.balanceHist = action.payload;
    },

    setStatus(state, action: PayloadAction<TranState["status"]>) {
      state.status = action.payload;
    },
    setError(state, action: PayloadAction<string | null>) {
      state.error = action.payload;
    },
  },
});

export const { setTran, setBalHist, setStatus, setError } =
  transactionSlice.actions;

export default transactionSlice.reducer;
