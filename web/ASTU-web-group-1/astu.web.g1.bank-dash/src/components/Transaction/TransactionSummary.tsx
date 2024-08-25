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
    <div className="grid grid-cols-6 md:grid-cols-8 lg:grid-cols-8 xl:grid-cols-12 ">
      <div className="flex items-center col-span-5 ">
        <div className="bg-[#DCFAF8] flex items-center rounded-xl mr-2">
          <SpotifyIcon className="w-10 h-10 shrink-0" />
        </div>
        <div>
          <p className="text-slate-950 text-xs lg:text-sm">{title}</p>
          <p className="text-blue-steel text-sm">{date}</p>
        </div>
      </div>
      <div className="hidden md:flex items-center text-blue-steel col-span-2  text-xs lg:text-sm">
        {reason}
      </div>
      <div className="hidden xl:flex items-center text-blue-steel col-span-2 text-xs lg:text-sm ">
        {accountNo}
      </div>
      <div className="hidden xl:flex items-center text-blue-steel col-span-2 text-xs lg:text-sm ">
        {status}
      </div>
      <div
        className={`flex items-center col-span-1 ${
          amount[0] == "-" ? "text-candyPink" : "text-mintGreen"
        } text-xs lg:text-sm`}
      >
        {amount}
      </div>
    </div>
  );
};

export default TransactionSummary;
