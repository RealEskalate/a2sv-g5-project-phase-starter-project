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

const CreditCards = () => {
  const { data: session } = useSession();
  const accessToken =
    "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJuYXR0eSIsImlhdCI6MTcyMzgzMDIxNiwiZXhwIjoxNzIzOTE2NjE2fQ.c5zYX74xJyowvSM8pmN4W8Aw6pMyiJjs9JOP__Cjy9J80EHlOS6gX2yJpcwSdBwF";

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
      <div className="cards-container w-full cente-Content flex flex-col gap-6">
        <p className="flex grow page text-xl font-semibold text-colorBody-1">
          My Cards
        </p>
        <div className="flex gap-6">
          <>
            {CardData?.slice(0, 3).map((item, index) => (
              <VisaCard
                key={index}
                data={item}
                isBlack={isBlack[index] || false}
                isFade={isFade[index] || false}
                isSimGray={isSimGray[index] || false}
              />
            ))}
          </>
        </div>
      </div>
      <div className="flex w-full gap-6 text-nowrap">
        <div className="expense flex w-[33%] flex-col gap-5 ">
          <h2 className="text-xl font-semibold text-colorBody-1">
            Card Expense Statistics
          </h2>
          <ExpenseChart />
        </div>
        <div className="cardlist w-[67%] flex flex-col gap-6">
          <h2 className="text-xl font-semibold text-colorBody-1">Card List</h2>
          <>
            {CardData.slice(0, 3).map((card, index) => (
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
            ))}
          </>

          {/* <CardList
                 key={index}
                 img="/assets/money.svg"
                 title={data.cardType}
                 desc="Secondary"
                 colOne="Bank"
                 descOne="DBL Bank"
                 colTwo="Card Number"
                 descTwo="**** **** 5600"
                 colThree="Namain Card"
                 descThree="William"
                 btn="View Details"
                 color="bg-blue-100"
               />; */}

          {/* <CardList
            img="/assets/money.svg"
            title="Card Type"
            desc="Secondary"
            colOne="Bank"
            descOne="DBL Bank"
            colTwo="Card Number"
            descTwo="**** **** 5600"
            colThree="Namain Card"
            descThree="William"
            btn="View Details"
            color="bg-blue-100"
          />
          <CardList
            img="/assets/moneyPink.svg"
            title="Card Type"
            desc="Secondary"
            colOne="Bank"
            descOne="BRC Bank"
            colTwo="Card Number"
            descTwo="**** **** 4300"
            colThree="Namain Card"
            descThree="Michel"
            btn="View Details"
            color="bg-pink-100"
          />
          <CardList
            img="/assets/moneyOrange.svg"
            title="Card Type"
            desc="Secondary"
            colOne="Bank"
            descOne="ABM Bank"
            colTwo="Card Number"
            descTwo="**** **** 7560"
            colThree="Namain Card"
            descThree="Edward"
            btn="View Details"
            color="bg-orange-100"
          /> */}
        </div>
      </div>
      <div className="flex w-full gap-10">
        <div className="w-[67%] flex flex-col gap-6">
          <h2 className="text-xl font-semibold text-colorBody-1">
            Add New Card
          </h2>
          <AddCard />
        </div>
        <div className="w-[33%] flex flex-col gap-6">
          <h2 className="text-xl font-semibold text-colorBody-1">
            Card Setting
          </h2>
          <div className="flex flex-col gap-2 rounded-3xl p-4 bg-white">
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
