import React from "react";
import TransactionSummary from "./TransactionSummary";

const transactions = [
  {
    title: "Spotify Subscription",
    date: "25 Jan 2021",
    reason: "Shopping",
    accountNo: "1234 ****",
    status: "Pending",
    amount: "-$150",
  },
  {
    title: "Mobile Service",
    date: "25 Jan 2021",
    reason: "Service",
    accountNo: "1234 ****",
    status: "Completed",
    amount: "-$340",
  },
  {
    title: "Emma Wilson",
    date: "25 Jan 2021",
    reason: "Transfer",
    accountNo: "1234 ****",
    status: "Completed",
    amount: "+$780",
  },
];

export default function LastTransaction() {
  return (
    <div className="w-full min-[890px]:w-2/3 mb-5">
      <p className=" text-[#333B69] pb-2 font-semibold">Last Transaction</p>
      <div className="bg-white rounded-2xl space-y-3 p-4 py-4 w-full">
        {transactions.map((ele) => (
          <TransactionSummary
            key={ele.title}
            title={ele.title}
            date={ele.date}
            reason={ele.reason}
            accountNo={ele.amount}
            status={ele.status}
            amount={ele.amount}
          />
        ))}
      </div>
    </div>
  );
}
