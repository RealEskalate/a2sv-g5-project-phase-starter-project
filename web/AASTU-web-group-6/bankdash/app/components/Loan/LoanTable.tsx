"use client";
import React, { useEffect, useState } from "react";
import Pagination from "../../components/Pagination";
import { LoanType } from "@/types/LoanValue";
import loanApi from "@/app/Services/api/loanApi";
import { useSession } from "next-auth/react";
import ApplyLoan from "./ApplyForLoan";

const LoanTable = () => {
  const [currentPage, setCurrentPage] = useState(0);
  const [loading, setLoading] = useState(true);
  const [currentLoan, setCurentLoan] = useState<LoanType[]>([]);
  const [sortCriteria, setSortCriteria] = useState<string | null>(null);
  const [sortOrder, setSortOrder] = useState<"asc" | "desc">("asc");
  const { data: session } = useSession();

  const accessToken = session?.accessToken as string;

  const sortLoans = (
    loans: LoanType[],
    criteria: string,
    order: "asc" | "desc"
  ) => {
    return loans.sort((a, b) => {
      let valueA, valueB;

      switch (criteria) {
        case "loanAmount":
          valueA = a.loanAmount;
          valueB = b.loanAmount;
          break;
        case "amountLeftToRepay":
          valueA = a.amountLeftToRepay;
          valueB = b.amountLeftToRepay;
          break;
        case "interestRate":
          valueA = a.interestRate;
          valueB = b.interestRate;
          break;
        case "installment":
          valueA = a.installment;
          valueB = b.installment;
          break;
        default:
          return 0;
      }

      return order === "asc" ? valueA - valueB : valueB - valueA;
    });
  };

  const fetchLoan = async () => {
    while (!accessToken) {
      await new Promise((resolve) => setTimeout(resolve, 120)); // Delay to wait for the token
    }
    let loanData = await loanApi.getLoan(accessToken, currentPage);

    if (sortCriteria) {
      loanData = sortLoans(loanData, sortCriteria, sortOrder);
    } else {
      loanData = sortLoans(loanData, "loanAmount", "desc");
    }

    setCurentLoan(loanData);
    setLoading(false);
  };

  useEffect(() => {
    setLoading(true);
    fetchLoan();
  }, [currentPage, session]);

  const handleSort = (criteria: string) => {
    const newOrder =
      sortCriteria === criteria && sortOrder === "asc" ? "desc" : "asc";
    setSortCriteria(criteria);
    setSortOrder(newOrder);

    const sortedLoans = sortLoans([...currentLoan], criteria, newOrder);
    setCurentLoan(sortedLoans);
  };

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
        <div className="flex justify-between">
          <h3 className="font-semibold text-[22px] text-[#343C6A] dark:text-gray-300">
            Active Loans Overview
          </h3>
          <ApplyLoan />
        </div>
        <div className="overflow-x-auto custom-scrollbar dark:bg-[#232328] w-full">
          <table className="border-separate border-spacing-y-4 font-[16px] w-full detail-table  min-w-[1000px] sm:min-w-full">
            <thead>
              <tr className="text-[#718EBF] text-left dark:text-gray-200">
                <th className="hidden lg:table-cell">SL No</th>
                <th
                  className="cursor-pointer"
                  onClick={() => handleSort("loanAmount")}
                >
                  Loan Money{" "}
                  {sortCriteria === "loanAmount" &&
                    (sortOrder === "asc" ? "↑" : "↓")}
                </th>
                <th
                  className="cursor-pointer"
                  onClick={() => handleSort("amountLeftToRepay")}
                >
                  Left to repay{" "}
                  {sortCriteria === "amountLeftToRepay" &&
                    (sortOrder === "asc" ? "↑" : "↓")}
                </th>
                <th className="hidden lg:table-cell">Duration</th>
                <th
                  className="hidden md:table-cell cursor-pointer"
                  onClick={() => handleSort("interestRate")}
                >
                  Interest rate{" "}
                  {sortCriteria === "interestRate" &&
                    (sortOrder === "asc" ? "↑" : "↓")}
                </th>
                <th
                  className="hidden md:table-cell cursor-pointer"
                  onClick={() => handleSort("installment")}
                >
                  Installment{" "}
                  {sortCriteria === "installment" &&
                    (sortOrder === "asc" ? "↑" : "↓")}
                </th>
                <th className="hidden md:table-cell">Status</th>
                <th>Repay</th>
              </tr>
            </thead>
            <tbody className="text-[#232323] p-8 space-y-4 dark:text-gray-300 z-1">
              {loading
                ? Array.from({ length: 5 }).map((_, index) => (
                    <tr key={index} className="animate-pulse">
                      <td className="flex gap-2 items-center">
                        <div className="h-6 bg-gray-300 rounded dark:bg-gray-600 w-3/4" />
                      </td>
                      <td className="hidden md:table-cell">
                        <div className="h-6 bg-gray-300 rounded dark:bg-gray-600 w-3/4" />
                      </td>
                      <td className="hidden md:table-cell">
                        <div className="h-6 bg-gray-300 rounded dark:bg-gray-600 w-3/4" />
                      </td>
                      <td className="hidden lg:table-cell">
                        <div className="h-6 bg-gray-300 rounded dark:bg-gray-600 w-3/4" />
                      </td>
                      <td className="hidden lg:table-cell">
                        <div className="h-6 bg-gray-300 rounded dark:bg-gray-600 w-3/4" />
                      </td>
                      <td>
                        <div className="h-6 bg-gray-300 rounded dark:bg-gray-600 w-3/4" />
                      </td>
                      <td className="hidden md:table-cell">
                        <div className="h-6 bg-gray-300 rounded dark:bg-gray-600 w-3/4" />
                      </td>
                      <td className="hidden md:table-cell">
                        <div className="h-6 bg-gray-300 rounded dark:bg-gray-600 w-3/4" />
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
                      <td
                        className={`hidden md:table-cell status-button ${
                          detail.activeLoneStatus === "approved"
                            ? "bg-[#96d754]" // Green for success
                            : detail.activeLoneStatus === "pending"
                            ? "bg-[#ffc87a]" // Yellow for pending
                            : detail.activeLoneStatus === "reject"
                            ? "bg-[#d9534f]" // Red for reject
                            : "bg-gray-300" // Default background if none of the statuses match
                        }`}
                      >
                        {detail.activeLoneStatus.charAt(0).toUpperCase() +
                          detail.activeLoneStatus.slice(1)}
                      </td>
                      <td>
                        <p className="table-button">Replay</p>
                      </td>
                    </tr>
                  ))}
              {!loading && (
                <tr className="font-medium text-[#FE5C73] text-[16px]">
                  <td className="hidden lg:table-cell">Total</td>
                  <td>${totalLoanAmount}</td>
                  <td>${totalLeftToRepay}</td>
                  <td className="hidden lg:table-cell">-</td>
                  <td className="hidden md:table-cell">-</td>
                  <td className="hidden md:table-cell">
                    ${Math.round(totalInstallment)}
                  </td>
                  <td className="hidden md:table-cell"></td>
                  <td className="hidden md:table-cell">-</td>
                </tr>
              )}
            </tbody>
          </table>
        </div>
      </div>
      <Pagination updatePage={updatePage} start={false} />
    </div>
  );
};

export default LoanTable;
