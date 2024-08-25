"use client";
import React from "react";
import CardListCard from "./CardiListCard";
import { useGetAllCardsQuery } from "@/lib/redux/slices/cardSlice";

const CardList = () => {
  const { data, isLoading } = useGetAllCardsQuery({ page: 1, size: 5 });
  if (isLoading) return <div>Loading</div>;
  const imageList = ["/assets/images/cardList.png"];
  const bankList = ["CBE", "DBL Bank", "BRC Bank", "ABM Bank"];
  return (
    <>
      <div className="flex flex-col  w-full  ">
        <p className="text-[#333B69] pb-2 font-semibold">Card List</p>
        <div className="h-80 lg:h-[16.5rem] xl:h-80 overflow-y-scroll whitespace-nowrap scroll-smooth">
          {data?.content?.map((card, index) => (
            <CardListCard
              key={card.id}
              cardType={card.cardType}
              bank={bankList[index % bankList.length]}
              cardNumber={card.semiCardNumber}
              imageUrl={imageList[index % imageList.length]}
              namainCard={card.cardHolder}
            />
          ))}
        </div>
      </div>
    </>
  );
};

export default CardList;
