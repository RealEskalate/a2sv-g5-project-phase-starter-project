import React from "react";

interface ServicesCardProps {
  img: string;
  title: string;
  desc: string;
}

const ServicesCard: React.FC<ServicesCardProps> = ({ img, title, desc }) => {
  return (
    <div className="flex items-center text-nowrap justify-center grow xxs:min-w-[300px] xxs:p-3 xs:px-24 md:p-3 lg:min-w-[30%] lg:p-4 gap-5 bg-white dark:bg-[#232328] rounded-3xl md:text-nowrap lg:text-wrap">
      <img className="" src={img} alt="" />
      <div>
        <p className="font-semibold text-nowrap md:text-[16px] lg:text-xl text-[#232323] dark:text-gray-300">
          {title}
        </p>
        <p className="text-nowrap md:text-[12px] lg:text-base text-[#718EBF] dark:text-gray-4f00">
          {desc}
        </p>
      </div>
    </div>
  );
};

export default ServicesCard;
