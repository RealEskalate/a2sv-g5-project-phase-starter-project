"use client";
import Image from "next/image";
import React, { useState } from "react";
import q1 from "../../../public/images/Card.svg";
import q2 from "../../../public/images/paypal.svg";
import q3 from "../../../public/images/Jemi.svg";

interface Type {
  typeOfTransfer: string;
  date: string;
  picture: string;
  amount: string;
}

interface Props {
  recents: Type[];
}

const RecentTransaction: React.FC<Props> = () => {
  const recents = [
    { typeOfTransfer: "Deposit from my Card", date: "28 January 2021", picture: q1, amount: "-$850" },
    { typeOfTransfer: "Deposit paypal", date: "25 January 2021", picture: q2, amount: "+$2,550" },
    { typeOfTransfer: "Jemi wilson", date: "21 January 2021", picture: q3, amount: "+$5,400" }
  ];

  const threeRecents = recents.slice(0, 3);

  return (
    <div className="flex flex-col w-full">
      <div>
        <h1 className="text-[22px] font-semibold text-[#343C6A] ml-2">
          Recent Transaction
        </h1>
      </div>
      <div className="mt-2 md:mt-3 py-5 px-6 flex flex-col bg-white rounded-3xl space-y-8 w-full">
        {threeRecents.map((item, index) => {
          const sign = item.amount[0];
          const amountStyle = `text-base font-medium ${sign === "+" ? "text-[#41D4A8]" : "text-[#FF4B4A]"}`;

          return (
            <div key={index} className="flex space-x-5 justify-between space-y-1 w-full">
              <div className="flex space-x-5">
                <Image
                  className="w-12 h-12"
                  src={item.picture}
                  alt="profile picture"
                />
                <div className="flex text-xs space-y-1 flex-col">
                  <h1 className="font-medium text-base text-[#232323]">{item.typeOfTransfer}</h1>
                  <h2 className="text-sm text-[#718EBF]">{item.date}</h2>
                </div>
              </div>
              <div className="flex justify-end ">
                <h3 className={amountStyle}>{item.amount}</h3>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default RecentTransaction;
