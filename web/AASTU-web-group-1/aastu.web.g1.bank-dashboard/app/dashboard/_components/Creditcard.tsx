import { FcSimCardChip } from "react-icons/fc";
import React from "react";
import Image from "next/image";
interface Props {
  isBlue:boolean
  balance: number;
  creditNumber: string;
  name: string;
  textColor: string;
}
const CreditCard = ({
  balance,
  isBlue,
  creditNumber,
  name,
  textColor,
}: Props) => {
  return (
    <div
      className={` min-w-[300px] h-[200px] ${
        isBlue ? "bg-gradient-to-r from-[#423fee] to-[#2723f1]" : "bg-white"
      } rounded-xl pt-3 space-y-5 border border-gray-300`}
    >
      <div className="flex justify-between px-5">
        <div className={`block ${textColor} space-y-[1px]`}>
          <p className="text-[11px]" style={{ fontWeight: 400 }}>
            Balance
          </p>
          <p className="text-[16px]" style={{ fontWeight: 700 }}>
            ${balance}
          </p>
        </div>
        <FcSimCardChip size={30} />
      </div>

      <div className="flex justify-between px-5">
        <div className="block space-y-[1px]">
          <p className="text-[10px] text-gray-400" style={{ fontWeight: 400 }}>
            CARD HOLDER
          </p>
          <p className={`text-[13px] ${textColor} `} style={{ fontWeight: 700 }}>
            {name}
          </p>
        </div>
        <div className="block space-y-[1px]">
          <p className="text-[10px] text-gray-400" style={{ fontWeight: 400 }}>
            VALID THRU
          </p>
          <p className={`text-[13px] ${textColor} `}style={{ fontWeight: 700 }}>
            12/22
          </p>
        </div>
      </div>

      <div className="relative">
        <div className="absolute top-0 left-0 w-full h-3/4 backdrop-blur-[2px] bg-gradient-to-b from-white/30 to-transparent border-t-gray-300"></div>

        <div className="relative flex justify-between px-5 items-center py-5">
          <p className={`text-[15px] ${textColor} style={{ fontWeight: 700 }`}>
            {creditNumber}
          </p>
          <Image
            src={`/images/intersection.png`}
            alt={"transaction"}
            width={27}
            height={18}
          />
        </div>
      </div>
    </div>
  );
};

export default CreditCard;
