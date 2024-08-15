"use client";
import { Inter } from "next/font/google";
import TransactionsTable from "../transactionsTable/TransactionsTable";
import { useState } from "react";

const transactions = [
  {
    id: 1,
    description: "Spotify Subscription",
    transactionId: "#12548796",
    type: "Shopping",
    card: "1234 ****",
    date: "28 Jan, 12.30 AM",
    amount: -2500,
  },
  {
    id: 2,
    description: "Freepik Sales",
    transactionId: "#12548796",
    type: "Transfer",
    card: "1234 ****",
    date: "25 Jan, 10.40 PM",
    amount: 750,
  },
  {
    id: 3,
    description: "Mobile Service",
    transactionId: "#12548796",
    type: "Service",
    card: "1234 ****",
    date: "20 Jan, 10.40 PM",
    amount: -150,
  },
  {
    id: 4,
    description: "Wilson",
    transactionId: "#12548796",
    type: "Transfer",
    card: "1234 ****",
    date: "15 Jan, 03.29 PM",
    amount: -1050,
  },
  {
    id: 5,
    description: "Emilly",
    transactionId: "#12548796",
    type: "Transfer",
    card: "1234 ****",
    date: "14 Jan, 10.40 PM",
    amount: 840,
  },
];

const deposit = transactions.filter((transaction) => transaction.amount >= 0);
const expense = transactions.filter((transaction) => transaction.amount < 0);

const inter = Inter({ subsets: ["latin"] });
const TransactionsDisplay = () => {
  const [chooseIndex, setChooseIndex] = useState(0);
  return (
    <div className={`flex flex-col gap-5 ${inter.className}`}>
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
      {chooseIndex == 0 ? (
        <TransactionsTable transactions={transactions} />
      ) : chooseIndex == 1 ? (
        <TransactionsTable transactions={deposit} />
      ) : (
        <TransactionsTable transactions={expense} />
      )}

      <div className="flex justify-end rounded-xl p-6 text-[#1814F3] gap-3">
        <button className="flex gap-1 items-center">
          <img
            src="/assets/transactionsDisplay/paginationLeftArrow.svg"
            alt="left-arrow"
          />{" "}
          Previous
        </button>
        <button className="rounded-xl px-4 py-2 bg-[#1814F3] text-white">
          1
        </button>
        <button>2</button>
        <button>3</button>
        <button>4</button>
        <button className="flex gap-1 items-center">
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
