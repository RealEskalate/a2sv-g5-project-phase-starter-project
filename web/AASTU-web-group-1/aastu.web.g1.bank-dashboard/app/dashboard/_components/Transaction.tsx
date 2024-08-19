import React from "react";
import Image from "next/image";
import { TransactionData } from "@/types";

// interface Props {
//   image: string;
//   transactionType: string;
//   date: string;
//   amount: string;
//   color: string;
// }
function Dateformat(dateString: string) {
  const date = new Date(dateString);
  const month = (date.getMonth() + 1).toString().padStart(2, "0");
  const year = date.getFullYear().toString().slice(-2);
  return `${year}-${month}-${date.getDate()}`;
}

export const Transaction = ({
 date,
  amount,
  type,
  description

}: TransactionData) => {

  
  return (
    <div className="flex justify-between">
      <div className="flex space-x-2">
        <div
          className={`inline-flex items-center justify-center 
                      ${
                      type === "shopping" ? "bg-yellow-100" : type === "transfer"?"bg-indigo-100" : "bg-green-100"}
                     rounded-full w-[35px] h-[35px]`}
        >
          <Image
            src={type === "shopping" ? "/icons/wallet.png" :
                 type === "transfer" ? "/icons/paypal.png" :
                 "/icons/dollarSign.png"}
            alt={`transation icon`}
            className="object-cover object-center"
            width={15}
            height={15}
          />
        </div>

        <div>
          <p className={`font-inter text-[14px] font-medium`}>{description}</p>
          <p className={`font-inter text-[12px] text-indigo-400 font-normal`}>
            {Dateformat(date)}
          </p>
        </div>
      </div>
      <div>
        <p
          className={`font-inter text-[11px] text-green-600`}
          style={{ fontWeight: 500 }}
        >
          ${Math.round(amount)}
        </p>
      </div>
    </div>
  );
};
