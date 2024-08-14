import Image from "next/image";
import React from "react";
import InvestmentCard from "./investmentCard";

const InvestmentItem = () => {
  return (
    <div className="flex flex-col md:flex-row gap-2 md:gap-2 justify-evenly w-full">
      <InvestmentCard
        image={"/assets/icons/moneyBag.svg"}
        name={"Total Invested Amount"}
        amount={`$${150000}`}
      />
      <InvestmentCard
        image={"/assets/icons/numberof-investment.svg"}
        name={"Number of Investments"}
        amount={`${1250}`}
      />
      <InvestmentCard
        image={"/assets/icons/rate-return.svg"}
        name={"Rate of Return"}
        amount={`+${5.8}%`}
      />
    </div>
  );
};

export default InvestmentItem;
