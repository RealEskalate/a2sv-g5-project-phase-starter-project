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
      <div className="flex flex-col justify-start lg:space-x-4 lg:flex-row lg:items-center ">
        <div className="flex flex-col w-full lg:w-2/3">
          <div className="flex justify-between items-center pb-4 text-[#343C6A] font-semibold">
            <p className="text-[22px]">My Cards</p>
            <a href="#" className="text-[17px]">
              See All
            </a>
          </div>
          <div className="flex space-x-4 w-full min-w-[400px] max-w-[730px] overflow-x-auto whitespace-nowrap scroll-smooth scrollbar-hide snap-mandatory">
            <MyCards />
            <MyCards />
          </div>
        </div>
        <div className=" lg:w-1/3 w-full">
          <p className="pb-4 text-[#343C6A] text-[22px] w-full font-semibold">
            Recent Transactions
          </p>
          <RecentTransaction />
        </div>
      </div>

      <div className="flex flex-col lg:flex-row lg:items-center lg:space-x-4">
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

      {/* <div className="mt-6">
        <p className="text-[#343C6A] text-[22px] font-semibold mb-2">Balance History</p>
        <BalanceHistory />
      </div> */}
    </div>
  );
};

export default Dashboard;
