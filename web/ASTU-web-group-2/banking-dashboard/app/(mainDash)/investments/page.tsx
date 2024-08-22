import React from "react";
import InfoboxForInvestementPage from "../../components/infobox/InfoboxForInvestementPage";
import MyExpenseChart from "../../components/charts/MyExpenseChart";
import Card from "../../components/card/Card";
import YearlyTotalInvestment from "../../components/charts/YearlyTotalInvestment";
import MonthlyRevenueChart from "../../components/charts/MonthlyRevenueChart";
import MyInvestment from "../../components/myInvestment/myInvestment";
import TrendingStock from "../../components/TrendingStock/TrendingStock";

const InvestmentsPage = () => {
  return (
    <div className="grid grid-cols-1 gap-2 pb-5">
      <InfoboxForInvestementPage />
      <div className="grid lg:grid-cols-2 max-md:grid-cols-1  gap-7 p-4 w-auto">
        <Card
          title="Yearly Total Investment"
          className=""
        >
          <YearlyTotalInvestment />
        </Card>
        <Card
          title="Monthly Revenue"
          className=""
        >
          <MonthlyRevenueChart />
        </Card>
      </div>
      <div className="grid lg:grid-cols-[6fr_4fr] max-md:grid-cols-1  gap-7 p-4 w-auto">
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
