import React from "react";
import InvestmentListCard from "./InvestmentListCard";

const InvestmentList = () => {
  return (
    <>
      <div className="flex flex-col items-start px-2 w-full md:w-3/5 space-y-4">
        <h1 className="text-deepNavy font-Inter py-2 font-[600] text-xl">My Investment</h1>
        <InvestmentListCard
          companyName="Apple Store"
          amount="54,000"
          returnValue="16%"
          imageUrl="/assets/images/redApple.png"
          sign="+"
          color="#16DBAA"
        />
        <InvestmentListCard
          companyName="Samsung Mobile"
          amount="25,300"
          returnValue="4%"
          imageUrl="/assets/images/google.png"
          sign="-"
          color="#FE5C73"
        />
        <InvestmentListCard
          companyName="Tesla Motors"
          amount="8,200"
          returnValue="25%"
          imageUrl="/assets/images/tesla.png"
          sign="+"
          color="#16DBAA"
        />
      </div>
    </>
  );
};

export default InvestmentList;
