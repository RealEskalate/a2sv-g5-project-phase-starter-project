import Image from "next/image";
import React from "react";

function LoanCard({
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
      <div className="flex bg-[#FFFFFF] rounded-2xl px-4 py-4 mr-2 md:pl-5 md:pr-9 md:py-5">
        <div className="flex items-center space-x-1 min-w-fit">
          <Image
            src={image}
            alt=""
            width={0}
            height={0}
            className="object-cover w-11 h-11 md:w-12 md:h-12 lg:w-14 lg:h-14"
          />
          <div className="space-y-1 max-w-full pr-10">
            <p className="flex text-13px lg:text-15px text-blue-steel">
              {name}
            </p>
            <p className="text-gray-dark text-12px lg:text-15px font-[600]">
              {amount}
            </p>
          </div>
        </div>
      </div>
    </>
  );
}

export default LoanCard;
