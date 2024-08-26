import Link from "next/link";
import React, { ReactNode } from "react";

interface Props {
  icon: ReactNode;
  data: Array<[string, string]>;
}
const CardSetting = ({ icon, data }: Props) => {
  return (
    <div className="flex gap-3 items-center p-3">
      <div className="w-10">{icon}</div>
      <div className="flex gap-8">
        {data.map((data, index) => {
          return (
            <div className={`${index > 0 && "hidden"} lg:block`} key={index}>
              <p className="text-[#232323] font-medium text-sm lg:text-base dark:text-[#d2d6ef]">
                {data[0]}
              </p>
              <p className="text-[#8297c0] text-xs">{data[1]}</p>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default CardSetting;
