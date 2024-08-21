"use client";
import React from "react";
import WeeklyActivityChart from "../components/charts/WeeklyActivityChart";
import Card from "../components/card/Card";
import CardForCreditCards from "../components/card/CardForCreditCards";
import CreditCard from "../components/creditCard/CreditCard";
import RecentTransaction from "../components/recent-transaction/RecentTransaction";
import ExpenseStatisticsChart from "../components/charts/ExpenseStatisticsChart";
import SendMoney from "../components/sendMoney/SendMoney";
import BalanceHistoryChart from "../components/charts/BalanceHistoryChart";
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

const page = () => {
  const { data: session, status } = useSession();
  const router = useRouter();

  useEffect(() => {}, [status, session]);
  console.log(session, status);
  if (!session?.user) router.push("/login");

  const [selectedCardIds, setSelectedCardIds] = useState<string[]>([]);

  const token = session?.user.accessToken || "";
  console.log("accesstoken: ", token);

  const {
    data: allCardsDataWithContent,
    isLoading: isLoadingAllCards,
    isError: isErrorAllCards,
  } = useGetAllCardInfoQuery({
    token: token,
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
  }, [allCardsDataWithContent]); // This effect runs only when allCardsDataWithContent changes

  const {
    data: cardInfoData,
    isLoading: isLoadingCardInfo,
    isError: isErrorCardInfo,
  } = useRetiriveCardInfoQuery(
    {
      id: selectedCardIds.length > 0 ? selectedCardIds[0] : "",
      token: token,
    },
    {
      skip: selectedCardIds.length === 0,
    }
  );

  if (isLoadingAllCards || isLoadingCardInfo) {
    return (
      <div>
        <div className="flex space-x-2 justify-center items-center bg-white h-screen">
          <span className="sr-only">Loading...</span>
          <div className="h-8 w-8 bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 rounded-full animate-bounce [animation-delay:-0.3s]"></div>
          <div className="h-8 w-8 bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 rounded-full animate-bounce [animation-delay:-0.15s]"></div>
          <div className="h-8 w-8 bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 rounded-full animate-bounce"></div>
        </div>
      </div>
    );
  }

  if (isErrorAllCards || isErrorCardInfo) {
    return <div>Error loading data</div>;
  }
  console.log("ehllo");
  console.log("the data we want to see: ", allCardsDataWithContent);
  console.log("the data we don't want to see: ", cardInfoData);
  const allCardsData = allCardsDataWithContent.content!;
  return (
    <div className="flex flex-col gap-2  pb-5">
      <div className="flex max-sm:flex-col justify-between">
        <CardForCreditCards
          className="flex flex-col lg:w-[730px] lg:h-[300px] max-md:w-[350px]"
          title="Credit Cards"
          button="See All"
          link="/credit-cards"
        >
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
                    cardType={index == 0 ? "primary" : "tertiary"}
                  />
                </div>
              ))}
            </div>
          ) : (
            <div>
              <div className="flex space-x-2 justify-center items-center bg-white h-screen">
                <span className="sr-only">Loading...</span>
                <div className="h-8 w-8 bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 rounded-full animate-bounce [animation-delay:-0.3s]"></div>
                <div className="h-8 w-8 bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 rounded-full animate-bounce [animation-delay:-0.15s]"></div>
                <div className="h-8 w-8 bg-gradient-to-r from-purple-400 via-pink-500 to-red-500 rounded-full animate-bounce"></div>
              </div>
            </div>
          )}
        </CardForCreditCards>
        <Card title="Recent Transactions" className="max-w-[350px]  h-auto">
          <RecentTransaction />
        </Card>
      </div>
      <div className="flex max-sm:flex-col justify-between">
        <Card
          title="Weekly Activity"
          className="flex flex-col lg:w-[75%] w-[350px] h-auto"
        >
          <WeeklyActivityChart />
        </Card>
        <Card title="Expense Statistics" className="max-w-[350px]  h-auto">
          <ExpenseStatisticsChart />
        </Card>
      </div>
      <div className="flex max-sm:flex-col gap-[30px]">
        <Card
          title="Quick Transfer"
          className="flex flex-col lg:w-1/3 w-[350px]"
        >
          <SendMoney />
        </Card>
        <Card
          title="Balance History"
          className="flex flex-col max-w-[730px] h-auto"
        >
          <BalanceHistoryChart />
        </Card>
      </div>
    </div>
  );
};

export default page;
