import React from 'react'
import { CiSaveDown1 } from "react-icons/ci";
import { CiSaveUp1 } from "react-icons/ci";

export interface TransactionProps {
  transactionId: string;
  type: "shopping" | "transfer" | "service" | "deposit";
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string;
}

export const ExpenseTable: React.FC<{ transactions: TransactionProps[] }> = ({ transactions }) => {
  return (
    <div className=" max-w-full  rounded-3xl shadow-md">
      <table className="w-full divide-y divide-gray-200 rounded-3xl  ">
        <thead className="bg-white ">
          <tr className="border-b-2">
            <th
              scope="col"
              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase  truncate max-w-[10rem] "
            >
              Description
            </th>
            <th
              scope="col"
              className="lg:px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase truncate max-w-[10rem] hidden md:table-cell"
            >
              Transaction ID
            </th>
            <th
              scope="col"
              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase hidden md:table-cell"
            >
              Type
            </th>
            <th
              scope="col"
              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase hidden md:table-cell"
            >
              Card
            </th>
            <th
              scope="col"
              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase hidden md:table-cell"
            >
              Date
            </th>
            <th
              scope="col"
              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase"
            >
              Amount
            </th>
            <th
              scope="col"
              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase hidden md:table-cell"
            >
              Receipt
            </th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {transactions.map((transaction) => (
            <tr key={transaction.transactionId} className="justify-between">
              {/* description */}
              <td className=" px-2 py-4  items-center text-sm">
                <div className="space-x-5 flex  max-w-[12rem]">
                  <span className="inline-block align-middle">
                    {transaction.amount > 0 ? (
                      <CiSaveUp1 size={20} />
                    ) : (
                      <CiSaveDown1 size={20} />
                    )}
                  </span>
                  <span
                    className="inline-block align-middle truncate max-w-[10rem]"
                    title={transaction.description}
                  >
                    {transaction.description}
                  </span>
                </div>
              </td>
              {/* Id */}
              <td className="lg:px-6 py-4 truncate max-w-[10rem] hidden md:table-cell text-sm">
                {transaction.transactionId}
              </td>
              {/* type */}
              <td className="lg:px-6 py-4 truncate max-w-[10rem] hidden md:table-cell text-sm">
                {transaction.type}
              </td>
              {/* card */}
              <td className="lg:px-6 py-4 truncate max-w-[10rem] hidden md:table-cell text-sm">
                {transaction.receiverUserName}
              </td>
              {/* date */}
              <td className="lg:px-6 py-4 truncate max-w-[10rem] hidden md:table-cell text-sm">
                {transaction.date}
              </td>
              {/* amount */}
              <td className="lg:px-6 py-4 truncate max-w-[10rem] text-sm">
                <p
                  className={`${
                    transaction.amount > 0 ? "text-green-500" : "text-red-500"
                  }`}
                >
                  ${transaction.amount}
                </p>
              </td>
              {/* receipt */}
              <td className="lg:px-2 py-4  truncate hidden md:table-cell">
                <button
                  className="bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-2 border border-blue-500 hover:border-transparent rounded-2xl
                    text-xs
                    "
                >
                  Download
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
