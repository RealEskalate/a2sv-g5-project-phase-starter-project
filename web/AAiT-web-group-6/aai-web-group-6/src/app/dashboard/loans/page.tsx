import Loan_box from "@/components/Loan_components/loan_box/Loan_box";
import Loan_table from "@/components/Loan_components/loan_table/Loan_table";
import React from "react";

const Loan = () => {
  return (
    <div className="flex flex-col gap-3 bg-[#F5F7FA] w-full max-w-full h-screen py-[1.5rem] px-[2rem] overflow-x-auto">
      <div className="h-auto">
        <Loan_box />
      </div>
      <h1 className="text-[#333B69] font-bold pt-[1rem] text-[14pt] md:text-[18pt]">Active Loan Overview</h1>
      <Loan_table />
    </div>
  );
};

export default Loan;
