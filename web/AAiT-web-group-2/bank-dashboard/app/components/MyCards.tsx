"use client";

import React from "react";
import ChipCard from "../../public/ChipCard";

interface MyCardsType {
  id: string;
  cardHolder: string;
  semiCardNumber: string;
  cardType: string;
  balance: number;
  expiryDate: string;
}

const data: MyCardsType[] = [
  {
    id: "string",
    cardHolder: "string",
    semiCardNumber: "string",
    cardType: "string",
    balance: 0,
    expiryDate: "2024-08-23T13:47:22.222Z",
  },
];

const MyCards = () => {
  return (

    <div
      className=" flex flex-col justify-between items-center
      text-white bg-my-card-bg-1 mx-auto w-my-card-width max-md:w-[260px] max-md:h-[170px] max-sm:w-[325px] h-[225px] 
      rounded-my-card-radius font-lat shadow-lg "
    >
      <div className="flex justify-between px-6 max-md:px-2 pt-6 max-md:pt-3 items-start w-full">
        <div className="flex flex-col items-start max-md:h-[33px]">
          <span className="font-extralight text-sm max-md:text-[11px] max-md:w-[51px]">Balance</span>
          <span className="text-lg w-[51px]">$5,756</span>

        </div>
        <ChipCard color="white" />
      </div>
      <div className="flex justify-between items-center px-6 w-full">
        <div className="flex flex-col items-start w-1/2">
          <span className="font-extralight text-[14px] sm:text-[16px]">
            CARD HOLDER
          </span>
          <span className="text-[16px] sm:text-[18px]">Eddy Cusuma</span>
        </div>
        <div className="flex flex-col items-start w-1/2">
          <span className="font-extralight text-[14px] sm:text-[16px]">
            VALID THRU
          </span>
          <span className="text-[16px] sm:text-[18px]">12/22</span>
        </div>
      </div>
      <div className="flex justify-between items-center bg-my-card-bg-2 h-[70px] w-full px-6 rounded-b-lg">
        <div className="text-[18px] sm:text-[26px]">3778 **** **** 1234</div>
        <div className="flex justify-center items-center -space-x-4">
          <div className="w-8 h-8 sm:w-6 sm:h-6 bg-white bg-opacity-50 rounded-full"></div>
          <div className="w-8 h-8 sm:w-6 sm:h-6 bg-white bg-opacity-50 rounded-full"></div>
        </div>
      </div>
    </div>
  );
};

export default MyCards;
