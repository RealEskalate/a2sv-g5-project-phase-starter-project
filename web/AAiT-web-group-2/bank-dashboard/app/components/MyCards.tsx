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
      text-white bg-my-card-bg-1 mx-auto w-my-card-width max-md:w-[260px] max-md:h-[170px] max-sm:w-[325px] h-[225px] 
      rounded-my-card-radius font-lat shadow-lg "
    >
      <div className="flex justify-between px-6 max-md:px-2 pt-6 max-md:pt-3 items-start w-full">
        <div className="flex flex-col items-start max-md:h-[33px]">
          <span className="font-extralight text-sm max-md:text-[11px] max-md:w-[51px]">Balance</span>
          <span className="text-lg w-[51px]">$5,756</span>
        </div>
        <ChipCard color="white"/>
      </div>
      <div className="flex justify-between items-center px-6 max-md:px-1 max-md:h-[30px]  w-full">
        <div className="flex flex-col items-start w-1/2">
          <span className="font-extralight text-sm">CARD HOLDER</span>
          <span className="max-md:text-[14px]">Eddy Cusuma</span>
        </div>
        <div className="flex flex-col items-start w-1/2">
          <span className="font-extralight text-sm">VALID THRU</span>
          <span>12/22</span>
        </div>
      </div>
      <div className="flex justify-between items-center bg-my-card-bg-2 h-my-card-height-2 max-md:h-[35px] w-full px-6 rounded-b-my-card-radius">
        <div className="text-xl max-md:text-[15px] max-md:w-[191px]">3778 **** **** 1234</div>
        <div className="flex justify-center items-center -space-x-4">
          <div className="w-8 max-md:w-6 h-8 max-md:h-6  bg-white bg-opacity-50 rounded-full"></div>
          <div className="w-8 max-md:w-6 h-8 max-md:h-6 bg-white bg-opacity-50 rounded-full"></div>
        </div>
      </div>
    </div>
  );
};

export default MyCards;
