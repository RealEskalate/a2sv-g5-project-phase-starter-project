
import { TransactionContent } from "@/types";
import Image from "next/image";
import { Dateformat } from "../transactions/component/utils";
import { useUser } from "@/contexts/UserContext";
export const Transaction = ({
  date,
  amount,
  type,
  description,
}: TransactionContent) => {
  const { isDarkMode } = useUser();
  return (
    <div
      className={`flex justify-between items-center space-x-3 ${
        isDarkMode ? "bg-gray-800 text-white" : "bg-white text-black"
      }`}
    >
      <div className="flex items-center space-x-3 min-w-0">
        <div
          className={`flex items-center justify-center rounded-full w-8 h-8 flex-shrink-0 ${
            type === "shopping"
              ? isDarkMode
                ? "bg-yellow-700"
                : "bg-yellow-100"
              : type === "transfer"
              ? isDarkMode
                ? "bg-indigo-700"
                : "bg-indigo-100"
              : isDarkMode
              ? "bg-green-700"
              : "bg-green-100"
          }`}
        >
          <Image
            src={
              type === "shopping"
                ? "/icons/wallet.png"
                : type === "transfer"
                ? "/icons/paypal.png"
                : "/icons/dollarSign.png"
            }
            alt={`transaction icon`}
            className="object-cover object-center"
            width={16}
            height={16}
          />
        </div>
        <div className="flex flex-col min-w-0">
          <p className="font-inter text-[14px] font-medium truncate">
            {description}
          </p>
          <p className="font-inter text-[12px] text-indigo-400 truncate">
            {Dateformat(date)}
          </p>
        </div>
      </div>
      <div className="flex justify-end items-center flex-shrink-0">
        <p className="font-inter text-[12px] md:text-[14px] text-green-600 font-semibold">
          ${Math.round(amount)}
        </p>
      </div>
    </div>
  );
};
