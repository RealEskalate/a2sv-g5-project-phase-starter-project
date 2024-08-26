"use client";
import React from "react";
import Card from "./LoanComponents/Card";
import { TableDemo } from "./LoanComponents/Table/columns";
import { loanTypes } from "@/constants";
import { getDetailData } from "@/lib/loanApies";
import { useUser } from "@/contexts/UserContext";
import { useState, useEffect } from "react";
import CardShimmer from "./LoanComponents/CardShimmer";

interface LoanTypes {
  personalLoan: number;
  businessLoan: number;
  corporateLoan: number;
}

const Loans = () => {
  const [data, setData] = useState<LoanTypes>();
  const [loading, setLoading] = useState(true);
  const { isDarkMode } = useUser();

  useEffect(() => {
    const fetchData = async () => {
      try {
        const detailData = await getDetailData();
        if (detailData == null) {
          throw new Error("failed to get data");
        }
        setData(detailData);
        setLoading(false);
      } catch (error) {
        console.error("An error occurred on card:", error);
      }
    };

    fetchData();
  }, []);

  return (
    <div className={`pt-5`}>
      {loading ? (
        <CardShimmer/>
      ) : (
        <div className="flex gap-2 overflow-x-scroll scrollbar-hidden my-3 scroll md:w-[80%] md:mx-auto rounded-full md:h-28">
          {loanTypes.map((item) => (
            <div className="p-0 m-0 my-auto" key={item.name}>
              {data && (
                <Card
                  name={item.name}
                  description={data[item.id as keyof LoanTypes]}
                  icon={item.icon}
                />
              )}
            </div>
          ))}
        </div>
      )}
      <div className="mt-5 scrollbar-hidden md:w-[80%] mx-auto">
        <h1 className={`text-md font-[500] md:font-[600]  ${isDarkMode ? "text-white" : "text-[#333B69]"} md:text-left px-2 md:text-[30px] my-3 w-[80%] `}>
          Active Loans Overview
        </h1>
        <TableDemo />
      </div>
    </div>
  );
};

export default Loans;
