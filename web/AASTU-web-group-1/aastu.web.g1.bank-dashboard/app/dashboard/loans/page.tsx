'use client'
import React from "react";
import Card from "./LoanComponents/Card";
import { TableDemo } from "./LoanComponents/Table/columns";
import { loanTypes } from "@/constants";
import { useUser } from "@/contexts/UserContext";
const Loans = () => {
  const { isDarkMode } = useUser();
  return (
    <div className={isDarkMode ? "bg-gray-700" : "bg-slate-100"}>
      <div className="flex gap-2 overflow-x-scroll scrollbar-hidden my-3 px-10">
        {loanTypes.map((item) => (
          <div className="my-2" key={item.name}>
            <Card {...item} />
          </div>
        ))}
      </div>
      <div className="mt-5 p-2">
        <h1
          className={`text-md font-[500] md:font-[600]  ${
            isDarkMode ? "text-white" : "text-[#333B69]"
          } md:text-left md:pl-20 md:text-[35px] pl-4`}
        >
          Active Loans Overview
        </h1>
        <TableDemo />
      </div>
    </div>
  );
};

export default Loans;
