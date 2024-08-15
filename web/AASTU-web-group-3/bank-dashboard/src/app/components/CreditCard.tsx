import React from "react";
import { simCard, masterCardLogo } from "../../../public/Icons";

type CardProps = {
  name: string;
  cardNumber: string;
  balance: string;
  validDate: string;
  backgroundImg: string;
};

const CreditCard: React.FC<CardProps> = ({
  name,
  cardNumber,
  validDate,
  balance,
  backgroundImg,
}) => {
  return (
    <div className={`w-2/5 h-56 m-auto bg-red-100 rounded-3xl relative text-white shadow-2xl transition-transform transform hover:scale-105`}>
      <img
        className="relative object-cover w-full h-full rounded-xl"
        src={backgroundImg}
        alt="Credit Card Background"
      />
      <div className="w-full px-8 absolute top-8">
        <div className="flex justify-between px-2">
          <div>
            <p className="font-light ">Balance</p>
            <p className="font-md">{balance}</p>
          </div>
          <img className="w-10 h-10" src={simCard} alt="Card Logo" />
        </div>
        <div className="flex justify-start gap-14 p-2">
          <div>
            <p className="font-light text-[#ffffff95] text-sm">CARD HOLDER</p>
            <p className="font-md">{name}</p>
          </div>
          <div>
            <p className="font-light text-xs text-[#ffffff95]">VALID THRU</p>
            <p className="font-md  text-sm">{validDate}</p>
          </div>
        </div>

            <div className="flex justify-between mt-2 p-2">
              <p className="font-md tracking-more-wider">{cardNumber}</p>
                <img className="w-8 h-8" src={masterCardLogo} alt="Card Logo" />
            </div>

        </div>
    </div>
  );
};

export default CreditCard;
