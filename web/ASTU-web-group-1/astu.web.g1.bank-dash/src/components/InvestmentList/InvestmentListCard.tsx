import React from "react";

type InvestmentListCardProps = {
  companyName: string;
  amount: string;
  imageUrl: string;
  returnValue: string;
  sign: string;
  color: string;
};

const InvestmentListCard: React.FC<InvestmentListCardProps> = ({
  companyName,
  amount,
  imageUrl,
  returnValue,
  sign,
  color,
}) => {
  return (
    <div className="flex  w-full bg-white rounded-[20px] p-2 ">
      <div className="flex flex-shrink-0 mr-4 items-center">
        <img
          className="lg:w-[55px] lg:h-[55px] w-[40px] h-[40px] sm:rounded-[20px]"
          src={imageUrl}
          alt="card image"
        />
      </div>
      <div className="flex-1 min-w-0 py-2 ">
        <p className="text-sm text-[#232323] truncate">{companyName}</p>
        <p className="text-xs text-[#718EBF] truncate ...">E-commerce, Marketplace</p>
      </div>
      <div className="hidden  md:block flex-1 min-w-0 py-2">
        <p className="text-sm text-[#232323] truncate">${amount}</p>
        <p className="text-xs text-[#718EBF] truncate ...">Envestment Value</p>
      </div>
      <div className=" min-w-0 py-2">
        <p className="text-sm text-[#232323] truncate" style={{ color }}>
          {sign}
          {returnValue}
        </p>
        <p className="text-xs text-[#718EBF] truncate">Return Value</p>
      </div>
    </div>
  );
};

export default InvestmentListCard;
