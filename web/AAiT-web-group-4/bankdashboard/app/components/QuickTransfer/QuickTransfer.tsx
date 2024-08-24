import React from "react";
import person from "@/public/person.svg";
import Image from "next/image";

const QuickTransfer = () => {
  return (
    <div className="px-6 py-9 bg-white rounded-lg w-full max-mobile:h-[276px] tablet:h-[220px] desktop:h-[276px]">
      <div className="flex items-center gap-6 h-1/2">
        <div className="flex flex-col items-center justify-evenly h-full">
          <Image src={person} alt="person" />
          <span className="flex flex-col items-center">
            <p>Livia bator</p>
            <p>CEO</p>
          </span>
        </div>

        <div className="flex flex-col items-center justify-evenly h-full">
          <Image src={person} alt="person" />
          <span className="flex flex-col items-center">
            <p>Livia bator</p>
            <p>CEO</p>
          </span>
        </div>

        <div className="flex flex-col items-center justify-evenly h-full">
          <Image src={person} alt="person" />
          <span className="flex flex-col items-center">
            <p>Livia bator</p>
            <p>CEO</p>
          </span>
        </div>

        <span className="flex justify-center items-center rounded-full bg-white  shadow-custom-shadow h-10 w-10" >
        <svg
          width="9"
          height="15"
          viewBox="0 0 9 15"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path d="M1 1L7.5 7.5L1 14" stroke="#718EBF" stroke-width="1.5" />
        </svg>

        </span>
        
      </div>
      <div className="flex">
        <p>Write amount</p>
        <input className="h-full bg-Very-Light-White" type="text" />
      </div>
    </div>
  );
};

export default QuickTransfer;
