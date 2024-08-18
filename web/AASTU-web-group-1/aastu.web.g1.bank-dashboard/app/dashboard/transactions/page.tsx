"use client";
import React, { useEffect, useState } from "react";
import CreditCard from "../_components/Credit_Card";
import { ExpenseChart } from "./component/ExpenseChart";
import { GrFormPrevious } from "react-icons/gr";
import { MdNavigateNext } from "react-icons/md";
import { ExpenseTable } from "./component/ExpenseTable";
import { getallTransactions } from "./component/getTransactions";
import { getIncomes } from "./component/getIncomes";
import { getExpenses } from "./component/getExpenses";
import { CardDetails, TransactionData } from "@/types";
import { getCreditCards } from "./component/getCreditCards";

const Transactions = () => {
  const rowsPerPage = 5;
  const totalPages = 10;
  
  const [activeTab, setActiveTab] = useState<"all" | "income" | "expense">("all");
  const [currentPage, setCurrentPage] = useState(1); 
  const [transactions, setTransactions] = useState<TransactionData[]>([]);
  const [creditCards, setCreditCards] = useState<CardDetails[]>([]);
  useEffect(() => {
    const fetchData = async () => {
      const res = await getCreditCards();
      setCreditCards(res || [])
    };
    fetchData();
  }, []);
 
  useEffect(() => {
    const fetchData = async () => {
      let data;
      if (activeTab === "all") {
        data = await getallTransactions(currentPage-1, rowsPerPage);
      } else if (activeTab === "income") {
        data = await getIncomes(currentPage-1, rowsPerPage);
      } else {
        data = await getExpenses(currentPage-1, rowsPerPage);
      }
      setTransactions(data || []);
    };
    fetchData();
  }, [currentPage, rowsPerPage, activeTab]);

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  const handleNextPage = () => {
    if (currentPage < totalPages) {
      setCurrentPage(currentPage + 1);
    }
  };

  const handlePreviousPage = () => {
    if (currentPage > 0) {
      setCurrentPage(currentPage - 1);
    }
  };

  const renderPageButtons = () => {

    const pagesToShow = totalPages < 4 ? totalPages : 4;

    const startPage = currentPage <= 2 || totalPages <= 4 ? 1 : currentPage > totalPages - 2 ? totalPages - 3: currentPage - 1;

    return Array.from({ length: pagesToShow }, (_, index) => {
      const page = startPage + index;
      return (
        <button
          key={page}
          onClick={() => handlePageChange(page)}
          className={`${
            page === currentPage ? "text-white bg-blue-600" : "text-blue-600"
          } hover:bg-blue-700 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 me-2 mb-2 align-middle`}
        >
          {page}
        </button>
      );
    });
  };

  return (
    <div className="p-5 space-y-5">
      <div className="md:flex sm:grid sm:grid-cols-2 space-y-5 md:space-y-0 ">
        <div className="md:w-2/3 space-y-5  ">
          <div className="flex justify-between font-inter text-[16px] font-semibold">
            <h4>My Cards</h4>
            <h4>+Add Card</h4>
          </div>
          <div className="flex space-x-5 overflow-x-auto [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]">
            {creditCards.map((card) => (
              <CreditCard
               id = {card.id}
                balance={card.balance}
                semiCardNumber={card.semiCardNumber}
                cardHolder={card.cardHolder}
                expiryDate={card.expiryDate}
                cardType={card.cardType}
              />
            ))}
         
          </div>
        </div>
        <div className="space-y-5 md:w-1/3">
          <div className="font-inter text-[16px] font-semibold">
            <h4>My Expense</h4>
          </div>
          <div className="px-2 md:p-3 bg-white rounded-xl md:shadow-lg md:border md:border-gray-300">
            <ExpenseChart />
          </div>
        </div>
      </div>

      <div className="p-5 space-y-5 w-[90%] items-center">
        <div className="space-y-5">
          <h4>Recent Transaction</h4>
          <div className="space-x-5 justify-start flex">
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
        <ExpenseTable transactions={transactions} />
      </div>

      <div className="justify-end flex items-center md:space-x-5  space-x-2">
        <button
          onClick={handlePreviousPage}
          className={`flex ${
            currentPage === 1 ? "text-gray-400" : "text-blue-600"
          }`}
          disabled={currentPage === 1}
        >
          <GrFormPrevious color="blue-600" size={30} />{" "}
          <p
            className={`align-middle ${
              currentPage === 1 ? "text-gray-400" : "text-blue-600"
            } `}
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
            } `}
          >
            Next
          </p>
          <MdNavigateNext color="blue-600" size={30} />
        </button>
      </div>
    </div>
  );
};

export default Transactions;
