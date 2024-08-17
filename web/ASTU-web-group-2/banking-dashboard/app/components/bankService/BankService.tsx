"use client";
import { Inter } from "next/font/google";
import { BankServiceType } from "./BankServiceMobile";
import { useRouter } from "next/navigation";

const inter = Inter({ subsets: ["latin"] });
const BankService = ({
  category,
  description,
  logoLink,
  details,
  action: { label, link },
}: BankServiceType) => {
  const router = useRouter();
  return (
    <div
      className={`${inter.className} flex justify-between bg-white rounded-[22px] h-fit p-5 items-center`}
    >
      <div className="flex gap-3 items-center">
        <img src={logoLink} alt="business-loans" />
        <div>
          <div className="font-medium ">{category}</div>
          <div className="font-normal text-[#718EBF]">{description}</div>
        </div>
      </div>
      {details.map((data,index) => (
        <div key = {index}>
          <div className="font-medium">{data.title}</div>
          <div className="font-normal text-[#718EBF]">{data.subtitle}</div>
        </div>
      ))}

      <button
        onClick={() => router.push(link)}
        className="h-fit px-5 border rounded-3xl border-[#718EBF] text-[#718EBF] hover:border-[#1814F3] hover:text-[#1814F3]"
      >
        {label}
      </button>
    </div>
  );
};

export default BankService;
