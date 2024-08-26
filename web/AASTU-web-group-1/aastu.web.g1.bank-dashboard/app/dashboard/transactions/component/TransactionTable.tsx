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
      setTotalPages(data?.totalPages || 1); 
      onLoadingComplete(false);
    } catch (error) {
      console.error("Error fetching transactions:", error);
    } finally {
      setLoading(false);
    }
  };


  fetchData();
}, [currentPage, activeTab, onLoadingComplete]);


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

   pageButtons.push(
     <button
       key={1}
       onClick={() => handlePageChange(1)}
       className={`hover:bg-blue-700 hover:text-white ${
         currentPage === 1
           ? "bg-blue-600 text-white"
           : isDarkMode
           ? "text-gray-300 bg-gray-900"
           : "text-blue-600"
       } 
      focus:ring-4 ${
        isDarkMode ? "focus:ring-gray-500" : "focus:ring-blue-300"
      } font-medium rounded-lg text-xs md:text-sm px-2 md:px-4 py-2 me-1 mb-1 md:me-2 md:mb-2 `}
     >
       1
     </button>
   );


   const startPage = Math.max(currentPage - 1, 2);
   const endPage = Math.min(startPage + maxVisiblePages - 1, totalPages - 1);
   if (startPage > 2) {
     pageButtons.push(
       <span key="dots-start" className="text-gray-500 text-xs md:text-sm">
         ...
       </span>
     );
   }

   for (let page = startPage; page <= endPage; page++) {
     pageButtons.push(
       <button
         key={page}
         onClick={() => handlePageChange(page)}
         className={`hover:bg-blue-700 hover:text-white ${
           page === currentPage
             ? "bg-blue-600 text-white"
             : isDarkMode
             ? "text-gray-300 bg-gray-900"
             : "text-blue-600 "
         } 
        focus:ring-4 ${
          isDarkMode ? "focus:ring-gray-500" : "focus:ring-blue-300"
        } font-medium rounded-lg text-xs md:text-sm px-2 md:px-4 py-2 me-1 mb-1 md:me-2 md:mb-2 `}
       >
         {page}
       </button>
     );
   }

   if (endPage < totalPages - 1) {
     pageButtons.push(
       <span key="dots-end" className="text-gray-500 text-xs md:text-sm">
         ...
       </span>
     );
   }

   if (totalPages > 1) {
     pageButtons.push(
       <button
         key={totalPages}
         onClick={() => handlePageChange(totalPages)}
         className={`hover:bg-blue-700 hover:text-white ${
           currentPage === totalPages
             ? "bg-blue-600 text-white"
             : isDarkMode
             ? "text-gray-300 bg-gray-900"
             : "text-blue-600 "
         } 
        focus:ring-4 ${
          isDarkMode ? "focus:ring-gray-500" : "focus:ring-blue-300"
        } font-medium rounded-lg text-xs md:text-sm px-2 md:px-4 py-2 me-1 mb-1 md:me-2 md:mb-2 `}
       >
         {totalPages}
       </button>
     );
   }

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
                         ? "border-b-blue-600 text-gray-300"
                         : "border-b-blue-600 text-[#1814F3]"
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
                      ? "border-b-blue-600 text-gray-300"
                      : "border-b-blue-600 text-[#1814F3]"
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
                      ? "border-b-blue-600 text-gray-300"
                      : "border-b-blue-600 text-[#1814F3]"
                    : "border-b-transparent text-[#718EBF]"
                } `}
              >
                Expense
              </p>
            </button>
          </div>
        </div>
        {loading ? (
          <div>
            {Array.from({ length: rowsPerPage }).map((_, index) => (
              <ShimmerRow key={index} />
            ))}
          </div>
        ) : (
          <ExpenseTable transactions={transactions} tab={activeTab} />
        )}
      </div>
      <div className={`flex justify-end items-center space-x-2  `}>
        <button
          onClick={handlePreviousPage}
          className={`flex items-center justify-center rounded-xl px-1 md:px-2 md:py-1 mb-1 md:me-2 md:mb-2 text-xs md:text-sm ${
            currentPage === 1
              ? "text-gray-400"
              : isDarkMode
              ? "text-gray-300  hover:text-white hover:bg-blue-600"
              : "text-blue-600 hover:text-white hover:bg-blue-600"
          } ${isDarkMode ? "bg-gray-800" : ""}`}
          disabled={currentPage === 1}
        >
          <GrFormPrevious size={30} />
          Previous
        </button>

        {renderPageButtons()}

        <button
          onClick={handleNextPage}
          className={`flex items-center justify-center rounded-xl  px-1 md:px-2 md:py-1 mb-1 md:me-2 md:mb-2 text-xs md:text-sm ${
            currentPage === totalPages
              ? "text-gray-400 "
              : isDarkMode
              ? "text-gray-300  hover:text-white hover:bg-blue-600"
              : "text-blue-600  hover:text-white hover:bg-blue-600"
          } ${isDarkMode ? "bg-gray-800" : ""} hover:bg-blue-100`}
          disabled={currentPage === totalPages}
        >
          Next
          <MdNavigateNext size={30} />
        </button>
      </div>
    </>
  );
};
