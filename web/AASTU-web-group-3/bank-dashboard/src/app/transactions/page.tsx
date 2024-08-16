import React from "react";
import AllTransactions from "../components/AllTransactions";
import MyExpenses from "../components/MyExpenses";
import CreditCard from "../components/CreditCard";
import DashboardBarChart from "../components/Chart/DashboardBarChart";

const TransactionsPage = () => {
  return (
    <section>
      <div className="credit-cards expenses flex flex-col lg:flex-row justify-between">
        <div className="lg:w-[65%]  rounded-xl bg-[#F5F7FA] p-1">
          <div className="credit-card-info flex justify-between px-4 h-20 items-center">
            <h1 className="font-semibold text-[#343C6A]">My cards</h1>
            <h1 className="font-semibold text-[#343C6A]">+Add Card</h1>
          </div>
          <div className="cards flex gap-2 lg:flex-row overflow-x-auto overflow-y-hidden no-scrollbar   h-56 lg:h-52  lg:justify-between lg:px-4 ">
            <div className="credit-card min-w-72 max-w-88 flex-shrink-0 lg:w-[45%]">
              <CreditCard
                name="Karthik P"
                balance="$5,756"
                cardNumber="3778 **** **** 1234"
                validDate="11/15"
                backgroundImg="bg-[linear-gradient(107.38deg,#2D60FF_2.61%,#539BFF_101.2%)]"
                textColor="text-white"
              />
            </div>
            <div className="credit-card min-w-72 max-w-88 flex-shrink-0 lg:w-[45%]">
              <CreditCard
                name="Karthik P"
                balance="$5,756"
                cardNumber="3778 **** **** 1234"
                validDate="11/15"
                backgroundImg="bg-white"
                textColor="text-black"
              />
            </div>
            <div className="credit-card min-w-72 max-w-88 flex-shrink-0 lg:w-[45%]">
              <CreditCard
                name="Karthik P"
                balance="$5,756"
                cardNumber="3778 **** **** 1234"
                validDate="11/15"
                backgroundImg="bg-[linear-gradient(107.38deg,#4C49ED_2.61%,#0A06F4_101.2%)]"
                textColor="text-white"
              />
            </div>

            {/* Add more cards here if needed */}
          </div>
        </div>
        <DashboardBarChart />
      </div>
      <AllTransactions />
    </section>
  );
};

export default TransactionsPage;
