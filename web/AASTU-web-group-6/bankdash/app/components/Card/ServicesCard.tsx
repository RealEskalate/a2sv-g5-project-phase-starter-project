import React from "react";

interface ServicesCardProps {
  img: string;
  title: string;
  desc: string;
}

const ServicesCard: React.FC<ServicesCardProps> = ({ img, title, desc }) => {
  return (
    <div className="flex items-center justify-center h-28 gap-5 bg-white border px-3 rounded-3xl ">
      <img className="p-2" src={img} alt="" />
      <div>
        <p className="font-semibold text-xl text-[#232323]">{title}</p>
        <p className="text-base text-[#718EBF]">{desc}</p>
      </div>
    </div>
  );
};

export default ServicesCard;
