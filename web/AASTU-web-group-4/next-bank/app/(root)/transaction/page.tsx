"use client";

import React, { Suspense } from "react";
import RecentTransactions from "@/components/RecentTransaction";
import ExpensesChart from "@/components/ExpensesCart";
import SlidingCards from "@/components/SlidingCards"; // Import the sliding cards component
import CreditCard from "@/components/CreditCard";

const Transaction: React.FC = () => {
  return (
    <div className=" w-[100%]">
      {/* Large Screens Layout */}
      <div className=" hidden lg:grid lg:grid-cols-2 lg:gap-5 lg:space-x-8 lg:pb-8 lg:ml-72">
        {/* Cards Section */}
        <div className="flex flex-col">
          <h1 className="text-2xl font-bold mb-4 dark:text-blue-500">
            My Cards
          </h1>
          <div className=" w-[100%] overflow-x-auto flex space-x-4">
            <div className="flex-1">
              {/* Card 1 Content */}

              <CreditCard backgroundColor="bg-blue-700" />
            </div>
            <div className="flex-1">
              {/* Card 2 Content */}
              <CreditCard backgroundColor="bg-purple-700" />
            </div>
          </div>
        </div>

        {/* Chart Section */}
        <div className="flex-1 ml-0">
          <h1 className="text-2xl font-bold dark:text-blue-500">My Expenses</h1>

          <div className="pl-8">
            <ExpensesChart />
          </div>
        </div>
      </div>

      {/* Mobile Layout */}
      <div className="lg:hidden mb-8">
        <h1 className="text-2xl font-bold mb-4 dark:text-blue-500">My Cards</h1>
        <div className="mb-8">
          <SlidingCards />
        </div>
        <h1 className="text-2xl font-bold mb-4 dark:text-blue-500">
          My Expenses
        </h1>
        <div className="w-full">
          <ExpensesChart />
        </div>
      </div>

      {/* Recent Transactions Section */}
      <div>
        <h1 className="text-2xl font-bold text-balance lg:text-center mb-4 dark:text-blue-500">
          Recent Transactions
        </h1>

        <RecentTransactions />
      </div>
    </div>
  );
};

export default Transaction;
