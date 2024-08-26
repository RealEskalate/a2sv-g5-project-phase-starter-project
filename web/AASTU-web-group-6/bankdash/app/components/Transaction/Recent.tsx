"use client";
import React, { useState, useEffect } from "react";
import { ArrowUpCircleIcon } from "@heroicons/react/24/outline";
import { TransactionType } from "@/types/TransactionValue";
import Download from "./Download";
import Pagination from "../Pagination";
import {
  getExpense,
  getIncome,
  getTransaction,
} from "@/app/Services/api/fetchTransaction";
import { useSession } from "next-auth/react";

const Recent = () => {
  const { data: session } = useSession();
  const accessToken = session?.accessToken as string;

  const [expenseData, setExpenseData] = useState<TransactionType[]>([]);
  const [incomeData, setIncomeData] = useState<TransactionType[]>([]);
  const [allData, setAllData] = useState<TransactionType[]>([]);
  const [selectedTab, setSelectedTab] = useState<"All" | "Income" | "Expense">(
    "All"
  );
  const [currentPage, setCurrentPage] = useState(0);
  const [reset, setReset] = useState(false);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    setReset(!reset);
  }, [selectedTab]);

  const fetchData = async () => {
    while (!accessToken) {
      await new Promise((resolve) => setTimeout(resolve, 100)); // Delay to wait for the token
    }

    setLoading(true);
    if (selectedTab === "Expense") {
      await fetchExpense();
    } else if (selectedTab === "Income") {
      await fetchIncome();
    } else if (selectedTab === "All") {
      await fetchAllTransactions();
    }
  };

  useEffect(() => {
    fetchData();
  }, [selectedTab, currentPage, accessToken]);

  const fetchExpense = async () => {
    const res = await getExpense(currentPage, accessToken);
    setExpenseData(res);
    setLoading(false);
  };

  const fetchIncome = async () => {
    const res = await getIncome(currentPage, accessToken);
    setIncomeData(res);
    setLoading(false);
  };

  const fetchAllTransactions = async () => {
    const res = await getTransaction(currentPage, accessToken);
    setAllData(res);
    setLoading(false);
  };

  const updatePage = (newPage: number) => {
    setCurrentPage(newPage);
  };

  const dataToDisplay =
    selectedTab === "Income"
      ? incomeData
      : selectedTab === "Expense"
      ? expenseData
      : allData;
  // console.log(dataToDisplay, "dis");
  return (
    <div className="space-y-7 my-6">
      <h3 className="font-semibold text-[22px] text-[#343C6A] dark:text-gray-300">
        Recent Transactions
      </h3>
      {/* Tab Selection */}
      <div className="flex gap-4 sm:gap-16 px-1 border-b text-[14px] sm:text-[16px] text-[#718EBF] font-semibold overflow-x-auto">
        {/* Tab Buttons */}
        <button
          className={`dark:text-gray-300 ${
            selectedTab === "All"
              ? "border-[#1814F3] border-b-[3px] pb-2 text-[#1814F3] whitespace-nowrap"
              : "text-[#232323]"
          }`}
          onClick={() => setSelectedTab("All")}
        >
          All Transactions
        </button>
        <button
          className={`dark:text-gray-300 ${
            selectedTab === "Income"
              ? "border-[#1814F3] border-b-[3px] pb-2 text-[#1814F3] whitespace-nowrap"
              : "text-[#232323]"
          }`}
          onClick={() => setSelectedTab("Income")}
        >
          Income
        </button>
        <button
          className={`dark:text-gray-300 ${
            selectedTab === "Expense"
              ? "border-[#1814F3] border-b-[3px] pb-2 text-[#1814F3] whitespace-nowrap"
              : "text-[#232323]"
          }`}
          onClick={() => setSelectedTab("Expense")}
        >
          Expense
        </button>
      </div>

      {/* Transaction Table */}
      <div className="w-full bg-white dark:bg-[#232328] rounded-[25px] px-8 py-6 overflow-x-auto custom-scrollbar">
        <table className="border-separate border-spacing-y-4 font-[16px] w-full min-w-[1000px] transaction-table sm:min-w-full">
          <thead>
            <tr className="text-[#718EBF] text-left dark:text-gray-200">
              <th>Description</th>
              <th className="hidden md:table-cell">Transaction ID</th>
              <th className="hidden md:table-cell">Type</th>
              <th className="hidden lg:table-cell">Card</th>
              <th className="hidden lg:table-cell">Date</th>
              <th>Amount</th>
              <th className="hidden md:table-cell">Receipt</th>
            </tr>
          </thead>
          <tbody className="text-[#232323] dark:text-gray-300 p-8 space-y-4">
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
              : dataToDisplay.map(
                  (transaction: TransactionType, index: number) => (
                    <tr key={index}>
                      <td className="flex gap-2 items-center">
                        <ArrowUpCircleIcon className="transaction-icon" />
                        {transaction.description}
                      </td>
                      <td className="hidden md:table-cell">
                        {transaction.transactionId}
                      </td>
                      <td className="hidden md:table-cell">
                        {transaction.type}
                      </td>
                      <td className="hidden lg:table-cell">
                        {transaction.type}
                      </td>
                      <td className="hidden lg:table-cell">
                        {transaction.date}
                      </td>
                      <td>
                        <p
                          className={
                            selectedTab === "Expense" ||
                            transaction.type !== "deposit"
                              ? "text-[#FE5C73]"
                              : "text-[#28A745] dark:text-green-400"
                          }
                        >
                          {selectedTab === "Expense" ||
                          transaction.type !== "deposit"
                            ? "-"
                            : "+"}
                          ${Math.abs(transaction.amount)}
                        </p>
                      </td>
                      <td className="hidden md:table-cell md:text-white">
                        <Download
                          transactionId={transaction.transactionId}
                          transaction={transaction}
                        />
                      </td>
                    </tr>
                  )
                )}
          </tbody>
        </table>
      </div>

      {/* Pagination */}
      <Pagination updatePage={updatePage} start={reset} />
    </div>
  );
};

export default Recent;
