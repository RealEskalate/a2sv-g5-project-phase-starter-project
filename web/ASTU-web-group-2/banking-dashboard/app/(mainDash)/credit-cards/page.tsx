"use client";
import React, { useEffect, useState } from "react";
import Card from "../../components/card/Card";
import CreditCard from "../../components/creditCard/CreditCard";
import ExpenseStatisticsChart from "../../components/charts/ExpenseStatisticsChart";
import CardList from "../../components/cardList/cardList";
import CardExpenseStatisticsChart from "../../components/charts/CardExpenseStatisticsChart";
import AddNewCard from "../../components/addNewCard/AddNewCard";
import CardSetting from "../../components/CardSetting/CardSetting";
import CardForCreditCards from "@/app/components/card/CardForCreditCards";
import CardDisplay, {
  CardData,
  formatCardNumber,
  card,
} from "../../components/cardDisplay/CardDisplay";
import {
  useGetAllCardInfoQuery,
  useRetiriveCardInfoQuery,
} from "@/lib/service/CardService";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";

const CreditCardsPage = () => {
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
        setSelectedCardIds(allCardsData.map((card: card) => card.id));
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
  const leng = allCardsData.length
  return (
    <div className="grid grid-cols-1 gap-2 pb-5">
      <div className="grid lg:grid-cols-1 p-4">
        <CardForCreditCards
          title="My Cards"
          className="overflow-x-auto"
          link="/credit-cards"
          button=""
        >
          <CardDisplay numofcard={leng}/>
        </CardForCreditCards>
      </div>
      <div className="grid lg:grid-cols-[4fr_6fr] max-md:grid-cols-1  gap-7 p-4">
        <Card
          title="Card Expense Statistics"
          className=""
        >
          <CardExpenseStatisticsChart />
        </Card>
        <Card
          title="Card List"
          className=""
        >
          <CardList />
        </Card>
      </div>
      <div className="grid lg:grid-cols-[6fr_4fr] max-md:grid-cols-1  gap-7 p-4">
        <Card
          title="Add New Card"
          className=""
        >
          <AddNewCard />
        </Card>
        <Card
          title="Card Setting"
          className=""
        >
          <CardSetting />
        </Card>
      </div>
    </div>
  );
};

export default CreditCardsPage;
