import React from "react";
import Image from "next/image";
import { useUser } from "@/contexts/UserContext";

export default function Card(props: any) {
  const { isDarkMode } = useUser();

  return (
    <div className={`h-[90px] w-[280px] md:w-[380px] ${isDarkMode ? "bg-gray-800" : "bg-white"} rounded-3xl m-2 flex justify-around`}>
      <Image src={props.icon} alt="" width={1} height={1} className={`my-auto h-[60%] w-[20%] p-3 md:p-4 size-[50px] rounded-full 
      ${ isDarkMode? "text-indigo-300 bg-[#718EBF33]" : "text-indigo-700 bg-[#577bb813]"}`} />
      <div className="pl-3 my-auto w-2/3">
        <h1 className={`font-[400] text-[12px] md:text-[17px] ${ isDarkMode ? "text-gray-300" : "text-[#718EBF]"}`}>
          {props.name }
        </h1>
        <p className={`font-[500] text-[12px] md:text-[14px] md:font-[600] ${isDarkMode ? "text-gray-400" : "text-gray-700"}`}>
          {props.description}
        </p>
      </div>
    </div>
  );
}
