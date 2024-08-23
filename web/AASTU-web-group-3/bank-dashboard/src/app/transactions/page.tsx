import React from "react";
import AllTransactions from "../components/transactions/AllTransactions";
import CreditCard from "../components/CreditCard";
import DashboardBarChart from "../components/Chart/DashboardBarChart";
import RecentTransactions from '../components/transactions/RecentTransactions';

const TransactionsPage = () => {
  return (
    <section className="xl:w-11/12 xl:mx-4">
      <div className="credit-cards expenses flex flex-col lg:flex-row justify-between lg:gap-1 xl:gap-6 ">
        <div className="lg:w-[65%] rounded-xl bg-[#F5F7FA]">
          <div className="credit-card-info flex justify-between  h-20 items-center ">
            <h1 className="font-semibold text-[#343C6A]">My cards</h1>
            <h1 className="font-semibold text-[#343C6A]">+Add Card</h1>
          </div>
          <div className="cards flex gap-5 lg:gap-1  lg:flex-row overflow-x-auto overflow-y-hidden no-scrollbar h-56 lg:justify-between xl:gap-10">
            <div className=" flex-shrink-0 min-w-60 w-73 lg:w-60 lg:h-48 xl:w-96  xl:h-56 items-center">
              <CreditCard
                name="Karthik P"
                balance="$5,756"
                cardNumber="3778 **** **** 1234"
                validDate="11/15"
                backgroundImg="bg-[linear-gradient(107.38deg,#2D60FF_2.61%,#539BFF_101.2%)]"
                textColor="text-white"
              />
            </div>
            <div className="flex-shrink-0 min-w-60  w-73 lg:w-60 lg:h-48 xl:w-96  xl:h-56">
              <CreditCard
                name="Karthik P"
                balance="$5,756"
                cardNumber="3778 **** **** 1234"
                validDate="11/15"
                backgroundImg="bg-white"
                textColor="text-black"
              />
            </div>
            <div className="flex-shrink-0 min-w-60 w-73 lg:w-60 lg:h-48 xl:w-96  xl:h-56">
              <CreditCard
                name="Karthik P"
                balance="$5,756"
                cardNumber="3778 **** **** 1234"
                validDate="11/15"
                backgroundImg="bg-[linear-gradient(107.38deg,#4C49ED_2.61%,#0A06F4_101.2%)]"
                textColor="text-white"
              />
            </div>
          </div>
        </div>
          <DashboardBarChart />
      </div>
      <RecentTransactions/>

    </section>
  );
};

export default TransactionsPage;
