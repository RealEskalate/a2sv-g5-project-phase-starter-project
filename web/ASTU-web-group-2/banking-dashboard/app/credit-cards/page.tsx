import React from "react";
import Card from "../components/card/Card";
import CreditCard from "../components/creditCard/CreditCard";

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
    </div>
  );
};

export default CreditCardsPage;
