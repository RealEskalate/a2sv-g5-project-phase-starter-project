"use client";
import { useUser } from "@/contexts/UserContext";
import getRandomBalance from "@/lib/api";
import { BalanceData } from "@/types";
import React, { useEffect, useState } from "react";
import { BalanceAreachart } from "../transactions/component/balanceChart";

// Enhanced Shimmer component for skeleton loading effect
const Shimmer = () => {
  return (
    <div className="animate-pulse space-y-4">
      <div className="h-6 bg-gray-300 rounded w-1/4"></div>{" "}
      {/* Simulate title */}
      <div className="h-48 bg-gradient-to-r from-gray-300 via-gray-200 to-gray-300 rounded-xl"></div>{" "}
      {/* Simulate chart area */}
    </div>
  );
};

export const BalanceHistory = ({
  onLoadingComplete,
}: {
  onLoadingComplete: any;
}) => {
  const { isDarkMode } = useUser();
  const [balanceHistory, setBalanceHistory] = useState<BalanceData[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const BalanceHistory = await getRandomBalance();
        setBalanceHistory(BalanceHistory || []);
        onLoadingComplete(false);
      } finally {
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
        p-5
      `}
      >
        {loading || balanceHistory.length === 0 ? (
          <Shimmer />
        ) : (
          <BalanceAreachart balanceHistory={balanceHistory} />
        )}
      </div>
    </div>
  );
};
