import React from "react";

interface ListCardProps {
  img: string;
  title: string;
  desc: string;
}

const SettingsCard: React.FC<ListCardProps> = ({ img, title, desc }) => {
  return (
    <div>
      <div>
        <div className="flex items-center w-80 bg-white pb-3">
          <img src={img} alt="" />
          <div className="ml-5">
            <p className="font-medium text-base text-[#232323] pb-1">{title}</p>
            <p className="text-[#718EBF] text-[15px]">{desc}</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default SettingsCard;
