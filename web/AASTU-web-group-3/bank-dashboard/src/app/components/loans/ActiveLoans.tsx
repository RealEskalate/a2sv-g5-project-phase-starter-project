'use client';
import React, { useEffect, useState } from 'react';
import { FaRegArrowAltCircleUp, FaRegArrowAltCircleDown } from 'react-icons/fa';
import { FaArrowLeft, FaArrowRight } from 'react-icons/fa';
import { useGetMyLoansQuery, useGetLoanDetailDataQuery } from '@/lib/redux/api/loansApi';
import { useDispatch } from 'react-redux';
import { setLoans } from '@/lib/redux/slices/loansSlice';

const ActiveLoans = () => {
  const [page, setPage] = useState(1); // Start from page 1
  const { data: loansData, isLoading: isLoadingLoans } = useGetMyLoansQuery({ page: page - 1, size: 10 });
  const { data: detailLoansData, isLoading: isLoadingDetail } = useGetLoanDetailDataQuery();
  const dispatch = useDispatch();

  useEffect(() => {
    if (loansData && loansData.success) {
      dispatch(setLoans(loansData.data.content));
    }
  }, [loansData, dispatch]);

  useEffect(() => {
    // Handle detail loans data if necessary
  }, [detailLoansData]);

  const handlePageClick = (newPage: number) => {
    if (newPage > 0 && newPage <= (loansData?.data.totalPages || 1)) {
      setPage(newPage);
    }
  };

  const renderPageButtons = () => {
    const totalPages = loansData?.data.totalPages || 1; // Ensure totalPages is at least 1

    const pageButtons = [];

    // Always show first page button and ellipsis if necessary
    if (page > 3) {
      pageButtons.push(
        <button
          key={1}
          onClick={() => handlePageClick(1)}
          className={`px-4 py-2 rounded mx-1 ${page === 1 ? 'bg-blue-600 text-white' : 'bg-white text-blue-600 border border-blue-600'}`}
        >
          1
        </button>
      );
      pageButtons.push(<span key="ellipsis-start" className="px-2">...</span>);
    }

    // Display the buttons for the current page and neighboring pages
    for (let i = Math.max(page - 1, 1); i <= Math.min(page + 1, totalPages); i++) {
      pageButtons.push(
        <button
          key={i}
          onClick={() => handlePageClick(i)}
          className={`px-4 py-2 rounded mx-1 ${i === page ? 'bg-blue-600 text-white' : 'bg-white text-blue-600 border border-blue-600'}`}
        >
          {i}
        </button>
      );
    }

    // Ellipsis if there are more pages after the current ones
    if (page + 1 < totalPages) {
      pageButtons.push(<span key="ellipsis-end" className="px-2">...</span>);
      pageButtons.push(
        <button
          key={totalPages}
          onClick={() => handlePageClick(totalPages)}
          className={`px-4 py-2 rounded mx-1 ${page === totalPages ? 'bg-blue-600 text-white' : 'bg-white text-blue-600 border border-blue-600'}`}
        >
          {totalPages}
        </button>
      );
    }

    return pageButtons;
  };

  return (
    <div>
      <h1 className="lg:m-10 text-2xl font-semibold my-4 dark:text-darkText">Active Loans Overview</h1>
      <section className="border-0 rounded-xl bg-white dark:bg-darkComponent shadow-md lg:mx-10 p-2">
        <div className="grid grid-cols-3 lg:grid-cols-7 font-medium text-sky-300 dark:text-darkAccent min-h-7 items-center border-b dark:border-darkPage mt-2 px-2">
          <div className="hidden md:block">Sl NO</div>
          <div>Loan Amount</div>
          <div>Left To Repay</div>
          <div className="hidden md:block">Duration</div>
          <div className="hidden md:block">Interest Rate</div>
          <div className="hidden md:block">Installment</div>
          <div className="justify-self-center">Repay</div>
        </div>
  
        {isLoadingLoans ? (
          <div className="text-center my-4 dark:text-darkText">Loading loans...</div>
        ) : (
          loansData?.data.content.map((loan, index) => (
            <div
              key={loan.serialNumber}
              className="grid grid-cols-3 lg:grid-cols-7 border-b dark:border-darkPage min-h-12 items-center dark:text-darkText"
            >
              <div className="hidden md:block">{(page - 1) * 10 + index + 1}</div>
              <div>${loan.loanAmount}</div>
              <div>${loan.amountLeftToRepay}</div>
              <div className="hidden md:block">{loan.duration} months</div>
              <div className="hidden md:block">{loan.interestRate}%</div>
              <div className={`hidden md:block`}>${loan.installment.toFixed(2)}/month</div>
              <div className="border border-blue-200 dark:border-darkAccent text-center p-1 rounded-lg justify-self-center hover:border-blue-700 dark:hover:border-darkAccent w-24 cursor-pointer">
                Repay
              </div>
            </div>
          ))
        )}
      </section>
  
      <div className="flex justify-end mt-4 mr-10 items-center">
        <button
          onClick={() => handlePageClick(page - 1)}
          disabled={page === 1}
          className={`px-4 py-2 rounded mx-1 flex items-center ${page === 1 ? 'text-gray-400 dark:text-gray-600' : 'text-blue-600 dark:text-darkAccent'}`}
        >
          <FaArrowLeft className="mr-2" />
          Previous
        </button>
        {renderPageButtons()}
        <button
          onClick={() => handlePageClick(page + 1)}
          disabled={page === (loansData?.data.totalPages || 1)}
          className={`px-4 py-2 rounded mx-1 flex items-center ${page === (loansData?.data.totalPages || 1) ? 'text-gray-400 dark:text-gray-600' : 'text-blue-600 dark:text-darkAccent'}`}
        >
          Next
          <FaArrowRight className="ml-2" />
        </button>
      </div>
    </div>
  );
  
};

export default ActiveLoans;
