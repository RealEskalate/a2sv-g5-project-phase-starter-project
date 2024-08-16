import React from "react";
import LoanCard from "./LoanCard";

const Loansitem = () => {
  return (
    <div
      className="flex overflow-x-auto justify-around overflow-clip whitespace-nowrap w-full"
      style={{
        scrollbarWidth: "none",
        msOverflowStyle: "none",
      }}
    >
      <LoanCard
        image={"/assets/icons/personal.svg"}
        name={"Personal Loans"}
        amount={`$${50000}`}
      />
      <LoanCard
        image={"/assets/icons/bag.svg"}
        name={"Corporate Loans"}
        amount={`$${100000}`}
      />
      <LoanCard
        image={"/assets/icons/businesstrack.svg"}
        name={"Business Loans"}
        amount={`$${500000}`}
      />
      <LoanCard
        image={"/assets/icons/customLoan.svg"}
        name={"Custom Loans"}
        amount={`Choose Money`}
      />
    </div>
  );
};

export default Loansitem;
