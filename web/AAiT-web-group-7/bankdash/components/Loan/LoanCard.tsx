import React from "react";

interface LoanCardProps {
  imageSrc: string;
  title: string;
  amount: string;
  bgColor: string;
}

const LoanCard: React.FC<LoanCardProps> = ({ imageSrc, title, amount ,bgColor }) => {
  return (
    <div className="w-[255px] h-[120px] rounded-3xl flex justify-center items-center bg-white">
      <div className="flex flex-row w-[201px] h-[70px] justify-between items-center">
        <div
          className={`rounded-full flex justify-center w-[70px] h-[70px] p-5 bg-[${bgColor}]`}
        >
          <img src={imageSrc} />
        </div>
        <div className="flex flex-col w-[116px] h-[51px] gap-1">
          <div className="font-[intel] text-[#718EBF]">{title}</div>
          <div className="text-[#232323] font-[intel] font-bold">{amount}</div>
        </div>
      </div>
    </div>
  );
};

export default LoanCard;
