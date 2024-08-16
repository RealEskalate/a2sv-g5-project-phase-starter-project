import React from "react";
import Card from "../components/card/Card";
import CreditCard from "../components/creditCard/CreditCard";
import ExpenseStatisticsChart from "../components/charts/ExpenseStatisticsChart";
import CardList from "../components/cardList/cardList";
import CardExpenseStatisticsChart from "../components/charts/CardExpenseStatisticsChart";
import AddNewCard from "../components/addNewCard/AddNewCard";
import CardSetting from "../components/CardSetting/CardSetting";

const CreditCardsPage = () => {
  return (
    <div className="flex flex-col gap-2">
      <div className="flex max-sm:flex-col gap-[30px]">
        <Card
          title="My Cards"
          className="flex flex-col w-full md:h-[299px] h-[254]"
        >
          <div className="flex  gap-[30px]">
            <div>
              <CreditCard
                balance={1250}
                cardHolder="John Doe"
                expiryDate="12/24"
                cardNumber="1234 5678 9012 3456"
                cardType="secondary" // Can be "primary", "secondary", or "tertiary"
              />
            </div>
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
        </Card>
      </div>
      <div className="flex max-sm:flex-col gap-[30px]">
        <Card
          title="Card Expense Statistics"
          className="flex flex-col max-w-[350px] lg:mx-auto h-auto"
        >
          <CardExpenseStatisticsChart />
        </Card>
        <Card
          title="Card List"
          className="flex flex-col max-w-[730px] lg:mx-auto h-auto"
        >
          <CardList />
        </Card>
      </div>
      <div className="flex max-sm:flex-col gap-[30px]">
        <Card
          title="Add New Card"
          className="flex flex-col min-w-[730px] lg:mx-auto h-auto"
        >
          <AddNewCard />
        </Card>
        <Card
          title="Card Setting"
          className="flex flex-col max-w-[350px] lg:mx-auto h-auto"
        >
          <CardSetting />
        </Card>
      </div>
    </div>
  );
};

export default CreditCardsPage;
