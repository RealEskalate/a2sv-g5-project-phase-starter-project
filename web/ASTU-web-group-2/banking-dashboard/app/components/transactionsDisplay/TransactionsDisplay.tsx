"use client";
import { Inter } from "next/font/google";
import TransactionsTable from "../transactionsTable/TransactionsTable";
import { useEffect, useState } from "react";
import { useSession } from "next-auth/react";
import { useGetAllTransactionQuery } from "@/lib/service/TransactionService";
import { TransactionType } from "../transaction/Transaction";
import RecentTransactionSkeleton from "../recent-transaction/RecentTransactionSkeleton";
import EmptyShow from "../emptyShowingImage/EmptyShow";

const inter = Inter({ subsets: ["latin"] });
const TransactionsDisplay = () => {
  const [chooseIndex, setChooseIndex] = useState(0);
  const [focusedPage, setFocusedPage] = useState(1);

  const { data: session, status } = useSession();

  useEffect(() => {}, [session, status]);

  const accessToken = session?.user.accessToken!;

  const { data: res, isLoading } = useGetAllTransactionQuery(accessToken);
  const pageLength = 1;

  const range = Array.from({ length: pageLength }, (_, i) => i + 1);

  if (isLoading) {
    return <RecentTransactionSkeleton />;
  }
  const data = res.data!.content!;

  let transactions: TransactionType[] = [];
  for (let transaction of data) {
    transactions.push({
      transactionId: transaction.transactionId,
      type: transaction.type,
      description: transaction.description,
      date: transaction.date,
      amount: transaction.amount,
    });
  }
  const deposit = transactions.filter(
    (transaction) => transaction.type.toLowerCase() == "deposit"
  );
  const expense = transactions.filter(
    (transaction) => transaction.type.toLowerCase() != "deposit"
  );

  return (
    <div className={`flex flex-col gap-5 w-full ${inter.className}`}>
      <div className="flex gap-12 text-[#718EBF] font-medium border-b border-[#EBEEF2] pl-6">
        <button
          onClick={() => setChooseIndex(0)}
          className={`${
            chooseIndex == 0 && "text-[#1814F3] border-[#1814F3] border-b-[3px]"
          }`}
        >
          All Transactions
        </button>
        <button
          onClick={() => setChooseIndex(1)}
          className={`${
            chooseIndex == 1 && "text-[#1814F3] border-[#1814F3] border-b-[3px]"
          }`}
        >
          Income
        </button>
        <button
          onClick={() => setChooseIndex(2)}
          className={`${
            chooseIndex == 2 && "text-[#1814F3] border-[#1814F3] border-b-[3px]"
          }`}
        >
          Expense
        </button>
      </div>
      <TransactionsTable
        transactions={
          chooseIndex == 0 ? transactions : chooseIndex == 1 ? deposit : expense
        }
      />
      <div className="flex justify-center">
        {chooseIndex == 0 && transactions.length == 0 && (
          // <img
          //   src="/assets/bankService/empty-image.png"
          //   alt="list empty"
          //   className="w-fit h-fit"
          // />
          <EmptyShow text="No Transaction available" />
        )}
        {chooseIndex == 1 && deposit.length == 0 && (
          // <img
          //   src="/assets/bankService/empty-image.png"
          //   alt="list empty"
          //   className="w-fit h-fit"
          // />
          <EmptyShow text="No deposit Found" />
        )}
        {chooseIndex == 2 && expense.length == 0 && (
          // <img
          //   src="/assets/bankService/empty-image.png"
          //   alt="list empty"
          //   className="w-fit h-fit"
          // />
          <EmptyShow text="No expense Found" />
        )}
      </div>

      <div className="flex justify-end rounded-xl p-6 text-[#1814F3] gap-3">
        <button
          className="flex gap-1 items-center"
          onClick={() => setFocusedPage((prev) => Math.max(1, prev - 1))}
        >
          <img
            src="/assets/transactionsDisplay/paginationLeftArrow.svg"
            alt="left-arrow"
          />{" "}
          Previous
        </button>
        {range.map((elm, index) => (
          <button
            key={index}
            onClick={() => setFocusedPage(elm)}
            className={`${
              focusedPage == elm &&
              "rounded-xl px-4 py-2 bg-[#1814F3] text-white"
            }`}
          >
            {elm}
          </button>
        ))}
        <button
          className="flex gap-1 items-center"
          onClick={() =>
            setFocusedPage((prev) => Math.min(pageLength, prev + 1))
          }
        >
          Next
          <img
            src="/assets/transactionsDisplay/paginationRightArrow.svg"
            alt="right-arrow"
          />
        </button>
      </div>
    </div>
  );
};

export default TransactionsDisplay;
