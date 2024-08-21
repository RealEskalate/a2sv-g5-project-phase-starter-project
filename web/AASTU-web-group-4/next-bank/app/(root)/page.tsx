"use client"

import React from "react";
import { Color } from "chart.js";
import { colors, logo } from "@/constants";
import DesktopCreditCart from "@/components/DesktopCreditCard";
import ResponsiveCreditCard from "@/components/CreditCard";
import RecentTransaction from "@/components/Recent Transaction";
import ExpensesChart from "@/components/ExpensesCart";
import { icons, Import } from "lucide-react";
import { text } from "stream/consumers";
import BarChart from "@/components/BarChart";
import PieChart from "@/components/PieChart";
import QuickTransfer from "@/components/QuickTransfer";
import LineChart from "@/components/LineChart";



const page = () => {
  return (
    <div className={`${colors.graybg} p-6 md:ml-64 md:max-w-full md:p-12`}>
      <div className="flex flex-col justify-between md:flex-row  gap-10 ">
        <div className=" py-4 md:w-3/5 md:max-w-full">
          <div className={`${colors.navbartext} flex justify-between `}>
            <h1 className="font-bold text-2xl">My Cards</h1>
            <p className="py-2"> See All</p>
          </div>

          <div className="max-w-[345px] md:max-w-full">
            <div className="flex gap-3 overflow-x-auto max-w-full md:w-auto">
              <div className=" py-3 ">
                <ResponsiveCreditCard
                  backgroundColor={colors.blue}
                />
              </div>
              <div className=" py-3 ">
              <ResponsiveCreditCard
                  backgroundColor={colors.white}
                />
              </div>
            </div>
          </div>
        </div>
        <div className="  md:w-2/5  flex flex-col ">
          <div className={`${colors.navbartext} flex justify-between  py-4`}>
            <h1 className="font-bold text-2xl">Recent Transaction</h1>
          </div>
          <div className="flex flex-col rounded-2xl pr-2 w-[100%]">
            <RecentTransaction/>
          </div>
        </div>
      </div>
      <div className=" w-[100%] flex flex-col justify-between  md:grid md:grid-cols-5 md:gap-10 ">
        <div className=" md:col-span-3 ">
          <div className={`${colors.navbartext} flex justify-between py-4`}>
            <h1 className="font-bold text-2xl">Weekly Activity</h1>
          </div>
          <div className="w-[100%]">
            <BarChart />
          </div>
        </div>
        <div className=" w-[100%] py-5 flex flex-col gap-5 md:col-span-2 ">
          <div className={`${colors.navbartext}`}>
            <h1 className="font-bold text-2xl">Expense Statstics</h1>
          </div>
          <div className="w-[100%] pr-6">
            <PieChart />
          </div>
        </div>
      </div>

      <div className="flex flex-col justify-between w-full  md:grid md:grid-cols-5 md:gap-10 ">
        <div className=" md:col-span-2 py-4  ">
          <div className={`${colors.navbartext} flex justify-between `}>
            <h1 className="font-bold text-2xl">Quick Transfer</h1>
          </div>
          <div className="flex  gap-3 ">
            <div className="flex py-3 ">
              {" "}
              <QuickTransfer />
            </div>
          </div>
        </div>
        <div className=" md:col-span-3 ">
          <div className={`${colors.navbartext} flex justify-between  py-4`}>
            <h1 className="font-bold text-2xl">Balance History</h1>
          </div>
          <div className="pr-6">
            <LineChart />
          </div>
        </div>
      </div>
    </div>
  );
};

export default page;



