import React from "react";
// import { CardData } from "@/types/cardData";
import { Card } from "@/lib/redux/types/cards";
import CardBox from "@/app/components/cardBox";
import {
  useDeleteCardMutation,
  useGetCardsQuery,
} from "@/lib/redux/api/cardsApi";
import Loading from "@/app/loading";

function formatCardNumber(cardNumber: string): string {
  const start = cardNumber.slice(0, 4);
  const end = cardNumber.slice(-4);
  return `${start} **** **** ${end}`;
}

const colorOptions = [
  { svgColor: "#396AFF", svgBgColor: "#E7EDFF" },
  { svgColor: "#FF82AC", svgBgColor: "#FFE0EB" },
  { svgColor: "#FFD700", svgBgColor: "#FFF5D9" },
];

const CardListPage = () => {
  const { data, error, isLoading } = useGetCardsQuery({ page: 0, size: 10 });

  if (isLoading) return <Loading/>;
  if (error) return <div>Error fetching cards</div>;

  // const cardData: CardData[] = [
  //   {
  //     cardType: "Secondary",
  //     bank: "DBL Bank",
  //     detailsLink: "/creditcardpage/abc",
  //     svgColor: "#396AFF",
  //     svgBgColor: "#E7EDFF",
  //     cardNumber: "**** **** **** 1234",
  //     NamainCard: "William",
  //   },
  //   {
  //     cardType: "Secondary",
  //     bank: "BRC Bank",
  //     detailsLink: "/creditcardpage/abc",
  //     svgColor: "#FF82AC",
  //     svgBgColor: "#FFE0EB",
  //     cardNumber: "**** **** **** 1234",
  //     NamainCard: "Michael",
  //   },
  //   {
  //     cardType: "Secondary",
  //     bank: "ABM Bank",
  //     detailsLink: "/creditcardpage/xyz",
  //     svgColor: "#FFD700",
  //     svgBgColor: "#FFF5D9",
  //     cardNumber: "**** **** **** 1234",
  //     NamainCard: "Edward",
  //   },
  // ];

  return (
    // const colorOption = colorOptions[index % colorOptions.length];
    <div className="card-list-container  ">
      {data?.content.map((card: Card, index: number) => {
        const colorOption = colorOptions[index % colorOptions.length]; // Cycle through color options

        return (
          <CardBox
            key={index}
            cardType={card.cardType}
            bank="CBE Bank"
            detailsLink="/creditcardpage/xyz"
            svgColor={colorOption.svgColor}
            svgBgColor={colorOption.svgBgColor}
            cardNumber={formatCardNumber(card.semiCardNumber)}
            NamainCard={card.cardHolder}
          />
        );
      })}
    </div>
  );
};

export default CardListPage;
