"use client";
import React, { useEffect, useState } from "react";
import Infobox from "../../components/infobox/Infobox";
import LastTransaction from "../../components/lastTransaction/LastTransaction";
import Card from "../../components/card/Card";
import CreditCard from "../../components/creditCard/CreditCard";
import CardForCreditCards from "../../components/card/CardForCreditCards";
import DebitCreditOverviewChart from "../../components/charts/DebitCreditOverviewChart";
import InvoicesSent from "../../components/InvoiceSend/InvoiceSent";
import {
  CardData,
  formatCardNumber,
  card,
} from "@/app/components/cardDisplay/CardDisplay";
import {
  useGetAllCardInfoQuery,
  useRetiriveCardInfoQuery,
} from "@/lib/service/CardService";
import { useSession } from "next-auth/react";
import CardSkeleton from "@/app/components/creditCard/CardSkeleton";

// import { useRouter } from "next/navigation";

const AccountsPage = () => {
  const { data: session, status } = useSession();
  // const router = useRouter();

  useEffect(() => {}, [status, session]);
  // console.log(session, status);
  

  const [selectedCardIds, setSelectedCardIds] = useState<string[]>([]);

  const token = session?.user.accessToken || "";
  // console.log("accesstoken: ", token);

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
        <div className="flex space-x-2 justify-center items-center bg-white h-screen"></div>
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
    <div className="grid grid-cols-1 gap-2 pb-5">
      <Infobox />
      <div className="flex max-sm:flex-col justify-between gap-8">
        <Card
          title="Last Transaction"
          className="flex flex-col md:h-[299px] h-[254] w-full"
        >
          <LastTransaction />
        </Card>
        <CardForCreditCards
          title="Credit Cards"
          button="See All"
          link="/credit-cards"
          className="flex flex-col lg:h-[300px] w-[350px]"
        >
          {allCardsData ? (
            <div>
              {allCardsData.slice(0, 1).map((card: CardData, index: number) => (
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
            <CardSkeleton />
          )}
        </CardForCreditCards>
      </div>
      <div className="grid lg:grid-cols-[3fr_2fr] max-md:grid-cols-1 gap-7 p-4 w-auto">
        <Card
          title="Debit & Credit Overview"
          className="w-full"
        >
          <DebitCreditOverviewChart />
        </Card>
        <Card
          title="Invoices Sent"
          className="w-full"
        >
          <InvoicesSent />
        </Card>
      </div>
    </div>
  );
};

export default AccountsPage;
