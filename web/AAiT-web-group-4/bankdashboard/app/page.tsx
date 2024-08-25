import React from "react";
import WeeklyActivity from "./components/Charts/WeeklyActivity";
import ExpensePolarChart from "./components/Charts/ExpensePolarChart";
import CreditCard from "./components/CreditCard/CreditCard";
import QuickTransfer from "./components/QuickTransfer/QuickTransfer";
import BalanceHistoryLineGraph from "./components/Charts/BalanceHistoryGraph";
import RecentTransactions from "./components/RecentTransactions/RecentTransactions";

const Page = () => {
  return (
    <div className="flex flex-col justify-between gap-5 w-full h-full bg-Very-Light-White max-mobile:bg-white py-6 px-5">
      <div className="flex max-mobile:flex-wrap justify-around w-full h-1/3 ">
        {/* cards */}
        <div className="flex flex-col justify-between  max-mobile:overflow-y-scroll mobile:w-[64%]  ">
          <div className=" flex w-full justify-between ">
            <h2 className="text-[22px] max-mobile:text-base text-Dark-Slate-Blue">My Cards</h2>
            <p className="text-base max-mobile:text-sm text-Dark-Slate-Blue">See all</p>
          </div>
          <div className="flex gap-6 w-full  h-full">
            <CreditCard />
            <CreditCard />
          </div>
        </div>

        {/* recent transactions */}
        <div className="flex  flex-col justify-between mobile:w-1/3  px-5 ">
          <h2 className="text-Dark-Slate-Blue text-xl mx-2">Recent transactions</h2>
          <RecentTransactions />
        </div>
      </div>
    

      <div className="flex flex-wrap  w-full h-1/3 justify-between items-center py- px-4">
        {/* weekly activity */}
        <div className="flex flex-col  gap-3 justify-between  mobile:w-[64%] h-full ">
          <h2 className="text-Dark-Slate-Blue text-lg my-2">Weekly Activity</h2>
          <WeeklyActivity />
        </div>

        {/* expense statistics */}
        <div className=" flex flex-col gap-3 justify-between mobile:w-1/3 h-full">
          <h2 className="text-Dark-Slate-Blue text-xl my-2">Expense Statstics</h2>
          <ExpensePolarChart sectors={["transfer" , "shopping" , "other" , "service"]} bgColors={["orange" , "grey" , "blue" , "purple"]} />
        </div>
      </div>

      <div className="flex flex-wrap  w-full h-1/3 justify-between items-center py- px-4 ">
        {/* Quick Transfers */}
        
        <div className="flex flex-col gap-3 justify-between  mobile:w-[43%]   max-mobile:w-full">
          <h2 className="text-Dark-Slate-Blue text-xl mb-3">Quick Transfer</h2>
          <QuickTransfer />
        </div>

        <div className="flex flex-col  h-full w-full max-mobile:flex-wrap gap-3 justify-around mobile:w-[55%] max-mobile:w-full">
          <h2 className="text-Dark-Slate-Blue text-xl ">Weekly Activity</h2>
          <BalanceHistoryLineGraph balanceHistory={[10 , 100 , 40 , 80 , 100 ,90 , 150]} />
        </div>
      
      </div>
    </div>
  );
};

export default Page;
