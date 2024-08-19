'use client';
import { useGetAllCardInfoQuery, useLazyRetiriveCardInfoQuery } from "@/lib/service/CardServices";
import React, { useState, useEffect } from "react";
import Link from "next/link";
import { useSession } from "next-auth/react";

const icons = [
  "/assets/cardlist/card1.svg",
  "/assets/cardlist/card2.svg",
  "/assets/cardlist/card3.svg",
];

interface Card {
  id: string;
  cardType: string;
  cardHolder: string;
}

interface FullCard extends Card {
  cardNumber: string;
  bank?: string; // Optional bank property
}

const CardList = () => {
  const session = useSession();
  const accessToken = session.data?.user.accessToken || "";
  
  const { data: cardsData, isLoading, error } = useGetAllCardInfoQuery(accessToken);
  const [cardDetails, setCardDetails] = useState<FullCard[]>([]);
  const [retrieveCardInfo, { data: cardDetailsData }] = useLazyRetiriveCardInfoQuery();

  useEffect(() => {
    const fetchFullCardDetails = async () => {
      if (cardsData && cardsData.length > 0) {
        const limitedData = cardsData.slice(0, 3); // Limit to the first 3 cards

        const fullCardsPromises = limitedData.map(async (card: Card) => {
          const { data: cardDetails } = await retrieveCardInfo({ token: accessToken, id: card.id });

          return {
            ...card,
            cardNumber: cardDetails?.cardNumber || "",
            bank: cardDetails?.bank || "DBM Bank", // Default bank value
          };
        });

        const fullCards = await Promise.all(fullCardsPromises);
        setCardDetails(fullCards);
        console.log("Fetched full cards:", fullCards);
      }
    };

    fetchFullCardDetails();
  }, [cardsData, accessToken, retrieveCardInfo]);

  if (isLoading) {
    return <div>Loading cards...</div>;
  }

  if (error) {
    return <div>Error fetching cards</div>;
  }

  if (!cardsData || cardsData.length === 0) {
    return <div>No cards available</div>;
  }

  const displayedCards = cardDetails.length > 0 ? cardDetails.slice(0, 3) : [];

  return (
    <div className="sm:w-[475px] md:w-[730px]">
      {displayedCards.length === 0 ? (
        <div>No card details available</div>
      ) : (
        displayedCards.map((card: FullCard, index) => (
          <div
            key={card.id}
            className="grid grid-flow-col h-[69px] lg:h-[90px] justify-between mb-[10px] sm:mb-[15px] items-center pl-[20px] bg-white rounded-3xl grid-col-12"
          >
            <div className="col-span-1">
              <img src={icons[index]} className="lg:w-[60px] w-[45px]" alt={`Card ${index}`} />
            </div>
            <div className="col-span-2">
              <p className="text-[14px] md:[text-[12px]] lg:text-[16px] text-[#333B69]">
                Card Type
              </p>
              <span className="text-[12px] sm:text-[15px] text-[#718EBF]">
                {card.cardType}
              </span>
            </div>
            <div className="col-span-[2.5]">
              <p className="text-[14px] md:[text-[12px]] lg:text-[16px] text-[#333B69]">
                Bank
              </p>
              <span className="text-[12px] md:[text-[12px]] lg:text-[16px] text-[#718EBF]">
                {card.bank}
              </span>
            </div>
            <div className="hidden col-span-[2.5] sm:block">
              <p className="text-[14px] md:text-[12px] lg:text-[16px] text-[#333B69] font-medium">
                Card Number
              </p>
              <span className="text-[12px] sm:text-[15px] text-[#718EBF]">
                {'*'.repeat(4)} {'*'.repeat(4)} {card.cardNumber.slice(-4)}
              </span>
            </div>
            <div className="hidden col-span-2 sm:block">
              <p className="text-[14px] sm:text-[16px] text-[#333B69] font-medium">
                Card Holder
              </p>
              <span className="text-[12px] sm:text-[15px] text-[#718EBF]">
                {card.cardHolder}
              </span>
            </div>
            <div className="col-span-2">
              <p className="text-[14px] sm:text-[16px] text-[#1814F3] font-medium">
                <Link href="#">View Detail</Link>
              </p>
            </div>
          </div>
        ))
      )}
    </div>
  );
};

export default CardList;
