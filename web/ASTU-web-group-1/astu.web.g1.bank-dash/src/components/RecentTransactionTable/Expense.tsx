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

const Expense = () => {
  const [currentPage, setCurrentPage] = useState(0);


  const { data, error, isLoading } = useGetTransactionExpenseQuery({ page: currentPage, size: 5 });
  console.log(data);

  if (isLoading) {
    return <div>Loading...</div>; // Display loading state
  }

  if (error || !data) {
    return <div>Opps Something happens</div>; // Display error state
  }
  const currentData = data?.data.content;
  const totalPages = data?.data.totalPages;


  return (
    <div>
      <div className="flex flex-col gap-4">
        
        {!currentData || currentData?.length == 0 ? (
          <div>No Expense found.</div>
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
                {currentData?.map((datax, index) => (
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
                        ? "-$" + -(datax.amount)
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

export default Expense;