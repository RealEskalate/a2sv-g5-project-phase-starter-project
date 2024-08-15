import React from "react";
import TransactionSummary from "./TransactionSummary";

export default function LastTransaction() {
  return (
    <div className="w-full md:w-2/3 ">
      <p className="font-semibold py-3 text-deepNavy text-lg ">
        Last Transaction
      </p>
      <div className="bg-white rounded-2xl space-y-4 p-4 w-full">
        <TransactionSummary
          title="Spotify Subscription"
          date="25 Jan 2021"
          reason="Shopping"
          accountNo="1234 ****"
          status="Pending"
          amount="-$150"
        />
        <TransactionSummary
          title="Mobile Service"
          date="25 Jan 2021"
          reason="Service"
          accountNo="1234 ****"
          status="Completed"
          amount="-$340"
        />
        <TransactionSummary
          title="Emma Wilson"
          date="25 Jan 2021"
          reason="Transfer"
          accountNo="1234 ****"
          status="Completed"
          amount="+$780"
        />
      </div>
    </div>
  );
}
