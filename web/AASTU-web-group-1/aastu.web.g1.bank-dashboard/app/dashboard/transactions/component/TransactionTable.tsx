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
    const maxVisiblePages = 3; // Number of pages to show before the ellipsis

    // Determine the range of pages to show
    const startPage = Math.max(currentPage - 1, 1);
    const endPage = Math.min(startPage + maxVisiblePages - 1, totalPages - 1);

    // Always show the first page when it's close enough
    if (currentPage <= maxVisiblePages) {
      for (let page = 1; page <= endPage; page++) {
        pageButtons.push(
          <button
            key={page}
            onClick={() => handlePageChange(page)}
            className={`${
              page !== currentPage
                ? isDarkMode
                  ? "text-gray-300 hover:bg-gray-700 hover:text-white bg-gray-900"
                  : "text-blue-600 hover:bg-blue-700 hover:text-white bg-white"
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
            className={`${
              page === currentPage
                ? "bg-blue-600 text-white"
                : isDarkMode
                ? "text-gray-300 hover:bg-gray-700 hover:text-white"
                : "text-blue-600 bg-white hover:bg-blue-700 hover:text-white"
            }
            focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 me-2 mb-2 shadow-lg`}
          >
            {page}
          </button>
        );
      }
    }

    // Always show the last page
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
        className={`${
          currentPage === totalPages
            ? "bg-blue-600 text-white"
            : "text-blue-600 hover:bg-blue-700 hover:text-white"
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
      <div className="space-y-5 w-[90%] items-center">
        <div className="space-y-5">
          <h4>Recent Transactions</h4>
          <div className="space-x-5 flex">
            <button onClick={() => setActiveTab("all")}>
              <p
                className={`border-b-2 text-xs ${
                  activeTab === "all"
                    ? "border-b-blue-600"
                    : "border-b-transparent"
                } ${isDarkMode ? "text-gray-300" : "text-gray-700"}`}
              >
                All Transactions
              </p>
            </button>
            <button onClick={() => setActiveTab("income")}>
              <p
                className={`border-b-2 text-xs ${
                  activeTab === "income"
                    ? "border-b-blue-600"
                    : "border-b-transparent"
                } ${isDarkMode ? "text-gray-300" : "text-gray-700"}`}
              >
                Income
              </p>
            </button>
            <button onClick={() => setActiveTab("expense")}>
              <p
                className={`border-b-2 text-xs ${
                  activeTab === "expense"
                    ? "border-b-blue-600"
                    : "border-b-transparent"
                } ${isDarkMode ? "text-gray-300" : "text-gray-700"}`}
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
      <div className="flex justify-end items-center space-x-2 m-5">
        <button
          onClick={handlePreviousPage}
          className={`flex items-center justify-center rounded-xl shadow-lg p-1 ${
            currentPage === 1
              ? "text-gray-400"
              : isDarkMode
              ? "text-gray-300 hover:bg-blue-600 hover:text-white"
              : "text-blue-600 hover:bg-blue-400 hover:text-black"
          } ${isDarkMode ? "bg-gray-800" : "bg-white"}`}
          disabled={currentPage === 1}
        >
          <GrFormPrevious size={30} />
          <p
            className={` ${
              currentPage === 1
                ? "text-gray-400"
                : isDarkMode
                ? "text-gray-300"
                : "text-blue-600"
            }`}
          >
            Previous
          </p>
        </button>

        {renderPageButtons()}

        <button
          onClick={handleNextPage}
          className={`flex items-center justify-center rounded-xl shadow-lg p-1 ${
            currentPage === totalPages
              ? "text-gray-400 "
              : isDarkMode
              ? "text-gray-300 hover:bg-blue-600 hover:text-white"
              : "text-blue-600 hover:bg-blue-400 hover:text-black"
          } ${isDarkMode ? "bg-gray-800" : "bg-white"} hover:bg-blue-100`}
          disabled={currentPage === totalPages}
        >
          <p
            className={` ${
              currentPage === totalPages ? "text-gray-400" : "text-blue-600"
            } ${isDarkMode ? "text-gray-300" : "text-gray-900"}`}
          >
            Next
          </p>
          <MdNavigateNext size={30} />
        </button>
      </div>
    </>
  );
};
