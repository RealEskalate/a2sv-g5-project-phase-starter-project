"use client";
import React, { useState } from "react";
import CreditCard from "../_components/Credit_Card";
import { ExpenseChart } from "./component/ExpenseChart";
import { GrFormPrevious } from "react-icons/gr";
import { MdNavigateNext } from "react-icons/md";

import { ExpenseTable} from "./component/ExpenseTable";
import { transactions } from "./component/transactionData";



const Transactions = () => {
  const rowsPerPage = 5;
  const [activeTab, setActiveTab] = useState<"all" | "income" | "expense">(
    "all"
  );
  const [currentPage, setCurrentPage] = useState(1);

  const filteredTransactions = transactions.filter((transaction) => {
    if (activeTab === "all") return true;
    if (activeTab === "income") return transaction.amount > 0;
    if (activeTab === "expense") return transaction.amount < 0;
    return false;
  });

  const totalPages = Math.ceil(filteredTransactions.length / rowsPerPage);

  // Get the transactions for the current page
  const paginatedTransactions = filteredTransactions.slice(
    (currentPage - 1) * rowsPerPage,
    currentPage * rowsPerPage
  );

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

  return (
    <div className="p-5 space-y-5 ">
      {/* First Row: My Cards and My expenses */}
      <div className="md:flex-row sm:grid sm:grid-cols-2 space-y-5 md:space-y-0 space-x-5">
        {/* My Cards Section */}
        <div className="space-y-5 mx-5">
          <div className="flex justify-between font-inter text-[16px] font-semibold">
            <h4>My Cards</h4>
            <h4>+Add Card</h4>
          </div>
          <div className="flex space-x-5 overflow-x-auto [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]">
            <CreditCard
              isBlue={true}
              balance={5894}
              creditNumber="3778*** ****1234"
              name="Eddy Cusuma"
              textColor="text-white"
            />
            <CreditCard
              isBlue={false}
              balance={3210}
              creditNumber="3778*** ****1234"
              name=" Sarah Johnson"
              textColor="text-black"
            />
            <CreditCard
              isBlue={true}
              balance={5894}
              creditNumber="3778*** ****1234"
              name="Eddy Cusuma"
              textColor="text-white"
            />
          </div>
        </div>
        {/* Expense Section */}
        <div className="space-y-5">
          <div className="font-inter text-[16px] font-semibold">
            <h4>My Expense</h4>
          </div>
          <div className="p-5 bg-white rounded-xl md:shadow-lg md:border md:border-gray-300">
            <ExpenseChart />
          </div>
        </div>
      </div>
      {/* Second Row: Recent Transaction */}
      <div className="space-y-5 mx-5  w-[80%] items-center">
        <div className="space-y-5 ">
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
        
        {/* table */}
        <ExpenseTable transactions={paginatedTransactions} />
      </div>
      {/* Third Row: Table Navigation */}
      {/* Pagination */}
      <div className="justify-end flex items-center space-x-5">
        <button
          onClick={handlePreviousPage}
          className={`flex ${currentPage === 1 ? "text-gray-400" : "text-blue-600"}`}
          disabled={currentPage === 1}
        >
          <GrFormPrevious color="blue-600" size={30} /> <p className="text-blue-600 align-middle">Previous</p>
        </button>
       
        {Array.from({ length: totalPages }, (_, index) => {
          const page = index + 1;
          return (
            <button
              key={page}
              onClick={() => handlePageChange(page)}
              className={`${
                page === currentPage
                  ? "text-white bg-blue-600"
                  : "text-blue-600"
              } hover:bg-blue-700 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 me-2 mb-2 align-middle`}
            >
              {page}
            </button>
          );
        })}

        <button
          onClick={handleNextPage}
          className={` flex ${
            currentPage === totalPages ? "text-gray-400" : "text-blue-600"
          }`}
          disabled={currentPage === totalPages}

        >
          <p className="text-blue-600 align-middle">Next</p>
          <MdNavigateNext color="blue-600" size={30} />
        </button>
      </div>
    </div>
  );
};

export default Transactions;
