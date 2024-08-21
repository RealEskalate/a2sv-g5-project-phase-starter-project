'use client';
import React, { useEffect } from 'react';
import { FaRegArrowAltCircleUp, FaRegArrowAltCircleDown } from 'react-icons/fa';
import { useGetMyLoansQuery, useGetLoanDetailDataQuery } from '@/lib/redux/api/loansApi';
import { useDispatch } from 'react-redux';
import { setLoans } from '@/lib/redux/slices/loansSlice';

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

const ActiveLoans = () => {
  const { data: loansData, isLoading: isLoadingLoans } = useGetMyLoansQuery();
  const {data:detailLoansData, isLoading: isLoadingDetail} = useGetLoanDetailDataQuery() 
  const dispatch = useDispatch();

  useEffect(() => {
    if (loansData && loansData.success) {
      dispatch(setLoans(loansData.data));
    }
  }, [loansData, dispatch]);
  useEffect(()=>{

  }, [detailLoansData])

  return (
    <div>
      <h1 className="lg:m-10 text-2xl font-semibold my-4">Active Loans Overview</h1>
      <section className="border-0 rounded-xl bg-white shadow-md lg:mx-10 p-2">
        <div className="grid grid-cols-3 lg:grid-cols-7 font-medium text-sky-300 min-h-7 items-center border-b mt-2 px-2">
          <div className="hidden md:block">Sl NO</div>
          <div>Loan Amount</div>
          <div>Left To Repay</div>
          <div className="hidden md:block">Duration</div>
          <div className="hidden md:block">Interest Rate</div>
          <div className="hidden md:block">Installment</div>
          <div className="justify-self-center">Repay</div>
        </div>

        {isLoadingLoans ? (
          <div className="text-center my-4">Loading loans...</div>
        ) : (
          loansData?.data.map((loan, index) => (
            <div
              key={loan.serialNumber}
              className="grid grid-cols-3 lg:grid-cols-7 border-b min-h-12 items-center"
            >
              <div className="hidden md:block">{index + 1}</div>
              <div>${loan.loanAmount}</div>
              <div>${loan.amountLeftToRepay}</div>
              <div className="hidden md:block">{loan.duration} months</div>
              <div className="hidden md:block">{loan.interestRate}%</div>
              <div className={`hidden md:block`}>${loan.installment.toFixed(2)}/month</div>
              <div className="border border-blue-200 text-center p-1 rounded-lg justify-self-center hover:border-blue-700 w-24 cursor-pointer">
                Repay
              </div>
            </div>
          ))
        )}
      </section>
    </div>
  );
};

export default ActiveLoans;
