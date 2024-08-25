"use client";
import React from 'react';

interface cardData{
    balance: string;
    cardHolder: string;
    cardNumber: string;
    validThru: string;
}

interface CardsProps {
    bgCol?: string;
    bbgCol?: string;
    imageSrc?: string;
    iconSrc?: string;
    isBlue?: boolean;
    cData1: cardData; //single object
    cData2: cardData; //single object
}


const Card: React.FC<CardsProps> = ({ bgCol, bbgCol, imageSrc, iconSrc, isBlue, cData1, cData2 }) => {
return (
    <div className={`bg-[#FFFFFF] w-[231px] h-[170px] rounded-[20px] border-[1px] border-[#DFEAF2] flex flex-col justify-between p-[14px] ${bgCol}`}>
      {/* top */}
        <div className='flex justify-between'>
            <div>
                <p className={`font-lato text-[11px] font-normal leading-[13.2px] ${isBlue ? "text-[#FFFFFF]" : "text-[#343C6A]"} text-[#718EBF] text-left`}>Balance</p>
                <p className={`font-lato text-[16px] font-semibold leading-[19.2px] text-left ${isBlue ? "text-[#FFFFFF]" : "text-[#343C6A]"}`}>{isBlue ? cData1.balance : cData2.balance}</p>
            </div>
            <div>
                <img src={imageSrc} alt="chip_card" />
            </div>
        </div>

      {/* middle */}
        <div className='flex justify-between'>
            <div>
                <p className={`font-lato text-[10px] font-normal leading-[12px] text-left text-[#718EBF]`}>CARD HOLDER</p>
                <p className={`text-[13px] font-medium leading-[15.6px] text-left ${isBlue ? "text-[#FFFFFF]" : "text-[#343C6A]"}`}>{isBlue ? cData1.cardHolder : cData2.cardHolder}</p>
            </div>
            <div>
                <p className={`font-lato text-[10px] font-normal leading-[12px] text-left text-[#718EBF]`}>VALID THRU</p>
                <p className={`font-lato text-[13px] font-semibold leading-[15.6px] text-left ${isBlue ? "text-[#FFFFFF]" : "text-[#343C6A]"}`}>{isBlue ? cData1.validThru : cData2.validThru}</p>
            </div>
        </div>

      {/* bottom */}
        <div className={`flex justify-between ${bbgCol} w-full h-[35px]`}>
            <div className='flex items-center'>
                <p className={`font-lato text-[15px] font-semibold leading-[18px] text-left ${isBlue ? "text-[#FFFFFF]" : "text-[#343C6A]"}`}>{isBlue ? cData1.cardNumber :cData2.cardNumber}</p>
            </div>
            <div className='flex items-center'>
                <img src={iconSrc} alt="chip_card-icon" />
            </div>
        </div>
    </div>
);
};

export default Card;
