"use client";
import { useUser } from "@/contexts/UserContext";
import React, { useEffect, useState } from "react";
import { Barchart } from "../transactions/component/weeklyActivityChart";
import { getExpenses, getIncomes } from "@/lib/api";
import { TransactionContent } from "@/types";

// Shimmer component for skeleton loading effect
const Shimmer = () => {
  return (
    <div className="animate-pulse flex flex-col space-y-4 p-5">
      <div className="h-6 bg-gray-300 rounded"></div>
      <div className="h-48 bg-gray-300 rounded"></div>
    </div>
  );
};

export const WeeklyActivity = ({ onLoadingComplete }: { onLoadingComplete: any }) => {
  const { isDarkMode } = useUser();
  const [weeklyIncome, setWeeklyIncome] = useState<TransactionContent[]>([]);
  const [weeklyWithdraw, setWeeklyWithdraw] = useState<TransactionContent[]>(
    []
  );
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const incomes = await getIncomes(0, 7);
        const withdraw = await getExpenses(0, 7);
        setWeeklyIncome(incomes?.content || []);
        setWeeklyWithdraw(withdraw?.content || []);
         onLoadingComplete(false);
           
      } finally {
       
      setLoading(false);
      }
    };
    fetchData();
  }, [onLoadingComplete]);

  return (
    <div className="md:w-2/3 space-y-5">
      <div className="font-inter text-[16px] font-semibold">
        <h4>Weekly Activity</h4>
      </div>
      <div
        className={`${
          isDarkMode
            ? "bg-gray-800 text-white border-gray-600"
            : "bg-white text-black "
        }  md:shadow-lg  rounded-xl `}
      >
        {loading || (weeklyIncome.length === 0 || weeklyWithdraw.length === 0) ? (
          <Shimmer />
        ) : (
          <Barchart
            weeklyDeposit={weeklyIncome}
            weeklyWithdraw={weeklyWithdraw}
          />
        )}
      </div>
    </div>
  );
};
