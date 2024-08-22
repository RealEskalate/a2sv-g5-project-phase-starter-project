"use client";
import React, { useEffect, useState } from "react";
import PieChart from "./PieChart";
import CreditCard from "./CreditCard";
import CardSettingList from "./CardSettingList";
import AddCardForm from "./AddCardForm";
import MainCreditCard from "./MainCreditCard";
import Card from "../components/Page2/Card";
import { Card as Card1 } from "../../types/cardController.Interface";
import { getCards } from "@/lib/api/cardController";
import { getSession } from "next-auth/react";
import router from "next/navigation";
import { useRouter } from "next/navigation";
import Refresh from "../api/auth/[...nextauth]/token/RefreshToken";

const HeadingTitle = ({ title }: { title: string }) => {
  return (
    <h1 className="text-[#343C6A] font-semibold lg:text-xl md:text-lg">
      {title}
    </h1>
  );
};

const CreditCards = () => {
  const [accessToken, setAccessToken] = useState<string>("");
  const [cards, setCards] = useState<Card1[]>([]);
  const [loading, setLoading] = useState(true);
  const router = useRouter();

  const convertToDate = (date: string) => {
    const year = date.slice(2, 4);
    const month = date.slice(5, 7);

    return month + "/" + year;
  };
  useEffect(() => {
    const fetchSession = async () => {
      const access_token = await Refresh();
      if (access_token) {
        setAccessToken(access_token);
      } else {
        router.push(
          `./api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`
        );
      }
    };

    fetchSession();
  }, [router]);

  useEffect(() => {
    if (accessToken == "") {
      return;
    }
    async function fetch() {
      const data = await getCards(accessToken, 0, 700);
      data.content.reverse();
      setCards(data.content);
      setLoading(false);
    }
    fetch();
  }, [accessToken]);

  const decideColor = (index: number) => {
    const remainder = index % 3;
    if (remainder == 0) {
      return ["from-[#0A06F4] to-[#0A06F4]", "text-white"];
    } else if (remainder == 1) {
      return ["from-[#4C49ED] to-[#4C49ED]", "text-white"];
    } else {
      return ["from-[#FFF] to-[#FFF]", "text-black"];
    }
  };
  const handleAddition = (card: Card1) => {
    const newCards = [card, ...cards];
    setCards(newCards);
  };
  if (loading) return null;

  return (
    <div className="bg-[#f5f7fb] w-full p-5 gap-5 flex flex-col">
      <div className="flex-col gap-5">
        <HeadingTitle title="My Cards" />

        <div className="flex overflow-scroll justify-between [&::-webkit-scrollbar]:hidden">
          {cards.map((card, index) => {
            const [bgColor, textColor] = decideColor(index);

            if (index <= 2) {
              return (
                <Card
                  balance={card.balance.toString()}
                  cardHolder={card.cardHolder}
                  validThru={convertToDate(card.expiryDate)}
                  cardNumber={card.id}
                  filterClass=""
                  bgColor={bgColor}
                  textColor={textColor}
                  iconBgColor="bg-opacity-10"
                  showIcon={true}
                  key={index}
                />
              );
            }
          })}
        </div>
      </div>
      <div className="flex flex-col gap-6 md:flex-row">
        <div className="flex flex-col gap-5 basis-5/12 ">
          <HeadingTitle title="Card Expense Statistics" />
          <PieChart />
        </div>
        <div className="flex flex-col gap-3 md:justify-between w-full h-full">
          <HeadingTitle title="Card List" />
          <div className="overflow-y-auto h-80 flex flex-col gap-4 [&::-webkit-scrollbar]:hidden">
            {cards.map((card, index) => {
              return (
                <CreditCard
                  icon={<img src="card1.svg" />}
                  linkUrl=""
                  data={[
                    ["Card Type", card.cardType],
                    ["Balance", card.balance.toString()],
                    ["Card Number", card.cardType],
                    ["Card Expiry", convertToDate(card.expiryDate)],
                  ]}
                  key={index}
                />
              );
            })}
          </div>
        </div>
      </div>

      <div className="flex flex-col gap-6 md:flex-row">
        <div className="flex flex-col gap-5">
          <HeadingTitle title="Add New Card" />
          <AddCardForm
            access_token={accessToken}
            handleAddition={handleAddition}
          />
        </div>
        <div className="flex flex-col gap-5 min-w-64 h-full">
          <HeadingTitle title="Card Setting" />
          <CardSettingList />
        </div>
      </div>
    </div>
  );
};

export default CreditCards;
