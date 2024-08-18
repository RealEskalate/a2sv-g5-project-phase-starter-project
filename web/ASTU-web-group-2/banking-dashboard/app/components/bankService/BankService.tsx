"use client";
import { Inter } from "next/font/google";
import { BankServiceType } from "./BankServiceMobile";
import { useRouter } from "next/navigation";

const inter = Inter({ subsets: ["latin"] });
const BankService = ({
  id,
  name,
  details,
  icon,
  numberOfUsers,
  status,
  type,
}: BankServiceType) => {
  const router = useRouter();
  return (
    <div
      className={`${inter.className} flex justify-between bg-white rounded-[22px] h-fit p-5 items-center`}
    >
      <div className="flex gap-3 items-center">
        <img src={icon} alt="business-loans" />
        <div>
          <div className="font-medium ">{name}</div>
          <div className="font-normal text-[#718EBF]">{details}</div>
        </div>
      </div>

      <div>
        <div className="font-medium">{numberOfUsers}</div>
        <div className="font-normal text-[#718EBF]">users</div>
      </div>
      <div>
        <div className="font-medium">{status}</div>
        <div className="font-normal text-[#718EBF]">status</div>
      </div>
      <div>
        <div className="font-medium">{type}</div>
        <div className="font-normal text-[#718EBF]">type</div>
      </div>

      <button className="h-fit px-5 border rounded-3xl border-[#718EBF] text-[#718EBF] hover:border-[#1814F3] hover:text-[#1814F3]">
        View Details
      </button>
    </div>
  );
};

export default BankService;
