import React from "react";
import Image from "next/image";

interface ListCardProps {
  img: string;
  title: string;
  desc: string;
  colOne: string;
  colTwo: string;
  colThree: string;
  descOne: string;
  descTwo: string;
  descThree: string;
  btn: string;
  color: string;
}

const CardList: React.FC<ListCardProps> = ({
  img,
  title,
  desc,
  colOne,
  colTwo,
  colThree,
  descOne,
  descTwo,
  descThree,
  btn,
  color,
}) => {
  return (
    <div className="relative flex w-full items-center justify-between bg-white gap-24 rounded-3xl p-4">
      <div className={`{icon rounded-2xl p-4 ${color}`}>
        <Image src={img} alt="" width={24} height={24} />
      </div>

      <div className="flex flex-col grow">
        <p className="font-medium text-base text-[#232323]">{title}</p>
        <p className="text-[#718EBF] text-[15px]">{desc}</p>
      </div>
      <div className="flex flex-col grow">
        <p className="font-medium text-base text-[#232323]">{title}</p>
        <p className="text-[#718EBF] text-[15px]">{desc}</p>
      </div>

      <button className="text-[#1814F3] text-[15px] w-36 font-semibold">
        {btn}
      </button>
    </div>
  );
};

export default CardList;
