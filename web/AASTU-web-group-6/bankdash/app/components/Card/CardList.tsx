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
    <div>
      <div className="ml-10 mb-5 flex items-center h-24 bg-white gap-24 border rounded-3xl">
        <div className={`{icon rounded-full p-3 ${color} ml-4`}>
          <Image src={img} alt="" width={24} height={24} />
        </div>
        <div>
          <p className="font-medium text-base text-[#232323] -ml-20">{title}</p>
          <p className="text-[#718EBF] -ml-20 text-[15px]">{desc}</p>
        </div>
        <div>
          <p className="font-medium text-base text-[#232323]">{colOne}</p>
          <p className="text-[#718EBF] text-[15px]">{descOne}</p>
        </div>
        <div>
          <p className="font-meduim text-base text-[#232323]">{colTwo}</p>
          <p className="text-[#718EBF] text-[15px]">{descTwo}</p>
        </div>
        <div>
          <p className="font-medium text-base text-[#232323]">{colThree}</p>
          <p className="text-[#718EBF] text-[15px]">{descThree}</p>
        </div>
        <button className="text-[#1814F3] text-[15px] w-36 h-8 mr-5 font-semibold">
          {btn}
        </button>
      </div>
    </div>
  );
};

export default CardList;
