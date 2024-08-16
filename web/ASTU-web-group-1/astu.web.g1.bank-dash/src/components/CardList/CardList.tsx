import React from "react";
import CardListCard from "./CardiListCard";

const CardList = () => {
  return (
    <>
      <div className="flex flex-col items-start  w-full space-y-2">
        <p className="text-[#333B69] pb-2 font-semibold">Card List</p>
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
