"use client";
import { Inter } from "next/font/google";
import { useRouter } from "next/navigation";
import { Router } from "next/router";

const inter = Inter({ subsets: ["latin"] });

export interface BankServiceType {
  id: string;
  name: string;
  details: string;
  icon: string;
  numberOfUsers: number;
  status: string;
  type: string;
}

const BankServiceMobile = ({ id, name, details, icon }: BankServiceType) => {
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
      <button className="h-fit px-5 border rounded-3xl border-[#718EBF] text-[#718EBF] hover:border-[#1814F3] hover:text-[#1814F3]">
        View Details
      </button>
    </div>
  );
};

export default BankServiceMobile;
