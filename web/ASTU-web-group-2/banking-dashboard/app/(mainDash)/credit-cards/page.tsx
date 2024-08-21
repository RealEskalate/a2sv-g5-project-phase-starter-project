import React from "react";
import Card from "../../components/card/Card";
import CreditCard from "../../components/creditCard/CreditCard";
import ExpenseStatisticsChart from "../../components/charts/ExpenseStatisticsChart";
import CardList from "../../components/cardList/cardList";
import CardExpenseStatisticsChart from "../../components/charts/CardExpenseStatisticsChart";
import AddNewCard from "../../components/addNewCard/AddNewCard";
import CardSetting from "../../components/CardSetting/CardSetting";
import CardForCreditCards from "@/app/components/card/CardForCreditCards";

const CreditCardsPage = () => {
  return (
    <div className="grid grid-cols-1 gap-2 pb-5">
      <div className="grid lg:grid-cols-1 p-4">
        <CardForCreditCards
          title="My Cards"
          className="overflow-x-auto"
          link="/credit-cards"
          button=""
        >
          <div className="grid grid-cols-3 justify-between w-full min-w-[1150px] overflow-x-auto rounded-3xl">
              <CreditCard
                balance={1250}
                cardHolder="John Doe"
                expiryDate="12/24"
                cardNumber="1234 5678 9012 3456"
                cardType="secondary" // Can be "primary", "secondary", or "tertiary"
              />
              <CreditCard
                balance={1250}
                cardHolder="John Doe"
                expiryDate="12/24"
                cardNumber="1234 5678 9012 3456"
                cardType="primary" // Can be "primary", "secondary", or "tertiary"
              />
              <CreditCard
                balance={1250}
                cardHolder="John Doe"
                expiryDate="12/24"
                cardNumber="1234 5678 9012 3456"
                cardType="tertiary" // Can be "primary", "secondary", or "tertiary"
              />
          </div>
        </CardForCreditCards>
      </div>
      <div className="grid lg:grid-cols-[4fr_6fr] max-md:grid-cols-1  gap-7 p-4">
        <Card
          title="Card Expense Statistics"
          className=""
        >
          <CardExpenseStatisticsChart />
        </Card>
        <Card
          title="Card List"
          className=""
        >
          <CardList />
        </Card>
      </div>
      <div className="grid lg:grid-cols-[6fr_4fr] max-md:grid-cols-1  gap-7 p-4">
        <Card
          title="Add New Card"
          className=""
        >
          <AddNewCard />
        </Card>
        <Card
          title="Card Setting"
          className=""
        >
          <CardSetting />
        </Card>
      </div>
    </div>
  );
};

export default CreditCardsPage;
