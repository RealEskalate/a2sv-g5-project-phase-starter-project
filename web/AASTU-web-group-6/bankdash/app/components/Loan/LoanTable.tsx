"use client";
import React, { useEffect, useState } from "react";
import Pagination from "../../components/Pagination";
import { LoanType } from "@/types/LoanValue";
import loanApi from "@/app/Services/api/loanApi";
import { useSession } from "next-auth/react";

const LoanTable = () => {
  const [currentPage, setCurrentPage] = useState(0);
  const [loading, setLoading] = useState(true);

  const [currentLoan, setCurentLoan] = useState<LoanType[]>([]);
  const { data: session } = useSession();

  const accessToken = session?.accessToken as string;

  const fetchLoan = async () => {
    while (!accessToken) {
      await new Promise((resolve) => setTimeout(resolve, 120)); // Delay to wait for the token
    }

    const loanData = await loanApi.getLoan(accessToken, currentPage);
    setCurentLoan(loanData);
    setLoading(false);
  };

  useEffect(() => {
    setLoading(true);
    fetchLoan();
  }, [currentPage, session]);

  const totalLoanAmount = currentLoan.reduce(
    (sum, loan) => sum + loan.loanAmount,
    0
  );
  const totalLeftToRepay = currentLoan.reduce(
    (sum, loan) => sum + loan.amountLeftToRepay,
    0
  );
  const totalInstallment = currentLoan.reduce(
    (sum, loan) => sum + loan.installment,
    0
  );
  const updatePage = (newPage: number = 0) => {
    setCurrentPage(newPage);
  };
  return (
    <div className="space-y mb-8">
      <div className="w-full bg-white rounded-[25px] px-8 py-6 dark:bg-[#232328]">
        <h3 className="font-semibold text-[22px] text-[#343C6A] dark:text-gray-300">
          Active Loans Overview
        </h3>
        <div className="overflow-x-auto custom-scrollbar dark:bg-[#232328] w-full">
          <table className="border-separate border-spacing-y-4 font-[16px] w-full detail-table  min-w-[1000px] sm:min-w-full">
            <thead>
              <tr className="text-[#718EBF] text-left dark:text-gray-200">
                <th className="hidden lg:table-cell">SL No</th>
                <th className="">Loan Money</th>
                <th className="">Left to repay</th>
                <th className="hidden lg:table-cell">Duration</th>
                <th className="hidden md:table-cell">Interest rate</th>
                <th className="hidden md:table-cell">Installment</th>
                <th>Repay</th>
              </tr>
            </thead>
            <tbody className="text-[#232323] p-8 space-y-4 dark:text-gray-300">
              {loading
                ? Array.from({ length: 5 }).map((_, index) => (
                    <tr key={index} className="animate-pulse">
                      <td className="flex gap-2 items-center">
                        <div className="w-4 h-4 bg-gray-300 rounded-sm dark:bg-gray-600" />
                        <div className="h-4 bg-gray-300 rounded dark:bg-gray-600 w-3/4" />
                      </td>
                      <td className="hidden md:table-cell">
                        <div className="h-4 bg-gray-300 rounded dark:bg-gray-600 w-3/4" />
                      </td>
                      <td className="hidden md:table-cell">
                        <div className="h-4 bg-gray-300 rounded dark:bg-gray-600 w-3/4" />
                      </td>
                      <td className="hidden lg:table-cell">
                        <div className="h-4 bg-gray-300 rounded dark:bg-gray-600 w-3/4" />
                      </td>
                      <td className="hidden lg:table-cell">
                        <div className="h-4 bg-gray-300 rounded dark:bg-gray-600 w-3/4" />
                      </td>
                      <td>
                        <div className="h-4 bg-gray-300 rounded dark:bg-gray-600 w-3/4" />
                      </td>
                      <td className="hidden md:table-cell">
                        <div className="h-4 bg-gray-300 rounded dark:bg-gray-600 w-3/4" />
                      </td>
                    </tr>
                  ))
                : currentLoan.map((detail, index) => (
                    <tr key={index}>
                      <td className="hidden lg:table-cell">
                        {detail.serialNumber}
                      </td>
                      <td>${detail.loanAmount}</td>
                      <td>${Math.round(detail.amountLeftToRepay)}</td>
                      <td className="hidden lg:table-cell">
                        {detail.duration} Months
                      </td>
                      <td className="hidden md:table-cell">
                        {detail.interestRate}%
                      </td>
                      <td className="hidden md:table-cell">
                        ${Math.round(detail.installment)}
                      </td>
                      <td>
                        <p className="table-button">Replay</p>
                      </td>
                    </tr>
                  ))}
              <tr className="font-medium text-[#FE5C73] text-[16px]">
                <td className="hidden lg:table-cell">Total</td>
                <td>${totalLoanAmount}</td>
                <td>${totalLeftToRepay}</td>
                <td className="hidden lg:table-cell">-</td>
                <td className="hidden md:table-cell">-</td>
                <td className="hidden md:table-cell">
                  ${Math.round(totalInstallment)}
                </td>
                <td className="hidden md:table-cell">-</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <Pagination updatePage={updatePage} start={false} />
    </div>
  );
};

export default LoanTable;
