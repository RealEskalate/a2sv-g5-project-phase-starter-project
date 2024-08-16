'use client'
import React, { useState } from "react";
import { colors } from "@/constants/index";
import Image from "next/image";

const dummyData = [
  { name: "Natnael Worku", position: "CEO", imageSrc: "/Images/pp.jpg" },
  { name: "John Doe", position: "CTO", imageSrc: "/Images/pp.jpg" },
  { name: "Jane Smith", position: "CFO", imageSrc: "/Images/pp.jpg" },
  { name: "Michael Johnson", position: "COO", imageSrc: "/Images/pp.jpg" },
  { name: "Emily Davis", position: "CMO", imageSrc: "/Images/pp.jpg" },
];

const QuickTransfer = () => {
  const [currentIndex, setCurrentIndex] = useState(0);

  const handleNext = () => {
    if (currentIndex < dummyData.length - 3) {
      setCurrentIndex(currentIndex + 1);
    }
  };

  const handlePrev = () => {
    if (currentIndex > 0) {
      setCurrentIndex(currentIndex - 1);
    }
  };

  return (
    <div className="flex flex-col items-start gap-3 w-[100%]">
      <h1 className={`font-semibold text-[18px] ${colors.navbartext}`}>
        Quick Transfer
      </h1>

      <div className="flex justify-between w-[100%] px-2">
        <div
          onClick={handlePrev}
          className={`w-[50px] h-[50px] rounded-full ${colors.white} flex justify-center items-center shadow-lg cursor-pointer`}
        >
          <span className="text-gray-500">&lt;</span>
        </div>

        <div className="flex overflow-hidden gap-2">
  {dummyData.slice(currentIndex, currentIndex + 3).map((item, index) => (
    <div key={index} className="flex flex-col gap-2 flex-1">
      <Image
        src={item.imageSrc}
        width={50}
        height={50}
        className="rounded-full"
        alt={item.name}
      />

      <div className="flex flex-col">
        <div
          className={`font-normal text-[12px] ${colors.textblack} whitespace-normal`}
        >
          {item.name}
        </div>
        <div
          className={`font-normal text-[12px] ${colors.textgray} whitespace-normal`}
        >
          {item.position}
        </div>
      </div>
    </div>
  ))}
</div>


        <div
          onClick={handleNext}
          className={`w-[50px] h-[50px] rounded-full ${colors.white} flex justify-center items-center shadow-lg cursor-pointer`}
        >
          <span className="text-gray-500">&gt;</span>
        </div>
      </div>

      <div className="grid grid-cols-3 w-[100%] justify-center items-center px-2">
        <p
          className={`text-center px-2 text-nowrap text-[12px] font-normal ${colors.textgray}`}
        >
          Write amount
        </p>
        <div
          className={`col-span-2 flex w-[100%] ${colors.lightblue} rounded-3xl`}
        >
          <input
            type="text"
            placeholder="535.35"
            className={`w-[50%] border-1 rounded-3xl py-2 border-black ${colors.lightblue} focus:${colors.lightblue} focus:outline-none px-3 focus:border-none`}
          />
          <button
            className={`${colors.blue} text-white w-[60%] py-2 rounded-3xl`}
          >
            Send ✈
          </button>
        </div>
      </div>
    </div>
  );
};

export default QuickTransfer;
