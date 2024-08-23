"use client";
import { CardDetails } from "@/types";
import Image from "next/image";
import React from "react";
import { FcSimCardChip } from "react-icons/fc";
import {
  formatDateString,
  maskCardNumber,
} from "../transactions/component/utils";
import { useUser } from "@/contexts/UserContext";

const CreditCard: React.FC<CardDetails> = ({
  cardHolder,
  semiCardNumber,
  cardType,
  balance,
  expiryDate,
}: CardDetails) => {
  let bgColor = "";
  let textColor = "";

  switch (cardType) {
    case "Visa":
      bgColor = "bg-blue-700";
      textColor = "text-white";
      break;
    case "MasterCard":
      bgColor = "bg-black";
      textColor = "text-white";
      break;
    case "American Express":
      bgColor = "bg-gray-300";
      textColor = "text-black";
      break;
    default:
      bgColor = "bg-white";
      textColor = "text-black";
  }

  const { isDarkMode } = useUser();

  return (
    <div
      className={`min-w-[300px] h-56  md:w-[350px] md:h-[220px] ${bgColor} rounded-3xl pt-3 sm:space-y-8 md:space-y-6 ${
        !isDarkMode ? "border border-gray-300" : ""
      }`}
    >
      <div className="flex justify-between px-5">
        <div className={`block ${textColor} space-y-1`}>
          <p className="text-[11px] md:text-[12px] font-lato font-normal">
            Balance
          </p>
          <p className="text-[16px] md:text-[18px] font-lato font-semibold">
            ${balance}
          </p>
        </div>
        <FcSimCardChip size={50} />
      </div>

      <div className="flex px-5">
        <div className="w-[60%] block space-y-1">
          <p className="text-[10px] md:text-[11px] text-gray-400 font-lato font-normal">
            CARD HOLDER
          </p>
          <p
            className={`text-[13px] md:text-[14px] ${textColor} font-lato font-semibold`}
          >
            {cardHolder}
          </p>
        </div>
        <div className="block space-y-1">
          <p className="text-[10px] md:text-[11px] text-gray-400 font-lato font-normal">
            Exp. Date
          </p>
          <p
            className={`text-[13px] md:text-[14px] ${textColor} font-lato font-semibold`}
          >
            {formatDateString(expiryDate)}
          </p>
        </div>
      </div>

     
        <div
          className={`flex justify-between px-5 items-center md:space-y-1  py-5 backdrop-blur-[3px] bg-gradient-to-r from-white/30 to-white/5 rounded-b-3xl`}
        >
          <p
            className={`text-[15px] md:text-[16px] ${textColor} font-lato font-semibold w-[80%]`}
          >
            {maskCardNumber(semiCardNumber)}
          </p>
          <Image
            src={`${
              bgColor !== "bg-white"
                ? "/icons/cardwhite.svg"
                : "/icons/cardgray.svg"
            }`}
            alt={"transaction"}
            width={30}
            height={20}
          />
        </div>
    
    </div>
  );
};

export default CreditCard;