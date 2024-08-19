import React from "react";

import WeeklyActivityChart from "./components/charts/WeeklyActivityChart";
import Card from "./components/card/Card";
import CardForCreditCards from "./components/card/CardForCreditCards";
import CreditCard from "./components/creditCard/CreditCard";
import RecentTransaction from "./components/recent-transaction/RecentTransaction";
import ExpenseStatisticsChart from "./components/charts/ExpenseStatisticsChart";
import SendMoney from "./components/sendMoney/SendMoney";
import BalanceHistoryChart from "./components/charts/BalanceHistoryChart";
const page = () => {
  return (
    <div className="flex flex-col gap-2 px-2 pb-5">
      <div className="flex max-sm:flex-col justify-between">
        <CardForCreditCards
          className="flex flex-col lg:w-[730px] lg:h-[300px] max-md:w-[350px]"
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
          className="max-w-[350px]  h-auto"
        >
          <RecentTransaction />
        </Card>
      </div>
      <div className="flex max-sm:flex-col justify-between">
        <Card
          title="Weekly Activity"
          className="flex flex-col lg:w-[730px] lg:h-[300px] max-md:w-[350px]"
        >
          <WeeklyActivityChart />
        </Card>
        <Card
          title="Expense Statistics"
          className="max-w-[350px]  h-auto"
        >
          <ExpenseStatisticsChart />
        </Card>
      </div>
      <div className="flex max-sm:flex-col gap-[30px]">
        <Card
          title="Quick Transfer"
          className="flex flex-col lg:w-1/3 w-full"
        >
          <SendMoney />
        </Card>
        <Card
          title="Balance History"
          className="flex flex-col lg:w-2/3 w-full h-[285px]"
        >
          <BalanceHistoryChart />
        </Card>
      </div>
    </div>
  );
};

export default page;
