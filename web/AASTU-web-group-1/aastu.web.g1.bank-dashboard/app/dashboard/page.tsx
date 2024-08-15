'use client'
import React, { useRef } from "react";
import {
  IoChevronForwardCircleOutline,
  IoChevronBackCircleOutline,
} from "react-icons/io5";

import { Transaction } from "./_components/Transaction";
import { Barchart } from "@/components/ui/Barchart";
import CreditCard from "./_components/CreditCard";
import { Profile } from "./_components/Profile";
import { PiTelegramLogoLight } from "react-icons/pi";
import { Areachart } from "@/components/ui/Areachart";
import { Pie_chart } from "@/components/ui/Pie_chart";

const MainDashboard = () => {
  const QuickTransferSection = useRef<HTMLDivElement|null>(null);
  const scrollCards = (scrollOffset:number)=>{
    if(QuickTransferSection.current){
    QuickTransferSection.current.scrollLeft+=scrollOffset;   
    }
   
  }
  return (
    <div className="p-5 md:pl-10 md:ml-auto space-y-5 ">
      {/* First Row: My Cards and Recent Transactions */}
      <div className="md:grid md:grid-cols-2 md:gap-10 space-y-5 md:space-y-0">
        {/* My Cards Section */}
        <div className="space-y-5">
          <div className="flex justify-between font-inter text-[16px] font-semibold">
            <h4>My Cards</h4>
            <h4>See All</h4>
          </div>
          <div className="flex space-x-5 overflow-x-auto [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]">
            <CreditCard
              color="bg-gradient-to-r from-[#423fee] to-[#2723f1]"
              balance={5894}
              creditNumber="3778*** ****1234"
              name="Eddy Cusuma"
              textColor="text-white"
            />
            <CreditCard
              color="bg-white"
              balance={3210}
              creditNumber="3778*** ****1234"
              name=" Sarah Johnson"
              textColor="text-black"
            />
          </div>
        </div>

        {/* Recent Transactions Section */}
        <div className="space-y-5">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Recent Transactions</h4>
          </div>
          <div className="space-y-2 md:p-5 bg-white rounded-xl md:shadow-lg md:border md:border-gray-300">
            <Transaction
              image="/icons/wallet.png"
              transactionType="Deposited from my"
              date="28 January 2021"
              amount="+$85"
              color="bg-yellow-100"
            />
            <Transaction
              image="/icons/paypal.png"
              transactionType="Deposited Paypal"
              date="28 January 2021"
              amount="+$85"
              color="bg-indigo-100"
            />
            <Transaction
              image="/icons/dollarSign.png"
              transactionType="Deposited from my"
              date="28 January 2021"
              amount="+$85"
              color="bg-green-100"
            />
          </div>
        </div>
      </div>

      {/* Second Row: Weekly Activity and Expense Statistics */}
      <div className="md:grid md:grid-cols-2 md:gap-10 space-y-5 md:space-y-0">
        {/* Weekly Activity Section */}
        <div className="space-y-5">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Weekly Activity</h4>
          </div>
          <div className="bg-white rounded-xl md:shadow-lg md:border md:border-gray-300 md:p-5 sm:min-w-[375px]">
            <Barchart />
          </div>
        </div>

        {/* Expense Statistics Section */}
        <div className="space-y-5">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Expense Statistics</h4>
          </div>
          <div className="bg-white rounded-xl md:shadow-lg md:border md:border-gray-300 md:p-5 sm:space-x-10">
            <Pie_chart />
          </div>
        </div>
      </div>

      {/* Third Row: Quick Transfer and Balance History */}
      <div className="md:flex md:gap-10 space-y-5 md:space-y-0">
        {/* Quick Transfer Section */}
        <div className="space-y-5 md:w-1/3">
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
              <Profile image="/images/avatar.png" name="Olivia Lia" job="CEO" />
              <Profile
                image="/images/avatar.png"
                name="Randy Press"
                job="Director"
              />
              <Profile
                image="/images/avatar.png"
                name="Workman"
                job="Designer"
              />
              <Profile image="/images/avatar.png" name="Olivia Lia" job="CEO" />
              <Profile
                image="/images/avatar.png"
                name="Randy Press"
                job="Director"
              />
              <Profile
                image="/images/avatar.png"
                name="Workman"
                job="Designer"
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
        <div className="space-y-5 md:w-2/3">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Balance History</h4>
          </div>
          <div className="bg-white rounded-xl md:shadow-lg md:border md:border-gray-300 p-5">
            <Areachart />
          </div>
        </div>
      </div>
    </div>
  );
};

export default MainDashboard;
