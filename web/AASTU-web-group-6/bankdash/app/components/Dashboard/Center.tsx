"use client";
import React from "react";
import Link from "next/link";
import VisaCard from "../Card/VisaCard";
import BarComp from "../Charts/BarComp";
import { PieComp } from "../Charts/PieComp";
import RecentTr from "./RecentTr";
import { useSession } from "next-auth/react";
import useCardDispatch from "@/app/Redux/Dispacher/useCardDispatch";
import { useAppSelector } from "@/app/Redux/store/store";
import { Card } from "@/app/Redux/slices/cardSlice";
import useTranDispatch from "@/app/Redux/Dispacher/useTranDispatch";
import {
  BalanceType,
  TransactionType,
} from "@/app/Redux/slices/TransactionSlice";
import { ShimmerVisaCard } from "../Shimmer/ShimmerVisa";
import loading from "@/app/loading";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCreditCard } from "@fortawesome/free-solid-svg-icons";

const Center = () => {
  const { data: session } = useSession();
  const accessToken = session?.accessToken as string;
  
  console.log(session?.accessToken, "token");
  // Update initial card and tran data using the custom hook
  useCardDispatch(accessToken);
  useTranDispatch(accessToken);
  const income = useAppSelector((state) => state.transactions.income);
  console.log(income, "hell");

  const CardData: Card[] = useAppSelector((state) => state.cards.cards);
  const TranData: TransactionType[] = useAppSelector(
    (state) => state.transactions.transactions
  );
  // console.log(TranData , "yyyyyyyyyyy")
  const cardColor = [false, true];

  return (
    <>
      <section className="flex gap-6 overflow-x-hidden xxs:flex-col lg:flex-row">
        <div className="cards-container xxs:w-full lg:w-[67%] cente-Content flex flex-col xxs:gap-4 md:gap-6 ">
          <div className="card-box flex w-full items-center justify-between">
            <h1 className="text-xl font-semibold text-colorBody-1 dark:text-gray-300">
              My Cards
            </h1>
            <Link
              href={""}
              className="text-base font-medium hover:underline dark:text-gray-300"
            >
              SeeAll
            </Link>
          </div>

          <div className="flex gap-6 py-2 grow scrollbar-hide xxs:justify-start xxs:overflow-x-auto sm:overflow-hidden sm:w-full sm:overflow-x-hidden">
            <>
              {CardData.length > 0 ? (
                CardData?.slice(0, 2).map((item, index) => (
                  <VisaCard
                    key={index}
                    data={item}
                    isBlack={cardColor[index] || false}
                    isFade={false}
                    isSimGray={false}
                  />
                ))
              ) : (
                <div className="w-full flex gap-6 ">
                  <ShimmerVisaCard />
                  <ShimmerVisaCard />
                </div>
              )}
            </>
          </div>
        </div>
        <RecentTr />
      </section>

      <section className="flex gap-6 grow xxs:flex-col  lg:flex-row">
        <div className="Weekly-container w-full cente-Content flex flex-col gap-6 ">
          <h1 className="flex grow page text-xl font-semibold text-colorBody-1 dark:text-gray-300">
            Weekly Activity
          </h1>
          <div className="flex flex-col min-w-[300px] w-full p-8 py-4 bg-white dark:bg-[#232328] rounded-3xl border-solid border-[0.1px] shadow-sm">
            <div className="indiColor w-full flex gap-4 justify-end">
              <div className="flex gap-2 p-4 items-center dark:text-gray-200">
                <div className="circle w-4 h-4 bg-[#1814F3]  rounded-full"></div>
                Deposit
              </div>
              <div className="flex gap-2 p-4 items-center dark:text-gray-200">
                <div className="circle w-4 h-4 bg-[#16DBCC] rounded-full"></div>
                Withdraw
              </div>
            </div>
            <div className="flex gap-2 ">
              <div className="leftCanva max-h-44 flex flex-col items-end justify-between text-sm text-[#718EBF]">
                <span>400</span>
                <span>300</span>
                <span>200</span>
                <span>100</span>
                <span>0</span>
              </div>
              <BarComp />
            </div>
          </div>
        </div>
        <div className="cards-container sm:w-full lg:w-[40%] flex flex-col gap-6">
          <h1 className="text-xl font-semibold text-colorBody-1 dark:text-gray-300">
            Expense Statistics
          </h1>
          <PieComp />
        </div>
      </section>
    </>
  );
};

export default Center;
