"use client";
import React, { useEffect } from "react";
import { infoboxListItemsInvestement } from "./infoboxListItemsInvestement";
import InfoboxCard from "./InfoboxCard";
import { useRouter } from "next/navigation";
import { useGetInvestmentHistoryQuery } from "@/lib/service/TransactionService";
import { useSession } from "next-auth/react";

const InfoboxForInvestementPage = () => {
  const router = useRouter();

  const { data: session, status } = useSession();

  useEffect(() => {}, [session, status]);

  const accessToken = session?.user.accessToken!;

  const { data: res, isLoading } = useGetInvestmentHistoryQuery(accessToken);

  if (isLoading) {
    return (
      <div className="flex justify-center items-center h-screen">
        <div className="w-16 h-16 border-t-4 border-b-4 border-blue-500 rounded-full animate-spin"></div>
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
    <div className="grid grid-cols-1 sm:grid-cols-3 md:grid-cols gap-4 p-4 w-auto">
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
