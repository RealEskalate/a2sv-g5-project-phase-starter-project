"use client";
import React, { useEffect, useState } from "react";
import CardForCreditCards from "../../components/card/CardForCreditCards";
import CreditCard from "../../components/creditCard/CreditCard";
import Card from "../../components/card/Card";
import MyExpenseChart from "../../components/charts/MyExpenseChart";
import TransactionsDisplay from "../../components/transactionsDisplay/TransactionsDisplay";
import { card, CardData, formatCardNumber } from "@/app/(mainDash)/page";
import { useSession } from "next-auth/react";
import { useRouter } from "next/navigation";
import {
  useGetAllCardInfoQuery,
  useRetiriveCardInfoQuery,
} from "@/lib/service/CardService";

const TransactionPage = () => {
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
          button="+ Add Card"
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
        <Card title="My Expense" className="w-[350px]  h-auto lg:pl-6 pl-0">
          <MyExpenseChart />
        </Card>
      </div>

      <Card title="Recent Transactions" className="flex flex-col w-[100%]">
        <TransactionsDisplay />
      </Card>
    </div>
  );
};

export default TransactionPage;
