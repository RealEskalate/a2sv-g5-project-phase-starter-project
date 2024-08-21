import Link from "next/link";
import React, { ReactNode } from "react";

interface Props {
  icon: ReactNode;
  data: Array<[string, string]>;
  linkUrl: string;
}
const CreditCard = ({ icon, data, linkUrl }: Props) => {
  return (
    <div className="flex justify-around items-center bg-white p-3 lg:p-4 rounded-2xl shadow-sm">
      {icon}
      <div className="flex gap-6">
        {data.map((data, index) => {
          return (
            <div className={`${index > 1 && "hidden"}  lg:block`} key={index}>
              <p className="text-[#232323] font-medium text-sm w-20 lg:text-base">
                {data[0]}
              </p>
              <p className="text-[#8297c0] text-xs lg:text-sm">{data[1]}</p>
            </div>
          );
        })}
      </div>
      <Link
        href={linkUrl}
        className="text-[#1814F3] font-medium text-xs lg:text-base"
      >
        View Detail
      </Link>
    </div>
  );
};

export default CreditCard;
