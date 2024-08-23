"use client";

import React from "react";
import MyCards from "../components/MyCards";
import RecentTransaction from "../components/RecentTransaction";
import WeeklyActivity from "../components/WeeklyActivity";
import PieChart from "../components/PieChart";
import QuickTransfer from "../components/QuickTransfer";

const Dashboard = () => {
  return (
    <div className="flex flex-col p-4 space-y-6">
      <div className="flex justify-between items-center"> 
      {/* lg:flex-row lg:space-x-6 */}
        <div className="flex flex-col w-2/3 space-y-3">
          <div className="flex justify-between items-center text-[#343C6A] font-semibold">
            <p className="text-[22px]">My Cards</p>
            <a href="#" className="text-[17px]">
              See All
            </a>
          </div>
          <div className="flex justify-between items-center">
            <MyCards />
            <MyCards />
          </div>
        </div>
        <div className="w-1/3">
          <p className="text-[#343C6A] text-[22px] font-semibold">
            Recent Transactions
          </p>
          <RecentTransaction />
        </div>
      </div>

      <div className="flex flex-col lg:flex-row lg:space-x-6">
        <div className="w-full lg:w-2/3">
          <p className="text-[#343C6A] text-[22px] font-semibold mb-2">
            Weekly Activity
          </p>
          <WeeklyActivity />
        </div>
        <div className="w-full lg:w-1/3">
          <p className="text-[#343C6A] text-[22px] font-semibold mb-4">
            Expense Statistics
          </p>
          <PieChart />
        </div>
      </div>

      <div>
        <p className="text-[#343C6A] text-[22px] font-semibold mb-2">
          Quick Transfer
        </p>
        <QuickTransfer />
      </div>

      {/* Uncomment when Balance History is ready */}
      {/* <div className="mt-6">
        <p className="text-[#343C6A] text-[22px] font-semibold mb-2">Balance History</p>
        <BalanceHistory />
      </div> */}
    </div>
  );
};

export default Dashboard;
