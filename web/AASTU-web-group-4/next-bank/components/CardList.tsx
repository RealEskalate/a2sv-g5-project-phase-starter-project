import React from "react";
import { colors, sidebarLinks } from "../constants/index";
import { CgCreditCard } from "react-icons/cg";
import { cardType } from "@/types/index";
import Image from "next/image";

interface Props {
  card_list: cardType[];
}

const CardList: React.FC<Props> = ({ card_list }) => {
  return (
    <div className="max-h-[400px] lg:w-[730px] md:w-[487px] w-[325] overflow-y-scroll pr-6 py-4 scrollbar-thin scrollbar-track-[#F5F7FA] dark:scrollbar-track-dark scrollbar-thumb-[#b5c2d9] scrollbar-thumb-rounded-full">
      {Array.isArray(card_list) && card_list.length > 0 ? (
        card_list.map((card: any, index: number) => (
        <div
          key={index}
          className="flex flex-row justify-between w-full bg-white p-4 mb-4 rounded-lg shadow-md dark:bg-dark dark:text-white"
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
      ))) : (
        <div className=" pr-6 py-32 bg-white w-full flex flex-col justify-center align-middle rounded-xl scrollbar-none">
          <Image
                src="/icons/null.png"
                width={80}
                height={80}
                alt="null"
                className="mx-auto pb-2 block"
              />
          <span className="mx-auto my-auto md:text-xl text-sm text-[#993d4b] font-bold">
            There is no cards for now!
            </span>
        </div>
      )}
    </div>
  );
};

export default CardList;

/* Group 343 */
