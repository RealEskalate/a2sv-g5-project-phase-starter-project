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

const Center = () => {
  const { data: session } = useSession();
  const accessToken =
    "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJuYXR0eSIsImlhdCI6MTcyMzgzMDIxNiwiZXhwIjoxNzIzOTE2NjE2fQ.c5zYX74xJyowvSM8pmN4W8Aw6pMyiJjs9JOP__Cjy9J80EHlOS6gX2yJpcwSdBwF";

  // Update initial card data using the custom hook
  useCardDispatch(accessToken);

  const CardData: Card[] = useAppSelector((state) => state.cards.cards);
  console.log("Fetched cards:", CardData);
  return (
    <>
      <section className="w-full flex gap-6">
        <div className="cards-container w-full cente-Content flex flex-col gap-6">
          <div className="card-box flex">
            <h1 className="flex grow page text-xl font-semibold text-colorBody-1">
              My Cards
            </h1>
            <Link href={""} className="text-base font-medium hover:underline">
              SeeAll
            </Link>
          </div>

          <div className="flex gap-6 grow">
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

      <section className="flex gap-6 grow">
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
        <div className="cards-container w-2/5 center-content flex flex-col gap-6">
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
