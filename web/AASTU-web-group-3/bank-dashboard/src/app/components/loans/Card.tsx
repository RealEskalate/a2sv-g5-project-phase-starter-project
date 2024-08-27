"use client";
import React, { ReactNode } from "react";
import { FaUser, FaBriefcase, } from "react-icons/fa";
import { FaScrewdriverWrench } from "react-icons/fa6";
import { useGetLoanDetailDataQuery } from "@/lib/redux/api/loansApi";
import Loading from "@/app/loading";

interface LoanProps {
  name: string;
  amount: number;
  color: string;
  icon: ReactNode;
}

const Card = () => {
  const { data: detailLoans, isLoading: isLoadingDetail } =
    useGetLoanDetailDataQuery();

  if (isLoadingDetail) {
    return <Loading />;
  }

  if (!detailLoans || !detailLoans.success) {
    return <div className="text-center">Failed to load loan details</div>;
  }

  const { personalLoan, businessLoan, corporateLoan } = detailLoans.data;

  return (
    <div className="card-holder flex gap-16 px-10 py-4">
      <div className="flex w-64 border-0 rounded-xl bg-white dark:bg-darkComponent min-h-32 gap-3 items-center">
        <div className="icons border-1 rounded-full ml-4 bg-gray-100 dark:bg-darkPage h-16 w-16 flex items-center justify-center">
          <FaUser color="#396AFF" size={30} />
        </div>
        <div className="info">
          <p className="text-[#718EBF] dark:text-darkText mt-2">Personal Loans</p>
          <p className="font-semibold text-xl dark:text-darkText">${personalLoan}</p>
        </div>
      </div>
  
      <div className="flex w-64 border-0 rounded-xl bg-white dark:bg-darkComponent min-h-32 gap-3 items-center">
        <div className="icons border-1 rounded-full ml-4 bg-[#FFF5D9] dark:bg-darkPage h-16 w-16 flex items-center justify-center">
          <FaBriefcase color="#FFBB38" size={30} />
        </div>
        <div className="info">
          <p className="text-[#718EBF] dark:text-darkText mt-2">Business Loans</p>
          <p className="font-semibold text-xl dark:text-darkText">${businessLoan}</p>
        </div>
      </div>
  
      <div className="flex w-64 border-0 rounded-xl bg-white dark:bg-darkComponent min-h-32 gap-3 items-center">
        <div className="icons border-1 rounded-full ml-4 bg-[#FFE0EB] dark:bg-darkPage h-16 w-16 flex items-center justify-center">
          <FaUser color="#FF82AC" size={30} />
        </div>
        <div className="info">
          <p className="text-[#718EBF] dark:text-darkText mt-2">Corporate Loans</p>
          <p className="font-semibold text-xl dark:text-darkText">${corporateLoan}</p>
        </div>
      </div>
  
      <div className="flex w-64 border-0 rounded-xl bg-white dark:bg-darkComponent min-h-32 gap-3 items-center">
        <div className="icons border-1 rounded-full ml-4 bg-[#DCFAF8] dark:bg-darkPage h-16 w-16 flex items-center justify-center">
          <FaScrewdriverWrench color="#16DBCC" size={30} />
        </div>
        <div className="info">
          <p className="text-[#718EBF] dark:text-darkText mt-2">Custom Loan</p>
          <p className="font-semibold text-xl dark:text-darkText">Choose Money</p>
        </div>
      </div>
    </div>
  );
  
};

export default Card;
