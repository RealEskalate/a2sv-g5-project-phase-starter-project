import React from "react";

import WeeklyActivityChart from "../../components/charts/WeeklyActivityChart";
import Card from "../../components/virtualCards/card/Card";
import CardForCreditCards from "../../components/virtualCards/card/CardForCreditCards";
import CreditCard from "../../components/virtualCards/creditCard/CreditCard";
import RecentTransaction from "../../components/recent-transaction/RecentTransaction";
import ExpenseStatisticsChart from "../../components/charts/ExpenseStatisticsChart";
import SendMoney from "../../components/sendMoney/SendMoney";
import BalanceHistoryChart from "../../components/charts/BalanceHistoryChart";
import CardDisplay from "../../components/virtualCards/cardDisplay/CardDisplay";

const page = () => {
  return (
    <div className="grid grid-cols-1 gap-2 pb-5">
      <div className="grid lg:grid-cols-[2fr_1fr] max-md:grid-cols-1  gap-7 p-4 w-auto">
        <CardForCreditCards
          className="overflow-x-auto scroll-hide"
          title="Credit Cards"
          button="See All"
          link="/credit-cards"
        >
          <CardDisplay numofcard={3} />
        </CardForCreditCards>
        <Card title="Recent Transactions" className=" h-auto">
          <RecentTransaction />
        </Card>
      </div>
      <div className="grid lg:grid-cols-[8fr_4fr] max-md:grid-cols-1  gap-7 p-4 w-auto">
        <Card title="Weekly Activity" className="w-fill h-full">
          <WeeklyActivityChart />
        </Card>
        <Card title="Expense Statistics" className="w-fill h-full">
          <ExpenseStatisticsChart />
        </Card>
      </div>
      <div className="grid lg:grid-cols-[2fr_3fr] max-md:grid-cols-1  gap-7 p-4 w-auto">
        <Card title="Quick Transfer" className="w-fill h-full">
          <SendMoney />
        </Card>
        <Card title="Balance History" className="">
          <BalanceHistoryChart />
        </Card>
      </div>
    </div>
  );
};

export default page;
