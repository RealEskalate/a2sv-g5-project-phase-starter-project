import React, { useEffect, useState } from "react";
import { Pie_chart } from "./Pie_chart";
import { TransactionContent } from "@/types";
import { getallTransactions } from "@/lib/api";

// Shimmer component for skeleton loading effect
const Shimmer = () => {
  return <div className="animate-pulse h-64 bg-gray-200 rounded-xl"></div>;
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
        const statistics = await getallTransactions(0, 100);
        console.log(statistics);
        setTransactions(statistics?.content || []);
      } finally {
        setLoading(false);
        onLoadingComplete(false);
      }
    };
    fetchData();
  }, [onLoadingComplete]);

  return (
    <div className="md:w-1/3 space-y-5">
      <div className="font-inter text-[16px] font-semibold">
        <h4>Expense Statistics</h4>
      </div>
      <div className="bg-white rounded-xl md:shadow-lg">
        {loading ? <Shimmer /> : <Pie_chart transactions={transactions} />}
      </div>
    </div>
  );
};
