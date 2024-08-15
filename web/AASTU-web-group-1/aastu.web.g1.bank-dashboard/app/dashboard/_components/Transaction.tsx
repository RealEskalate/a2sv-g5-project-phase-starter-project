import React from "react";
import Image from "next/image";

interface Props {
  image: string;
  transactionType: string;
  date: string;
  amount: string;
  color: string;
}
export const Transaction = ({
  image,
  transactionType,
  date,
  amount,
  color,
}: Props) => {
  
  return (
    <div className="flex justify-between">
      <div className="flex space-x-2">
        <div
          className={`inline-flex items-center justify-center ${color} rounded-full w-[50px] h-[50px]`}
        >
          <Image
            src={image}
            alt={`transation icon`}
            className="object-cover object-center"
            width={20}
            height={20}
          />
        </div>

        <div>
          <p className={`font-inter text-[14px]`} style={{ fontWeight: 500 }}>
            {transactionType}
          </p>
          <p
            className={`font-inter text-[12px] text-indigo-400`}
            style={{ fontWeight: 400 }}
          >
            {date}
          </p>
        </div>
      </div>
      <div>
        <p
          className={`font-inter text-[11px] text-green-600`}
          style={{ fontWeight: 500 }}
        >
          {amount}
        </p>
      </div>
    </div>
  );
};
