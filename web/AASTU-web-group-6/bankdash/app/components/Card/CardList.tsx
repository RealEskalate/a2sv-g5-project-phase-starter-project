import React from "react";
import Image from "next/image";

interface ListCardProps {
  img: string;
  title: string;
  desc: string;
  colOne: string;
  colTwo: string;
  colThree: string;
  descOne: string;
  descTwo: string;
  descThree: string;
  btn: string;
  color: string;
}

const CardList: React.FC<ListCardProps> = ({
  img,
  title,
  desc,
  colOne,
  colTwo,
  colThree,
  descOne,
  descTwo,
  descThree,
  btn,
  color,
}) => {
  return (
    <div>
      <div className="flex items-center justify-between h-24 bg-white dark:bg-[#232328] gap-5 lg:gap-5 xl:gap-5 border rounded-3xl">
        <div className={`{icon rounded-full p-3 ${color} ml-4`}>
          <img src={img} alt="" />
        </div>
        <div className="w-[15%]">
          <p className="font-medium text-base lg:text-[12px] xl:text-base text-[#232323]">
            {title}
          </p>
          <p className="text-[#718EBF] text-[15px] lg:text-[12px] xl:text-[15px]">
            {desc}
          </p>
        </div>
        <div className="w-[15%]">
          <p className="font-medium text-base lg:text-[12px] xl:text-base text-[#232323]">
            {colOne}
          </p>
          <p className="text-[#718EBF] text-[15px] lg:text-[12px] xl:text-[15px]">
            {descOne}
          </p>
        </div>
        <div className="hidden lg:block w-[25%]">
          <p className="font-meduim text-base lg:text-[12px] xl:text-base text-[#232323] dark:text-gray-300">
            {colTwo}
          </p>
          <p className="text-[#718EBF] text-[15px] lg:text-[12px] xl:text-[15px] dark:text-gray-200">
            {descTwo}
          </p>
        </div>
        <div className="hidden xl:block w-[15%]">
          <p className="font-medium text-base lg:text-[12px] xl:text-base text-[#232323] dark:text-gray-300">
            {colThree}
          </p>
          <p className="text-[#718EBF] text-[15px] lg:text-[12px] xl:text-[15px] dark:text-gray-400">
            {descThree}
          </p>
        </div>
        <button className="text-[#1814F3] text-[15px] lg:text-[11px] xl:text-[15px] w-36 h-8 mr-5 font-semibold">
          {btn}
        </button>
      </div>
    </div>
  );
};

export default CardList;
