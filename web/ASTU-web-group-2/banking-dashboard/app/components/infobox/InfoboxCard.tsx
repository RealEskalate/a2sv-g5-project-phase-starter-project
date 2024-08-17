import Image from "next/image";
import React from "react";

const InfoboxCard = ({
  name,
  icon,
  value,
  classNameText1 = "text-[#718EBF] font-normal text-sm max-lg:text-[12px] max-md:text-sm",
  classNameText2="text-[#232323] font-semibold text-lg max-lg:text-[12px] max-md:text-lg",
}: {
  name: string;
  icon: string;
  value: string;
  classNameText1?: string;
  classNameText2?: string;
}) => {
  return (
    <div className="flex items-center rounded-[20px] bg-white h-[90px] shadow-md">
      <div className="flex items-center p-2 gap-3 w-full">
        <div className="flex-shrink-0">
          <img
            src={icon}
            alt="my-balance"
            className="w-[50px] h-[50px] max-lg:w-[30px] max-lg:h-[30px]  max-md:w-[50px] max-md:h-[50px]"
            // width={50}
            // height={50}
          />
        </div>
        <div className="flex flex-col justify-center ">
          <p className={classNameText1}>
            {name}
          </p>
          <p className={classNameText2}>
            {value}
          </p>
        </div>
      </div>
    </div>
  );
};

export default InfoboxCard;
