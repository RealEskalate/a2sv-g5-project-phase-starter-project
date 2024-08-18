import Link from "next/link";
import Image from "next/image";
import { ServiceType } from "@/types/serviceCard";
import { menuItems } from "@/../../public/Icons";


export default function ServiceCard({
    name,
    details,
    numberOfUsers,
    status,
    type,
    icon,
}: ServiceType) {
  return (
    <div className="body flex  md:w-auto w-auto h-auto p-2 border-[1px] rounded-[10px] m-2 bg-white">
      <div className="flex items-center rounded-2xl px-5 bg-[#FFE0EB]">
        <Image width={18} height={18} src={menuItems[5].icon} alt="aastu"/>
      </div>
      <div className="right w-full flex justify-between items-center p-2">
        <div className="md:w-1/4">
          <div className="font-normal ">{name}</div>
          <div className="font-normal text-xs text-[#718EBF]">{details}</div>
        </div>
        <div className="hidden md:block md:w-1/6">
          <div className="font-medium text-sm md:text-[12px]">Status</div>
          <div className="font-normal text-xs text-[#718EBF]">{status}</div>
        </div>
        <div className="hidden md:block md:w-1/6">
          <div className="font-medium text-sm md:text-[12px]"> Type</div>
          <div className="font-normal text-xs text-[#718EBF]">{type}</div>
        </div>
        <div className="hidden md:block md:w-1/6">
          <div className="font-medium text-sm md:text-[12px]">Number of users</div>
          <div className="font-normal text-xs text-[#718EBF]">{numberOfUsers}</div>
        </div>
        {/* <div className="md:px-4 md:py-1 md:border  md:border-[#718EBF] md:rounded-full hover:border-[#1814F3] text-center"> */}
          <Link
            href={"/services"}
            className="md:px-4 md:py-2 md:border  md:border-[#718EBF] md:rounded-full hover:border-[#1814F3] text-center font-normal text-[11px] text-[#1814F3] md:text-[#718EBF] hover:text-[#1814F3]"
          >
            View Details
          </Link>
        </div>
      </div>
    // </div>
  );
}
