import React from "react";
import Image from "next/image";
import { useUser } from "@/contexts/UserContext";

export default function InvestmentCard(props: any) {
  const { isDarkMode } = useUser();

  return (
    <div
      className={`flex rounded-3xl mb-2 w-[90%] md:w-full h-[100px] mx-auto my-2 p-2 ${
        isDarkMode ? "bg-gray-800" : "bg-white"
      } shadow-md`}
    >
      <div className="my-auto w-[15%] md:w-[10%] h-auto">
        <Image
          src={props.icon}
          alt=""
          width={1}
          height={1}
          className="my-auto size-[25px] md:size-[50px] rounded-full"
        />
      </div>
      <div className="my-auto w-[50%] md:w-[30%]">
        <h1 className={`${isDarkMode ? "text-gray-200" : "text-black"}`}>
          {props.name}
        </h1>
        <div className="flex flex-row gap-1 flex-wrap">
          {props.type.map((item: any) => (
            <p
              key={item}
              className={`font-[400] text-[13px] md:text-[14px] text-wrap w-[100px] ml-0 ${
                isDarkMode ? "text-gray-400" : "text-blue-900"
              }`}
            >
              {item}
            </p>
          ))}
        </div>
      </div>
      <div className="my-auto w-[30%] hidden lg:table-cell">
        <h1
          className={`${
            isDarkMode ? "text-gray-200" : "text-black"
          } text-[13px] md:text-[14px]`}
        >
          {props.investmentValue}
        </h1>
        <p
          className={`${
            isDarkMode ? "text-gray-400" : "text-blue-800"
          } text-[12px] md:text-[14px]`}
        >
          Investment Value
        </p>
      </div>
      <div className="my-auto w-[40%] md:w-[30%] pl-16 md:pl-0">
        <h1
          className={`${isDarkMode ? "text-emerald-500" : "text-emerald-700"} text-[14px]`}
        >
          {props.returnValue}
        </h1>
        <p
          className={`${
            isDarkMode ? "text-gray-400" : "text-blue-800"
          } hidden md:table-cell my-auto text-[14px]`}
        >
          Return Value
        </p>
      </div>
    </div>
  );
}
