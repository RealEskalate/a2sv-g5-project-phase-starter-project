import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';

interface Transaction {
  name: string;
  date: string;
  amount: string;
  image: string;
}

interface TransactionsState {
  data: Transaction[];
  loading: boolean;
  error: string | null;
}

const initialState: TransactionsState = {
  data: [],
  loading: false,
  error: null,
};

// Replace with your actual API endpoint
const API_URL = 'https://bank-dashboard-6acc.onrender.com/transactions';

export const getRecentTransactions = createAsyncThunk(
  'transactions/getRecentTransactions',
  async (_, { rejectWithValue }) => {
    try {
      const response = await axios.get(API_URL, {
        headers: {
          'Authorization': `Bearer ${process.env.NEXT_PUBLIC_API_TOKEN}`, // Use environment variable
        }
      });
      return response.data;
    } catch (error) {
      // Handle the error properly
      if (axios.isAxiosError(error) && error.response) {
        return rejectWithValue(error.response.data);
      }
      return rejectWithValue('An unknown error occurred');
    }
  }
);

const transactionsSlice = createSlice({
  name: 'transactions',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(getRecentTransactions.pending, (state) => {
        state.loading = true;
      })
      .addCase(getRecentTransactions.fulfilled, (state, action) => {
        state.loading = false;
        state.data = action.payload;
      })
      .addCase(getRecentTransactions.rejected, (state, action) => {
        state.loading = false;
        // Ensure that the action.payload is a string
        state.error = action.payload as string;
      });
  },
});

export default transactionsSlice.reducer;
