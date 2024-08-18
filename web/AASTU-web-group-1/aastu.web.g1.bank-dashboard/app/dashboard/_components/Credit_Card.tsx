import { FcSimCardChip } from "react-icons/fc";
import React from "react";
import Image from "next/image";
import { CardDetails } from "@/types";
import formatDateString from "../transactions/component/formatDateString";
import maskCardNumber from "../transactions/component/maskCardNumber";




const CreditCard: React.FC<CardDetails> = ({
  cardHolder,semiCardNumber,cardType,balance,expiryDate
}:CardDetails) => {

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
       bgColor = "bg-white"; // Default white background
       textColor = "text-black";
   }
  return (
    <div
      className={` min-w-[300px] h-55 ${
        bgColor
      } rounded-2xl pt-3 space-y-5 border border-gray-300`}
    >
      <div className="flex justify-between px-5">
        <div className={`block ${textColor} space-y-[1px]`}>
          
          <p className="text-[11px] font-lato font-normal ">
            Balance
          </p>
          <p className="text-[16px] font-lato font-semibold" >
            ${balance}
          </p>
        </div>
        <FcSimCardChip size={30} />
      </div>

      <div className="flex px-5">
        <div className="w-[60%] block space-y-[1px]">
          <p className="text-[10px] text-gray-400 gont-lato font-normal" >
            CARD HOLDER
          </p>
          <p
            className={`text-[13px] ${textColor} font-lato font-semibold`}
          >
            {cardHolder}
          </p>
        </div>
        <div className="block space-y-[1px]">
          <p className="text-[10px] text-gray-400 font-lato font-normal" >
            Exp. Date
          </p>
          <p
            className={`text-[13px] ${textColor} font-lato font-semibold`}
        
          >
            {formatDateString(expiryDate)}
          </p>
        </div>
      </div>

      <div className="relative">
        <div className={`absolute top-0 left-0 w-full h-1/3 backdrop-blur-[1px] bg-gradient-to-b from-white/30 to-transparent
           ${bgColor ==="bg-white" ?"border-t" :"" }`}></div>

        <div className="relative flex justify-between px-5 items-center py-5">
          <p className={`text-[15px] ${textColor} font-lato font-semibold w-[80%]`}>
            {maskCardNumber(semiCardNumber)}
          </p>
          <Image
            src={`${bgColor !== "bg-white" ? "/icons/cardwhite.svg":"/icons/cardgray.svg"}`}
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
