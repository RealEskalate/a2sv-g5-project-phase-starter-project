"use client";
import React, { useEffect } from "react";
import CreditCard from "../CreditCard";
import DashboardBarChart from "../Chart/DashboardBarChart";
import { useDispatch, useSelector } from "react-redux";
import Link from "next/link";

import { useGetCardsQuery } from "@/lib/redux/api/cardsApi";
// import { setCards, setLoading, setError } from '@/lib/redux/slices/cardsSlice'
import { RootState } from "@/lib/redux/store";
import { setCards, setLoading, setError } from "@/lib/redux/slices/cardsSlice";
import { cardStyles } from "@/lib/CardColor";

const CreditCardTransaction = () => {
  const dispatch = useDispatch();
  const { cards, loading, error } = useSelector(
    (state: RootState) => state.cards
  );

  const {
    data: cardsData,
    isLoading: cardsLoading,
    isError: errorCard,
  } = useGetCardsQuery({ size: 5, page: 0 });

  useEffect(() => {
    dispatch(setLoading(cardsLoading));
    if (cardsData) {
      dispatch(setCards(cardsData.content));
    }
    if (errorCard) {
      dispatch(setError("Error on fetching data"));
    }
  }, [cardsData, errorCard, cardsLoading, dispatch]);

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;
  return (
    <div className="lg:w-[65%] xl:w-[68%] rounded-xl bg-[#F5F7FA]">
      <div className="credit-card-info flex justify-between  h-16 items-center ">
        <h1 className="font-semibold text-[#343C6A]">My cards</h1>
        <Link href="/creditcardpage#add-new-card">
          <h1 className="text-[#2D60FF]">+ Add Card</h1>
        </Link>
        {/* <h1 className="font-semibold text-[#343C6A]">+Add Card</h1> */}
      </div>
      <div className="cards flex gap-5 lg:gap-1  lg:flex-row overflow-x-auto overflow-y-hidden no-scrollbar h-56 lg:justify-between xl:gap-10">
      {/* <div className="creditcards flex  gap-5 lg:flex-row overflow-x-auto overflow-y-hidden no-scrollbar  h-56  lg:justify-start lg:px-4 "> */}
            {cardsData?.content.map((card, index: number) => {
              const style =
                cardStyles[card.cardType as keyof typeof cardStyles] ||
                cardStyles.Primary;

              return (
                <div
                  key={index}
                  className="credit-card min-h-80 w-[360px] max-w-72 md:max-w-96 flex-shrink-0"
                >
                  <CreditCard
                    name={card.cardHolder}
                    balance={String(card.balance)}
                    cardNumber={card.semiCardNumber}
                    validDate={card.expiryDate}
                    backgroundImg={style.backgroundImg}
                    textColor={style.textColor}
                  />
                </div>
              );
            })}
          {/* </div> */}
      </div>
    </div>
  );
};

export default CreditCardTransaction;
