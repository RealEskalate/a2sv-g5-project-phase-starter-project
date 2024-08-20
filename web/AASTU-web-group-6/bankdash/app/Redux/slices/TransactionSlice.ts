// src/features/cardSlice.ts
import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { string } from "zod";

export interface TransactionType {
  transactionId: string;
  type: string;
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string | null;
}
export interface BalanceType {
  time: string;
  value: string;
}

interface TranState {
  transactions: TransactionType[];
  balanceHist: BalanceType[];
  expense: TransactionType[];
  income: TransactionType[];
  status: "idle" | "loading" | "succeeded" | "failed";
  error: string | null;
}

const initialState: TranState = {
  transactions: [],
  balanceHist: [],
  expense: [],
  income: [],
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
    setBalHist(state, action: PayloadAction<BalanceType[]>) {
      state.balanceHist = action.payload;
    },
    setExpense(state, action: PayloadAction<TransactionType[]>) {
      state.expense = action.payload;
    },
    setIncome(state, action: PayloadAction<TransactionType[]>) {
      state.income = action.payload;
    },

    setStatus(state, action: PayloadAction<TranState["status"]>) {
      state.status = action.payload;
    },
    setError(state, action: PayloadAction<string | null>) {
      state.error = action.payload;
    },
  },
});

export const {
  setTran,
  setBalHist,
  setExpense,
  setIncome,
  setStatus,
  setError,
} = transactionSlice.actions;

export default transactionSlice.reducer;
