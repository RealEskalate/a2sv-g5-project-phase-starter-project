import React from "react";
import CardListCard from "./CardiListCard";

const CardList = () => {
  return (
    <>
      <div className="flex flex-col  w-full  ">
        <p className="text-[#333B69] pb-2 font-semibold">Card List</p>
        <div className="h-80 lg:h-[16.5rem] xl:h-80 overflow-y-scroll whitespace-nowrap scroll-smooth">
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
          <CardListCard
            cardType="Secondary"
            bank="BRC Bank"
            cardNumber="**** **** 4300"
            imageUrl="/assets/images/cardList.png"
            namainCard="Michael"
          />
          <CardListCard
            cardType="Secondary"
            bank="BRC Bank"
            cardNumber="**** **** 4300"
            imageUrl="/assets/images/cardList.png"
            namainCard="Michael"
          />
        </div>
      </div>
    </>
  );
};

export default CardList;
