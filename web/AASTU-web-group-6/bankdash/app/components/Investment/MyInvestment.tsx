import React from "react";
interface props {
  color: string;
  title: string;
  icon: string;
  amount: string;
  titleRe: string;
  returnValue: string;
  returnRe: string;
  Envestment: string;
}
const MyInvestment = ({
  color,
  title,
  icon,
  amount,
  titleRe,
  returnValue,
  returnRe,
  Envestment,
}: props) => {
  const textColor = returnValue.startsWith('-') ? 'text-[#FE5C73]' : 'text-[#16DBAA]';
  return (
    <div className="border border-solid rounded-[20px] flex items-center justify-around py-4 mb-5 bg-white shadow-lg border-gray-300">
    <div className="flex items-center gap-2 w-[40%]">
      <div
        className="border border-solid rounded-2xl  w-[75px] h-[75px] flex justify-center items-center"
        style={{ borderColor: color, backgroundColor: color }}
      >
        <img src={icon} />
      </div>
      <div className="flex flex-col ">
        <p className="font-inter text-base font-medium text-[#232323]">{title}</p>
        <p className="font-inter text-[15px] font-normal text-[#718EBF]">{titleRe}</p>
      </div>
      </div>
      <div className=" hidden lg:block  flex-col w-[20%]">
        <p className="font-inter text-base font-medium text-[#232323]">{amount}</p>
        <p className="font-inter text-[15px] font-normal text-[#718EBF]">{Envestment}</p>
      </div>
      <div className="flex flex-col w-[20%]">
        <p className={` font-medium text-base font-inter ${textColor}`}>{returnValue}</p>
        <p className="hidden lg:block font-inter text-[15px] font-normal text-[#718EBF]">{returnRe}</p>
      </div>

    </div>
  );
};

export default MyInvestment;
