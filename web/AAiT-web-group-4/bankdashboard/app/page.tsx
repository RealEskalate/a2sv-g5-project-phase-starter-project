import React from "react";
import WeeklyActivity from "./components/Charts/WeeklyActivity";
import ExpensePolarChart from "./components/Charts/ExpensePolarChart";
import CreditCard from "./components/CreditCard/CreditCard";
import QuickTransfer from "./components/QuickTransfer/QuickTransfer";
import BalanceHistoryLineGraph from "./components/Charts/BalanceHistoryGraph";
import RecentTransactions from "./components/RecentTransactions/RecentTransactions";
import MyExpenseGraph from "./components/Charts/MyExpenseGraph";
const Page = () => {
  return (
    <div className="flex flex-col justify-between gap-5 h-full bg-Very-Light-White max-mobile:bg-white py-6 px-5">
      <div className="flex max-mobile:flex-wrap justify-around w-full">
        {/* cards */}
        <div className="flex flex-col justify-between  max-mobile:overflow-y-scroll ">
          <div className=" flex w-full justify-between py-1 ">
            <h2 className="text-[22px] max-mobile:text-base">My Cards</h2>
            <p className="text-base max-mobile:text-sm">See all</p>
          </div>
          <div className="flex gap-6 w-full max-desktop:overflow-x-scrollmax-mobile:h-[210px] max-mobile:w-[560px] mobile:w-[487px] mobile:h-[208px] tablet:h-[282px] tablet:w-[720px]">
            <CreditCard />
            <CreditCard />
          </div>
        </div>

        {/* recent transactions */}
        <div className="flex flex-col justify-between w-auto ">
          <h2 className="text-Dark-Slate-Blue text-xl py-3">Recent transactions</h2>
          <RecentTransactions />
        </div>
      </div>

      <div className="flex flex-wrap w-full h-auto justify-around items-center py-2">
        {/* weekly activity */}
        <div className="flex flex-col  gap-3 justify-between ">
          <h2 className="text-Dark-Slate-Blue text-lg ">Weekly Activity</h2>
          <WeeklyActivity />
        </div>

        {/* expense statistics */}
        <div className=" flex flex-col gap-3 justify-between max-tablet:w-[231px] tablet:w-[367px] h-auto">
          <h2 className="text-Dark-Slate-Blue text-xl">Expense Statstics</h2>
          <ExpensePolarChart sectors={["transfer" , "shopping" , "other" , "service"]} bgColors={["orange" , "grey" , "blue" , "purple"]} />
        </div>
      </div>

      <div className="flex flex-wrap w-full h-auto justify-between items-center py-2">
        {/* Quick Transfers */}
        <div className="flex flex-col gap-3 justify-between  max-mobile:w-[445px] max-mobile:h-[323px] mobile:w-[231px] mobile:h-[170px] tablet:w-[445px] tablet:h-[323px]">
          <h2 className="text-Dark-Slate-Blue text-xl ">Quick Transfer</h2>
          <QuickTransfer />
        </div>

        <div className="flex flex-col  gap-3 justify-around ">
          <h2 className="text-Dark-Slate-Blue text-xl ">Weekly Activity</h2>
          <BalanceHistoryLineGraph balanceHistory={[10 , 100 , 40 , 80 , 100 ,90 , 150]} />
        </div>
      
      </div>
    </div>
  );
};

export default Page;
