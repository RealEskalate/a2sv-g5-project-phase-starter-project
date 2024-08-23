"use client";
import { useUser } from "@/contexts/UserContext";
import { getallTransactions } from "@/lib/api";
import { TransactionContent } from "@/types";
import React, { useEffect, useState } from "react";
import { Transaction } from "./Transaction";

// Shimmer component for skeleton loading effect
const Shimmer = () => {
  return (
    <div className="animate-pulse flex space-x-4">
      <div className="rounded-full bg-gray-300 h-12 w-12"></div>
      <div className="flex-1 space-y-4 ">
        <div className="h-4 bg-gray-300 rounded w-3/4"></div>
        <div className="h-4 bg-gray-300 rounded"></div>
      </div>
    </div>
  );
};

export const RecentTransaction = ({onLoadingComplete}:{onLoadingComplete:any}) => {
  const { isDarkMode } = useUser();
  const [recentTransactions, setRecentTransactions] = useState<
    TransactionContent[]
  >([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const recent = await getallTransactions(0, 3);
        setRecentTransactions(recent?.content || []);
         onLoadingComplete(false);
           setLoading(false);
      } finally {
      }
    };
    fetchData();
  }, [onLoadingComplete]);

  return (
    <div className=" md:space-y-5 w-full md:w-1/3">
      <div className="font-inter text-[16px] font-semibold">
        <h4>Recent Transactions</h4>
      </div>
      <div
        className={`space-y-5  md:p-5 p-3 md:h-[200px] lg:w-[365px] lg:h-[220px] ${
          isDarkMode
            ? "bg-gray-800 text-white border-gray-600"
            : "bg-white text-black"
        }
        rounded-xl
        md:shadow-lg
          `}
      >
        {loading
          ? [1, 2, 3].map((index) => <Shimmer key={index} />)
          : recentTransactions.map((transaction) => (
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
  );
};
