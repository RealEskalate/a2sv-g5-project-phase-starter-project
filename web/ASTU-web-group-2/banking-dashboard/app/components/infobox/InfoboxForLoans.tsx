"use client";
import React, { useEffect } from "react";
import InfoboxCard from "./InfoboxCard";
import { infoboxForLoans } from "./infoboxListItemsLoans";
import { useSession } from "next-auth/react";
import { useGetMyLoansDetailQuery } from "@/lib/service/LoanService";

const InfoboxForLoans = () => {

  const { data: session, status } = useSession();

  useEffect(() => {}, [session, status]);

  const accessToken = session?.user.accessToken!;

  const { data: res, isLoading } = useGetMyLoansDetailQuery(accessToken);

  if (isLoading) {
    return (
      <div className="flex justify-center items-center h-screen">
        <div className="w-16 h-16 border-t-4 border-b-4 border-blue-500 rounded-full animate-spin"></div>
      </div>
    );
  }

  const data = res.data!;

  const list = [
    `$${data.personalLoan}`,
    `$${data.corporateLoan}`,
    `$${data.businessLoan}`,
    "Choose Money",
  ];
  return (
    <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4 p-4">
      {infoboxForLoans.map((item, index) => (
        <InfoboxCard
          key={index}
          name={item.name}
          icon={item.icon}
          value={list[index]}
        />
      ))}
    </div>
  );
};

export default InfoboxForLoans;
