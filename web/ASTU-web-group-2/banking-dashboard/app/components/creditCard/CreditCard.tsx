"use client";
import React, { useState } from "react";
import Image from "next/image";
{
  /* <CreditCard
        balance={1250}
        cardHolder="John Doe"
        expiryDate="12/24"
        cardNumber="1234 5678 9012 3456"
        cardType="tertiary" // Can be "primary", "secondary", or "tertiary"
      /> */
}

interface Props {
  balance: number;
  cardHolder: string;
  expiryDate: string;
  cardNumber: string;
  cardType: "primary" | "secondary" | "tertiary";
}

const CreditCard = ({
  balance,
  cardHolder,
  cardNumber,
  expiryDate,
  cardType,
}: Props) => {
  // Format the card number
  const formatCardNumber = (number: string) => {
    if (number.length >= 8) {
      const firstPart = number.slice(0, 4);
      const lastPart = number.slice(-4);
      // const hiddenPartone = "*".repeat(8);
      const hiddenPart = "**** ****";
      return `${firstPart} ${hiddenPart} ${lastPart}`;
    }
    return number;
  };

  // Handle card number hover reveal
  const [isHovered, setIsHovered] = useState(false);

  const cardBackground =
    cardType === "primary"
      ? "bg-credit-card-gradient"
      : cardType === "secondary"
      ? "bg-secondary-credit-card"
      : "bg-white border border-[#DFEAF2] border-b-0";

  const cardFooterBackground =
    cardType === "primary"
      ? "bg-gradient-to-b to-[#0A06F4] from-[#4C49ED]"
      : cardType === "secondary"
      ? "bg-gradient-to-b to-[#2D60FF] from-[#539BFF]"
      : "bg-white border border-[#DFEAF2] border-t-0";

  const textColor = cardType === "tertiary" ? "#343c6a" : "#ffffff";
  const labelColor = cardType === "tertiary" ? "#718EBF" : "#ffffffb3";
  const chipIcon =
    cardType === "tertiary"
      ? "/assets/credit-card/teritary_chip_card.svg"
      : "/assets/credit-card/Chip_Card.svg";
  const ellipseIcon =
    cardType === "tertiary"
      ? "/assets/credit-card/teritary_elipse.svg"
      : "/assets/credit-card/elipse.svg";

  return (
    <div className="flex flex-col">
      <div
        className={`w-[350px] h-[165px] ${cardBackground}  rounded-t-[25px]  `}
      >
        <div
          className={`font-normal leading-[14.4px] pl-[26px] pt-[24px] flex gap-[202px]`}
          style={{ color: textColor }}
        >
          <div>
            <p
              className={`text-[12px] leading-[14.4px] font-normal max-sm:text-[11px] max-sm:leading-[13.2px]`}
              style={{ color: labelColor }}
            >
              Balance
            </p>
            <p
              className={`font-semibold text-[20px] leading-[24px] max-sm:text-[16px] max-sm:leading-[19.2px]`}
              style={{ color: textColor }}
            >
              ${balance}
            </p>
          </div>
          <div className="">
            <Image alt="icon" src={chipIcon} width={34.77} height={34.77} />
          </div>
        </div>
        <div className="flex justify-between pt-[33px] pl-[24px] pr-[24px]">
          <div>
            <p
              className={`font-normal leading-[14.4px] text-[12px] max-sm:text-[10px]   max-sm:leading-[12px]`}
              style={{ color: labelColor }}
            >
              CARD HOLDER
            </p>
            <p
              className={`font-semibold text-[15px] max-sm:text-[13px] max-sm:leading-[15.6px]  leading-[18px]`}
              style={{ color: textColor }}
            >
              {cardHolder}
            </p>
          </div>
          <div>
            <p
              className={`text-xs font-normal leading-[14.4px] max-sm:text-[10px]   max-sm:leading-[12px]`}
              style={{ color: labelColor }}
            >
              VALID THRU
            </p>
            <p
              className={`font-semibold text-[15px] leading-[18px] max-sm:text-[13px] max-sm:leading-[15.6px]`}
              style={{ color: textColor }}
            >
              {expiryDate}
            </p>
          </div>
        </div>
      </div>
      <div
        className={`w-[350px] h-[70px] ${cardFooterBackground} rounded-b-[25px] flex  gap-[32px] items-center justify-between px-[26px]`}
      >
        <p
          className={`font-semibold  ${
            isHovered ? "text-[20px]" : "text-[20px]"
          } max-sm:text-[10px] leading-[26.4px] `}
          style={{ color: textColor }}
          onMouseEnter={() => setIsHovered(true)}
          onMouseLeave={() => setIsHovered(false)}
        >
          {isHovered ? cardNumber : formatCardNumber(cardNumber)}
        </p>
        <div>
          <Image alt="icon" src={ellipseIcon} width={44} height={44} />
        </div>
      </div>
    </div>
  );
};

export default CreditCard;
