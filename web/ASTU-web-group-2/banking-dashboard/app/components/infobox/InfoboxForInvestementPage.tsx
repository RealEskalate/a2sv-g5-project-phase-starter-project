"use client";
import React, { useEffect } from "react";
import { infoboxListItemsInvestement } from "./infoboxListItemsInvestement";
import InfoboxCard from "./InfoboxCard";
import { useRouter } from "next/navigation";
import { useGetInvestmentHistoryQuery } from "@/lib/service/TransactionService";
import { useSession } from "next-auth/react";
import BalanceCardSkeleton from "./BalanceCardSkeleton";

const InfoboxForInvestementPage = () => {
  const router = useRouter();

  const { data: session, status } = useSession();

  useEffect(() => {}, [session, status]);

  const accessToken = session?.user.accessToken!;

  const { data: res, isLoading } = useGetInvestmentHistoryQuery(accessToken);

  if (isLoading) {
    return (
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 p-4">
      <BalanceCardSkeleton />
      <BalanceCardSkeleton />
      <BalanceCardSkeleton />
      

      </div>
    );
  }

  const data = res.data!;
  const totalInvestment = data.totalInvestment;
  const rateOfReturn = data.rateOfReturn;

  const list = [
    `$${Math.round(totalInvestment * 100) / 100}`,
    `${Math.round(Math.random() * 1000)}`,
    `+${Math.round(rateOfReturn * 100) / 100}%`,
  ];
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 p-4">
      {infoboxListItemsInvestement.map((item, index) => (
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

export default InfoboxForInvestementPage;
