import React, { useEffect, useState } from "react";
import { Pie_chart } from "./Pie_chart";
import { TransactionContent } from "@/types";
import { getallTransactions } from "@/lib/api";
import { useUser } from "@/contexts/UserContext";

// Enhanced Shimmer component for skeleton loading effect
const Shimmer = () => {
  return (
    <div className="animate-pulse space-y-4 ">
      <div className="h-6 bg-gray-300 rounded w-1/4"></div>{" "}
      {/* Simulate title */}
      <div className="flex justify-center">
        <div className="h-48 w-48 bg-gradient-to-r from-gray-300 via-gray-200 to-gray-300 rounded-full"></div>{" "}
        {/* Simulate pie chart */}
      </div>
    </div>
  );
};

export const ExpenseStatistics = ({
  onLoadingComplete,
}: {
  onLoadingComplete: any;
}) => {
  const [transactions, setTransactions] = useState<TransactionContent[]>([]);
  const [loading, setLoading] = useState(true);
  const {isDarkMode} = useUser()
  useEffect(() => {
    const fetchData = async () => {
      try {
        const statistics = await getallTransactions(0, 100);
        setTransactions(statistics?.content || []);
        onLoadingComplete(false);
      } finally {
        setLoading(false);
      }
    };
    fetchData();
  }, [onLoadingComplete]);
  
  return (
    <div className="md:w-1/3 space-y-5">
      <div className="font-inter text-[16px] font-semibold">
        <h4>Expense Statistics</h4>
      </div>
      <div className={`  rounded-xl md:shadow-lg `}>
        {loading || transactions.length === 0 ? (
          <Shimmer />
        ) : (
          <Pie_chart transactions={transactions} />
        )}
      </div>
    </div>
  );
};
