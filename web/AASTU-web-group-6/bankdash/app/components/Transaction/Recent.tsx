import React, { useState } from "react";
import { ArrowUpCircleIcon } from "@heroicons/react/24/outline";
import { TransactionType } from "@/types/TransactionValue";
import { useAppSelector } from "@/app/Redux/store/store";
import Download from "./Download";

const Recent = () => {
  const [selectedTab, setSelectedTab] = useState<"All" | "Income" | "Expense">(
    "All"
  );
  const TranData: TransactionType[] = useAppSelector(
    (state) => state.transactions.transactions
  );
  const expenseData: TransactionType[] = useAppSelector(
    (state) => state.transactions.expense
  );
  const incomeData: TransactionType[] = useAppSelector(
    (state) => state.transactions.income
  );

  const dataToDisplay =
    selectedTab === "Income"
      ? incomeData
      : selectedTab === "Expense"
      ? expenseData
      : TranData;

  return (
    <div className="space-y-7 my-6">
      <h3 className="font-semibold text-[22px] text-[#343C6A] dark:text-gray-300">
        Recent Transactions
      </h3>
      <div className="flex gap-4 sm:gap-16 px-1 border-b text-[14px] sm:text-[16px] text-[#718EBF] font-semibold overflow-x-auto">
        <button
          className={`dark:text-gray-300 ${
            selectedTab === "All"
              ? "border-[#1814F3] border-b-[3px] pb-2 text-[#1814F3] whitespace-nowrap"
              : "text-red-100"
          }`}
          onClick={() => setSelectedTab("All")}
        >
          All Transaction
        </button>
        <button
          className={`dark:text-gray-300 ${
            selectedTab === "Income"
              ? "border-[#1814F3] border-b-[3px] pb-2 text-[#1814F3] whitespace-nowrap"
              : "text-red-100"
          }`}
          onClick={() => setSelectedTab("Income")}
        >
          Income
        </button>
        <button
          className={`dark:text-gray-300 ${
            selectedTab === "Expense"
              ? "border-[#1814F3] border-b-[3px] pb-2 text-[#1814F3] whitespace-nowrap"
              : "text-red-100"
          }`}
          onClick={() => setSelectedTab("Expense")}
        >
          Expense
        </button>
      </div>
      {/* Add a wrapper with overflow-x-auto to enable horizontal scrolling */}
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
            {dataToDisplay.map(
              (transaction: TransactionType, index: number) => (
                <tr key={index}>
                  <td className="flex gap-2 items-center">
                    <ArrowUpCircleIcon className="transaction-icon" />
                    {transaction.description}
                  </td>
                  <td className="hidden md:table-cell">
                    {transaction.transactionId}
                  </td>{" "}
                  {/* Hidden on mobile, visible on tablets and larger */}
                  <td className="hidden md:table-cell">
                    {transaction.type}
                  </td>{" "}
                  {/* Hidden on mobile, visible on tablets and larger */}
                  <td className="hidden lg:table-cell">
                    {transaction.type}
                  </td>{" "}
                  {/* Hidden on mobile, visible on tablets and larger */}
                  <td className="hidden lg:table-cell">
                    {transaction.date}
                  </td>{" "}
                  {/* Hidden on mobile, visible on tablets and larger */}
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
    </div>
  );
};

export default Recent;
