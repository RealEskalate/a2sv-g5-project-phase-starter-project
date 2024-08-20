import React, { useEffect, useState } from "react";
import { CiSaveDown1, CiSaveUp1 } from "react-icons/ci";

import { TransactionData, UserData } from "@/types";
import { getCurrentUser } from "@/lib/api";



interface Props {
  transactions: TransactionData[];
  tab: string;
}

export const ExpenseTable: React.FC<Props> = ({ transactions, tab }: Props) => {
  const [currentUser, setCurrentUser] = useState<UserData | null>(null);

  useEffect(() => {
    getCurrentUser().then((user) => {
      if (user) {
        setCurrentUser(user);
      }
    });
  }, []);

  return (
    <div className="w-full border border-gray-300 rounded-3xl bg-white shadow-lg   ">
      <table className="w-full divide-y divide-gray-200 rounded-3xl overflow-hidden ">
        <thead className=" ">
          <tr className="border-b-2 ">
            <th
              scope="col"
              className="lg:px-5 py-3 md:px-2 text-left text-xs font-medium text-gray-500 uppercase truncate md:max-w-[6rem] lg:max-w-[10rem]"
            >
              Description
            </th>
            <th
              scope="col"
              className="lg:px-5 py-3 md:px-2 text-left text-xs font-medium text-gray-500 uppercase truncate hidden md:table-cell md:max-w-[6rem]"
            >
              Transaction ID
            </th>
            <th
              scope="col"
              className="lg:px-5 py-3 md:px-2 text-left text-xs font-medium text-gray-500 uppercase hidden md:table-cell md:max-w-[6rem]"
            >
              Type
            </th>
            <th
              scope="col"
              className="lg:px-5 py-3 md:px-2 text-left text-xs font-medium text-gray-500 uppercase hidden md:table-cell md:max-w-[6rem]"
            >
              Card
            </th>
            <th
              scope="col"
              className="lg:px-5 py-3 md:px-2 text-left text-xs font-medium text-gray-500 uppercase hidden md:table-cell md:max-w-[6rem]"
            >
              Date
            </th>
            <th
              scope="col"
              className="lg:px-5 py-3 md:px-2 text-left text-xs font-medium text-gray-500 uppercase md:max-w-[4rem]"
            >
              Amount
            </th>
            <th
              scope="col"
              className="lg:px-5 py-3 md:px-2 text-left text-xs font-medium text-gray-500 uppercase hidden md:table-cell md:max-w-[6rem]"
            >
              Receipt
            </th>
          </tr>
        </thead>
        <tbody className=" divide-y divide-gray-200">
          {transactions.map((transaction) => (
            <tr key={transaction.transactionId}>
              {/* Description */}
              <td className="lg:px-5 py-4 px-2 items-center text-sm truncate lg:max-w-[10rem] md:max-w-[6rem]">
                <div className="space-x-2 flex md:max-w-[6rem] lg:max-w-[10rem]">
                  <span className="inline-block align-middle">
                    {tab === "income" ||
                    transaction.senderUserName !== currentUser?.username ? (
                       <CiSaveDown1 size={20} />
                    ) : ( <CiSaveUp1 size={20} />
                    
                    )}
                  </span>
                  <span
                    className="inline-block align-middle truncate"
                    title={transaction.description}
                  >
                    {transaction.description}
                  </span>
                </div>
              </td>
              {/* ID */}
              <td className="lg:px-5 py-4 md:px-2 text-sm truncate hidden md:table-cell md:max-w-[4rem]">
                {transaction.transactionId}
              </td>
              {/* Type */}
              <td className="lg:px-5 py-4 md:px-2 text-sm truncate hidden md:table-cell md:max-w-[6rem]">
                {transaction.type}
              </td>
              {/* Receiver */}
              <td className="lg:px-5 py-4 md:px-2 text-sm truncate hidden md:table-cell md:max-w-[6rem]">
                {transaction.receiverUserName !== null
                  ? transaction.receiverUserName
                  : "unknown"}
              </td>
              {/* Date */}
              <td className="lg:px-5 py-4 md:px-2 text-sm truncate hidden md:table-cell md:max-w-[6rem]">
                {transaction.date}
              </td>
              {/* Amount */}
              <td className="lg:px-5 py-4 md:px-2 text-sm truncate md:max-w-[4rem]">
                <p
                  className={
                    tab === "income" ||
                    transaction.senderUserName !== currentUser?.username
                      ? "text-green-500"
                      : "text-red-500"
                  }
                >
                  ${transaction.amount}
                </p>
              </td>
              {/* Receipt */}
              <td className="lg:px-5 py-4 md:px-2 text-xs  hidden md:table-cell md:max-w-[3rem]">
                <button className="bg-transparent hover:bg-gradient-to-r hover:from-blue-400 hover:to-blue-600 text-blue-700 font-semibold py-2 px-2 border border-blue-500 rounded-full shadow-md hover:shadow-lg transition-all duration-300 ease-in-out transform hover:-translate-y-1 trunicate">
                  Download
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};
