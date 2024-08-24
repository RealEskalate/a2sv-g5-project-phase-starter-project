import React, { useEffect, useState } from "react";
import { CiSaveDown1, CiSaveUp1 } from "react-icons/ci";

import { TransactionContent, UserData } from "@/types";
import { getCurrentUser } from "@/lib/api";
import { useUser } from "@/contexts/UserContext";

interface Props {
  transactions: TransactionContent[];
  tab: string;
}

export const ExpenseTable: React.FC<Props> = ({ transactions, tab }: Props) => {
  const [currentUser, setCurrentUser] = useState<UserData | null>(null);
  const { isDarkMode } = useUser();

  useEffect(() => {
    getCurrentUser().then((user) => {
      if (user) {
        setCurrentUser(user);
      }
    });
  }, []);

  return (
    <div
      className={`w-full border ${
        isDarkMode ? "border-gray-700 bg-gray-800" : "border-gray-300 bg-white"
      } rounded-3xl shadow-lg`}
    >
      <table
        className={`w-full divide-y ${
          isDarkMode ? "divide-gray-600" : "divide-gray-200"
        } rounded-3xl overflow-hidden`}
      >
        <thead>
          <tr
            className={`${
              isDarkMode ? "border-b-2 border-gray-700" : "border-b-2"
            }`}
          >
            <th
              scope="col"
              className={`lg:px-5 py-3 md:px-2 text-left text-xs font-medium uppercase truncate md:max-w-[6rem] lg:max-w-[10rem] ${
                isDarkMode ? "text-gray-400" : "text-gray-500"
              }`}
            >
              Description
            </th>
            <th
              scope="col"
              className={`lg:px-5 py-3 md:px-2 text-left text-xs font-medium uppercase truncate hidden md:table-cell md:max-w-[6rem] ${
                isDarkMode ? "text-gray-400" : "text-gray-500"
              }`}
            >
              Transaction ID
            </th>
            <th
              scope="col"
              className={`lg:px-5 py-3 md:px-2 text-left text-xs font-medium uppercase hidden md:table-cell md:max-w-[6rem] ${
                isDarkMode ? "text-gray-400" : "text-gray-500"
              }`}
            >
              Type
            </th>
            <th
              scope="col"
              className={`lg:px-5 py-3 md:px-2 text-left text-xs font-medium uppercase hidden md:table-cell md:max-w-[6rem] ${
                isDarkMode ? "text-gray-400" : "text-gray-500"
              }`}
            >
              Card
            </th>
            <th
              scope="col"
              className={`lg:px-5 py-3 md:px-2 text-left text-xs font-medium uppercase hidden md:table-cell md:max-w-[6rem] ${
                isDarkMode ? "text-gray-400" : "text-gray-500"
              }`}
            >
              Date
            </th>
            <th
              scope="col"
              className={`lg:px-5 py-3 md:px-2 text-left text-xs font-medium uppercase md:max-w-[4rem] ${
                isDarkMode ? "text-gray-400" : "text-gray-500"
              }`}
            >
              Amount
            </th>
            <th
              scope="col"
              className={`lg:px-5 py-3 md:px-2 text-left text-xs font-medium uppercase hidden md:table-cell md:max-w-[6rem] ${
                isDarkMode ? "text-gray-400" : "text-gray-500"
              }`}
            >
              Receipt
            </th>
          </tr>
        </thead>
        <tbody
          className={`divide-y ${
            isDarkMode ? "divide-gray-700" : "divide-gray-200"
          }`}
        >
          {transactions.map((transaction) => (
            <tr key={transaction.transactionId}>
              {/* Description */}
              <td
                className={`lg:px-5 py-4 px-2 items-center text-sm truncate lg:max-w-[10rem] md:max-w-[6rem] ${
                  isDarkMode ? "text-gray-300" : "text-gray-700"
                }`}
              >
                <div className="space-x-2 flex md:max-w-[6rem] lg:max-w-[10rem]">
                  <span className="inline-block align-middle">
                    {tab === "income" ||
                    transaction.senderUserName !== currentUser?.username ? (
                      <CiSaveDown1 size={20} color="red"/>
                    ) : (
                      <CiSaveUp1 size={20} color="green"/>
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
              <td
                className={`lg:px-5 py-4 md:px-2 text-sm truncate hidden md:table-cell md:max-w-[4rem] ${
                  isDarkMode ? "text-gray-300" : "text-gray-700"
                }`}
              >
                {transaction.transactionId}
              </td>
              {/* Type */}
              <td
                className={`lg:px-5 py-4 md:px-2 text-sm truncate hidden md:table-cell md:max-w-[6rem] ${
                  isDarkMode ? "text-gray-300" : "text-gray-700"
                }`}
              >
                {transaction.type}
              </td>
              {/* Receiver */}
              <td
                className={`lg:px-5 py-4 md:px-2 text-sm truncate hidden md:table-cell md:max-w-[6rem] ${
                  isDarkMode ? "text-gray-300" : "text-gray-700"
                }`}
              >
                {transaction.receiverUserName !== null
                  ? transaction.receiverUserName
                  : "unknown"}
              </td>
              {/* Date */}
              <td
                className={`lg:px-5 py-4 md:px-2 text-sm truncate hidden md:table-cell md:max-w-[6rem] ${
                  isDarkMode ? "text-gray-300" : "text-gray-700"
                }`}
              >
                {transaction.date}
              </td>
              {/* Amount */}
              <td
                className={`lg:px-5 py-4 md:px-2 text-sm truncate md:max-w-[4rem] ${
                  isDarkMode ? "text-gray-300" : "text-gray-700"
                }`}
              >
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
              <td
                className={`lg:px-5 py-4 md:px-2 text-xs hidden md:table-cell md:max-w-[3rem] ${
                  isDarkMode ? "text-gray-300" : "text-gray-700"
                }`}
              >
                <button className="bg-transparent hover:bg-gradient-to-r hover:from-blue-400 hover:to-blue-600 text-blue-700 font-semibold py-2 px-2 border border-blue-500 rounded-full shadow-md hover:shadow-lg transition-all duration-300 ease-in-out transform hover:-translate-y-1 truncate">
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
