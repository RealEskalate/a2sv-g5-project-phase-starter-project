import React from "react";
import Image from "next/image";

interface prop {
  title: string;
  amount: number;
  image: string;
  color: string
}

const InfoCard = ({ image, color, title, amount }: prop) => {
  const formattedAmount = amount.toLocaleString();
  
  return (
      <div className="flex bg-white">
        <div className={`${color} bg-opacity-25 font-semibold py-1 px-2 rounded-lg  text-sm w-[45px]`}>
          <Image
          src={image}
          alt={title}
          width={20}
          height={20}
          />
        </div>
        <div>
            <div className="text-[#718EBF]">
                {title}
            </div>
            <div>
               ${formattedAmount}
            </div>
        </div>
      </div>
  );
};

export default InfoCard;
