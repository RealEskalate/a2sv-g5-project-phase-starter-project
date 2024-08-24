import Image from "next/image";
import React from "react";

interface Props {
  iconUrl: string;
}
const BankServiceCard = ({ iconUrl }: Props) => {
  return (
    <div className="min-w-[320px] bg-[#F5F7FA]">
      <div className="bg-white rounded-2xl h-[90px] flex justify-start items-center p-10">
        <Image width={70} height={70} src={iconUrl} alt="" className="pr-5" />
        <div>
          <p className="font-semibold text-base sm:text-xl sm:mb-auto mb-0.5">
            Business loans
          </p>
          <p className="font-light text-[#718EBF] text-xs sm:text-sm">
            it is a long established
          </p>
        </div>
      </div>
    </div>
  );
};

export default BankServiceCard;
