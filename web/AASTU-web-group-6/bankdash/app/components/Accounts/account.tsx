import React from "react";
import Image from "next/image";
interface props {
  title: string;
  amount: string;
  icon: string;
  color: string;
  width: string;
}
const Card = ({ title, amount, icon, color, width }: props) => {
  return (
    <div
      className={`flex ${width}  justify-center items-center rounded-3xl p-4 gap-2 lg:gap-7 bg-white dark:bg-[#232328] min-w-[170px]`}
    >
      <div
        className="border  flex justify-center items-center rounded-full w-[45px] h-[45px] lg:w-[70px] lg:h-[70px]"
        style={{ backgroundColor: color, borderColor: color }}
      >
        <Image src={icon} width={24} height={24} alt="" />
      </div>

      <div>
        <p className="text-[#718EBF] font-normal text-base font-inter dark:text-gray-300">
          {title}
        </p>
        <p className="text-[#232323] font-semibold text-2xl font-inter dark:text-gray-200">
          {amount}
        </p>
      </div>
    </div>
  );
};

export default Card;
