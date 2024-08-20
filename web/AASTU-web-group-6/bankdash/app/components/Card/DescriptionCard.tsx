import React from "react";

interface DescriptionCardProps {
  img: string;
  title: string;
  desc: string;
  colOne: string;
  colTwo: string;
  colThree: string;
  descOne: number;
  descTwo: string;
  descThree: string;
  btn: string;
  color: string;
}

const DescriptionCard: React.FC<DescriptionCardProps> = ({
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
  console.log(img);
  return (
    <div className="lg:ml-10 mb-5 flex items-center justify-between h-24 bg-white gap-5 border rounded-3xl lg:w-[1300px]">
      <div className={`{icon rounded-full p-3 ${color} ml-4`}>
        <img src={img} alt="" />
      </div>
      <div>
        <p className="font-semibold text-base text-[#232323] lg:-ml-36">
          {title}
        </p>
        <p className="text-[#718EBF] lg:-ml-36 text-[15px]">{desc}</p>
      </div>
      <div className="hidden lg:block">
        <p className="font-semibold text-base text-[#232323]">{colOne}</p>
        <p className="text-[#718EBF] text-[15px]">{descOne}</p>
      </div>
      <div className="hidden lg:block">
        <p className="font-semibold text-base text-[#232323]">{colTwo}</p>
        <p className="text-[#718EBF] text-[15px]">{descTwo}</p>
      </div>
      <div className="hidden lg:block">
        <p className="font-semibold text-base text-[#232323]">{colThree}</p>
        <p className="text-[#718EBF] text-[15px]">{descThree}</p>
      </div>
      <button className="text-[#1814F3] font-semibold text-[15px] lg:border lg:border-[#718EBF] lg:w-36 h-8 rounded-[50px] mr-8">
        {btn}
      </button>
    </div>
  );
};

export default DescriptionCard;
