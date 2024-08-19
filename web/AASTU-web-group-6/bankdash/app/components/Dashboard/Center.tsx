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
import { TransactionType } from "@/app/Redux/slices/TransactionSlice";

const Center = () => {
  const { data: session } = useSession();
  const accessToken =
    "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJuYXR0eSIsImlhdCI6MTcyMzkxNjcyMiwiZXhwIjoxNzI0MDAzMTIyfQ.p9aXcFbPCTl88kFVBDa5JuDj03jg1KEjfo82GCAPY2kQo8c9x4gwNNYai4isHVZM";

  // Update initial card and tran data using the custom hook
  useCardDispatch(accessToken);
  useTranDispatch(accessToken);

  const CardData: Card[] = useAppSelector((state) => state.cards.cards);
  const TranData: TransactionType[] = useAppSelector(
    (state) => state.transactions.transactions
  );
  const balanceHist: TransactionType[] = useAppSelector(
    (state) => state.transactions.balanceHist
  );

  return (
    <>
      <section className="flex gap-6 sm:flex-col lg:flex-row">
        <div className="cards-container sm:w-full lg:w-[67%] cente-Content flex flex-col gap-6">
          <div className="card-box flex w-full items-center justify-between">
            <h1 className="text-xl font-semibold text-colorBody-1">My Cards</h1>
            <Link href={""} className="text-base font-medium hover:underline">
              SeeAll
            </Link>
          </div>

          <div className="flex gap-6 grow w-full">
            <VisaCard
              data={CardData[0]}
              isBlack={false}
              isFade={false}
              isSimGray={false}
            />
            <VisaCard
              data={CardData[1]}
              isBlack={true}
              isFade={false}
              isSimGray={false}
            />
          </div>
        </div>
        <RecentTr />
      </section>

      <section className="flex gap-6 grow sm:flex-col lg:flex-row">
        <div className="Weekly-container w-full cente-Content flex flex-col gap-6 ">
          <h1 className="flex grow page text-xl font-semibold text-colorBody-1">
            Weekly Activity
          </h1>
          <div className="flex flex-col min-w-[300px] w-full p-8 py-4 bg-white rounded-3xl border-solid border-[0.1px] shadow-sm">
            <div className="indiColor w-full flex gap-4 justify-end">
              <div className="flex gap-2 p-4 items-center">
                <div className="circle w-4 h-4 bg-[#1814F3]  rounded-full"></div>
                Diposit
              </div>
              <div className="flex gap-2 p-4 items-center">
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
        <div className="cards-container sm:w-full lg:w-[40%] items-center flex flex-col gap-6">
          <h1 className="text-xl font-semibold text-colorBody-1">
            Expense Statistics
          </h1>
          <PieComp />
        </div>
      </section>
    </>
  );
};

export default Center;
