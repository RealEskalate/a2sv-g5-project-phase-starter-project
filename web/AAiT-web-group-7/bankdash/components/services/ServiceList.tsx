import React from "react";

interface LoanCardProps {
  imageSrc: string;
  title: string;
  bgColor: string;
}

const ServiceList: React.FC<LoanCardProps> = ({ imageSrc, title, bgColor }) => {
  return (
    <div className="flex justify-around items-center gap-8 p-[15px] rounded-2xl bg-white">
      <div
        style={{ backgroundColor: bgColor }}
        className={`rounded-full flex justify-center p-4`}
      >
        <img src={imageSrc} />
      </div>
        <div className="flex flex-col gap-1">
          <div className="font-semibold text-sm">{title}</div>
          <div className="text-sm text-[#718EBF]">It is a long established</div>
        </div>
        <div className="flex flex-col  gap-1">
          <div className="font-semibold">Lorem Ipsum</div>
          <div className="text-sm text-[#718EBF]">Many publishing</div>
        </div>
        <div className="flex flex-col  gap-1">
          <div className="font-semibold">Lorem Ipsum</div>
          <div className="text-sm text-[#718EBF]">Many publishing</div>
        </div>
        <div className="flex flex-col gap-1">
          <div className="font-semibold">Lorem Ipsum</div>
          <div className="text-sm text-[#718EBF]">Many publishing</div>
        </div>
        <button className="border text-sm rounded-full border-[#718EBF] text-[#718EBF] px-4 py-2">
          View Details
        </button>
    </div>
  );
};

export default ServiceList;
