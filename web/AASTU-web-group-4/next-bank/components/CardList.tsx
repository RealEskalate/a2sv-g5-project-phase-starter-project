import React from "react";
import { colors, sidebarLinks } from "../constants/index";
import { CgCreditCard } from "react-icons/cg";
import { cardType } from "@/types";

interface Props {
  card_list: cardType[];
}

const CardList: React.FC<Props> = ({ card_list }) => {
  return (
    <div className="max-h-[400px] lg:w-[730px] md:w-[487px] w-[325] overflow-y-scroll pr-6 py-4 scrollbar-thin scrollbar-thumb-rounded scrollbar-thumb-blue-400">
      {card_list.map((card: any, index: number) => (
        <div
          key={index}
          className="flex flex-row justify-between w-full bg-white p-4 mb-4 rounded-lg shadow-md "
        >
          {/* Icon */}
          <div className="h-12 w-12 rounded-xl bg-[#E7EDFF] flex items-center justify-center">
            <CgCreditCard className="w-7 h-7" />
          </div>

          {/* Card Type */}
          <div className="flex flex-col">
            <p className="text-sm font-medium">Card Type</p>
            <p className={`${colors.textgray} text-xs `}>{card.cardType}</p>
          </div>

          {/* Bank */}
          <div className="flex flex-col">
            <p className="text-sm font-medium">Bank</p>
            <p className={`${colors.textgray} text-xs`}>Next-bank</p>
          </div>

          {/* Card Number - Hide on small screens */}
          <div className="flex flex-col">
            <p className="text-sm font-medium hidden sm:block">Card Number</p>
            <p className={`${colors.textgray} text-xs hidden sm:block`}>
              {card.semiCardNumber}
            </p>
          </div>

          {/* Card Name - Hide on small screens */}
          <div className="flex flex-col">
            <p className="text-sm font-medium hidden sm:block">Card Name</p>
            <p className={`${colors.textgray} text-xs hidden sm:block`}>
              {card.cardHolder.split(' ')[0]}
            </p>
          </div>

          {/* View Details Link */}
          <a
            href="#"
            className="text-[#1814F3] font-semibold text-sm sm:text-xs"
          >
            View Details
          </a>
        </div>
      ))}
    </div>
  );
};

export default CardList;

/* Group 343 */
