'use client'
import React, { useState, useEffect } from 'react';
import { FaRegArrowAltCircleUp, FaRegArrowAltCircleDown } from "react-icons/fa";
import { useDispatch, useSelector } from "react-redux";
import {
  useGetAllTransactionsQuery,
  useGetIncomeTransactionsQuery,
  useGetExpenseTransactionsQuery,
} from "@/lib/redux/api/transactionsApi";
import {
  setAllTransactions,
  setIncomeTransactions,
  setExpenseTransactions,
  setLoading,
  setError,
} from "@/lib/redux/slices/transactionsSlice";
import { RootState } from "@/lib/redux/store";

const formatDate = (dateString: string) => {
  const date = new Date(dateString);

  const formattedDate = date.toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric',
  });

  const formattedTime = date.toLocaleTimeString('en-US', {
    hour: '2-digit',
    minute: '2-digit',
    hour12: true,
  });

  return `${formattedDate}, ${formattedTime}`;
};

const RecentTransactions = () => {
  const [activeTab, setActiveTab] = useState('all transactions');

  const dispatch = useDispatch();
  const { allTransactions, incomeTransactions, expenseTransactions, loading, error } = useSelector(
    (state: RootState) => state.transactions
  );

  const {
    data: allData,
    isLoading: isLoadingAll,
    isError: isErrorAll,
  } = useGetAllTransactionsQuery({ size: 10, page: 0 });

  const {
    data: incomeData,
    isLoading: isLoadingIncome,
    isError: isErrorIncome,
  } = useGetIncomeTransactionsQuery({ size: 10, page: 0 });

  const {
    data: expenseData,
    isLoading: isLoadingExpense,
    isError: isErrorExpense,
  } = useGetExpenseTransactionsQuery({ size: 10, page: 0 });

  useEffect(() => {
    dispatch(setLoading(isLoadingAll || isLoadingIncome || isLoadingExpense));

    if (allData) {
      dispatch(setAllTransactions(allData.data.content));
    }
    if (incomeData) {
      dispatch(setIncomeTransactions(incomeData.data.content));
    }
    if (expenseData) {
      dispatch(setExpenseTransactions(expenseData.data.content));
    }

    if (isErrorAll || isErrorIncome || isErrorExpense) {
      console.error("Error fetching data", {
        isErrorAll,
        isErrorIncome,
        isErrorExpense,
      });
      dispatch(setError("Error loading transactions"));
    }
  }, [allData, incomeData, expenseData, isLoadingAll, isLoadingIncome, isLoadingExpense, isErrorAll, isErrorIncome, isErrorExpense, dispatch]);

  const renderContent = () => {
    switch (activeTab) {
      case 'all transactions':
        return allTransactions;
      case 'income':
        return incomeTransactions;
      case 'expenses':
        return expenseTransactions;
      default:
        return [];
    }
  };

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className='mt-4 mx-auto'>
      <h1 className='text-xl font-semibold mb-2'>Recent Transactions</h1>
      <div className='border-b flex justify-start gap-4 mb-2'>
        <div
          onClick={() => setActiveTab('all transactions')}
          className={`cursor-pointer text-sm xl:text-xl 
            ${activeTab === 'all transactions' ? "text-blue-600 border-blue-600 border-b" : "text-[#718EBF]"}
          `}
        >
          All transactions
        </div>
        <div
          onClick={() => setActiveTab('income')}
          className={`cursor-pointer text-sm xl:text-xl 
            ${activeTab === 'income' ? "text-blue-600 border-blue-600 border-b" : "text-[#718EBF]"}
          `}
        >
          Income
        </div>
        <div
          onClick={() => setActiveTab('expenses')}
          className={`cursor-pointer text-sm xl:text-xl 
            ${activeTab === 'expenses' ? "text-blue-600 border-blue-600 border-b" : "text-[#718EBF]"}
          `}
        >
          Expenses
        </div>
      </div>
      <section className="border-0 rounded-xl bg-white shadow-md p-2">
        <div className="hidden lg:grid lg:grid-cols-7 xl:grid-cols-9 font-medium text-sky-300 text-xs h-7 items-center border-b mt-2">
          <div className="lg:col-span-2">Description</div>
          <div className="hidden xl:block xl:col-span-2">Transaction Id</div>
          <div>Type</div>
          <div>Sender</div>
          <div>Date</div>
          <div>Amount</div>
          <div className="justify-self-center">Receipt</div>
        </div>
        {renderContent().map((transaction, index) => (
          <div
            key={index}
            className="grid grid-cols-7 xl:grid-cols-9 border-b min-h-12 items-center text-xs lg:font-medium xl:text-[18px]"
          >
            <div className="flex items-center gap-2 col-span-5 lg:col-span-2 lg:font-medium">
              {transaction.amount < 0 ? (
                <FaRegArrowAltCircleUp
                  color="#718EBF"
                  className="text-4xl md:text-xl xl:text-3xl"
                />
              ) : (
                <FaRegArrowAltCircleDown
                  color="#718EBF"
                  className="text-4xl md:text-xl xl:text-3xl"
                />
              )}
              <span>{transaction.description}</span>
            </div>
            <div className="hidden xl:block xl:col-span-2">{transaction.transactionId}</div>
            <div className="hidden lg:block">{transaction.type}</div>
            <div className="hidden lg:block">{transaction.senderUserName}</div>
            <div className="hidden lg:block">{formatDate(transaction.date)}</div>
            <div
              className={`col-span-2 lg:col-span-1 justify-self-end lg:justify-self-auto ${
                transaction.amount < 0 ? "text-red-500" : "text-green-500"
              }`}
            >
              {transaction.amount < 0 ? "-" : "+"}${Math.abs(transaction.amount)}
            </div>
            <div className="hidden lg:block border p-1 rounded-lg w-auto justify-self-center hover:border-blue-400 hover:cursor-pointer">Download</div>
          </div>
        ))}
      </section>
    </div>
  );
};

export default RecentTransactions;
