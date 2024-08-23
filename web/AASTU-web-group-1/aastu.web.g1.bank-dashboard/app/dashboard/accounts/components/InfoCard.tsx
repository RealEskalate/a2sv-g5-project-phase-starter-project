import React from "react";
import Image from "next/image";
import { useUser } from "@/contexts/UserContext";

interface Prop {
  title: string;
  amount: number;
  image: string;
  color: string;
}

const InfoCard = ({ image, color, title, amount }: Prop) => {
  const formattedAmount = amount.toLocaleString();
  const { isDarkMode } = useUser();

  // Conditional styles based on dark mode
  const containerClass = isDarkMode
    ? "bg-gray-800 text-gray-100"
    : "bg-white text-gray-800";
  const colorClass = isDarkMode
    ? `${color} bg-opacity-40` // slightly different opacity for dark mode
    : `${color} bg-opacity-25`;
  const textColor = isDarkMode ? "text-gray-400" : "text-[#718EBF]";

  return (
    <div className={`flex p-4 rounded-lg ${containerClass}`}>
      <div
        className={`flex items-center justify-center ${colorClass} font-semibold gap-2  rounded-lg text-sm w-[45px] h-[45px]  `}
      >
        <Image src={image} alt={title} width={20} height={20} />
      </div>
      <div className="ml-4">
        <div className={`text-sm ${textColor}`}>{title}</div>
        <div className="text-lg font-semibold">${formattedAmount}</div>
      </div>
    </div>
  );
};

export default InfoCard;
