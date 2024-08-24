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
    cData1: cardData;
    cData2: cardData;
}



const Card = () => {

    const  bgCol="bg-custom-gradient"
    const   bbgCol="bg-bottom-gradient"
    const   isBlue=false //this is for card-color
    const   imageSrc= "./images/Chip_Card.png"
    const   iconSrc= "/images/Group 17.png"
    //  const   imageSrc = "/images/Chip_Card.png", iconSrc = "/images/Group 17.png"

const card_data = [
  { balance: "$5,708", cardHolder: "Eddy Cusuma", cardNumber: "3778 **** **** 1234", validThru: "12/26" },
  { balance: "$908", cardHolder: "Jatani Iya", cardNumber: "3778 **** **** 5678", validThru: "11/25" },
  { balance: "$2,908", cardHolder: "Dida Jaldo", cardNumber: "3778 **** **** 9101", validThru: "02/30" },
  { balance: "$1,500", cardHolder: "Mikael Tesfaye", cardNumber: "3778 **** **** 1121", validThru: "05/27" },
  { balance: "$3,200", cardHolder: "Sara Yared", cardNumber: "3778 **** **** 3141", validThru: "08/28" },
  { balance: "$4,750", cardHolder: "Liya Alemu", cardNumber: "3778 **** **** 5161", validThru: "09/29" }
];
const cData1=card_data[0]
const cData2=card_data[1]



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
