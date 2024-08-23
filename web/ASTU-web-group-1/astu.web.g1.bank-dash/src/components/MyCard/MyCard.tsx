"use client";
import { CardContentType } from "@/types/card.types";
import Image from "next/image";
import React from "react";

const MyCard = ({
  content,
  index,
}: {
  content: CardContentType;
  index: number;
}) => {
  const balance = content.balance;
  const cardHolder = content.cardHolder;
  const expiryDate = content.expiryDate;
  const cardNumber = content.semiCardNumber;

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
    {
      bg: "bg-gray-800",
      text: "text-white",
      gradient: "bg-gradient-to-b from-gray-700 to-gray-900",
      icon: "/assets/icons/chip-card-white.svg",
      border: "",
    },
    {
      bg: "bg-indigo-500",
      text: "text-gray-100",
      gradient: "bg-gradient-to-b from-indigo-400 to-indigo-600",
      icon: "/assets/icons/chip-card-white.svg",
      border: "",
    },
  ];
  const idx = index % cardStyles.length;

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
    const formattedCardNumber = cardNumber.padStart(16, "0");
    const part1 = formattedCardNumber.slice(0, 4);
    const part2 = formattedCardNumber.slice(4, 8).replace(/\d/g, "*");
    const part3 = formattedCardNumber.slice(8, 12).replace(/\d/g, "*");
    const part4 = formattedCardNumber.slice(12, 16);
    return `${part1} ${part2} ${part3} ${part4}`;
  }

  function makeCardNumberVisible(cardNumber: string): string {
    const formattedCardNumber = cardNumber.padStart(16, "0");
    const part1 = formattedCardNumber.slice(0, 4);
    const part2 = formattedCardNumber.slice(4, 8);
    const part3 = formattedCardNumber.slice(8, 12);
    const part4 = formattedCardNumber.slice(12, 16);
    return `${part1} ${part2} ${part3} ${part4}`;
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
          className={`${gradient} ${border} group cursor-pointer rounded-b-3xl flex justify-between px-6 py-3 `}
        >
          <div className="flex items-center">
            <span className="text-16px font-text-navy group-hover:hidden ">
              {formatCardNumber(cardNumber)}
            </span>
            <span className="text-16px font-text-navy hidden group-hover:block">
              {makeCardNumberVisible(cardNumber)}
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
