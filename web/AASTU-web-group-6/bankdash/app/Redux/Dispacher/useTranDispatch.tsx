"use client";
import { useEffect } from "react";
import { useAppDispatch } from "@/app/Redux/store/store";
import TransactionService from "@/app/Services/api/transactionApi";
import {
  setTran,
  setStatus,
  setBalHist,
  setError,
  setExpense,
  setIncome,
} from "../slices/TransactionSlice";

const useTranDispatch = (accessToken: string) => {
  const dispatch = useAppDispatch();

  useEffect(() => {
    const fetchInitialCards = async () => {
      try {
        dispatch(setStatus("loading"));
        const total_expense : any =  await TransactionService.getOneExpenseData(accessToken);
        console.log(total_expense , "vvvvvvvvv")
        const total_income: any =  await TransactionService.getOneIncomeData(accessToken);
        console.log(total_income , "wwwwwww")
        const transaction: any = await TransactionService.getTransactions(
          accessToken
        );
        console.log(transaction , "transaction")
        const history: any = await TransactionService.balanceHistory(
          accessToken
        );
        // console.log(history , "his")
        const expense: any = await TransactionService.getExpenseData(
          accessToken , total_expense.totalPages
        );
        console.log(total_income.totalPages , "total")
        const income: any = await TransactionService.getIncomeData(accessToken , total_income.totalPages);
        console.log(income , "inc")
        // console.log(transaction , "transaction")
        // console.log(history , "history")
        // console.log(expense , "dddddd")
        if (transaction) {
          dispatch(setTran(transaction.content));
          dispatch(setBalHist(history));
          dispatch(setExpense(expense.content));
          dispatch(setIncome(income.content));
          dispatch(setStatus("succeeded"));
        }
      } catch (error) {
        dispatch(setError("Failed to fetch transaction"));
        dispatch(setStatus("failed"));
      }
    };

    fetchInitialCards();
  }, [dispatch, accessToken]);
};

export default useTranDispatch;
