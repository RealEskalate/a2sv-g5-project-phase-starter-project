'use client'

import React from 'react';
import RecentTransactions from '@/components/RecentTransaction';
import ExpensesChart from '@/components/ExpensesCart';
import SlidingCards from '@/components/SlidingCards'; // Import the sliding cards component
import DesktopCreditCard from '@/components/DesktopCreditCard';

const Transaction: React.FC = () => {
  return (
    <div className="p-4">
      {/* Large Screens Layout */}
      <div className="hidden lg:flex lg:space-x-8 lg:mb-8 lg:ml-72">
        {/* Cards Section */}
        <div className="flex-1">
          <h1 className="text-2xl font-bold mb-4">My Cards</h1>
          <div className="flex space-x-4">
            <div className="flex-1">
              {/* Card 1 Content */}
              <DesktopCreditCard bgColor="bg-blue-700" textColor="text-white" />
            </div>
            <div className="flex-1">
              {/* Card 2 Content */}
              <DesktopCreditCard bgColor="bg-green-500" textColor="text-white" />
            </div>
          </div>
        </div>

        {/* Chart Section */}
        <div className="flex-1">
          <h1 className="text-2xl font-bold mb-4">My Expenses</h1>
          <div className="pl-8">
            <ExpensesChart />
          </div>
        </div>
      </div>

      {/* Mobile Layout */}
      <div className="lg:hidden mb-8">
        <h1 className="text-2xl font-bold mb-4">My Cards</h1>
        <div className="mb-8">
          <SlidingCards />
        </div>
        <h1 className="text-2xl font-bold mb-4">My Expenses</h1>
        <div className="w-full">
          <ExpensesChart />
        </div>
      </div>

      {/* Recent Transactions Section */}
      <div>
        <h1 className="text-2xl font-bold text-balance lg:text-center mb-4">Recent Transactions</h1>
        <RecentTransactions />
      </div>
    </div>
  );
};

export default Transaction;
