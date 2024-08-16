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
    <div className="flex flex-col gap-2">
      <InfoboxForInvestementPage />
      <div className="flex max-sm:flex-col gap-[30px]">
        <Card
          title="Yearly Total Investment"
          className="w-[540px] h-[329px] "
        >
          <YearlyTotalInvestment />
        </Card>
        <Card
          title="Monthly Revenue"
          className="w-[540px] h-[329px]"
        >
          <MonthlyRevenueChart />
        </Card>
      </div>
      <div className="flex max-sm:flex-col gap-[30px]">
        <Card
          title="My Investment"
          className=""
        >
          <MyInvestment />
        </Card>
        <Card
          title="Trending Stock"
          className=""
        >
          <TrendingStock />
        </Card>
      </div>
    </div>
  );
};

export default InvestmentsPage;
