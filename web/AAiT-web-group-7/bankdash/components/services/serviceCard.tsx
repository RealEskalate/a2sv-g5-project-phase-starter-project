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
    <div className="rounded-3xl flex justify-between items-center bg-white py-4 px-16 text-sm">
      <div className="flex justify-between items-center gap-4">
        <div
          style={{ backgroundColor: bgColor }}
          className={`rounded-full flex justify-center p-5 `}
        >
          <img src={imageSrc} />
        </div>
        <div className="flex flex-col">
          <div className="text-[#232323] font-semibold">{title}</div>
          <div className="text-[#718EBF]">{decription}</div>
        </div>
      </div>
    </div>
  );
};

export default ServiceCard;
