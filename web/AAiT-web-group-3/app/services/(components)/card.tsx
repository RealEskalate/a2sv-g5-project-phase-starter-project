import React from "react";
import { IconType } from "react-icons";

interface Cardprops {
  icon: IconType;
  title: string;
  sub_title: string;
  icon_bg?: string;
  icon_color?: string;
}

const card = ({ icon: Icon, title, sub_title,icon_bg,icon_color }: Cardprops) => {
  return (
    <div className=" inline-flex   gap-4  items-center rounded-[25px] bg-white px-16 py-7">
      <div className="rounded-full w-[45px] h-[45px] md:w-[50] md:h-[50]  flex items-center justify-center whitespace-nowrap" style={{backgroundColor:icon_bg}}>
        <span style={{color:icon_color}}>
          <Icon className="w-[20px] h-[20px] " />
        </span>
      </div>
      <div>
        <p className="font-inter font-semibold text-[20px] leading-[24.2px] text-[#232323] whitespace-nowrap">
          {title}
        </p>
        <p className="font-normal text-[#718EBF] text-[16px] leading-[19.36px] mt-2 whitespace-nowrap ">
          {sub_title}
        </p>
      </div>
    </div>
  );
};

export default card;
