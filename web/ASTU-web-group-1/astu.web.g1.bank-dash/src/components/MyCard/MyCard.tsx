"use client";
import { CardResponseType } from "@/types/card.types";
import Image from "next/image";
import React from "react";

interface MyCardPropsType {
  props: CardResponseType;
  index: number;
}

// {
//   props: { id, cardHolder, semiCardNumber, cardType, expiryDate },
//   index,
// }: MyCardPropsType

// use the above comment in the props of the MyCard function

const MyCard = () => {

  const balance = "5,758";
  const cardHolder = "Eddy Cusuma";
  const expiryDate = "2022-12-12";
  const cardNumber = "37781234";
  const idx = 0;
  
  const cardStyles = [
    {
      bg: "bg-grad-end",
      text: "text-white",
      gradient: "bg-gradient-to-b from-grad-start to-grad-end",
      icon: "/assets/icons/chip-card-white.svg",
      border: "",
    },
    {
      bg: "bg-white",
      text: "navy",
      gradient: "",
      icon: "/assets/icons/chip-card-black.svg",
      border: "border-t",
    },
  ];


  const containerBG = cardStyles[idx].bg;
  const textColor = cardStyles[idx].text;
  const icon = cardStyles[idx].icon;
  const gradient = cardStyles[idx].gradient;
  const border = cardStyles[idx].border;

  function formatTime(timeString: string): string {
    const month = timeString.slice(5, 7);
    const year = timeString.slice(2, 4);
    return `${month}/${year}`;
  }
  function formatCardNumber(cardNumber: string): string {
    return `${cardNumber.slice(0, 4)} **** **** ${cardNumber.slice(4)}`;
  }
  return (
    <div>
      <div
        className={`${containerBG} ${textColor} w-[280px] h-[175px]   border rounded-3xl flex flex-col justify-between`}
      >
        <div className="flex flex-col  px-4  pt-4 h-full">
          <div className="flex justify-between ">
            <div className="flex flex-col">
              <span className="text-12px">Balance</span>
              <span className="text-14px font-semibold">$ {balance}</span>
            </div>
            <Image src={icon} alt="chip_card" width={35} height={35} />
          </div>
          <div className=""></div>
          <div className="flex h-full pb-3">
            <div className="flex flex-1 flex-col justify-center">
              <span className="text-12px">CARD HOLDER</span>
              <span className="text-14px font-semibold">{cardHolder}</span>
            </div>
            <div className="flex flex-1 flex-col justify-around items-center">
              <div className="flex flex-col">
                <span className="text-12px">VALID THRU</span>
                <span className="text-14px  font-semibold">
                  {formatTime(expiryDate)}
                </span>
              </div>
            </div>
          </div>
        </div>
        <div
          className={`${gradient} ${border} rounded-b-3xl flex justify-between px-6 py-3 `}
        >
          <div className="flex items-center">
            <span className="text-16px font-text-navy">
              {formatCardNumber(cardNumber)}
            </span>
          </div>
          <div className="flex pr-[15px]">
            <div className="flex bg-[#9199AF] h-8 w-8 rounded-full opacity-50"></div>
            <div className="flex bg-[#9199AF] h-8 w-8 rounded-full opacity-50 -mx-[15px]"></div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default MyCard;