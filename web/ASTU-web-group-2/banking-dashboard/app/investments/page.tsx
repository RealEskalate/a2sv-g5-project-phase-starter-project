import React from "react";
import InfoboxForInvestementPage from "../components/infobox/InfoboxForInvestementPage";
import MyExpenseChart from "../components/charts/MyExpenseChart";
import Card from "../components/card/Card";
import YearlyTotalInvestment from "../components/charts/YearlyTotalInvestment";
import MonthlyRevenueChart from "../components/charts/MonthlyRevenueChart";
import MyInvestment from "../components/myInvestment/myInvestment";
import TrendingStock from "../components/TrendingStock/TrendingStock";

const InvestmentsPage = () => {
  return (
    <div className="flex flex-col gap-2 pb-5">
      <InfoboxForInvestementPage />
      <div className="flex max-sm:flex-col justify-between max-sm:w-[350px]">
        <Card
          title="Yearly Total Investment"
          className="max-sm:w-full w-[540px]"
        >
          <YearlyTotalInvestment />
        </Card>
        <Card
          title="Monthly Revenue"
          className="max-sm:w-full w-[540px]"
        >
          <MonthlyRevenueChart />
        </Card>
      </div>
      <div className="flex max-sm:flex-col justify-between max-sm:w-[350px]">
        <Card
          title="My Investment"
          className=""
        >
          <MyInvestment />
        </Card>
        <Card
          title="Trending Stock"
          className="max-sm:w-full"
        >
          <TrendingStock />
        </Card>
      </div>
    </div>
  );
};

export default InvestmentsPage;
