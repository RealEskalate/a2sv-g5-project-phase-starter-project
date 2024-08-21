import Image from "next/image";
import { TopPageCardType } from "@/types/serviceCard";

const TopPageCard = ({title,subTitle,bgColor,svgIcon,}: TopPageCardType) => {
  return (
    <div className="w-11/12 bg-white p-2 sm:p-1 rounded-2xl">
      <div className="flex justify-between items-center">
        <div className={`p-5  rounded-full ${bgColor}`}>
          <Image src={svgIcon} alt=";" width={25} height={20} />
        </div>
        <div className="p-3 w-3/4">
          <div className="text-[#718EBF] text-sm">{title}</div>
          <div className="font-semibold text-sm ">{subTitle}</div>
        </div>
      </div>
    </div>
  );
};

export default TopPageCard;
