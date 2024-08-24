import React from "react";

interface LoanCardProps {
  imageSrc: string;
  title: string;
  decription: string;
  bgColor: string;
}

const ServiceCard: React.FC<LoanCardProps> = ({
  imageSrc,
  title,
  decription,
  bgColor,
}) => {
  return (
    <div className="w-[350px] h-[120px] rounded-3xl flex justify-center items-center bg-white">
      <div className="flex flex-row w-[246px] h-[70px]  justify-between items-center">
        <div
          style={{ backgroundColor: bgColor }}
          className={`rounded-full flex justify-center w-[70px] h-[70px] p-5 `}
        >
          <img src={imageSrc} />
        </div>
        <div className="flex flex-col w-[154px] h-[51px] gap-1">
          <div className="font-[intel] text-[#232323] font-bold">{title}</div>
          <div className="font-[intel] text-[#718EBF]">{decription}</div>
        </div>
      </div>
    </div>
  );
};

export default ServiceCard;
