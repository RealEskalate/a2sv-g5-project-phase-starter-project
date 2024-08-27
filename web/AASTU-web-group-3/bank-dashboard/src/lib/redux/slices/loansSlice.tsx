import { createSlice, PayloadAction } from '@reduxjs/toolkit';
import { LoanResponse } from '@/lib/redux/api/loansApi';

interface LoansState {
  loans: LoanResponse['data'][];
  selectedLoan: LoanResponse['data'] | null;
}

const initialState: LoansState = {
  loans: [],
  selectedLoan: null,
};

const loansSlice = createSlice({
  name: 'loans',
  initialState,
  reducers: {
    setLoans(state, action: PayloadAction<LoanResponse['data'][]>) {
      state.loans = action.payload;
    },
    setSelectedLoan(state, action: PayloadAction<LoanResponse['data']>) {
      state.selectedLoan = action.payload;
    },
  },
});

export const { setLoans, setSelectedLoan } = loansSlice.actions;
export default loansSlice.reducer;
