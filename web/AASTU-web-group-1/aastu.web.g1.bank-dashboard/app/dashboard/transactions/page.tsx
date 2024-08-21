"use client";

import { useState, useEffect } from "react";
import { GrFormPrevious } from "react-icons/gr";
import { MdNavigateNext } from "react-icons/md";
import CreditCard from "../_components/Credit_Card";
import { ExpenseChart } from "./component/ExpenseChart";
import { ExpenseTable } from "./component/ExpenseTable";
import { CardDetails, TransactionContent, TransactionData } from "@/types";
import {
  getallTransactions,
  getCreditCards,
  getExpenses,
  getIncomes,
} from "@/lib/api";
import { Loading } from "../_components/Loading";
import { useUser } from "@/contexts/UserContext";
import Link from "next/link";

const Transactions = () => {
  const { isDarkMode } = useUser();
  const rowsPerPage = 5;
  const [loading, setLoading] = useState(true);
  const [activeTab, setActiveTab] = useState<"all" | "income" | "expense">(
    "all"
  );
  const [currentPage, setCurrentPage] = useState(1);
  const [transactions, setTransactions] = useState<TransactionContent[]>([]);
  const [creditCards, setCreditCards] = useState<CardDetails[]>([]);
  const [expenses, setExpenses] = useState<TransactionContent[]>([]);
  const [totalPages, setTotalPages] = useState<number>(5);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const [cards, initialExpenses] = await Promise.all([
          getCreditCards(0, 2),
          getExpenses(0, 6),
        ]);
        setCreditCards(cards?.content || []);
        setExpenses(initialExpenses?.content || []);
      } catch (error) {
        console.error("Error fetching data:", error);
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, []);

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
          className={`${
            page === currentPage ? " bg-blue-600" : "text-blue-600"
          } hover:bg-blue-700 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 me-2 mb-2 ${
            isDarkMode ? "bg-gray-800 text-white" : "bg-white "
          }`}
        >
          {page}
        </button>
      );
    });
  };

  if (loading) {
    return <Loading />;
  }

  return (
    <div
      className={`space-y-5 p-5 ${
        isDarkMode ? "bg-gray-700 text-gray-200" : "bg-[#F5F7FA] text-gray-900"
      }`}
    >
      <div className="md:flex  sm:grid-cols-2 md:gap-5 space-y-5 md:space-y-0">
        <div className="md:w-2/3 space-y-5">
          <div className="flex justify-between font-inter text-[16px] font-semibold ">
            <h4>My Cards</h4>
            <h4>
              <Link href="/dashboard/credit-cards/#add-card">+Add Card</Link>
            </h4>
          </div>
          <div className="flex space-x-5 overflow-x-auto [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]">
            {creditCards.map((card) => (
              <CreditCard
                key={card.id}
                id={card.id}
                balance={card.balance}
                semiCardNumber={card.semiCardNumber}
                cardHolder={card.cardHolder}
                expiryDate={card.expiryDate}
                cardType={card.cardType}
              />
            ))}
          </div>
        </div>
        <div className="md:w-1/3 md:space-y-5 w-full">
          <div className="font-inter text-[16px] font-semibold">
            <h4>My Expense</h4>
          </div>
          <div
            className={`rounded-xl  pt-1 ${
              isDarkMode ? "bg-gray-800" : "bg-white s"
            }hadow-lg`}
          >
            <ExpenseChart expenses={expenses} />
          </div>
        </div>
      </div>

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
    </div>
  );
};

export default Transactions;
