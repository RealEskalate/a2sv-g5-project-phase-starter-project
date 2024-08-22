"use client";
import React, { useEffect, useState } from "react";
import TableButton from "../TableButton/TableButton";
import RecentTransactionDescription from "./RecentTransactionDescription";
import { ChevronLeft, ChevronRight } from "lucide-react";
import {
  useGetAllTransactionsQuery,
  useGetTransactionExpenseQuery,
  useGetTransactionIncomeQuery,
} from "@/lib/redux/slices/transactionSlice";
import Pagination from "./Pagination";
import { TransactionDataType } from "@/types/transaction.types";
import { number } from "zod";

const RecentTransactionTable = () => {
  const [currentPage, setCurrentPage] = useState(0);
  const [currentButton, setCurrentButton] = useState("all-transaction");

  let queryResult;
  if (currentButton === "all-transaction") {
    queryResult = useGetAllTransactionsQuery({ page: currentPage, size: 5 });
  } else if (currentButton === "income") {
    queryResult = useGetTransactionIncomeQuery({ page: currentPage, size: 5 });
  } else {
    queryResult = useGetTransactionExpenseQuery({ page: currentPage, size: 5 });
  }

  const { data, error, isLoading } = queryResult;
  console.log(data);

  if (isLoading) {
    return <div>Loading...</div>; // Display loading state
  }

  const currentData = data?.data.content;
  const totalPages = data?.data.totalPages;
  if (!currentData) {
    return <div>No transactions found.</div>; // Handle case when no data is returned
  }

  const handelAllTransaction = () => {
    setCurrentButton("all-transaction");
  };

  const handelIncome = () => {
    setCurrentButton("income");
  };

  const handelExpense = () => {
    setCurrentButton("expense");
  };

  return (
    <div>
      <div className="flex flex-col gap-4">
        <h1 className="text-16px md:text-18px xl:text-22px text-[#333B69] font-semibold">
          Recent Transaction
        </h1>
        <div className="flex flex-row max-[640px]:justify-between  md:gap-[60px] lg:gap-[80px] text-blue-steel">
          <div className="flex flex-col gap-2">
            <button
              className={`border-none px-[11px] whitespace-nowrap ${
                currentButton === "all-transaction" && "text-blue-bright"
              }`}
              onClick={handelAllTransaction}
            >
              All Transactions
            </button>
            {currentButton === "all-transaction" && (
              <div className="bg-blue-bright h-1 rounded-t-full"></div>
            )}
          </div>
          <div className="flex flex-col gap-2">
            <button
              className={`border-none px-[11px] whitespace-nowrap ${
                currentButton === "income" && "text-blue-bright"
              }`}
              onClick={handelIncome}
            >
              Incomes
            </button>
            {currentButton === "income" && (
              <div className="bg-blue-bright h-1 rounded-t-full"></div>
            )}
          </div>
          <div className="flex flex-col gap-2">
            <button
              className={`border-none px-[11px] whitespace-nowrap ${
                currentButton === "expense" && "text-blue-bright"
              }`}
              onClick={handelExpense}
            >
              Expense
            </button>
            {currentButton === "expense" && (
              <div className="bg-blue-bright h-1 rounded-t-full"></div>
            )}
          </div>
        </div>
        {currentData.length == 0 ? (
          <div>No transactions found.</div>
        ) : (
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
                  <th scope="col" className="hidden lg:table-cell pb-2">
                    Type
                  </th>
                  <th scope="col" className="hidden lg:table-cell pb-2">
                    Card
                  </th>
                  <th scope="col" className="hidden lg:table-cell pb-2">
                    Date
                  </th>
                  <th scope="col" className="hidden md:table-cell pb-2">
                    Amount
                  </th>
                  <th scope="col" className="hidden lg:table-cell pb-2 w-fit">
                    Recipt
                  </th>
                </tr>
              </thead>
              <tbody className="text-12px xl:text-16px text-gray-dark cursor-pointer  hover:bg-gray-100 dark:hover:bg-gray-700">
                {currentData.map((datax, index) => (
                  <tr
                    key={index}
                    className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-100 dark:hover:bg-gray-700"
                  >
                    <td className="py-3">
                      <RecentTransactionDescription
                        amount={datax.amount}
                        description={datax.description}
                      />
                    </td>
                    <td className="hidden md:table-cell py-3">
                      {datax.transactionId}
                    </td>
                    <td className="hidden lg:table-cell py-3">{datax.type}</td>
                    <td className="hidden lg:table-cell py-3">
                      {"1234" + " ****"}
                    </td>
                    <td className="hidden lg:table-cell py-3">{datax.date}</td>
                    <td
                      className={`py-3 ${
                        datax.amount < 0 ? "text-candyPink" : "text-mintGreen"
                      }`}
                    >
                      {datax.amount < 0
                        ? "-$" + datax.amount
                        : "+$" + datax.amount}
                    </td>
                    <td className="hidden lg:table-cell py-3 w-24 md:w-32">
                      <TableButton
                        text="Download"
                        classname="px-6 text-[#123288] border-[#123288]"
                      />
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
      {typeof totalPages === "number" && totalPages > 0 && (
        <Pagination
          totalPages={totalPages ? totalPages : 0}
          currentPage={currentPage}
          setCurrentPage={setCurrentPage}
        />
      )}
    </div>
  );
};

export default RecentTransactionTable;