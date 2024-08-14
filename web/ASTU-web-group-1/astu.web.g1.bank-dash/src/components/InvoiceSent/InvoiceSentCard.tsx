import React from "react";

type InvoiceSentCardProps = {
  name: string;
  time: string;
  amount: number;
  imageUrl: string;
};

const InvoiceSentCard: React.FC<InvoiceSentCardProps> = ({
  name,
  time,
  amount,
  imageUrl,
}) => {
  return (
    <div className="flex items-center">
      <div className="flex-shrink-0">
        <img
          className="lg:w-[50px] lg:h-[50px] w-[45px] h-[45px] lg:rounded-[20px] rounded-[12px]"
          src={imageUrl}
          alt={`${name} image`}
        />
      </div>
      <div className="flex-1 min-w-0 ms-4">
        <p className="text-sm font-medium text-[#B1B1B1] ">{name}</p>
        <p className="text-sm text-[#718EBF] dark:text-gray-400">{time} ago</p>
      </div>
      <div className="inline-flex items-center text-base text-[#718EBF] ">
        ${amount}
      </div>
    </div>
  );
};

export default InvoiceSentCard;
