import React from "react";
import Link from "next/link";
import VisaCard from "../Card/VisaCard";
import BarComp from "../Charts/BarComp";

const Center = () => {
  return (
    <section className="w-3/4 flex flex-col gap-6">
      <div className="cards-container w-full cente-Content flex flex-col gap-6">
        <div className="flex">
          <h1 className="flex grow page text-xl font-semibold text-colorBody-1">
            My Cards
          </h1>
          <Link href={""} className="text-base font-medium hover:underline">
            SeeAll
          </Link>
        </div>
        <div className="flex gap-6 grow">
          <VisaCard isBlack={false} />
          <VisaCard isBlack={true} />
        </div>
      </div>
      <div className="Weekly-container w-full cente-Content flex flex-col gap-6 ">
        <h1 className="flex grow page text-xl font-semibold text-colorBody-1">
          Weekly Activity
        </h1>
        <div className="flex w-full h-80 gap-6 bg-white rounded-3xl border-solid border-2 border-gray-200 ">
          <BarComp />
        </div>
      </div>
    </section>
  );
};

export default Center;
