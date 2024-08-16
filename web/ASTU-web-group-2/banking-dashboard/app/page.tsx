import React from "react";
import WeeklyActivityChart from "./components/charts/WeeklyActivityChart";
import Card from "./components/card/Card";
import CardForCreditCards from "./components/card/CardForCreditCards";
import CreditCard from "./components/creditCard/CreditCard";
import Page from "./components/TrendingStock/Page";
import RecentTransaction from "./components/recent-transaction/RecentTransaction";
import ExpenseStatisticsChart from "./components/charts/ExpenseStatisticsChart";
import SendMoney from "./components/sendMoney/SendMoney";
import BalanceHistoryChart from "./components/charts/BalanceHistoryChart";
const page = () => {
  return (
    <div className="flex flex-col gap-2">
      <div className="flex max-sm:flex-col gap-[30px]">
        <CardForCreditCards
          title="Credit Cards"
          button="See All"
          link="/credit-cards"
        >
          <div className="flex  gap-[30px]">
            <div>
              <CreditCard
                balance={1250}
                cardHolder="John Doe"
                expiryDate="12/24"
                cardNumber="1234 5678 9012 3456"
                cardType="primary" // Can be "primary", "secondary", or "tertiary"
              />
            </div>
            <div>
              <CreditCard
                balance={1250}
                cardHolder="John Doe"
                expiryDate="12/24"
                cardNumber="1234 5678 9012 3456"
                cardType="tertiary" // Can be "primary", "secondary", or "tertiary"
              />
            </div>
          </div>
        </CardForCreditCards>
        <Card
          title="Recent Transactions"
          className="max-w-[350px] lg:mx-auto h-[367px]"
        >
          <RecentTransaction />
        </Card>
      </div>
      <div className="flex max-sm:flex-col gap-[30px]">
        <Card
          title="Weekly Activity"
          className="flex flex-col min-w-[730px] lg:mx-auto h-auto"
        >
          <WeeklyActivityChart />
        </Card>
        <Card
          title="Expense Statistics"
          className="flex flex-col max-w-[350px] lg:mx-auto h-auto"
        >
          <ExpenseStatisticsChart />
        </Card>
      </div>
      <div className="flex max-sm:flex-col gap-[30px]">
        <Card
          title="Quick Transfer"
          className="flex flex-col max-w-[350px] lg:mx-auto h-auto"
        >
          <SendMoney />
        </Card>
        <Card
          title="Balance History"
          className="flex flex-col max-w-[730px] lg:mx-auto h-auto"
        >
          <BalanceHistoryChart />
        </Card>
      </div>
    </div>
  );


};

export default page;
