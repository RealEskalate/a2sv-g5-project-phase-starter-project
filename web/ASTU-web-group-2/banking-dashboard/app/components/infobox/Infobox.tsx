"use client";
import React, { useEffect } from "react";
import InfoboxCard from "./InfoboxCard";
import { infoboxListItems } from "./infoboxListItems";
import { useRouter } from "next/navigation";
import { useSession } from "next-auth/react";
import { useGetAllTransactionQuery } from "@/lib/service/TransactionService";
import { useGetCurrentUserQuery } from "@/lib/service/UserService";
import BalanceCardSkeleton from "./BalanceCardSkeleton";

const Infobox = () => {
  const router = useRouter();

  const { data: session, status } = useSession();

  useEffect(() => {}, [session, status]);

  const accessToken = session?.user.accessToken!;

  const { data: res, isLoading } = useGetAllTransactionQuery(accessToken);
  const { data: resUserCurrent, isLoading: userIsLoading } =
    useGetCurrentUserQuery(accessToken);

  // if (isLoading || userIsLoading) {
  if (isLoading){
    return (
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4 p-4">
      <BalanceCardSkeleton />
      <BalanceCardSkeleton />
      <BalanceCardSkeleton />
      <BalanceCardSkeleton />

      </div>
    );
  }

  const data = res.data!.content!;
  const balance = resUserCurrent.data!.accountBalance!;
  let transactions: { type: string; amount: number }[] = [];
  for (let transaction of data) {
    transactions.push({
      type: transaction.type,
      amount: transaction.amount,
    });
  }
  const deposit = transactions
    .filter((transaction) => transaction.type.toLowerCase() == "deposit")
    .reduce((prev, currTransaction) => prev + currTransaction.amount, 0);
  const expense = transactions
    .filter((transaction) => transaction.type.toLowerCase() != "deposit")
    .reduce((prev, currTransaction) => prev + currTransaction.amount, 0);
  const list = [balance, deposit, expense, balance];
  return (
    <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4 p-4">
      {infoboxListItems.map((item, index) => (
        <InfoboxCard
          key={index}
          name={item.name}
          icon={item.icon}
          value={`$${list[index]}`}
        />
      ))}
    </div>
  );
};

export default Infobox;
