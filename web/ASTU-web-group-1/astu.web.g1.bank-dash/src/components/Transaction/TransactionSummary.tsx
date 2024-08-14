import React from "react";
import SpotifyIcon from "../../../public/assets/icons/spotify-icon.svg";
import Image from "next/image";
import { table } from "console";

interface TransactionSummaryType {
  title: string;
  date: string;
  reason: string;
  accountNo: string;
  status: string;
  amount: string;
}

const TransactionSummary = ({
  title,
  date,
  reason,
  accountNo,
  status,
  amount,
}: TransactionSummaryType) => {
  return (
    <div className="grid grid-cols-6 ">
      <div className="flex items-center col-span-2 ">
        <div className="bg-[#DCFAF8] flex items-center rounded-xl mr-2">
          <SpotifyIcon className="w-10 h-10 shrink-0" />
        </div>
        <div>
          <p className="font-semibold">{title}</p>
          <p className=" text-gray-500">{date}</p>
        </div>
      </div>
      <div className="flex items-center col-span-1  text-gray-500">
        {reason}
      </div>
      <div className="flex items-center col-span-1  text-gray-500">
        {accountNo}
      </div>
      <div className="flex items-center col-span-1 text-gray-500">{status}</div>
      <div
        className={`flex items-center col-span-1 ${
          amount[0] == "-" ? "text-red-600" : "text-blue-steel"
        }`}
      >
        {amount}
      </div>
    </div>
  );
};

export default TransactionSummary;
