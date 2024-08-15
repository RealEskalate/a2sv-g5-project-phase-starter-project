import { title } from "process";
import React from "react";
interface props {
  title: string;
  date: string;
  type: string;
  account_no: string;
  status: string;
  amount: string;
  color:string;
  icon:string
}
const Last_trans = ({
  title,
  date,
  type,
  account_no,
  status,
  amount,
  color,
  icon
}: props) => {
  const textColor = amount.startsWith('-') ? 'text-[#FE5C73]' : 'text-[#16DBAA]';
  return (
    <div className="flex gap-10 items-center mb-5 pl-4 ">
      <div className="border border-solid rounded-2xl  w-[55px] h-[55px] flex justify-center items-center" style={{borderColor: color , backgroundColor:color}}>
    <img src={icon}/>
      </div>
      <div className="flex flex-col  w-[25%]">
        <p className="font-inter font-medium text-base text-[#232323] ">{title}</p>
        <p className="font-inter font-normal text-[15px] text-[#718EBF]">{date}</p>
      </div>
      <p className="font-inter font-normal text-base text-[#718EBF] w-[15%]">{type}</p>
      <p className="font-inter font-normal text-base text-[#718EBF] w-[10%]">{account_no}</p>
      <p className="font-inter font-normal text-base text-[#718EBF] w-[10%]"> {status}</p>
      <p className={`font-inter font-medium text-base ${textColor}`}>{amount}</p>
    </div>
  );
};

export default Last_trans;
