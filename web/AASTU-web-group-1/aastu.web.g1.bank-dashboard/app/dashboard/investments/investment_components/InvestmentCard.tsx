import React from "react";
import Image from "next/image";
import { useUser } from "@/contexts/UserContext";

export default function InvestmentCard(props: any) {
  const { isDarkMode } = useUser();

  return (
    <div
      className={`flex rounded-3xl mb-2 w-[90%] md:w-full h-[100px] mx-auto p-2 ${
        isDarkMode ? "bg-gray-800" : "bg-white"
      }`}
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
              className={`font-[400] text-[14px] md:text-[16px] text-wrap w-[100px] ml-0 ${
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
          } text-[16px] md:text-[20px]`}
        >
          {props.investmentValue}
        </h1>
        <p
          className={`${
            isDarkMode ? "text-gray-400" : "text-blue-800"
          } text-[10px] md:text-[16px]`}
        >
          Investment Value
        </p>
      </div>
      <div className="my-auto w-[40%] md:w-[30%] pl-16 md:pl-0">
        <h1
          className={`${isDarkMode ? "text-emerald-500" : "text-emerald-700"}`}
        >
          {props.returnValue}
        </h1>
        <p
          className={`${
            isDarkMode ? "text-gray-400" : "text-blue-800"
          } hidden md:table-cell my-auto`}
        >
          Return Value
        </p>
      </div>
    </div>
  );
}
