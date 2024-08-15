import React from "react";
import { FcSimCardChip } from "react-icons/fc";
import Image from "next/image";
import { Transaction } from "./_components/Transaction";
import { Barchart } from "@/components/ui/Barchart";
import { Creditcard } from "./_components/Creditcard";
import { Profile } from "./_components/Profile";
import { PiTelegramLogoLight } from "react-icons/pi";
import { Areachart } from "@/components/ui/Areachart";
import { Piechart } from "@/components/ui/Piechart";

const MainDashboard = () => {
  return (
    <div className="m-2 md:m-4 space-y-5">
      {/* First Row: My Cards and Recent Transactions */}
      <div className="md:grid md:grid-cols-2 md:gap-5 space-y-5 md:space-y-0">
        {/* My Cards Section */}
        <div className="space-y-5">
          <div className="flex justify-between font-inter text-[16px] font-semibold">
            <h4>My Cards</h4>
            <h4>See All</h4>
          </div>
          <div className="flex space-x-5 overflow-x-auto [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]">
            <Creditcard />
            <div className="w-[265px] min-w-[265px] h-[170px] bg-white rounded-xl pt-3 space-y-5 shadow-lg border border-gray-300">
              <div className="flex justify-between px-5">
                <div className="text-gray-700 space-y-[1px]">
                  <p className="font-lato text-[11px] font-normal">Balance</p>
                  <p className="font-lato text-[16px] font-bold">$3,210</p>
                </div>
                <FcSimCardChip size={30} />
              </div>
              <div className="flex justify-between px-5">
                <div className="space-y-[1px]">
                  <p className="font-lato text-[10px] text-gray-400 font-normal">
                    CARD HOLDER
                  </p>
                  <p className="font-lato text-[13px] text-gray-700 font-bold">
                    Sarah Johnson
                  </p>
                </div>
                <div className="space-y-[1px]">
                  <p className="font-lato text-[10px] text-gray-400 font-normal">
                    VALID THRU
                  </p>
                  <p className="font-lato text-[13px] text-gray-700 font-bold">
                    11/23
                  </p>
                </div>
              </div>
              <div className="relative">
                <div className="flex justify-between px-5 items-center py-1 border-t">
                  <p className="font-lato text-[15px] text-gray-700 font-bold">
                    1234****5678
                  </p>
                  <Image
                    src={`/images/intersection.png`}
                    alt={"transaction"}
                    width={27}
                    height={18}
                  />
                </div>
              </div>
            </div>
          </div>
        </div>

        {/* Recent Transactions Section */}
        <div className="space-y-5">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Recent Transactions</h4>
          </div>
          <div className="space-y-2 p-5 bg-white rounded-xl shadow-lg border border-gray-300 max-w-[450px]">
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
      <div className="md:grid md:grid-cols-2 md:gap-5 space-y-5 md:space-y-0">
        {/* Weekly Activity Section */}
        <div className="space-y-5">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Weekly Activity</h4>
          </div>
          <div className="bg-white rounded-xl shadow-lg border border-gray-300 p-5">
            <Barchart />
          </div>
        </div>

        {/* Expense Statistics Section */}
        <div className="space-y-5">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Expense Statistics</h4>
          </div>
          <div className="bg-white rounded-xl shadow-lg border border-gray-300 p-5">
            <Piechart />
          </div>
        </div>
      </div>

      {/* Third Row: Quick Transfer and Balance History */}
      
      <div className="space-y-5">
        <div className="font-inter text-[16px] font-semibold">
          <h4>Quick Transfer</h4>
        </div>
        <div className="bg-white rounded-xl shadow-lg border border-gray-300 p-5 space-y-5">
          <div className="flex space-x-5 overflow-x-auto [&::-webkit-scrollbar]:hidden [-ms-overflow-style:none] [scrollbar-width:none]">
            <Profile image="/images/avatar.png" name="Olivia Lia" job="CEO" />
            <Profile
              image="/images/avatar.png"
              name="Randy Press"
              job="Director"
            />
            <Profile image="/images/avatar.png" name="Workman" job="Designer" />
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
              <button className="bg-[#1814F3] text-white rounded-full px-4 h-[40px] ml-2 flex items-center">
                Send <PiTelegramLogoLight />
              </button>
            </div>
          </div>
        </div>

        {/* Balance History Section */}
        <div className="space-y-2">
          <div className="font-inter text-[16px] font-semibold">
            <h4>Balance History</h4>
          </div>
          <div className="bg-white rounded-xl shadow-lg border border-gray-300 p-5">
            <Areachart />
          </div>
        </div>
      </div>
    </div>
  );
};

export default MainDashboard;
