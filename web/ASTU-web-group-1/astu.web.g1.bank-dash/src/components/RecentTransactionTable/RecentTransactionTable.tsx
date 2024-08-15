"use client";
import React, { useState } from "react";
import TableButton from "../TableButton/TableButton";
import RecentTransactionDescription from "./RecentTransactionDescription";
const data = [
  {
    id: 1,
    description: "Payment for order",
    transactionId: "TRX123456",
    type: "Credit",
    card: "123456",
    date: "2024-08-14",
    amount: 50.0,
    receiptUrl: "/download-receipt/TRX123456",
  },
  {
    id: 2,
    description: "Refund for order",
    transactionId: "TRX654321",
    type: "Debit",
    card: "567890",
    date: "2024-08-13",
    amount: -20.0,
    receiptUrl: "/download-receipt/TRX654321",
  },
  {
    id: 3,
    description: "Subscription renewal",
    transactionId: "TRX987654",
    type: "Credit",
    card: "987654",
    date: "2024-08-12",
    amount: 100.0,
    receiptUrl: "/download-receipt/TRX987654",
  },
  {
    id: 4,
    description: "Payment for order",
    transactionId: "TRX112233",
    type: "Credit",
    card: "432198",
    date: "2024-08-11",
    amount: 75.5,
    receiptUrl: "/download-receipt/TRX112233",
  },
  {
    id: 5,
    description: "Donation to Charity",
    transactionId: "TRX334455",
    type: "Debit",
    card: "334455",
    date: "2024-08-10",
    amount: 25.0,
    receiptUrl: "/download-receipt/TRX334455",
  },
  {
    id: 6,
    description: "Grocery Purchase",
    transactionId: "TRX556677",
    type: "Credit",
    card: "556677",
    date: "2024-08-09",
    amount: 120.75,
    receiptUrl: "/download-receipt/TRX556677",
  },
  {
    id: 7,
    description: "Online Course Payment",
    transactionId: "TRX778899",
    type: "Credit",
    card: "778899",
    date: "2024-08-08",
    amount: 200.0,
    receiptUrl: "/download-receipt/TRX778899",
  },
];

const RecentTransactionTable = () => {
  const [currentPage, setCurrentPage] = useState(1);
  const rowsPerPage = 5;

  // Calculate total pages
  const totalPages = Math.ceil(data.length / rowsPerPage);

  // Get current page data
  const currentData = data.slice(
    (currentPage - 1) * rowsPerPage,
    currentPage * rowsPerPage
  );

  const handlePreviousPage = () => {
    if (currentPage > 1) {
      setCurrentPage(currentPage - 1);
    }
  };

  const handleNextPage = () => {
    if (currentPage < totalPages) {
      setCurrentPage(currentPage + 1);
    }
  };

  return (
    <div>
      <div className="flex flex-col gap-4">
      <h1 className="text-16px md:text-18px xl:text-22px text-[#333B69] font-semibold">
        Recent Transaction
      </h1>
      <div className="relative overflow-x-auto bg-white px-4 md:px-6 pt-5 md:pt-6 rounded-2xl md:rounded-2xl">
        <table className="bg-white px-5 lg:px-11 w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
          <thead className=" text-12px md:text-16px font-Lato font-medium text-blue-steel bg-white border-b">
            <tr className="">
              <th scope="col" className="hidden md:table-cell pb-2">
                Description
              </th>
              <th scope="col" className=" hidden md:table-cell pb-2">
                Transacton ID
              </th>
              <th scope="col" className="hidden md:table-cell pb-2">
                Type
              </th>
              <th scope="col" className="hidden md:table-cell pb-2">
                Card
              </th>
              <th scope="col" className="hidden md:table-cell pb-2">
                Date
              </th>
              <th scope="col" className="hidden md:table-cell pb-2">
                Amount
              </th>
              <th scope="col" className="hidden md:table-cell pb-2 w-fit">
                Recipt
              </th>
            </tr>
          </thead>
          <tbody className="text-12px xl:text-16px text-gray-dark cursor-pointer  hover:bg-gray-100 dark:hover:bg-gray-700">
            {currentData.map((data, index) => (
              <tr
                key={index}
                className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-100 dark:hover:bg-gray-700"
              >
                <td className="py-3"><RecentTransactionDescription amount={data.amount} description={data.description}/></td>
                <td className="hidden md:table-cell py-3">{data.transactionId}</td>
                <td className="hidden md:table-cell py-3">{data.type}</td>
                <td className="hidden md:table-cell py-3">{data.card.slice(0,4) + " ****"}</td>
                <td className="hidden md:table-cell py-3">
                  {data.date}
                </td>
                <td className={`py-3 ${data.amount>0?"text-candyPink" : "text-mintGreen"}`}>
                  {data.amount>0?"+$"+data.amount:"-$"+data.amount}
                </td>
                <td className="hidden md:table-cell py-3 w-24 md:w-32">
                  <TableButton text="Download" classname="px-6 text-[#123288] border-[#123288]" />
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
      <nav
        className="flex items-center justify-end pt-4"
      >
        <ul className="inline-flex items-center -space-x-px text-sm h-8">
          <li>
            <button
              onClick={handlePreviousPage}
              disabled={currentPage === 1}
              className="flex items-center justify-center px-3 h-8 ms-0 leading-tight text-gray-500 bg-white border border-gray-300 rounded-l-lg hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
            >
              Previous
            </button>
          </li>
          {[...Array(totalPages)].map((_, index) => (
            <li key={index}>
              <button
                onClick={() => setCurrentPage(index + 1)}
                className={`flex items-center justify-center px-3 h-8 leading-tight ${
                  currentPage === index + 1
                    ? "text-blue-600 border border-gray-300 bg-blue-50 hover:bg-blue-100 hover:text-blue-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white"
                    : "text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
                }`}
              >
                {index + 1}
              </button>
            </li>
          ))}
          <li>
            <button
              onClick={handleNextPage}
              disabled={currentPage === totalPages}
              className="flex items-center justify-center px-3 h-8 leading-tight text-gray-500 bg-white border border-gray-300 rounded-r-lg hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white"
            >
              Next
            </button>
          </li>
        </ul>
      </nav>
    </div>
  );
};

export default RecentTransactionTable;
