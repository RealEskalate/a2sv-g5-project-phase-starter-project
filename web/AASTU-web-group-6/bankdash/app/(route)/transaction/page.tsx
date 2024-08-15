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
    <div className="space-y- bg-[#F5F7FA] px-10">
      <div className="flex justify-between gap-8 py-4">
        <div className=" w-full">
          <p className="font-semibold text-[22px] text-[#343C6A]">My Cards</p>
        </div>
        <div className="h-[100%] text-right w-full">
          <p className="font-semibold text-[17px] text-[#343C6A] text-right">
            + Add Card
          </p>
        </div>
        <div className="w-full text-right">
          <p className="font-semibold text-[22px] text-[#343C6A]">My Expense</p>
        </div>
      </div>

      <div className="flex gap-8 w-full">
        <div className="w-full">
          <VisaCard isBlack={false} />
        </div>
        <div className="w-full">
          <VisaCard isBlack={true} />
        </div>
        <div className="w-full">
          <BarGraph />
        </div>
      </div>
      <Recent transactions={transactions} />
      <Pagination />
    </div>
  );
};

export default Transaction;
