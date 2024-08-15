import { Inter } from "next/font/google";
import Transaction, { TransactionType } from "../transaction/Transaction";
import React from "react";
interface TransactionTypeArray {
  transactions: TransactionType[];
}

const inter = Inter({ subsets: ["latin"] });
const TransactionsTable: React.FC<TransactionTypeArray> = ({transactions}) => {
  return (
    <div className="overflow-x-auto">
      <table className=" bg-white max-md:hidden rounded-3xl shadow-sm">
        <thead>
          <tr className={`text-left text-[#718EBF] ${inter.className}`}>
            <th className="py-3 px-6 font-medium">Description</th>
            <th className="py-3 px-6 font-medium">Transaction ID</th>
            <th className="py-3 px-6 font-medium">Type</th>
            <th className="py-3 px-6 font-medium">Card</th>
            <th className="py-3 px-6 font-medium">Date</th>
            <th className="py-3 px-6 font-medium">Amount</th>
            <th className="py-3 px-6 font-medium">Receipt</th>
          </tr>
        </thead>
        <tbody>
          {transactions.map((transaction: TransactionType) => (
            <Transaction {...transaction} key={transaction.id} />
          ))}
        </tbody>
      </table>

      {/** Mobile View */}
      <div className=" bg-white rounded-2xl pt-2 px-4 md:hidden">
        {transactions.map((transaction: TransactionType) => (
          <div
            key={transaction.id}
            className="border-b border-[#E6EFF5] py-4 flex items-center justify-between gap-1"
          >
            <div className="flex items-center gap-3">
              <img
                src={
                  transaction.amount < 0
                    ? "/assets/transaction/withdraw.svg"
                    : "/assets/transaction/deposit.svg"
                }
                alt="icon"
                className="w-6 h-6"
              />
              <div>
                <div className="font-semibold">{transaction.description}</div>
                <div className="text-sm text-gray-500">{transaction.date}</div>
              </div>
            </div>
            <div
              className={`text-sm ${
                transaction.amount < 0 ? "text-red-500" : "text-green-500"
              }`}
            >
              {transaction.amount < 0
                ? `-$${Math.abs(transaction.amount)}`
                : `+$${transaction.amount}`}
            </div>
          </div>
        ))}
        
      </div>
    </div>
  );
};

export default TransactionsTable;
