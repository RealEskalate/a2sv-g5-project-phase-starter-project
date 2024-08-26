import React, { useEffect, useState } from "react";
import { Pie_chart } from "./Pie_chart";
import { TransactionContent } from "@/types";
import { getallTransactions } from "@/lib/api";
import { useUser } from "@/contexts/UserContext";


const Shimmer = () => {
  const {isDarkMode} = useUser();
  return (
    <div
      className={`animate-pulse space-y-4 p-5 ${
        isDarkMode ? "bg-gray-900" : "bg-white"
      }  rounded-2xl`}
    >
      <div className="h-6 bg-gray-300 rounded w-1/4 "></div>

      <div className="flex justify-center">
        <div className="h-52 w-52 bg-gradient-to-r from-gray-300 via-gray-200 to-gray-300 rounded-full"></div>{" "}
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
  useEffect(() => {
    const fetchData = async () => {
      try {
        const statistics = await getallTransactions(0, 50);
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
        <h4 className="lg:text-[22px] md:text-lg text-base">
          Expense Statistics
        </h4>
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
