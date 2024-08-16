import React from "react";
import Image from "next/image";
import { simCard, masterCardLogo,masterCardLogoDarker,simCardDarker } from "../../../public/Icons";

type CardProps = {
  name: string;
  cardNumber: string;
  balance: string;
  validDate: string;
  backgroundImg: string;
  textColor: string;
};

const CreditCard: React.FC<CardProps> = ({
  name,
  cardNumber,
  validDate,
  balance,
  backgroundImg,
  textColor,
}) => {

 const  masterCardDarker =  backgroundImg === 'bg-white'?masterCardLogoDarker:masterCardLogo
 const  sCard =  backgroundImg === 'bg-white'?simCardDarker:simCard
 const  cardText =  backgroundImg === 'bg-white'?"text-[#718EBF]":"text-[#ffffff95]"
  return (
    <div className={`w-full h-48 m-auto ${backgroundImg} ${textColor} rounded-3xl relative  shadow-2xl`}>
      <div className="w-full absolute top-4">
        <div className="flex justify-between px-2 py-1">
          <div>
            <p className={` lg:font-bold text-xs ${cardText}`}>Balance</p>
            <p className="font-semibold">{balance}</p>
          </div>
          <Image width={34} height={34} src={sCard} alt="Card Logo" />
        </div>
        <div className="flex justify-start gap-10 py-2 px-3">
          <div>
            <p className={`font-thin ${cardText} text-[10px]`}>CARD HOLDER</p>
            <p className="font-semibold text-sm lg:text-sm ">{name}</p>
          </div>
          <div>
            <p className={`font-thin ${cardText} text-sm`}>VALID THRU</p>
            <p className="font-semibold  text-lg ">{validDate}</p>
          </div>
        </div>

            <div className="flex justify-between mt-2 p-2">
              <p className="font-semibold tracking-wider">{cardNumber}</p>
                <Image width= {24} height={24}  src={masterCardDarker} alt="Card Logo" />
            </div>

        </div>
    </div>
  );
};

export default CreditCard;
