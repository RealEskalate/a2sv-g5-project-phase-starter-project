import { useUser } from '@/contexts/UserContext';
import React, { useEffect, useState } from 'react'
import { ExpenseTable } from './ExpenseTable';
import { TransactionContent, TransactionData } from '@/types';
import { getallTransactions, getExpenses, getIncomes } from '@/lib/api';
import { MdNavigateNext } from "react-icons/md";
import { GrFormPrevious } from "react-icons/gr";
export const TransactionTable = () => {
  const { isDarkMode } = useUser();
  const [activeTab, setActiveTab] = useState<"all" | "income" | "expense">(
    "all"
  );
  const [currentPage, setCurrentPage] = useState(1);
    const [transactions, setTransactions] = useState<TransactionContent[]>([]);
    const [totalPages, setTotalPages] = useState<number>(5);
      const rowsPerPage = 5;
  // Fetch transactions based on active tab and current page
  useEffect(() => {
    const fetchData = async () => {
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
      } catch (error) {
        console.error("Error fetching transactions:", error);
      }
    };
    fetchData();
  }, [currentPage, activeTab]);
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
    const pagesToShow = Math.min(totalPages, 4);
    const startPage =
      currentPage <= 2 || totalPages <= 4
        ? 1
        : currentPage > totalPages - 2
        ? totalPages - 3
        : currentPage - 1;

    return Array.from({ length: pagesToShow }, (_, index) => {
      const page = startPage + index;
      return (
        <button
          key={page}
          onClick={() => handlePageChange(page)}
          className={` ${isDarkMode ? "bg-gray-800 text-white" : "bg-white "}${
            (page === currentPage) && (!isDarkMode) ? " bg-blue-600" : "text-blue-600"
          } hover:bg-blue-700 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 me-2 mb-2`}
        >
          {page}
        </button>
      );
    });
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
        <ExpenseTable transactions={transactions} tab={activeTab} />
      </div>
      <div className="flex justify-end items-center space-x-2">
        <button
          onClick={handlePreviousPage}
          className={`flex rounded-xl ${
            currentPage === 1 ? "text-gray-400" : "text-blue-600"
          } ${isDarkMode ? "bg-gray-800" : "bg-white"}`}
          disabled={currentPage === 1}
        >
          <GrFormPrevious size={30} />
          <p
            className={` m-2  ${
              currentPage === 1 ? "text-gray-400" : "text-blue-600"
            } ${isDarkMode ? "text-gray-300" : "text-gray-900"}`}
          >
            Previous
          </p>
        </button>

        {renderPageButtons()}

        <button
          onClick={handleNextPage}
          className={`flex rounded-xl ${
            currentPage === totalPages ? "text-gray-400" : "text-blue-600"
          } ${isDarkMode ? "bg-gray-800" : "bg-white"}`}
          disabled={currentPage === totalPages}
        >
          <p
            className={`m-2 ${
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
}
