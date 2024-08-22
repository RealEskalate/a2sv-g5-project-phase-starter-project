"use client";
import { useEffect, useRef, useState } from "react";

import {
  BalanceData,
  CardDetails,
  QuickTransferData,
  TransactionContent,
  TransactionData,
} from "@/types";

import getRandomBalance, {
  addTransactions,
  getallTransactions,
  getCreditCards,
  getExpenses,
  getIncomes,
  getQuickTransfer,
} from "@/lib/api";
import { Loading } from "./_components/Loading";
import { useUser } from "@/contexts/UserContext";
import { toast } from "sonner";
import { Cards } from "./_components/Cards";
import { RecentTransaction } from "./_components/RecentTransaction";
import { WeeklyActivity } from "./_components/WeeklyActivity";
import { ExpenseStatistics } from "./_components/ExpenseStatistics";
import { BalanceHistory } from "./_components/BalanceHistory";
import { QuickTransfer } from "./_components/QuickTransfer";
const MainDashboard = () => {
  const { isDarkMode } = useUser();
  const [loading, setLoading] = useState(true);

  const [sendLoading, setSendLoading] = useState(false);

  // const handleSend = async () => {
  //   if (selectedProfile) {
  //     console.log("Sending to:", selectedProfile.username, "Amount:", amount);
  //     setSendLoading(true);
  //     const result: boolean | undefined = await addTransactions({
  //       type: "transfer",
  //       amount: parseInt(amount),
  //       receiverUserName: selectedProfile.username,
  //       description: "Quick Transfer",
  //     });
  //     setLoading(false);
  //     if (result) {
  //       toast("sucess sending");
  //     } else {
  //       toast("failed sending");
  //     }
  //     setLoading(true);
  //   }
  // };

  useEffect(() => {
    const fetchData = async () => {
      try {
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, []);

  if (loading) {
    return <Loading />;
  }
  
  return (
    <div
      className={` relative p-10 space-y-5 ${
        isDarkMode ? "bg-gray-700 text-white" : "bg-[#F5F7FA] text-black"
      }`}
    >
      {/* First Row: My Cards and Recent Transactions */}
      <div className="md:flex sm:grid-cols-2 md:gap-5 space-y-5 md:space-y-0">
        <Cards />
        <RecentTransaction />
      </div>

      {/* Second Row: Weekly Activity and Expense Statistics */}
      <div className="md:flex sm:grid-cols-2 md:gap-5 space-y-5 md:space-y-0">
        <WeeklyActivity />
        <ExpenseStatistics />
      </div>

      {/* Third Row: Quick Transfer and Balance History */}
      <div className="md:grid md:grid-cols-[1fr,2fr] md:gap-10 space-y-5 md:space-y-0">
        <QuickTransfer />
        <BalanceHistory />
      </div>
      {/* {sendLoading && (
        <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 backdrop-blur-md z-50">
          <div role="status">
            <svg
              aria-hidden="true"
              className="w-12 h-12 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600"
              viewBox="0 0 100 101"
              fill="none"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                fill="currentColor"
              />
              <path
                d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                fill="currentFill"
              />
            </svg>
            <span className="sr-only">Loading...</span>
          </div>
        </div>
      )} */}
    </div>
  );
};

export default MainDashboard;
