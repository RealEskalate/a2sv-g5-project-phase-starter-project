import React from "react";
import CardBox from "../cardBox/page";

interface CardData {
  cardType: string;
  bank: string;
  detailsLink: string;
  svgColor: string;
  svgBgColor: string;
  cardNumber: string;
  NamainCard: string;
}

const CardListPage = () => {
  const cardData: CardData[] = [
    {
      cardType: "Secondary",
      bank: "DBL Bank",
      detailsLink: "/creditcardpage/abc",
      svgColor: "#396AFF",
      svgBgColor: "#E7EDFF",
      cardNumber: "**** **** **** 1234",
      NamainCard: "William",
    },
    {
      cardType: "Secondary",
      bank: "BRC Bank",
      detailsLink: "/creditcardpage/abc",
      svgColor: "#FF82AC",
      svgBgColor: "#FFE0EB",
      cardNumber: "**** **** **** 1234",
      NamainCard: "Michael",
    },
    {
      cardType: "Secondary",
      bank: "ABM Bank",
      detailsLink: "/creditcardpage/xyz",
      svgColor: "#FFD700",
      svgBgColor: "#FFF5D9",
      cardNumber: "**** **** **** 1234",
      NamainCard: "Edward",
    },
  ];

  return (
    <div className="card-list-container  md:flex md:flex-col ">
      {cardData.map((card, index) => (
        <CardBox
          key={index}
          cardType={card.cardType}
          bank={card.bank}
          detailsLink={card.detailsLink}
          svgColor={card.svgColor}
          svgBgColor={card.svgBgColor}
          cardNumber={card.cardNumber}
          NamainCard={card.NamainCard}
        />
      ))}
    </div>
  );
};

export default CardListPage;
