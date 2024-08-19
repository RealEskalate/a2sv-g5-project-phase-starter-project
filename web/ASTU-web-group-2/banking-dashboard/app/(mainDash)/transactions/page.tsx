import React from "react";
import CardForCreditCards from "../../components/card/CardForCreditCards";
import CreditCard from "../../components/creditCard/CreditCard";
import Card from "../../components/card/Card";
import MyExpenseChart from "../../components/charts/MyExpenseChart";
import TransactionsDisplay from "../../components/transactionsDisplay/TransactionsDisplay";

const TransactionPage = () => {
  return (
    <div className="flex flex-col gap-2">
      <div className="flex max-sm:flex-col gap-[30px]">
        <CardForCreditCards
          title="Credit Cards"
          button="+ Add Card"
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
        <Card title="My Expense" className="max-w-[350px] lg:mx-auto h-auto">
          <MyExpenseChart />
        </Card>
      </div>
     
        <Card
          title="Recent Transactions"
          className="flex flex-col max-sm:w-fit"
        >
          <TransactionsDisplay />
        </Card>
    </div>
  );
};

export default TransactionPage;
