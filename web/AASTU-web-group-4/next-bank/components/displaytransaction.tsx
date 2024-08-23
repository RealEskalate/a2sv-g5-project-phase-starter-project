import React from "react";
import { FaArrowCircleUp, FaArrowCircleDown } from "react-icons/fa";

import Pagination from "@/components/Pagination";
const displaytransaction = (alltransaction: any, type: any) => {
  return (
    <>
      {alltransaction.length == 0 ? (
        <> 
          <table className="min-w-full bg-white rounded-lg shadow-md border border-gray-200">
            <thead className="bg-blue-50">
              <tr>
                <th className="p-4 text-left text-sm font-semibold text-gray-700">
                  Description
                </th>
                <th className="p-4 text-left text-sm font-semibold text-gray-700">
                  Transaction ID
                </th>
                <th className="p-4 text-left text-sm font-semibold text-gray-700">
                  Type
                </th>
                <th className="p-4 text-left text-sm font-semibold text-gray-700">
                  Card
                </th>
                <th className="p-4 text-left text-sm font-semibold text-gray-700">
                  Date
                </th>
                <th className="p-4 text-left text-sm font-semibold text-gray-700">
                  Amount
                </th>
                <th className="p-4 text-left text-sm font-semibold text-gray-700">
                  Receipt
                </th>
              </tr>
            </thead>
          </table>

          <div className="  p-8 text-center text-gray-500 text-xl font-semibold">
            No data to display
          </div>
        </>
      ) : (
        <div>
          {/* Transactions Table for Desktop/Tablets */}
          <div className="hidden md:block overflow-x-auto">
            <table className="min-w-full bg-white rounded-lg shadow-md border border-gray-200">
              <thead className="bg-blue-50">
                <tr className="dark:text-blue-500 text-gray-700">
                  <th className="p-4 text-left text-sm font-semibold">
                    Description
                  </th>
                  <th className="p-4 text-left text-sm font-semibold">
                    Transaction ID
                  </th>
                  <th className="p-4 text-left text-sm font-semibold">
                    Type
                  </th>
                  <th className="p-4 text-left text-sm font-semibold">
                    Card
                  </th>
                  <th className="p-4 text-left text-sm font-semibold">
                    Date
                  </th>
                  <th className="p-4 text-left text-sm font-semibold">
                    Amount
                  </th>
                  <th className="p-4 text-left text-sm font-semibold">
                    Receipt
                  </th>
                </tr>
              </thead>
              <tbody>
                {alltransaction.map((transaction: any, index: any) => (
                  <tr
                    key={transaction.transactionId}
                    className={`border-b border-gray-200 ${
                      index % 2 === 0 ? "bg-gray-50" : "bg-white"
                    } hover:bg-gray-100 transition-colors duration-300 dark:bg-dark text-gray-900 dark:text-white`}
                  >
                    <td className="p-4 flex items-center text-sm text-gray-700 dark:text-white truncate">
                      {transaction.type === "deposit" || type === "income" ? (
                        <FaArrowCircleDown className="text-green-500 text-lg mr-2" />
                      ) : (
                        <FaArrowCircleUp className="text-red-500 text-lg mr-2" />
                      )}
                      {transaction.description}
                    </td>
                    <td className="p-4 text-sm dark:text-white">
                      {transaction.transactionId}
                    </td>
                    <td className="p-4 text-sm dark:text-white">
                      {transaction.type}
                    </td>
                    <td className="p-4 text-sm dark:text-white">Card Name</td>
                    <td className="p-4 text-sm dark:text-white">
                      {transaction.date}
                    </td>
                    <td
                      className={`p-4 text-sm ${
                        transaction.type === "deposit"
                          ? "text-green-500"
                          : "text-red-500"
                      }`}
                    >
                      {transaction.type === "deposit" ? (
                        <div>+{transaction.amount}$</div>
                      ) : (
                        <div>-{transaction.amount}$</div>
                      )}
                    </td>
                    <td className="p-4">
                      <button className="text-blue-500 text-sm hover:underline">
                        Download
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>

          {/* Transactions for Mobile View */}
          <div className="block md:hidden">
            {alltransaction.map((transaction: any) => (
              <div
                key={transaction.transactionId}
                className="flex justify-between bg-white p-4 mb-2 rounded-lg shadow-sm border border-gray-200 items-center dark:bg-dark text-gray-900 dark:text-white"
              >
                <div>
                  <div className="flex items-center mb-2">
                    {transaction.type === "deposit" ? (
                      <FaArrowCircleDown className="text-green-500 text-xl mr-2" />
                    ) : (
                      <FaArrowCircleUp className="text-red-500 text-xl mr-2" />
                    )}
                    <span className="font-semibold">
                      {transaction.description}
                    </span>
                  </div>
                  <div className="text-[12px] text-gray-400 mb-1 pl-5">
                    {transaction.date}
                  </div>
                </div>
                <div>
                  <div
                    className={`font-bold ${
                      transaction.type === "deposit"
                        ? "text-green-500"
                        : "text-red-500"
                    }`}
                  >
                    {transaction.type === "deposit" ? (
                      <div>+{transaction.amount}$</div>
                    ) : (
                      <div>-{transaction.amount}$</div>
                    )}
                  </div>
                </div>
              </div>
            ))}
          </div>

          {/* Pagination */}
        </div>
      )}
    </>
  );
};

export default displaytransaction;
