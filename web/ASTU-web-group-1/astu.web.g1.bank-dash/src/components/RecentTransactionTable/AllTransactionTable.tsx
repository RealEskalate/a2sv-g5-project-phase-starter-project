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
import RecentTransctionSkeleton from "../AllSkeletons/RecentTransactionSkeleton/recentTransactionSkeleton";

const AllTransactionTable = () => {
  const [currentPage, setCurrentPage] = useState(0);

  const { data, isLoading, isFetching } = useGetAllTransactionsQuery({
    page: currentPage,
    size: 5,
  });

  if (isLoading || isFetching) {
    return <RecentTransctionSkeleton />; // Display loading state
  }

  const currentData = data?.data.content;
  const totalPages = data?.data.totalPages || 0;

  return (
    <div>
      <div className="flex flex-col gap-4">
        {currentData?.length == 0 ? (
          <div className="p-5">No transactions found.</div>
        ) : (
          <div className="relative overflow-x-auto bg-white px-2 pt-3 md:pt-6 rounded-2xl md:rounded-2xl">
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
                        amount={datax.type}
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
                        datax.type !== "deposit" ? "text-candyPink" : "text-mintGreen"
                      }`}
                    >
                      {datax.type !== "deposit"
                        ? "-$" + (datax.amount)
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
      {totalPages && (
        <Pagination
          totalPages={totalPages}
          currentPage={currentPage}
          setCurrentPage={setCurrentPage}
        />
      )}
    </div>
  );
};

export default AllTransactionTable;
