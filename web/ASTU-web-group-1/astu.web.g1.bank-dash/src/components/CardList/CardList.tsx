import React from "react";
import CardListCard from "./CardiListCard";

const CardList = () => {
  return (
    <>
      <div className="flex flex-col items-start px-4 w-full md:w-2/3 space-y-2">
        <h1 className="text-[#333B69] py-2 font-semibold">Card List</h1>
        <CardListCard
          cardType="Secondary"
          bank="DBL Bank"
          cardNumber="**** **** 5600"
          imageUrl="/assets/images/cardList.png"
          namainCard="William"
        />
        <CardListCard
          cardType="Secondary"
          bank="BRC Bank"
          cardNumber="**** **** 4300"
          imageUrl="/assets/images/cardList.png"
          namainCard="Michael"
        />
        <CardListCard
          cardType="Secondary"
          bank="ABM Bank"
          cardNumber="**** **** 7560"
          imageUrl="/assets/images/cardList.png"
          namainCard="Edward"
        />
      </div>
    </>
  );
};

export default CardList;
