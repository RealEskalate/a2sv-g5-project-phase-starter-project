import React from "react";

import WeeklyActivityChart from "../components/charts/WeeklyActivityChart";
import Card from "../components/card/Card";
import CardForCreditCards from "../components/card/CardForCreditCards";
import CreditCard from "../components/creditCard/CreditCard";
import RecentTransaction from "../components/recent-transaction/RecentTransaction";
import ExpenseStatisticsChart from "../components/charts/ExpenseStatisticsChart";
import SendMoney from "../components/sendMoney/SendMoney";
import BalanceHistoryChart from "../components/charts/BalanceHistoryChart";
const page = () => {
  return (
    <div className="grid grid-cols-1 gap-2 pb-5">
      <div className="grid lg:grid-cols-[2fr_1fr] max-md:grid-cols-1  gap-7 p-4 w-auto">
        <CardForCreditCards
          className="overflow-x-auto"
          title="Credit Cards"
          button="See All"
          link="/credit-cards"
        >
          <div className="flex  gap-7  rounded-lg overflow-x-auto scrollbar-hide">
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
            <div>
              <CreditCard
                balance={1250}
                cardHolder="John Doe"
                expiryDate="12/24"
                cardNumber="1234 5678 9012 3456"
                cardType="secondary" // Can be "primary", "secondary", or "tertiary"
              />
            </div>
            
          </div>
        </CardForCreditCards>
        <Card
          title="Recent Transactions"
          className=" h-auto"
        >
          <RecentTransaction />
        </Card>
      </div>
      <div className="grid lg:grid-cols-[8fr_4fr] max-md:grid-cols-1  gap-7 p-4 w-auto">
        <Card
          title="Weekly Activity"
          className="w-fill h-full"
        >
          <WeeklyActivityChart />
        </Card>
        <Card
          title="Expense Statistics"
          className="w-fill h-full"
        >
          <ExpenseStatisticsChart />
        </Card>
      </div>
      <div className="grid lg:grid-cols-[2fr_3fr] max-md:grid-cols-1  gap-7 p-4 w-auto">
        <Card
          title="Quick Transfer"
          className="w-fill h-full"
        >
          <SendMoney />
        </Card>
        <Card
          title="Balance History"
          className=""
        >
          <BalanceHistoryChart />
        </Card>
      </div>
    </div>
  );
};

export default page;
