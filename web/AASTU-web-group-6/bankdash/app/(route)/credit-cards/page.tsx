"use client";
import React from "react";
import CardList from "@/app/components/Card/CardList";
import AddCard from "@/app/components/Card/AddCard";
import ExpenseChart from "@/app/components/Charts/ExpenseChart";
import SettingsCard from "@/app/components/Card/SettingsCard";
import VisaCard from "@/app/components/Card/VisaCard";
import useCardDispatch from "@/app/Redux/Dispacher/useCardDispatch";
import { useAppSelector } from "@/app/Redux/store/store";
import { useSession } from "next-auth/react";
import { Card } from "@/app/Redux/slices/cardSlice";

const CreditCards = () => {
  const { data: session } = useSession();
  const accessToken =
    "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJuYXR0eSIsImlhdCI6MTcyMzgzMDIxNiwiZXhwIjoxNzIzOTE2NjE2fQ.c5zYX74xJyowvSM8pmN4W8Aw6pMyiJjs9JOP__Cjy9J80EHlOS6gX2yJpcwSdBwF";

  const CardData: Card[] = useAppSelector((state) => state.cards.cards);

  return (
    <>
      <div>
        <p className="text-[#333B69] font-semibold text-[22px] pb-5">
          My Cards
        </p>
        <div className="flex gap-10 pb-5">
          <VisaCard
            data={CardData[0]}
            isBlack={false}
            isFade={true}
            isSimGray={false}
          />
          <VisaCard
            data={CardData[1]}
            isBlack={false}
            isFade={false}
            isSimGray={false}
          />
          <VisaCard
            data={CardData[2]}
            isBlack={true}
            isFade={false}
            isSimGray={false}
          />
        </div>
      </div>
      <div className="flex flex-row">
        <div>
          <p className="text-[#333B69] font-semibold text-[22px] pb-5">
            Card Expense Statistics
          </p>
          <ExpenseChart />
        </div>
        <div>
          <p className="text-[#333B69] font-semibold text-[22px] ml-10 pb-5">
            Card List
          </p>
          <>
            {CardData.slice(0, 3).map((card) => (
              <div key={card.id}>
                <CardList
                  img="/assets/money.svg"
                  title={card.cardType}
                  desc="Secondary"
                  colOne="Bank"
                  descOne={card.cardNumber}
                  colTwo="Card Number"
                  descTwo={`**** **** ${card.semiCardNumber}`}
                  colThree="Name on Card"
                  descThree={card.cardHolder}
                  btn="View Details"
                  color="bg-blue-100"
                />
              </div>
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
      <div className="flex w-[95%] gap-10">
        <div>
          <p className="text-[#333B69] font-semibold text-[22px] pb-5">
            Add New Card
          </p>
          <AddCard />
        </div>
        <div>
          <p className="text-[#333B69] font-semibold text-[22px] pb-5">
            Card Setting
          </p>
          <div className="border rounded-3xl p-4 bg-white">
            <SettingsCard
              img="assets/block.svg"
              title="Block Card"
              desc="Instantly block your card"
            />
            <SettingsCard
              img="assets/lock.svg"
              title="Change Pin Code"
              desc="Choose another pin code"
            />
            <SettingsCard
              img="assets/google.svg"
              title="Add to Google Pay"
              desc="Withdraw without any card"
            />
            <SettingsCard
              img="assets/apple.svg"
              title="Add to Apple Pay"
              desc="Withdraw without any card"
            />
            <SettingsCard
              img="assets/apple.svg"
              title="Add to Apple Store"
              desc="Withdraw without any card"
            />
          </div>
        </div>
      </div>
    </>
  );
};

export default CreditCards;
