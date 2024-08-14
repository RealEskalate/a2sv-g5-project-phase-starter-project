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
        <div className="flex gap-3 items-center min-w-fit">
          <Image
            src={image}
            alt=""
            width={70}
            height={70}
            className="object-cover w-11 h-11 md:w-12 md:h-12 lg:w-14 lg:h-14"
          />
          <div className="space-y-1 p-1 max-w-fit">
            <p
              className="flex text-13px lg:text-15px text-blue-steel overflow-clip"
              style={{
                display: "-webkit-box",
                WebkitLineClamp: 2,
                WebkitBoxOrient: "vertical",
                overflow: "hidden",
                textOverflow: "ellipsis",
              }}
            >
              {name}
            </p>
            <p className="flex text-gray-dark text-15px lg:text-17px font-[600]">
              {amount}
            </p>
          </div>
        </div>
      </div>
    </>
  );
}

export default InvestmentCard;
