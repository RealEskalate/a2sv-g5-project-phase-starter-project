import React from "react";

interface LoanCardProps {
  imageSrc: string;
  title: string;
  bgColor: string;
}

const ServiceList: React.FC<LoanCardProps> = ({ imageSrc, title, bgColor }) => {
  return (
    <div className="flex justify-center items-center gap-[20px] p-[15px] rounded-2xl bg-white">
      <div className={`bg-[${bgColor}] w-[60px] h-[60px] rounded-3xl flex justify-center p-[18px]`}>
        <img src={imageSrc} />
      </div>
      <div className=" gap-[73px] w-[995px] h-[46px] flex flex-row items-center ">
        <div className="flex flex-col gap-[1px]">
          <div className="font-semibold text-sm">{title}</div>
          <div className="text-sm text-[#718EBF]">It is a long established</div>
        </div>
        <div className="flex flex-col  gap-[1px]">
          <div className="font-semibold">Lorem Ipsum</div>
          <div className="text-sm text-[#718EBF]">Many publishing</div>
        </div>
        <div className="flex flex-col  gap-[1px]">
          <div className="font-semibold">Lorem Ipsum</div>
          <div className="text-sm text-[#718EBF]">Many publishing</div>
        </div>
        <div className="flex flex-col gap-[1px]">
          <div className="font-semibold">Lorem Ipsum</div>
          <div className="text-sm text-[#718EBF]">Many publishing</div>
        </div>
        <button className="border text-sm w-[150px] h-[35px] rounded-full  border-[#718EBF] text-[#718EBF] px-[31px]">
          View Details
        </button>
      </div>
    </div>
  );
};

export default ServiceList;
