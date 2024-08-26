import React from "react";
import { colors } from "@/constants";
import Image from "next/image";

const CardSettingLine = ({
  icon,
  title,
  description,
  background,
}: {
  icon: string;
  title: string;
  description: string;
  background: string;
}) => {
  return (
    <div className="flex p-0.5 justify-between">
      <div className=" flex gap-2 md:gap-3 lg:gap-4">
        <div
          className={`rounded-full flex items-center justify-center p-4 h-20 w-20`}
        >
          <Image
            src={icon}
            alt="go"
            width={120}
            height={120}
            className={`${background} p-2 rounded-xl items-center`}
          />
        </div>
        <div className="my-2">
          <p className="md:text-[16px] sm:text-[13px] font-medium pb-1">
            {title}
          </p>
          <p
            className={`${colors.textgray} md:text-[13px] sm:text-[11px] text-start`}
          >
            {description}
          </p>
        </div>
      </div>
    </div>
  );
};

export default CardSettingLine;
