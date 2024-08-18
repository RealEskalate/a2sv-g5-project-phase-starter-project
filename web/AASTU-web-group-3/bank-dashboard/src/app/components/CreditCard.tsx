import React from "react";
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
    <div className={`w-full h-full m-auto ${backgroundImg} ${textColor} rounded-3xl relative shadow-2xl`}>
      <div className="w-full px-8 absolute top-8">
        <div className="flex justify-between px-2 py-1">
          <div>
            <p className={`font-bold text-xs ${cardText}  `}>Balance</p>
            <p className="font-semibold">{balance}</p>
          </div>
          <img className="w-10 h-10" src={sCard} alt="Card Logo" />
        </div>
        <div className="flex justify-start gap-14 py-2 px-3">
          <div>
            <p className={`font-thin ${cardText} text-xs`}>CARD HOLDER</p>
            <p className="font-semibold ">{name}</p>
          </div>
          <div>
            <p className={`font-thin ${cardText} text-xs`}>VALID THRU</p>
            <p className="font-semibold  text-lg ">{validDate}</p>
          </div>
        </div>

            <div className="flex justify-between mt-2 p-2">
              <p className="font-semibold tracking-wider">{cardNumber}</p>
                <img className="w-8 h-8" src={masterCardDarker} alt="Card Logo" />
            </div>

        </div>
    </div>
  );
};

export default CreditCard;
