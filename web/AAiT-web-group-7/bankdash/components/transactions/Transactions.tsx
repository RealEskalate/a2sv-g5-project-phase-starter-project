// This coomponent will be refactored after we have the API
"use client";
import { useState } from "react";
import BalanceCard from "../commonalities/BalanceCard";
import ExpenseChart from "./ExpenseChart";
import Pagination from "./Pagenation";

const Transactions = () => {
  const [selected, setSelected] = useState("All Transactions");
  const [currentPage, setCurrentPage] = useState(1);
  const totalPages = 10;

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };
  const handleClick = (link: string) => {
    setSelected(link);
  };

  return (
    <div className="w-full h-fit bg-[#F5F7FA] flex flex-col gap-6 p-6">
      <div className="flex flex-wrap gap-6">
        <div className="flex flex-col gap-4">
          <div className="flex justify-between">
            <h5 className="text-[22px] font-semibold text-[#343C6A]">
              My Cards
            </h5>
            <button className="text-[17px] font-semibold text-[#343C6A]">
              + Add Card
            </button>
          </div>
          <div className="flex gap-5">
            <BalanceCard property="blue" />
            <BalanceCard property="white" />
          </div>
        </div>
        <div className="flex flex-col gap-4">
          <div className="text-[22px] font-semibold text-[#343C6A]">
            My Expenses
          </div>
          <div className="w-[300px] h-[225px] ">
            <ExpenseChart />
          </div>
        </div>
      </div>

      <h5 className="text-xl font-semibold text-[#343C6A]">
        Recent Transactions
      </h5>
      <div className="flex gap-10 ">
        <div className="flex flex-col gap-1 justify-center">
          <div
            className={`text-md text-center cursor-pointer ${
              selected === "All Transactions"
                ? "text-[#1814F3]"
                : "text-[#718EBF]"
            }`}
            onClick={() => handleClick("All Transactions")}
          >
            All Transactions
          </div>
          {selected === "All Transactions" && (
            <div className="h-[3px] rounded-t-[10px] bg-[#1814F3]"></div>
          )}
        </div>
        <div className="flex flex-col gap-1 justify-center">
          <div
            className={`text-md text-center cursor-pointer ${
              selected === "Incomes" ? "text-[#1814F3]" : "text-[#718EBF]"
            }`}
            onClick={() => handleClick("Incomes")}
          >
            Incomes
          </div>
          {selected === "Incomes" && (
            <div className="h-[3px] rounded-t-[10px] bg-[#1814F3]"></div>
          )}
        </div>
        <div className="flex flex-col gap-1 justify-center">
          <div
            className={`text-md text-center cursor-pointer ${
              selected === "Expenses" ? "text-[#1814F3]" : "text-[#718EBF]"
            }`}
            onClick={() => handleClick("Expenses")}
          >
            Expenses
          </div>
          {selected === "Expenses" && (
            <div className="h-[3px] rounded-t-[10px] bg-[#1814F3]"></div>
          )}
        </div>
      </div>
      <div className="p-8 rounded-[25px] bg-[#FFFF]">
        <table className="w-full h-[397px]">
          <thead className="text-[#718EBF]">
            <tr className="text-left">
              <th className="text-[16px] font-medium text-[#718EBF]">
                Description
              </th>
              <th className="text-[16px] font-medium text-[#718EBF]">
                Transaction ID
              </th>
              <th className="text-[16px] font-medium text-[#718EBF]">Type</th>
              <th className="text-[16px] font-medium text-[#718EBF]">Card</th>
              <th className="text-[16px] font-medium text-[#718EBF]">Date</th>
              <th className="text-[16px] font-medium text-[#718EBF]">Amount</th>
              <th className="text-[16px] font-medium text-[#718EBF]">
                Receipt
              </th>
            </tr>
          </thead>
          <tbody>
            {Array(4)
              .fill(null)
              .map((_, index) => (
                <tr key={index} className="text-sm">
                  <td>
                    <div className="flex items-center">
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        strokeWidth="1.5"
                        stroke="currentColor"
                        className="text-[#718EBF] w-[30px] h-[30px]"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          d="m9 12.75 3 3m0 0 3-3m-3 3v-7.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
                        />
                      </svg>
                      <p className="text-[16px] text-[#232323]">
                        Spotify Subscription
                      </p>
                    </div>
                  </td>
                  <td>
                    <p className="text-[16px] text-[#232323]">#123546435</p>
                  </td>
                  <td>
                    <p className="text-[16px] text-[#232323]">Shopping</p>
                  </td>
                  <td>
                    <p className="text-[16px] text-[#232323]">123***</p>
                  </td>
                  <td>
                    <p className="text-[16px] text-[#232323]">
                      28 Jan, 12:30AM
                    </p>
                  </td>
                  <td className="text-red-700">-$2,500</td>
                  <td>
                    <button className="w-[100px] h-[35px] rounded-full border border-[#123288] text-[#1814F3]">
                      Download
                    </button>
                  </td>
                </tr>
              ))}
          </tbody>
        </table>
      </div>
      <div className="flex justify-end  rounded-[10px] text-[#1814F3]">
        <Pagination
          currentPage={currentPage}
          totalPages={totalPages}
          onPageChange={handlePageChange}
        />
      </div>
    </div>
  );
};

export default Transactions;
