import React from "react";
import { ArrowUpCircleIcon } from "@heroicons/react/24/outline";
import Transaction from "@/app/(route)/transaction/page";

interface LoanTableProps {
  transactions: TransactionValue[];
}

const Recent = ({ transactions }: LoanTableProps) => {
  return (
    <div className="space-y-7 my-6">
      <h3 className="font-semibold text-[22px] text-[#343C6A]">
        Recent Transactions
      </h3>
      <div className="flex gap-16 px-1 border-b text-[16px] text-[#718EBF] font-semibold">
        <p className="border-[#1814F3] border-b-[3px] pb-2 text-[#1814F3]">
          All Transactions
        </p>
        <p>Income</p>
        <p>Expense</p>
      </div>
      <div className="w-full bg-white rounded-[25px] px-8 py-6">
        <table className="border-separate border-spacing-y-4 font-[16px] w-full transaction-table">
          <thead>
            <tr className="text-[#718EBF] text-left">
              <th>Description</th>
              <th>Transaction ID</th>
              <th>Type</th>
              <th>Card</th>
              <th>Date</th>
              <th>Amount</th>
              <th>Receipt</th>
            </tr>
          </thead>
          <tbody className="text-[#232323] p-8 space-y-4">
            {transactions.map((transaction, index) => (
              <tr key={index}>
                <td className="flex gap-2 items-center">
                  <ArrowUpCircleIcon className="transaction-icon" />
                  {transaction.description}
                </td>
                <td>{transaction.id}</td>
                <td>{transaction.type}</td>
                <td>{transaction.card}</td>
                <td>{transaction.date}</td>
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
                <td>
                  <p className="table-button"> Download </p>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default Recent;
