"use client";
import React from "react";
import CardList from "@/app/components/Card/CardList";
import AddCard from "@/app/components/Card/AddCard";
import ExpenseChart from "@/app/components/Charts/ExpenseChart";
import SettingsCard from "@/app/components/Card/SettingsCard";
import VisaCard from "@/app/components/Card/VisaCard";
import { useAppSelector } from "@/app/Redux/store/store";
import { useSession } from "next-auth/react";
import { Card } from "@/app/Redux/slices/cardSlice";
import { ShimmerVisaCard } from "@/app/components/Shimmer/ShimmerVisa";
import ShimmerCard from "@/app/components/Shimmer/SimmerCard";

const CreditCards = () => {
  const { data: session } = useSession();
  const accessToken = session?.accessToken as string;
  const CardData: Card[] = useAppSelector((state) => state.cards.cards);
  const imgCont = [
    "assets/block-card-blue-icon.svg",
    "assets/block-card-pink-icon.svg",
    "assets/block-card-orange-icon.svg",
  ];
  const bgCont = ["bg-[#E7EDFF]", "bg-[#FFE0EB]", "bg-[#FFF5D9]"];
  const isBlack = [false, false, true];
  const isFade = [true, false, false];
  const isSimGray = [false, false, true];

  return (
    <div className="w-[96%] flex flex-col grow gap-6 p-8 pt-6">
      <div className="cards-container w-full flex flex-col gap-6">
        <p className="flex grow page text-xl font-semibold text-colorBody-1 dark:text-gray-300">
          My Cards
        </p>
        <div className="flex gap-6 overflow-x-auto sm:scroll-snap-x overflow-y-hidden scrollbar-hide">
          {CardData.length > 0 ? (
            CardData?.slice(0, 3).map((item, index) => (
              <VisaCard
                key={index}
                data={item}
                isBlack={isBlack[index] || false}
                isFade={isFade[index] || false}
                isSimGray={isSimGray[index] || false}
                className="flex-shrink-0 sm:w-auto"
              />
            ))
          ) : (
            <div className="w-full flex gap-6 ">
              <ShimmerVisaCard />
              <ShimmerVisaCard />
              <ShimmerVisaCard />
            </div>
          )}
        </div>
      </div>
      <div className="flex flex-col lg:flex-row w-full gap-6 text-nowrap">
        <div className="expense flex w-full lg:w-[33%] flex-col gap-5 ">
          <h2 className="text-xl font-semibold text-colorBody-1 dark:text-gray-300">
            Card Expense Statistics
          </h2>
          <ExpenseChart />
        </div>
        <div className="cardlist w-full lg:w-[60%] flex flex-col gap-6">
          <h2 className="text-xl font-semibold text-colorBody-1 dark:text-gray-300">
            Card List
          </h2>
          <div className="flex flex-col gap-6">
            {CardData.length > 0 || undefined ? (
              CardData.slice(0, 3).map((card, index) => (
                <CardList
                  key={card.id}
                  img={imgCont[index]}
                  title={card.cardType}
                  desc="Secondary"
                  colOne="Bank"
                  descOne={card.cardNumber}
                  colTwo="Card Number"
                  descTwo={`**** **** ${card.semiCardNumber}`}
                  colThree="Name on Card"
                  descThree={card.cardHolder}
                  btn="View Details"
                  color={bgCont[index]}
                />
              ))
            ) : (
              <div className="w-full flex flex-col gap-6 ">
                <ShimmerCard />
                <ShimmerCard />
                <ShimmerCard />
              </div>
            )}
          </div>
        </div>
      </div>
      <div className="flex flex-col lg:flex-row w-full gap-10">
        <div className="w-full lg:w-[67%] flex flex-col gap-6">
          <h2 className="text-xl font-semibold text-colorBody-1 dark:text-gray-300">
            Add New Card
          </h2>
          <AddCard />
        </div>
        <div className="w-full lg:w-[40%] flex flex-col gap-6">
          <h2 className="text-xl font-semibold text-colorBody-1 dark:text-gray-300">
            Card Setting
          </h2>
          <div className="flex flex-col gap-2 rounded-3xl p-4 bg-white dark:bg-[#232328]">
            <SettingsCard
              img="assets/block-card-orange-icon.svg"
              title="Block Card"
              desc="Instantly block your card"
              bg="bg-[#FFF5D9]"
            />
            <SettingsCard
              img="assets/lock-icon.svg"
              title="Change Pin Code"
              desc="Choose another pin code"
              bg="bg-[#E7EDFF]"
            />
            <SettingsCard
              img="assets/google-icon.svg"
              title="Add to Google Pay"
              desc="Withdraw without any card"
              bg="bg-[#FFE0EB]"
            />
            <SettingsCard
              img="assets/apple-icon.svg"
              title="Add to Apple Pay"
              desc="Withdraw without any card"
              bg="bg-[#DCFAF8]"
            />
            <SettingsCard
              img="assets/apple-icon.svg"
              title="Add to Apple Store"
              desc="Withdraw without any card"
              bg="bg-[#DCFAF8]"
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default CreditCards;
