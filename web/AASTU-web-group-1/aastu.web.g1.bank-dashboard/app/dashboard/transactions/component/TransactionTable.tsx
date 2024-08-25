import { useUser } from "@/contexts/UserContext";
import React, { useEffect, useState } from "react";
import { ExpenseTable } from "./ExpenseTable";
import { TransactionContent, TransactionData } from "@/types";
import { getallTransactions, getExpenses, getIncomes } from "@/lib/api";
import { MdNavigateNext } from "react-icons/md";
import { GrFormPrevious } from "react-icons/gr";
import { ShimmerRow } from "../../_components/Shimmer";

export const TransactionTable = ({
  onLoadingComplete,
}: {
  onLoadingComplete: any;
}) => {
  const { isDarkMode } = useUser();
  const [activeTab, setActiveTab] = useState<"all" | "income" | "expense">(
    "all"
  );
  const [currentPage, setCurrentPage] = useState(1);
  const [transactions, setTransactions] = useState<TransactionContent[]>([]);
  const [totalPages, setTotalPages] = useState<number>(5);
  const [loading, setLoading] = useState(true);
  const [dataFetched, setDataFetched] = useState(false);
  const rowsPerPage = 5;

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      onLoadingComplete(true);

      try {
        let data: TransactionData | undefined;
        if (activeTab === "all") {
          data = await getallTransactions(currentPage - 1, rowsPerPage);
        } else if (activeTab === "income") {
          data = await getIncomes(currentPage - 1, rowsPerPage);
        } else {
          data = await getExpenses(currentPage - 1, rowsPerPage);
        }
        setTransactions(data?.content || []);
        setTotalPages(data?.totalPages || 7);
        setDataFetched(true); // Mark data as fetched
        onLoadingComplete(false);
      } catch (error) {
        console.error("Error fetching transactions:", error);
      } finally {
        setLoading(false);
      }
    };

    // Only fetch data if not already fetched
    if (!dataFetched) {
      fetchData();
    }
  }, [currentPage, activeTab, onLoadingComplete, dataFetched]);

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  const handleNextPage = () => {
    if (currentPage < totalPages) {
      setCurrentPage(currentPage + 1);
    }
  };

  const handlePreviousPage = () => {
    if (currentPage > 1) {
      setCurrentPage(currentPage - 1);
    }
  };

  const renderPageButtons = () => {
    const pageButtons = [];
    const maxVisiblePages = 3;

    const startPage = Math.max(currentPage - 1, 1);
    const endPage = Math.min(startPage + maxVisiblePages - 1, totalPages - 1);

    if (currentPage <= maxVisiblePages) {
      for (let page = 1; page <= endPage; page++) {
        pageButtons.push(
          <button
            key={page}
            onClick={() => handlePageChange(page)}
            className={`hover:bg-blue-700 hover:text-white ${
              page !== currentPage
                ? isDarkMode
                  ? "text-gray-300   bg-gray-900"
                  : "text-blue-600  bg-white"
                : isDarkMode
                ? "bg-blue-600 text-white"
                : "bg-blue-600 text-white"
            } 
            focus:ring-4 ${
              isDarkMode ? "focus:ring-gray-500" : "focus:ring-blue-300"
            } font-medium rounded-lg text-sm px-4 py-2 me-2 mb-2 shadow-lg`}
          >
            {page}
          </button>
        );
      }
    } else {
      for (let page = startPage; page <= endPage; page++) {
        pageButtons.push(
          <button
            key={page}
            onClick={() => handlePageChange(page)}
            className={`hover:bg-blue-700 hover:text-white ${
              page === currentPage
                ? "bg-blue-600 text-white"
                : isDarkMode
                ? "text-gray-300 "
                : "text-blue-600 bg-white "
            }
            focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 me-2 mb-2 shadow-lg`}
          >
            {page}
          </button>
        );
      }
    }

    if (endPage < totalPages - 1) {
      pageButtons.push(
        <span key="dots-end" className="text-gray-500">
          ...
        </span>
      );
    }

    pageButtons.push(
      <button
        key={totalPages}
        onClick={() => handlePageChange(totalPages)}
        className={` text-white${
          currentPage === totalPages
            ? "bg-blue-600 "
            : " hover:bg-blue-700 hover:text-white"
        } ${isDarkMode ? "bg-gray-800" : "bg-white"} 
        focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 me-2 mb-2 shadow-lg`}
      >
        {totalPages}
      </button>
    );

    return pageButtons;
  };

  return (
    <>
      <div className="space-y-5 w-[100%] items-center my-5">
        <div className={`space-y-5 `}>
          <h4 className="font-inter  font-semibold lg:text-[22px] md:text-lg text-base">
            Recent Transactions
          </h4>
          <div
            className={`md:space-x-20 space-x-5 flex flex-grow border-b-2 ${
              isDarkMode ? "border-[#ebeef23a]" : "border-[#EBEEF2]"
            }`}
          >
            <button onClick={() => setActiveTab("all")}>
              <p
                className={`border-b-4  font-inter  font-medium lg:text-base text-xs
                   ${
                     activeTab === "all"
                       ? isDarkMode
                         ? "border-b-blue-600 text-gray-300" // Dark mode when activeTab is "all"
                         : "border-b-blue-600 text-[#1814F3]" // Light mode when activeTab is "all"
                       : "border-b-transparent text-[#718EBF]"
                   }
                  
                `}
              >
                All Transactions
              </p>
            </button>
            <button onClick={() => setActiveTab("income")}>
              <p
                className={`border-b-4 font-medium lg:text-base text-xs  ${
                  activeTab === "income"
                    ? isDarkMode
                      ? "border-b-blue-600 text-gray-300" // Dark mode when activeTab is "all"
                      : "border-b-blue-600 text-[#1814F3]" // Light mode when activeTab is "all"
                    : "border-b-transparent text-[#718EBF]"
                }`}
              >
                Income
              </p>
            </button>
            <button onClick={() => setActiveTab("expense")}>
              <p
                className={`border-b-4 font-medium lg:text-base text-xs  ${
                  activeTab === "expense"
                    ? isDarkMode
                      ? "border-b-blue-600 text-gray-300" // Dark mode when activeTab is "all"
                      : "border-b-blue-600 text-[#1814F3]" // Light mode when activeTab is "all"
                    : "border-b-transparent text-[#718EBF]"
                } `}
              >
                Expense
              </p>
            </button>
          </div>
        </div>
        {loading && !dataFetched ? (
          <div>
            {Array.from({ length: rowsPerPage }).map((_, index) => (
              <ShimmerRow key={index} />
            ))}
          </div>
        ) : (
          <ExpenseTable transactions={transactions} tab={activeTab} />
        )}
      </div>
      <div className="flex justify-end items-center space-x-2 ">
        <button
          onClick={handlePreviousPage}
          className={`flex items-center justify-center rounded-xl shadow-lg px-2 py-1 me-2 mb-2  ${
            currentPage === 1
              ? "text-gray-400"
              : isDarkMode
              ? "text-gray-300  hover:text-white hover:bg-blue-600"
              : "text-blue-600 hover:text-white hover:bg-blue-600"
          } ${isDarkMode ? "bg-gray-800" : "bg-white"}`}
          disabled={currentPage === 1}
        >
          <GrFormPrevious size={30} />
          Previous
        </button>

        {renderPageButtons()}

        <button
          onClick={handleNextPage}
          className={`flex items-center justify-center rounded-xl shadow-lg px-2 py-1 me-2 mb-2 ${
            currentPage === totalPages
              ? "text-gray-400 "
              : isDarkMode
              ? "text-gray-300  hover:text-white hover:bg-blue-600"
              : "text-blue-600  hover:text-white hover:bg-blue-600"
          } ${isDarkMode ? "bg-gray-800" : "bg-white"} hover:bg-blue-100`}
          disabled={currentPage === totalPages}
        >
          Next
          <MdNavigateNext size={30} />
        </button>
      </div>
    </>
  );
};
