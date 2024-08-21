import React from "react";

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
    <div className="mb-5 flex items-center grow justify-between h-24 bg-white dark:bg-[#232328] gap-5 border rounded-3xl xs:w-full md:w-full lg:w-full">
      <div className={`{icon rounded-full p-3 ${color} ml-4`}>
        <img src={img} alt="" />
      </div>
      <div className="flex flex-col">
        <p className="font-semibold text-base text-[#232323] dark:text-gray-300 ">
          {title}
        </p>
        <p className="text-[#718EBF] text-[15px] dark:text-gray-400">{desc}</p>
      </div>
      <div className=" flex flex-col hidden lg:flex">
        <p className="font-semibold text-base text-[#232323] dark:text-gray-300">
          {colOne}
        </p>
        <p className="text-[#718EBF] text-[15px] dark:text-gray-400">
          {descOne}
        </p>
      </div>
      <div className="flex flex-col hidden lg:flex">
        <p className="font-semibold text-base text-[#232323] dark:text-gray-300">
          {colTwo}
        </p>
        <p className="text-[#718EBF] text-[15px] dark:text-gray-400">
          {descTwo}
        </p>
      </div>
      <div className="flex flex-col hidden lg:flex">
        <p className="font-semibold text-base text-[#232323] dark:text-gray-300">
          {colThree}
        </p>
        <p className="text-[#718EBF] text-[15px] dark:text-gray-400">
          {descThree}
        </p>
      </div>
      <button className="text-[#1814F3] font-semibold text-[15px] lg:border lg:border-[#718EBF] lg:w-36 h-8 rounded-[50px] mr-8">
        {btn}
      </button>
    </div>
  );
};

export default DescriptionCard;
