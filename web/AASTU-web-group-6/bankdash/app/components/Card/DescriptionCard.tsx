import React from "react";

interface DescriptionCardProps {
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
  return (
    <div className="ml-10 mb-5 flex items-center justify-between h-24 bg-white gap-24 border rounded-3xl">
      <div className={`{icon rounded-full p-3 ${color} ml-4`}>
        <img src={img} alt="" />
      </div>
      <div>
        <p className="font-semibold text-base text-[#232323] -ml-32">{title}</p>
        <p className="text-[#718EBF] -ml-32 text-[15px]">{desc}</p>
      </div>
      <div>
        <p className="font-semibold text-base text-[#232323]">{colOne}</p>
        <p className="text-[#718EBF] text-[15px]">{descOne}</p>
      </div>
      <div>
        <p className="font-semibold text-base text-[#232323]">{colTwo}</p>
        <p className="text-[#718EBF] text-[15px]">{descTwo}</p>
      </div>
      <div>
        <p className="font-semibold text-base text-[#232323]">{colThree}</p>
        <p className="text-[#718EBF] text-[15px]">{descThree}</p>
      </div>
      <button className="text-[#718EBF] text-[15px] border border-[#718EBF] w-36 h-8 rounded-[50px] mr-8">
        {btn}
      </button>
    </div>
  );
};

export default DescriptionCard;
