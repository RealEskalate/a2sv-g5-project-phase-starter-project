import React from "react";
import Image from "next/image";
interface Card1Props {

  text: string;
  num: number | string;
  img: string;
}
const Card1 = ({text, num, img}: Card1Props) => {
  return (
    <div className="rounded-2xl p-4 mx-2 shadow-sm bg-white w-[100%] dark:bg-[#050914] dark:border dark:border-[#333B69] my-2">
      <div className=" flex justify-center items-center">
        <div className="flex justify-center items-center">
          <div className="w-20">
            <Image src={img} width={55} height={37} alt="image" />
          </div>
          <div className="pl-4 ">
            <div className="text-[#718EBF] dark:text-[#9faaeb] text-sm">
              {text}
            </div>
            <div className="text-[#232323] dark:text-white font-semibold text-lg">
              {num}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Card1;
