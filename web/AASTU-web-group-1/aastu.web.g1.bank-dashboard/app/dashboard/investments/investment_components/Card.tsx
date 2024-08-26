import React from "react";
import Image from "next/image";
import { useUser } from "@/contexts/UserContext";

export default function Card(props: any) {
  const { isDarkMode } = useUser();

  return (
    <div className={`flex px-2 pl-5 h-[90px] w-[90%] md:w-[400px] rounded-2xl mx-auto md:mx-0 ${isDarkMode ? "bg-gray-800 shadow-xl shadow-slate-700" : "bg-white shadow-lg"}`}>
      <Image src={props.icon} alt="" width={1} height={1} className={`my-auto h-[50%] w-[20%] p-3 size-[50px] rounded-full ${
          isDarkMode ? "text-indigo-300 bg-[#718EBF33]" : "text-indigo-700 bg-[#577bb813]"}`}/>
      <div className="pl-4 my-auto w-[70%]">
        <h1 className={`font-[400] text-[14px] md:text-[16px] ${isDarkMode ? "text-gray-300" : "text-[#718EBF]"}`}>
          {props.name}
        </h1>
        <p className={`font-[500] text-[12px] md:text-[14px] md:font-[600] pl-1 ${isDarkMode ? "text-gray-400" : "text-gray-700"}`}>
          {props.amount||0}
        </p>
      </div>
    </div>
  );
}
