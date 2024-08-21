"use client";
import React from "react";
import WeeklyActivityChart from "../components/charts/WeeklyActivityChart";
import Card from "../components/card/Card";
import CardForCreditCards from "../components/card/CardForCreditCards";
import RecentTransaction from "../components/recent-transaction/RecentTransaction";
import ExpenseStatisticsChart from "../components/charts/ExpenseStatisticsChart";
import SendMoney from "../components/sendMoney/SendMoney";
import BalanceHistoryChart from "../components/charts/BalanceHistoryChart";
import CardDisplay from "../components/cardDisplay/CardDisplay";

const page = () => {
  return (
    <div className="flex flex-col gap-2  pb-5">
      <div className="flex max-sm:flex-col justify-between">
        <CardForCreditCards
          className="flex flex-col lg:w-[730px] lg:h-[300px] max-md:w-[350px]"
          title="Credit Cards"
          button="See All"
          link="/credit-cards"
        >
          <CardDisplay />
        </CardForCreditCards>
        <Card title="Recent Transactions" className="max-w-[350px]  h-auto">
          <RecentTransaction />
        </Card>
      </div>
      <div className="flex max-sm:flex-col justify-between">
        <Card
          title="Weekly Activity"
          className="flex flex-col lg:w-[75%] w-[350px] h-auto"
        >
          <WeeklyActivityChart />
        </Card>
        <Card title="Expense Statistics" className="max-w-[350px]  h-auto">
          <ExpenseStatisticsChart />
        </Card>
      </div>
      <div className="flex max-sm:flex-col gap-[30px]">
        <Card
          title="Quick Transfer"
          className="flex flex-col lg:w-1/3 w-[350px]"
        >
          <SendMoney />
        </Card>
        <Card
          title="Balance History"
          className="flex flex-col max-w-[730px] h-auto"
        >
          <BalanceHistoryChart />
        </Card>
      </div>
    </div>
  );
};

export default page;
