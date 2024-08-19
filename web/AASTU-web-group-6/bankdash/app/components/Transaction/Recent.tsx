import React from "react";
import { ArrowUpCircleIcon } from "@heroicons/react/24/outline";
import fetchTransaction from "@/app/Services/api/transactionApi";
import { TransactionResponse, TransactionType } from "@/types/TransactionValue";

const Recent = ({ data }: { data: TransactionType[] }) => {
  return (
    <div className="space-y-7 my-6">
      <h3 className="font-semibold text-[22px] text-[#343C6A]">
        Recent Transactions
      </h3>
      <div className="flex gap-4 sm:gap-16 px-1 border-b text-[14px] sm:text-[16px] text-[#718EBF] font-semibold overflow-x-auto">
        <p className="border-[#1814F3] border-b-[3px] pb-2 text-[#1814F3] whitespace-nowrap">
          All Transactions
        </p>
        <p className="whitespace-nowrap">Income</p>
        <p className="whitespace-nowrap">Expense</p>
      </div>
      {/* Add a wrapper with overflow-x-auto to enable horizontal scrolling */}
      <div className="w-full bg-white rounded-[25px] px-8 py-6 overflow-x-auto custom-scrollbar">
        <table className="border-separate border-spacing-y-4 font-[16px] w-full min-w-[900px] transaction-table">
          <thead>
            <tr className="text-[#718EBF] text-left">
              <th>Description</th> {/* Always visible */}
              <th className="hidden md:table-cell">Transaction ID</th>{" "}
              {/* Hidden on mobile, visible on tablets and larger */}
              <th className="hidden md:table-cell">Type</th>{" "}
              {/* Hidden on mobile, visible on tablets and larger */}
              <th className="hidden md:table-cell">Card</th>{" "}
              {/* Hidden on mobile, visible on tablets and larger */}
              <th className="hidden md:table-cell">Date</th>{" "}
              {/* Hidden on mobile, visible on tablets and larger */}
              <th>Amount</th> {/* Always visible */}
              <th className="hidden md:table-cell">Receipt</th>{" "}
              {/* Hidden on mobile, visible on tablets and larger */}
            </tr>
          </thead>
          <tbody className="text-[#232323] p-8 space-y-4">
            {data.map((transaction: TransactionType, index: number) => (
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
                <td className="hidden md:table-cell">
                  {transaction.type}
                </td>{" "}
                {/* Hidden on mobile, visible on tablets and larger */}
                <td className="hidden md:table-cell">
                  {transaction.date}
                </td>{" "}
                {/* Hidden on mobile, visible on tablets and larger */}
                <td>
                  <p
                    className={
                      transaction.amount < 0
                        ? "text-[#FE5C73]"
                        : "text-[#28A745]"
                    }
                  >
                    {transaction.amount < 0 ? "-" : "+"}$
                    {Math.abs(transaction.amount)}
                  </p>
                </td>
                <td className="hidden md:table-cell">
                  <p className="table-button">Download</p>
                </td>{" "}
                {/* Hidden on mobile, visible on tablets and larger */}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default Recent;
