"use client";
import { useUser } from "@/contexts/UserContext";
import getRandomBalance from "@/lib/api";
import { BalanceData } from "@/types";
import React, { useEffect, useState } from "react";
import { BalanceAreachart } from "../transactions/component/balanceChart";

// Shimmer component for skeleton loading effect
const Shimmer = () => {
  return <div className="animate-pulse h-64 bg-gray-200 rounded-xl"></div>;
};

export const BalanceHistory = ({ onLoadingComplete }: { onLoadingComplete: any }) => {
  const { isDarkMode } = useUser();
  const [balanceHistory, setBalanceHistory] = useState<BalanceData[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const BalanceHistory = await getRandomBalance();
        setBalanceHistory(BalanceHistory || []);
      } finally {
        onLoadingComplete(false);
        setLoading(false);
      }
    };
    fetchData();
  }, [onLoadingComplete]);

  return (
    <div className="space-y-5">
      <div className="font-inter text-[16px] font-semibold">
        <h4>Balance History</h4>
      </div>
      <div
        className={`
        ${isDarkMode ? "bg-gray-800 shadow-md" : "bg-white shadow-lg"}
        rounded-xl
        md:shadow
        transition-all
        duration-300
      `}
      >
        {loading ? (
          <Shimmer />
        ) : (
          <BalanceAreachart balanceHistory={balanceHistory} />
        )}
      </div>
    </div>
  );
};
