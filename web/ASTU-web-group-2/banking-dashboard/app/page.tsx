import React from "react";
import CreditCard from "./components/creditCard/CreditCard";
const page = () => {
  return (
    <div>
      <CreditCard
        balance={1250}
        cardHolder="John Doe"
        expiryDate="12/24"
        cardNumber="1234 5678 9012 3456"
        cardType="tertiary" // Can be "primary", "secondary", or "tertiary"
      />
    </div>
  );
};

export default page;
