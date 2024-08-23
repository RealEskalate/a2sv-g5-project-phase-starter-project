import React from "react";
import CardImage from "@/public/Group.svg";
import Image from "next/image";
import whiteChip from "@/public/Chip_white.svg";
import BlackChip from "@/public/Chip_black.svg";
import union from "@/public/union.svg";

const BalanceCard = ({ property }: { property: string }) => {
  const chip = property == "blue" ? whiteChip : BlackChip;
  const unionImage = property == "blue" ? CardImage : union;
  return (
    <div
      className={`w-[300px]  mx-auto md:mx-0 md:w-[330px] rounded-3xl px-5 py-3 space-y-5 shadow-md ${
        property == "blue"
          ? "bg-gradient-to-l from-[#0a06f4] to-[#4c49ed] text-white"
          : "bg-white"
      }`}
    >
      <div className="flex justify-between items-center">
        <div className="">
          <p className="text-[#718EBF] text-xs">Balance</p>
          <span className="font-semibold">$5,756</span>
        </div>
        <Image src={chip} alt="chip image" />
      </div>
      <div className="flex justify-between items-center">
        <div className="text-sm">
          <p className="text-[#718EBF] text-xs">CARD HOLDER</p>
          <span className="font-semibold">Eddy Cusuma</span>
        </div>
        <div className="text-sm">
          <p className="text-xs text-[#718EBF]">VALID THRU</p>
          <span className="font-semibold">12/22</span>
        </div>
      </div>
      <div
        className={` ${
          property == "blue" ? "bg-[#4845e8]" : "border-t bg-white"
        } flex justify-between items-center p-3  rounded-3xl`}
      >
        <p className="font-semibold">3778**** **** 1234</p>
        <Image src={unionImage} alt="cardimage" className="text-black" />
      </div>
    </div>
  );
};

export default BalanceCard;
