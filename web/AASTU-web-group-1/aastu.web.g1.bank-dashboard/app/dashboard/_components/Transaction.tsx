import { TransactionData } from "@/types";
import Image from "next/image";
import { Dateformat } from "../transactions/component/utils";

export const Transaction = ({
  date,
  amount,
  type,
  description,
}: TransactionData) => {
  return (
    <div className="flex justify-between items-center space-x-3">
      <div className="flex items-center space-x-3 min-w-0">
        <div
          className={`flex items-center justify-center ${
            type === "shopping"
              ? "bg-yellow-100"
              : type === "transfer"
              ? "bg-indigo-100"
              : "bg-green-100"
          } rounded-full w-8 h-8 flex-shrink-0`}
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
