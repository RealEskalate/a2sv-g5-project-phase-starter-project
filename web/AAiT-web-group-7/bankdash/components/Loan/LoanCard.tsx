import React from "react";

interface LoanCardProps {
  imageSrc: string;
  title: string;
  amount: string;
  bgColor: string;
}

const LoanCard: React.FC<LoanCardProps> = ({
  imageSrc,
  title,
  amount,
  bgColor,
}) => {
  return (
    <div className="p-5 rounded-3xl flex justify-center items-center bg-white text-sm">
      <div className="flex flex-row gap-4 justify-between items-center">
        <div
          style={{ backgroundColor: bgColor }}
          className={`rounded-full flex justify-center w-14 h-14 p-5 `}
        >
          <img src={imageSrc} />
        </div>
        <div className="flex flex-col item-start">
          <div className="text-[#718EBF]">{title}</div>
          <div className="text-[#232323] font-semibold">{amount}</div>
        </div>
      </div>
    </div>
  );
};

export default LoanCard;
