import React from "react";
import InvestmentItem from "@/components/InvestmentItems/InvestmentItem";
import InvestmentList from "@/components/InvestmentList/InvestmentList";
import TrendingList from "@/components/TrendingStock/trendingList";
import YearlyTotalInvestment from "@/components/Charts/YearlyTotalInvestment";
import MonthlyRevenue from "@/components/Charts/MonthlyRevenue";
import StoreProvider from "@/providers/StoreProvider";

export default function page() {
  return (
    <div>
      <StoreProvider>
        <InvestmentItem />
      </StoreProvider>
      <div className="md:flex space-x-4">
        <YearlyTotalInvestment />
        <MonthlyRevenue />
      </div>
      <div className="flex flex-col md:flex-row ">
        <InvestmentList />
        <TrendingList />
      </div>
    </div>
  );
}
