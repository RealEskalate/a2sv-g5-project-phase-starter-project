import React from "react";
import Image from "next/image";
import { useUser } from "@/contexts/UserContext";

export default function Card(props: any) {
  const { isDarkMode } = useUser();

  return (
    <div className={`h-[80px] w-[280px] md:w-[380px] ${isDarkMode ? "bg-gray-800 shadow-md shadow-blue-900" : "bg-white shadow-md"} rounded-3xl m-2 flex justify-around`}>
      <Image src={props.icon} alt="" width={1} height={1} className={`my-auto h-[55%] w-[15%] p-3 rounded-full 
      ${ isDarkMode? "text-indigo-400 bg-[#718EBF33]" : "text-indigo-700 bg-[#577bb813]"}`} />
      <div className="pl-3 my-auto w-2/3">
        <h1 className={`font-[400] text-[12px] md:text-[15px] ${ isDarkMode ? "text-white" : "text-[#718EBF]"}`}>
          {props.name }
        </h1>
        <p className={`font-[500] text-[12px] md:text-[14px] md:font-[600] ${isDarkMode ? "text-white" : "text-gray-700"}`}>
          {props.description||0}
        </p>
      </div>
    </div>
  );
}
