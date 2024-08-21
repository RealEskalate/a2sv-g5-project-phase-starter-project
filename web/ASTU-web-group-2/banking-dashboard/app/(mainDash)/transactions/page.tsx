"use client";
import CardForCreditCards from "../../components/card/CardForCreditCards";
import Card from "../../components/card/Card";
import MyExpenseChart from "../../components/charts/MyExpenseChart";
import TransactionsDisplay from "../../components/transactionsDisplay/TransactionsDisplay";
import CardDisplay from "@/app/components/cardDisplay/CardDisplay";

const TransactionPage = () => {
  return (
    <div className="flex flex-col gap-2  pb-5">
      <div className="flex max-sm:flex-col justify-between">
        <CardForCreditCards
          className="flex flex-col lg:w-[730px] lg:h-[300px] max-md:w-[350px]"
          title="Credit Cards"
          button="+ Add Card"
          link="/credit-cards"
        >
          <CardDisplay />
        </CardForCreditCards>
        <Card title="My Expense" className="w-[350px]  h-auto lg:pl-6 pl-0">
          <MyExpenseChart />
        </Card>
      </div>

      <Card title="Recent Transactions" className="flex flex-col w-[100%]">
        <TransactionsDisplay />
      </Card>
    </div>
  );
};

export default TransactionPage;
