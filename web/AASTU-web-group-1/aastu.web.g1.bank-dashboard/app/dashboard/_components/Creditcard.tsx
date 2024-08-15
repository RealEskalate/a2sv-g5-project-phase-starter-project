import React from 'react'
import { FcSimCardChip } from "react-icons/fc";
import Image from 'next/image';
interface Props{
  color:string,
  creditNumber:string
  balance:number,
  name:string
  textColor:string
}

export const Creditcard = ({color,creditNumber,balance,name,textColor}:Props) => {
  return (
    <div className={`max-w-[350px] max-h-[235px] min-w-[260px] ${color} w-[265px]  h-[170px]  rounded-xl pt-3 space-y-5 border border-gray-300`}>
      <div className="flex justify-between px-5">
        <div className={`block ${textColor} space-y-[1px]`}>
          <p
            className={`font-lato text-[11px]`}
            style={{ fontWeight: 400 }}
          >
            Balance
          </p>
          <p
            className={`text-[16px] font-semibold`}
          >
            ${balance}
          </p>
        </div>
        <FcSimCardChip size={30} />
      </div>

      <div className="flex justify-between px-5">
        <div className={`block space-y-[1px]`}>
          <p
            className={`font-lato text-[10px] text-gray-400`}
            style={{ fontWeight: 400 }}
          >
            CARD HOLDER
          </p>
          <p
            className={`font-lato text-[13px] ${textColor} font-semibold`}
          >
            {name}
          </p>
        </div>
        <div className={`block space-y-[1px]`}>
          <p
            className={`font-lato text-[10px] text-gray-400`}
            style={{ fontWeight: 400 }}
          >
            VALID THRU
          </p>
          <p
            className={`font-lato text-[13px] ${textColor} font-semibold`}
           
          >
            End 12/22
          </p>
        </div>
      </div>

      <div className="relative">
        <div className="absolute top-0 left-0 w-full h-3/4 backdrop-blur-[2px] bg-gradient-to-b from-white/20 to-transparent border-t border-gray-200 "></div>

        <div className="relative flex justify-between px-5 items-center py-1">
          <p
            className={`font-lato text-[15px] ${textColor} font-semibold`}
          >
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
}
