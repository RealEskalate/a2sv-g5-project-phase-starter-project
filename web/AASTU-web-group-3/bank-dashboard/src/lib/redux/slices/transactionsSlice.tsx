import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Transaction } from "../types/transactions";

interface TransactionsState {
  allTransactions: Transaction[];
  incomeTransactions: Transaction[];
  expenseTransactions: Transaction[];
  loading: boolean;
  error: string | null;
}

const initialState: TransactionsState = {
  allTransactions: [],
  incomeTransactions: [],
  expenseTransactions: [],
  loading: false,
  error: null,
};

const transactionsSlice = createSlice({
  name: "transactions",
  initialState,
  reducers: {
    setAllTransactions: (state, action: PayloadAction<Transaction[]>) => {
      state.allTransactions = action.payload;
    },
    setIncomeTransactions: (state, action: PayloadAction<Transaction[]>) => {
      state.incomeTransactions = action.payload;
    },
    setExpenseTransactions: (state, action: PayloadAction<Transaction[]>) => {
      state.expenseTransactions = action.payload;
    },
    setLoading: (state, action: PayloadAction<boolean>) => {
      state.loading = action.payload;
    },
    setError: (state, action: PayloadAction<string | null>) => {
      state.error = action.payload;
    },
  },
});

export const {
  setAllTransactions,
  setIncomeTransactions,
  setExpenseTransactions,
  setLoading,
  setError,
} = transactionsSlice.actions;

export default transactionsSlice.reducer;
