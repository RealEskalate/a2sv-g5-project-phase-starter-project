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

const Recent_transaction = () => {
  const recents = [
    { typeOfTransfer: "Deposit from my Card", date: "28 January 2021", picture: q1, amount: "-$850" },
    { typeOfTransfer: "Deposit paypal", date: "25 January 2021", picture: q2, amount: "+$2,550" },
    { typeOfTransfer: "Jemi wilson", date: "21 January 2021", picture: q3, amount: "+$5,400" }
  ];
  const [currIndex, setCurrIndex] = useState(0);

  const threeRecents = recents.slice(0, 3);

  return (
    <div className="flex md:w-3/12  flex-col">
      <div>
        <h1 className="text-[#343C6A] font-semibold ml-2">
          Recent Transaction
        </h1>
      </div>
      <div className="mt-2 md:mt-3 py-5 px-6 flex flex-col bg-white rounded-3xl space-y-8">
        {threeRecents.map((item, index) => (
          <div key={index} className="flex space-x-5 items-center space-y-1">
            <Image
              className="w-10 h-10"
              src={item.picture}
              alt="profile picture"
            />
            <div className="flex text-xs space-y-1 flex-col">
              <h1 className="font-semibold">{item.typeOfTransfer}</h1>
              <h2>{item.date}</h2>
            </div>
            <div className="flex justify-end">
              <h3 className="text-xs">{item.amount}</h3>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Recent_transaction;
