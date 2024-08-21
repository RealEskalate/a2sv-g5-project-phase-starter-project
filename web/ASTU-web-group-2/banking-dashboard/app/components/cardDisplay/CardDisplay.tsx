"use client";
import React from "react";
import Card from "../card/Card";
import CardForCreditCards from "../card/CardForCreditCards";
import CreditCard from "../creditCard/CreditCard";
import { useState, useEffect } from "react";
import {
  useGetAllCardInfoQuery,
  useRetiriveCardInfoQuery,
} from "@/lib/service/CardService";

import { useSession } from "next-auth/react";
import { useRouter } from "next/navigation";

export interface card {
  card: string;
  id: string;
}

export interface CardData {
  id: string;
  cardHolder: string;
  semiCardNumber: string;
  cardType: "primary" | "secondary" | "tertiary";
  balance: number;
  expiryDate: string;
}

export const formatCardNumber = (number: string): string => {
  return number.replace(/(\d{4})(?=\d)/g, "$1 ");
};

const CardDisplay = () => {
  const { data: session, status } = useSession();
  const router = useRouter();

  useEffect(() => {
    if (!session?.user) router.push("/login");
  }, [status, session]);

  const [selectedCardIds, setSelectedCardIds] = useState<string[]>([]);

  const token = session?.user.accessToken || "";
  console.log("Access token: ", token);

  const {
    data: allCardsDataWithContent,
    isLoading: isLoadingAllCards,
    isError: isErrorAllCards,
  } = useGetAllCardInfoQuery({
    token,
    size: 10,
  });

  useEffect(() => {
    if (allCardsDataWithContent) {
      const allCardsData = allCardsDataWithContent.content;
      if (allCardsData) {
        setSelectedCardIds(
          allCardsData.slice(0, 2).map((card: card) => card.id)
        );
      }
    }
  }, [allCardsDataWithContent]);

  const {
    data: cardInfoData,
    isLoading: isLoadingCardInfo,
    isError: isErrorCardInfo,
  } = useRetiriveCardInfoQuery(
    {
      id: selectedCardIds.length > 0 ? selectedCardIds[0] : "",
      token,
    },
    {
      skip: selectedCardIds.length === 0,
    }
  );

  if (isLoadingAllCards || isLoadingCardInfo) {
    return (
      <div className="flex space-x-2 justify-center items-center bg-white h-screen">
        <span className="sr-only">Loading...</span>
        <div className="h-8 w-8 bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 rounded-full animate-bounce [animation-delay:-0.3s]"></div>
        <div className="h-8 w-8 bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 rounded-full animate-bounce [animation-delay:-0.15s]"></div>
        <div className="h-8 w-8 bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 rounded-full animate-bounce"></div>
      </div>
    );
  }

  if (isErrorAllCards || isErrorCardInfo) {
    return <div>Error loading data</div>;
  }

  const allCardsData = allCardsDataWithContent.content!;
  console.log("The data we want to see: ", allCardsDataWithContent);
  console.log("The data we don't want to see: ", cardInfoData);

  return (
    <div className="flex flex-col gap-2 pb-5">
      <div className="flex max-sm:flex-col justify-between">
        {allCardsData ? (
          <div className="flex gap-[30px]">
            {allCardsData.slice(0, 2).map((card: CardData, index: number) => (
              <div key={index}>
                <CreditCard
                  balance={card.balance}
                  cardHolder={card.cardHolder}
                  expiryDate={new Date(card.expiryDate).toLocaleDateString()}
                  cardNumber={formatCardNumber(
                    cardInfoData?.cardNumber || card.semiCardNumber
                  )}
                  cardType={index === 0 ? "primary" : "tertiary"}
                />
              </div>
            ))}
          </div>
        ) : (
          <div className="flex space-x-2 justify-center items-center bg-white h-screen">
            <span className="sr-only">Loading...</span>
            <div className="h-8 w-8 bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 rounded-full animate-bounce [animation-delay:-0.3s]"></div>
            <div className="h-8 w-8 bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 rounded-full animate-bounce [animation-delay:-0.15s]"></div>
            <div className="h-8 w-8 bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 rounded-full animate-bounce"></div>
          </div>
        )}
      </div>
    </div>
  );
};

export default CardDisplay;
