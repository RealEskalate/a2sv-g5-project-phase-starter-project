import React from "react";
import Image from "next/image";
import {
  simCard,
  masterCardLogo,
  masterCardLogoDarker,
  simCardDarker,
} from "../../../public/Icons";
import { CardProps } from "../../types/creditCardData";

function formatDate(dateString: string): string {
  const date = new Date(dateString);
  const month = (date.getMonth() + 1).toString().padStart(2, "0"); // Add 1 because months are zero-indexed
  const year = date.getFullYear().toString().slice(-2); // Get the last two digits of the year
  return `${month}/${year}`;
}

function formatCardNumber(cardNumber: string): string {
  const start = cardNumber?.slice(0, 4);
  const end = cardNumber?.slice(-4);
  return `${start} **** **** ${end}`;
}

function formatBalance(balance: number): string {
  return new Intl.NumberFormat("en-US", {
    style: "currency",
    currency: "USD",
    minimumFractionDigits: 2,
  }).format(balance);
}

export default function CreditCard({
  name,
  cardNumber,
  validDate,
  balance,
  backgroundImg,
  textColor,
}: CardProps) {
  const masterCardDarker =
    backgroundImg === "bg-white" ? masterCardLogoDarker : masterCardLogo;
  const sCard = backgroundImg === "bg-white" ? simCardDarker : simCard;
  const cardText =
    backgroundImg === "bg-white" ? "text-[#718EBF]" : "text-[#ffffff95]";
  return (
    <div
      className={` w-full h-52 xl:h-52  ${backgroundImg} ${textColor} rounded-3xl relative overflow-hidden`}
    >
      <div className="w-full flex flex-col justify-around absolute top-4  h-44">
        <div className="flex justify-between px-5 py-1">
          <div>
            <p className={` lg:font-bold text-xs ${cardText}`}>Balance</p>
            <p className="font-semibold">{formatBalance(parseInt(balance))}</p>
          </div>
          <Image width={34} height={34} src={sCard} alt="Card Logo" />
        </div>
        <div className="flex justify-start gap-10 py-2 px-5">
          <div>
            <p className={`font-thin ${cardText} text-[10px]`}>CARD HOLDER</p>
            <p className="font-semibold text-sm lg:text-sm ">{name}</p>
          </div>

          <div>
            <p className={`font-thin ${cardText} text-sm`}>VALID THRU</p>
            <p className="font-semibold  text-lg ">{formatDate(validDate)}</p>
          </div>
        </div>

        <div className="flex justify-between items-end  mt-2 px-5">
          <p className="font-semibold tracking-wider">
            {formatCardNumber(cardNumber)}
          </p>
          <Image
            width={24}
            height={24}
            src={masterCardDarker}
            alt="Card Logo"
          />
        </div>
        <div className="absolute bottom-0 w-full h-12 bg-gradient-to-b from-gray-400/40 to-gray-400/0"></div>
      </div>
    </div>
  );
}
