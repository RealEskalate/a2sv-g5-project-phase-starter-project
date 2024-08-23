"use client";
import React, { useEffect } from "react";
import InfoboxCard from "./InfoboxCard";
import { infoboxForLoans } from "./infoboxListItemsLoans";
import { useSession } from "next-auth/react";
import { useGetMyLoansDetailQuery } from "@/lib/service/LoanService";
import BalanceCardSkeleton from "./BalanceCardSkeleton";

const InfoboxForLoans = () => {

  const { data: session, status } = useSession();

  useEffect(() => {}, [session, status]);

  const accessToken = session?.user.accessToken!;

  const { data: res, isLoading } = useGetMyLoansDetailQuery(accessToken);

  if (isLoading) {
    return (
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 p-4">
      <BalanceCardSkeleton />
      <BalanceCardSkeleton />
      <BalanceCardSkeleton />
      <BalanceCardSkeleton />

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
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 p-4">
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
