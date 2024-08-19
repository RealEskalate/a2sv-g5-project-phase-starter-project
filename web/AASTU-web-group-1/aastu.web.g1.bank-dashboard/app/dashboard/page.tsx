"use client";
import { useEffect, useRef, useState } from "react";
import {
  IoChevronBackCircleOutline,
  IoChevronForwardCircleOutline,
} from "react-icons/io5";


import { Barchart } from "@/app/dashboard/transactions/component/weeklyActivityChart";
import { Pie_chart } from "@/components/ui/Pie_chart";
import { PiTelegramLogoLight } from "react-icons/pi";
import CreditCard from "./_components/Credit_Card";
import { Profile } from "./_components/Profile";
import { Transaction } from "./_components/Transaction";
import { getCreditCards } from "./transactions/component/getCreditCards";
import { CardDetails, TransactionData } from "@/types";
import { BalanceAreachart  } from "./transactions/component/balanceChart";
import { getallTransactions } from "./transactions/component/getTransactions";

const MainDashboard = () => {

  const QuickTransferSection = useRef<HTMLDivElement | null>(null);

  const scrollCards = (scrollOffset: number) => {
    if (QuickTransferSection.current) {
      QuickTransferSection.current.scrollLeft += scrollOffset;
    }
  };

   const [creditCards, setCreditCards] = useState<CardDetails[]>([]);
   const [recentTransactions, setRecentTransactions] = useState<TransactionData[]>([]);
   useEffect(() => {
     const fetchData = async () => {
       const res = await getCreditCards();
       const recent = await getallTransactions(0,3);
       setRecentTransactions(recent || []);
       setCreditCards(res || []);
     };
     fetchData();
   }, []);

  return (
    <div className=" p-5 space-y-5 ">
      {/* First Row: My Cards and Recent Transactions */}
      <div className="md:flex sm:grid-cols-2 md:gap-5 space-y-5 md:space-y-0  ">
        {/*  My Cards Section  */}
        <div className="md:w-2/3 space-y-5 ">
          <div className="flex justify-between font-inter text-[16px] font-semibold">
            <h4>My Cards</h4>
            <h4>See All</h4>
          </div>
          <div className="flex space-x-5 overflow-x-auto [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]">
            {creditCards.map((card) => (
              <CreditCard
                key={card.id} 
                id={card.id}
                balance={card.balance}
                semiCardNumber={card.semiCardNumber}
                cardHolder={card.cardHolder}
                expiryDate={card.expiryDate}
                cardType={card.cardType}
              />
            ))}
          </div>
        </div>

        {/* Recent Transactions Section */}
        <div className="space-y-5 md:w-1/3">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Recent Transactions</h4>
          </div>
          <div className="space-y-5 md:p-5 bg-white rounded-xl md:shadow-lg md:border md:border-gray-300">
            {recentTransactions.map((transaction) => (
              <Transaction
                key={transaction.transactionId}
                date={transaction.date}
                amount={transaction.amount}
                description={transaction.description}
                type={transaction.type}
                transactionId={transaction.transactionId}
                senderUserName={transaction.senderUserName}
                receiverUserName={transaction.receiverUserName}
              />
            ))}
          
          </div>
        </div>
      </div>

      {/* Second Row: Weekly Activity and Expense Statistics */}
      <div className="md:flex sm:grid-cols-2 md:gap-5 space-y-5 md:space-y-0 ">
        {/* Weekly Activity Section */}
        <div className="md:w-2/3 space-y-5 ">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Weekly Activity</h4>
          </div>
          <div className="bg-white rounded-xl md:shadow-lg md:border md:border-gray-300  ">
            <Barchart />
          </div>
        </div>

        {/* Expense Statistics Section */}
        <div className="md:w-1/3 space-y-5 ">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Expense Statistics</h4>
          </div>
          <div className="bg-white rounded-xl md:shadow-lg md:border md:border-gray-300   ">
            <Pie_chart />
          </div>
        </div>
      </div>

      {/* Third Row: Quick Transfer and Balance History */}
      <div className="md:grid md:grid-cols-[1fr,2fr] md:gap-10 space-y-5 md:space-y-0">
        {/* Quick Transfer Section */}
        <div className="space-y-5 ">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Quick Transfer</h4>
          </div>
          <div className="bg-white rounded-xl md:shadow-lg md:border md:border-gray-300 p-5 space-y-5">
            <div>
              <button
                className="float-right hover:bg-blue-500 rounded-xl"
                onClick={() => scrollCards(200)}
              >
                <IoChevronForwardCircleOutline size={20} />
              </button>
              <button
                className="float-left hover:bg-blue-500 rounded-xl"
                onClick={() => scrollCards(-200)}
              >
                <IoChevronBackCircleOutline size={20} />
              </button>
            </div>

            <div
              ref={QuickTransferSection}
              className="flex max-w-[300px] space-x-5 overflow-x-auto [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]"
            >
              <Profile
                image="/images/avatar1.svg"
                name="Olivia Lia"
                job="CEO"
              />
              <Profile
                image="/images/avatar2.svg"
                name="Randy Press"
                job="Director"
              />
              <Profile
                image="/images/avatar3.svg"
                name="Workman"
                job="Designer"
              />
              <Profile
                image="/images/avatar4.svg"
                name="Patricia Lia"
                job="CEO"
              />
              <Profile
                image="/images/avatar1.svg"
                name="Randy Press"
                job="Director"
              />
            </div>
            <div className="flex space-x-10 h-[40px] items-center">
              <h4 className="font-inter text-[12px] text-[#718EBF]">
                Write Amount
              </h4>
              <div className="bg-gray-200 rounded-3xl flex items-center">
                <input
                  type="text"
                  className="w-[90px] h-[40px] rounded-full bg-gray-200 px-3"
                  placeholder="525.50"
                />
                <button className="bg-[#1814F3] text-white rounded-full px-4 h-[40px] ml-2 flex items-center space-x-2">
                  <p>Send </p>
                  <PiTelegramLogoLight />
                </button>
              </div>
            </div>
          </div>
        </div>

        {/* Balance History Section */}
        <div className="space-y-5 ">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Balance History</h4>
          </div>
          <div className="bg-white rounded-xl md:shadow-lg md:border md:border-gray-300 p-5">
            <BalanceAreachart />
          </div>
        </div>
      </div>
    </div>
  );
};

export default MainDashboard;
