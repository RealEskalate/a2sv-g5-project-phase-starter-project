"use client";

import React from "react";
import ChipCard from "../../public/ChipCard";

interface MyCardsType {
  balance: number;
  name: string;
  valid_through: string;
  account: string;
  bg_color_1: string;
  bg_color_2: string;
  chip_color: string;
  font_color: string;
}

const dummy_data: MyCardsType[] = [
  {
    balance: 5_756,
    name: "Eddy Cusuma",
    valid_through: "12/22",
    account: "3778 **** **** 1234",
    bg_color_1: "my-card-bg-1",
    bg_color_2: "my-card-bg-2",
    chip_color: "white",
    font_color: "white",
  },
  {
    balance: 5_756,
    name: "Eddy Cusuma",
    valid_through: "12/22",
    account: "3778 **** **** 1234",
    bg_color_1: "white",
    bg_color_2: "white",
    chip_color: "black",
    font_color: "custom-font-color",
  },
];

const MyCards = () => {
  return (
    <div
      className=" flex flex-col justify-between items-center
      text-white bg-my-card-bg-1 min-w-[235px] min-h-[170px] max-w-[350px] max-h-[235px] rounded-[25px] shadow-lg"
    >
      <div className="flex justify-between px-6 pt-6 items-start w-full">
        <div className="flex flex-col items-start">
          <span className="font-extralight text-sm text-[clamp(14px,1vw,13px)] ">Balance</span>
          <span className="text-[clamp(19px,1vw,24px)]">$5,756</span>
        </div>
        <ChipCard color="white"/>
      </div>
      <div className="flex justify-between items-center px-6 w-full">
        <div className="flex flex-col items-start w-1/2">
          <span className="font-extralight text-[clamp(14px,1vw,13px)]">CARD HOLDER</span>
          <span className="text-[clamp(16px,1vw,18px)]">Eddy Cusuma</span>
        </div>
        <div className="flex flex-col items-start w-1/2">
          <span className="font-extralight text-[clamp(14px,1vw,13px)]">VALID THRU</span>
          <span className="text-[clamp(16px,1vw,18px)]">12/22</span>
        </div>
      </div>
      <div className="flex justify-between items-center bg-my-card-bg-2 min-h-[70px] max-h-[70px] w-full px-6 rounded-b-my-card-radius">
        <div className="text-[clamp(18px,1vw,26px)]">3778 **** **** 1234</div>
        <div className="flex justify-center items-center -space-x-4">
          <div className="w-8 max-md:w-6 h-8 max-md:h-6  bg-white bg-opacity-50 rounded-full"></div>
          <div className="w-8 max-md:w-6 h-8 max-md:h-6 bg-white bg-opacity-50 rounded-full"></div>
        </div>
      </div>
    </div>
  );
};

export default MyCards;
