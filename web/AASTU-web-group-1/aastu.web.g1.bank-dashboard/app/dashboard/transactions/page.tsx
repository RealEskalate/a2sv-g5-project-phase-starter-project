"use client";

import { useState, useEffect } from "react";
import { GrFormPrevious } from "react-icons/gr";
import { MdNavigateNext } from "react-icons/md";
import CreditCard from "../_components/Credit_Card";
import { ExpenseChart } from "./component/ExpenseChart";
import { ExpenseTable } from "./component/ExpenseTable";
import { CardDetails, TransactionData } from "@/types";
import { getallTransactions, getCreditCards, getExpenses, getIncomes } from "@/lib/api";

const Transactions = () => {
  const rowsPerPage = 5;
  const totalPages = 10; 
  const [loading, setLoading] = useState(true);
  const [activeTab, setActiveTab] = useState<"all" | "income" | "expense">(
    "all"
  );
  const [currentPage, setCurrentPage] = useState(1);
  const [transactions, setTransactions] = useState<TransactionData[]>([]);
  const [creditCards, setCreditCards] = useState<CardDetails[]>([]);
  const [expenses, setExpenses] = useState<TransactionData[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const [cards, initialExpenses] = await Promise.all([
          getCreditCards(),
          getExpenses(0, 6),
        ]);
        setCreditCards(cards || []);
        setExpenses(initialExpenses || []);
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
        let data;
        if (activeTab === "all") {
          data = await getallTransactions(currentPage - 1, rowsPerPage);
        } else if (activeTab === "income") {
          data = await getIncomes(currentPage - 1, rowsPerPage);
        } else {
          data = await getExpenses(currentPage - 1, rowsPerPage);
        }
        setTransactions(data || []);
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
            page === currentPage ? "text-white bg-blue-600" : "text-blue-600"
          } hover:bg-blue-700 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 me-2 mb-2`}
        >
          {page}
        </button>
      );
    });
  };

  if (loading) {
    return (
      <div className="flex justify-center items-center h-screen">
        <div className="animate-spin rounded-full h-32 w-32 border-t-4 border-dotted border-blue-600"></div>
      </div>
    );
  }

  return (
    <div className="p-5 space-y-5 lg:p-10">
      <div className="lg:flex md:grid md:grid-cols-2 gap-5 space-y-5 md:space-y-0">
        <div className="lg:w-2/3 space-y-5 ">
          <div className="flex justify-between font-inter text-[16px] font-semibold mx-3">
            <h4>My Cards</h4>
            <h4>+Add Card</h4>
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
        <div className="lg:w-1/3 space-y-5">
          <div className="font-inter text-[16px] font-semibold">
            <h4>My Expense</h4>
          </div>
          <div className="rounded-xl p-2 pt-1">
            <ExpenseChart expenses={expenses} />
          </div>
        </div>
      </div>

      <div className="space-y-5 w-full items-center">
        <div className="space-y-5">
          <h4>Recent Transactions</h4>
          <div className="space-x-5 flex">
            <button onClick={() => setActiveTab("all")}>
              <p
                className={`border-b-2 text-xs ${
                  activeTab === "all"
                    ? "border-b-blue-600"
                    : "border-b-transparent"
                }`}
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
                }`}
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
                }`}
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
          className={`flex ${
            currentPage === 1 ? "text-gray-400" : "text-blue-600"
          }`}
          disabled={currentPage === 1}
        >
          <GrFormPrevious size={30} />{" "}
          <p
            className={`align-middle ${
              currentPage === 1 ? "text-gray-400" : "text-blue-600"
            }`}
          >
            Previous
          </p>
        </button>

        {renderPageButtons()}

        <button
          onClick={handleNextPage}
          className={`flex ${
            currentPage === totalPages ? "text-gray-400" : "text-blue-600"
          }`}
          disabled={currentPage === totalPages}
        >
          <p
            className={`align-middle ${
              currentPage === totalPages ? "text-gray-400" : "text-blue-600"
            }`}
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
