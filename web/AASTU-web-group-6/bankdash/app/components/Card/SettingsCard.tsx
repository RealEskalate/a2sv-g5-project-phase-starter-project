import React from "react";
import Image from "next/image";

interface ListCardProps {
  img: string;
  title: string;
  desc: string;
  bg: string;
}

const SettingsCard: React.FC<ListCardProps> = ({ img, title, desc, bg }) => {
  return (
    <div className="flex items-center w-80 bg-white p-2">
      <div
        className={`flex item-center justify-center box w-14 h-14 rounded-xl ${bg}`}
      >
        <Image src={img} width={24} height={24} alt="" />
      </div>
      <div className="ml-5">
        <p className="font-medium text-base text-[#232323] pb-1">{title}</p>
        <p className="text-[#718EBF] text-[15px]">{desc}</p>
      </div>
    </div>
  );
};

export default SettingsCard;
