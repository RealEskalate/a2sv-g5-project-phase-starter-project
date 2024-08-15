import React from "react";
import BarGraph from "@/app/components/Transaction/BarGraph";
import Recent from "@/app/components/Transaction/Recent";
import Pagination from "@/app/components/Transaction/Pagination";
import VisaCard from "@/app/components/Card/VisaCard";

const transactions = [
  {
    description: "Spotify Subscription",
    id: "#123456789",
    type: "Shopping",
    card: "123****",
    date: "28 Jan, 12:36 PM",
    amount: -25.99,
    receipt: "Download",
  },
  {
    description: "Freelance Payment",
    id: "#987654321",
    type: "Income",
    card: "456****",
    date: "27 Jan, 10:45 AM",
    amount: 250.0,
    receipt: "Download",
  },
  // Add more transactions as needed
];

const Transaction = () => {
  return (
    <div className="space-y-6 bg-[#F5F7FA] px-4 sm:px-6 md:px-8 lg:px-10">
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8 w-full">
        <div className="lg:col-span-2 py-4 overflow-x-auto scrollbar-hide">
          <div className="flex justify-between mb-5">
            <p className="font-semibold text-xl sm:text-2xl text-[#343C6A]">
              My Cards
            </p>
            <p className="font-semibold text-base sm:text-lg text-[#343C6A] cursor-pointer">
              + Add Card
            </p>
          </div>
          <div className="overflow-x-auto scrollbar-hide">
            <div className="flex gap-8 min-w-[650px] min-h-[170px]">
              <VisaCard isBlack={false} />
              <VisaCard isBlack={true} />
            </div>
          </div>
        </div>
        <div className="w-full">
          <p className="font-semibold text-xl sm:text-2xl text-[#343C6A] mb-1 py-4">
            My Expense
          </p>
          <BarGraph />
        </div>
      </div>
      <Recent transactions={transactions} />
      <Pagination />
    </div>
  );
};

export default Transaction;
