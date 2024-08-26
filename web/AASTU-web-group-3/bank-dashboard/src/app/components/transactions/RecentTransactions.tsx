'use client'
import React, { useState, useEffect } from 'react';
import { FaRegArrowAltCircleUp, FaRegArrowAltCircleDown } from "react-icons/fa";
import { FaLessThan, FaGreaterThan } from "react-icons/fa6";
import { useDispatch, useSelector } from "react-redux";
import {
  useGetAllTransactionsQuery,
  useGetIncomeTransactionsQuery,
  useGetExpenseTransactionsQuery,
} from "@/lib/redux/api/transactionsApi";
import {
  useGetCardsQuery
} from '@/lib/redux/api/cardsApi'
import {
  setAllTransactions,
  setIncomeTransactions,
  setExpenseTransactions,
  setLoading,
  setError,
} from "@/lib/redux/slices/transactionsSlice";

import { setCards } from '@/lib/redux/slices/cardsSlice';
import { RootState } from "@/lib/redux/store";
import Loading from '@/app/loading';

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
  const [allPage, setAllPage] = useState(1);
  const [incomePage, setIncomePage] = useState(1);
  const [expensePage, setExpensePage] = useState(1);

  const dispatch = useDispatch();
  const { allTransactions, incomeTransactions, expenseTransactions, loading, error } = useSelector(
    (state: RootState) => state.transactions
  );

  const {cards} = useSelector((state:RootState)=>state.cards)

  const {
    data: allData,
    isLoading: isLoadingAll,
    isError: isErrorAll,
  } = useGetAllTransactionsQuery({ size: 5, page: allPage - 1 });

  const {
    data:cardsData,
    isLoading:cardLoading,
    isError:cardError
  } = useGetCardsQuery({size:5, page:0})

  const {
    data: incomeData,
    isLoading: isLoadingIncome,
    isError: isErrorIncome,
  } = useGetIncomeTransactionsQuery({ size: 5, page: incomePage - 1 });

  const {
    data: expenseData,
    isLoading: isLoadingExpense,
    isError: isErrorExpense,
  } = useGetExpenseTransactionsQuery({ size: 5, page: expensePage - 1 });

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
    if (cardsData){
      dispatch(setCards(cardsData.content))

    }

    if (isErrorAll || isErrorIncome || isErrorExpense||cardError) {
      console.error("Error fetching data", {
        isErrorAll,
        isErrorIncome,
        isErrorExpense,
      });
      dispatch(setError("Error loading transactions"));
    }
  }, [allData, incomeData, expenseData, cardsData,  isLoadingAll, isLoadingIncome, isLoadingExpense, isErrorAll, isErrorIncome, isErrorExpense,cardError, dispatch]);

  const handlePrevPage = () => {
    if (activeTab === 'all transactions' && allPage > 1) {
      setAllPage(allPage - 1);
    } else if (activeTab === 'income' && incomePage > 1) {
      setIncomePage(incomePage - 1);
    } else if (activeTab === 'expenses' && expensePage > 1) {
      setExpensePage(expensePage - 1);
    }
  };

  const handleNextPage = () => {
    if (activeTab === 'all transactions' && allPage < allData!.data.totalPages) {
      setAllPage(allPage + 1);
    } else if (activeTab === 'income' && incomePage < incomeData!.data.totalPages) {
      setIncomePage(incomePage + 1);
    } else if (activeTab === 'expenses' && expensePage < expenseData!.data.totalPages) {
      setExpensePage(expensePage + 1);
    }
  };

  const handlePageClick = (pageNumber: number) => {
    if (activeTab === 'all transactions') {
      setAllPage(pageNumber);
    } else if (activeTab === 'income') {
      setIncomePage(pageNumber);
    } else if (activeTab === 'expenses') {
      setExpensePage(pageNumber);
    }
  };

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

  const renderPageButtons = () => {
    const totalPages = activeTab === 'all transactions' 
      ? allData?.data.totalPages 
      : activeTab === 'income' 
      ? incomeData?.data.totalPages 
      : expenseData?.data.totalPages;
    
    if (!totalPages) return null;
  
    const currentPage = activeTab === 'all transactions' 
      ? allPage 
      : activeTab === 'income' 
      ? incomePage 
      : expensePage;
  
    const pageButtons = [];
    for (let i = 1; i <= totalPages; i++) {
      pageButtons.push(
        <button
          key={i}
          onClick={() => handlePageClick(i)}
          className={`px-4 py-2 rounded mx-1 
            ${i === currentPage ? 'bg-blue-600 text-white' : 'bg-white text-blue-600 border border-blue-600'}
          `}
        >
          {i}
        </button>
      );
    }
    return pageButtons;
  };
  
  

  // if (loading) return <Loading />;
  if (error) return <div>{error}</div>;

  return (
    <div className="mt-4 mx-auto">
      <h1 className="text-xl font-semibold mb-4 text-[#343C6A] dark:text-white ">Recent Transactions</h1>
      <div className="border-b border-gray-200 dark:border-darkBorder flex justify-start gap-4 mb-2">
        <div
          onClick={() => setActiveTab('all transactions')}
          className={`cursor-pointer text-sm xl:text-xl 
            ${
              activeTab === 'all transactions'
                ? 'text-blue-600 border-blue-600 border-b dark:text-darkPrimary dark:border-darkPrimary'
                : 'text-[#718EBF] dark:text-lightText'
            }
          `}
        >
          All transactions
        </div>
        <div
          onClick={() => setActiveTab('income')}
          className={`cursor-pointer text-sm xl:text-xl 
            ${
              activeTab === 'income'
                ? 'text-blue-600 border-blue-600 border-b dark:text-darkPrimary dark:border-darkPrimary'
                : 'text-[#718EBF] dark:text-lightText'
            }
          `}
        >
          Income
        </div>
        <div
          onClick={() => setActiveTab('expenses')}
          className={`cursor-pointer text-sm xl:text-xl 
            ${
              activeTab === 'expenses'
                ? 'text-blue-600 border-blue-600 border-b dark:text-darkPrimary dark:border-darkPrimary'
                : 'text-[#718EBF] dark:text-lightText'
            }
          `}
        >
          Expenses
        </div>
      </div>
      <section className="border-0 rounded-xl bg-white dark:bg-darkComponent shadow-md p-2">
  <div className="hidden lg:grid lg:grid-cols-7 xl:grid-cols-9 font-medium text-sky-300 dark:text-lightText text-xs lg:text-sm xl:text-base h-7 items-center border-b dark:border-darkBorder mt-2">
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
      className="grid grid-cols-7 xl:grid-cols-9 border-b dark:border-darkBorder min-h-12 items-center text-xs lg:text-sm xl:text-base dark:text-lightText"
    >
      <div className="flex items-center gap-2 col-span-5 lg:col-span-2 lg:font-medium">
        {transaction.amount < 0 ? (
          <FaRegArrowAltCircleUp
            color="#718EBF"
            className="text-4xl md:text-xl xl:text-3xl dark:text-lightText"
          />
        ) : (
          <FaRegArrowAltCircleDown
            color="#718EBF"
            className="text-4xl md:text-xl xl:text-3xl dark:text-lightText"
          />
        )}
        <span>{transaction.description}</span>
      </div>
      <div className="hidden xl:block xl:col-span-2">{transaction.transactionId}</div>
      <div className="hidden lg:block">{transaction.type}</div>
      <div className="hidden lg:block">{transaction.senderUserName}</div>
      <div className="hidden lg:block sm:text-xs">{formatDate(transaction.date)}</div>
      <div
        className={`col-span-2 lg:col-span-1 justify-self-end lg:justify-self-auto ${
          transaction.amount < 0 ? 'text-red-500 dark:text-red-400' : 'text-green-500 dark:text-green-400'
        }`}
      >
        {transaction.amount < 0 ? '-' : '+'}${Math.abs(transaction.amount)}
      </div>
      <div className="hidden lg:block first-line:col-span-2 lg:col-span-1 justify-self-end lg:justify-self-auto xl:justify-self-center border border-blue-300 hover:border-blue-500 px-3 py-1 rounded-xl hover:cursor-pointer ">
        Download
      </div>
    </div>
  ))}
</section>

      <div className="flex justify-end space-x-2 items-center mt-4">
        <button
          onClick={handlePrevPage}
          disabled={
            (activeTab === 'all transactions' && allPage === 1) ||
            (activeTab === 'income' && incomePage === 1) ||
            (activeTab === 'expenses' && expensePage === 1)
          }
          className="px-4 py-2 bg-blue-600 dark:bg-darkPrimary rounded disabled:bg-blue-400 text-white flex gap-1 items-center"
        >
          <FaLessThan />
          <span> Previous </span>
        </button>
        <div className="flex">{renderPageButtons()}</div>
        <button
          onClick={handleNextPage}
          disabled={
            (activeTab === 'all transactions' && allPage === allData?.data.totalPages) ||
            (activeTab === 'income' && incomePage === incomeData?.data.totalPages) ||
            (activeTab === 'expenses' && expensePage === expenseData?.data.totalPages)
          }
          className="px-4 py-2 bg-blue-600 dark:bg-darkPrimary rounded disabled:bg-blue-400 text-white flex gap-1 items-center"
        >
          <span> Next </span>
          <FaGreaterThan />
        </button>
      </div>
    </div>
  );
  
};

export default RecentTransactions;
