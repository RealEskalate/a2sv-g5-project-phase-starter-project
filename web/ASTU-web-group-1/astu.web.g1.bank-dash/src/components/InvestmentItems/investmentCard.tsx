import Image from "next/image";
import React from "react";

function InvestmentCard({
  image,
  name,
  amount,
}: {
  image: string;
  name: string;
  amount: string;
}) {
  return (
    <>
      <div className="flex bg-[#FFFFFF] rounded-3xl md:w-1/3 p-5 mx-1">
        <div className="flex gap-2 items-center w-full">
          <Image
            src={image}
            alt=""
            width={70}
            height={70}
            className="object-cover w-11 h-11 md:w-12 md:h-12 lg:w-14 lg:h-14"
          />
          <div className="space-y-1 p-1 overflow-hidden">
            <p className="flex text-13px lg:text-15px text-blue-steel overflow-clip truncate ...">
              {name}
            </p>
            <p className="flex text-gray-dark text-13px lg:text-16px font-[600] truncate ...">
              {amount}
            </p>
          </div>
        </div>
      </div>
    </>
  );
}

export default InvestmentCard;
