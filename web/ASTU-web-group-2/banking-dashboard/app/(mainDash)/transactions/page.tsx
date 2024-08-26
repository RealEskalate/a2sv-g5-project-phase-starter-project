import React from "react";
import CardForCreditCards from "../../components/virtualCards/card/CardForCreditCards";
import CreditCard from "../../components/virtualCards/creditCard/CreditCard";
import Card from "../../components/virtualCards/card/Card";
import MyExpenseChart from "../../components/charts/MyExpenseChart";
import TransactionsDisplay from "../../components/transactions/transactionsDisplay/TransactionsDisplay";
import CardDisplay from "@/app/components/virtualCards/cardDisplay/CardDisplay";

const TransactionPage = () => {
  return (
    <div className="grid grid-cols-1 pb-5 w-full">
      <div className="grid lg:grid-cols-[3fr_1fr] max-md:grid-cols-1  gap-7 p-4 w-auto">
        <CardForCreditCards
          className="overflow-x-auto "
          title="Credit Cards"
          button="+ Add Card"
          link="/credit-cards"
        >
          <CardDisplay numofcard={3}/>
        </CardForCreditCards>
        <Card title="My Expense" className="">
          <MyExpenseChart />
        </Card>
      </div>

      <Card title="Recent Transactions" className="w-full">
        <TransactionsDisplay />
      </Card>
    </div>
  );
};

export default TransactionPage;
