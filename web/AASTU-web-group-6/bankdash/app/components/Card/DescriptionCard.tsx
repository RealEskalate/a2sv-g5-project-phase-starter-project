import React from "react";
import Image from "next/image";

interface DescriptionCardProps {
  img: string;
  title: string;
  desc: string;
  colOne: string;
  colTwo: string;
  colThree: string;
  descOne: number;
  descTwo: string;
  descThree: string;
  btn: string;
  color: string;
}

const DescriptionCard: React.FC<DescriptionCardProps> = ({
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
  console.log(img);
  return (
    <div className="mb-5 flex items-center text-nowrap grow justify-between p-5 bg-white dark:bg-[#232328] gap-5 rounded-3xl xs:w-full md:w-full lg:w-full">
      <div className={`{icon rounded-[12px]  p-3 ${color} ml-4`}>
        <img src={img} alt="" className="min-h-6 min-w-6" />
      </div>
      <div className="flex flex-col gap-1">
        <p className="font-semibold text-base text-[#232323] dark:text-gray-300 ">
          {title}
        </p>
        <p className="text-[#718EBF] text-[15px] dark:text-gray-400">{desc}</p>
      </div>
      <div className=" flex flex-col gap-1 xxs:hidden md:flex">
        <p className="font-semibold text-base text-[#232323] dark:text-gray-300">
          {colOne}
        </p>
        <p className="text-[#718EBF] text-[15px] dark:text-gray-400">
          {descOne}
        </p>
      </div>
      <div className="flex flex-col gap-1 xxs:hidden md:flex">
        <p className="font-semibold text-base text-[#232323] dark:text-gray-300">
          {colTwo}
        </p>
        <p className="text-[#718EBF] text-[15px] dark:text-gray-400">
          {descTwo}
        </p>
      </div>
      <div className="flex flex-col gap-1 xxs:hidden md:flex">
        <p className="font-semibold text-base text-[#232323] dark:text-gray-300">
          {colThree}
        </p>
        <p className="text-[#718EBF] text-[15px] dark:text-gray-400">
          {descThree}
        </p>
      </div>
      <button className="text-[#1814F3] font-medium text-[15px] lg:border lg:border-[#1814F3] px-6 py-2 rounded-[50px] mr-8">
        {btn}
      </button>
    </div>
  );
};

export default DescriptionCard;
