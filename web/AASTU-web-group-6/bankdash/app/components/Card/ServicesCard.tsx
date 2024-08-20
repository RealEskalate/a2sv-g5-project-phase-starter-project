import React from "react";

interface ServicesCardProps {
  img: string;
  title: string;
  desc: string;
}

const ServicesCard: React.FC<ServicesCardProps> = ({ img, title, desc }) => {
  return (
    <div className="flex items-center justify-center grow xxs:min-w-[300px] xxs:p-3 xs:px-24 md:p-3 lg:p-4 gap-5 bg-white border rounded-3xl md:text-nowrap lg:text-wrap">
      <img className="" src={img} alt="" />
      <div>
        <p className="font-semibold md:text-[16px] lg:text-xl text-[#232323]">{title}</p>
        <p className="xs:text-wrap md:text-[12px] lg:text-base text-[#718EBF]">{desc}</p>
      </div>
    </div>
  );
};

export default ServicesCard;
