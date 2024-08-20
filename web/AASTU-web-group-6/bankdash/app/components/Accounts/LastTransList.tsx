"use client";
import React, { useState, useEffect } from "react";
import Last_trans from "./Last_trans";
import { useAppSelector } from "@/app/Redux/store/store";
import { TransactionType } from "@/types/TransactionValue";

interface LastTransData {
  transactionId: string;
  type: string;
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string;
}
function formatAmount(amount:number) {
  if (amount >= 1_000_000_000_000) {
    
    return `${(amount / 1_000_000_000_000).toFixed(1)}T`;
  } else if (amount >= 1_000_000_000) {
    
    return `${(amount / 1_000_000_000).toFixed(1)}B`;
  } else if (amount >= 1_000_000) {
    
    return `${(amount / 1_000_000).toFixed(1)}M`;
  } else if (amount >= 1_000) {
    
    return `${(amount / 1_000).toFixed(1)}k`;
  } else {
    
    return amount.toLocaleString();
  }
}

const LastTransList = () => {
  const [data, setData] = useState<LastTransData[]>([]);
  const [error, setError] = useState<string | null>(null);

  const expense: TransactionType[] = useAppSelector(
    (state) => state.transactions.expense
  );
  const income: TransactionType[] = useAppSelector(
    (state) => state.transactions.income
  );

  useEffect(() => {
    const getData = () => {
      try {
        const incomeData = (income ?? []).map((income: TransactionType) => ({
          ...income,
          amount: Math.abs(income.amount), // Ensure the amount is positive for income
        }));

        const expenseData = (expense ?? []).map((expense: TransactionType) => ({
          ...expense,
          amount: -Math.abs(expense.amount), // Ensure the amount is negative for expense
        }));

        const combinedTransactions: any = [...incomeData, ...expenseData];

        combinedTransactions.sort(
          (a: any, b: any) =>
            new Date(b.date).getTime() - new Date(a.date).getTime()
        );

        setData(combinedTransactions?.slice(0, 3)); // Display the most recent 3 transactions
      } catch (error) {
        setError("Error fetching data");
        alert("Error fetching data");
      }
    };

    getData(); // Call the function to fetch data
  }, [income, expense]); // Dependencies added to re-run effect when income or expense changes

  if (error) {
    return <div>{error}</div>;
  }

  if (!data || data.length === 0) {
    return <div>No transactions available</div>;
  }

  return (
    <div>
      {data.map((transaction) => (
        <Last_trans
          key={transaction.transactionId}
          transactionId={transaction.transactionId}
          type={transaction.type}
          senderUserName={transaction.senderUserName}
          description={transaction.description}
          date={transaction.date}
          amount={formatAmount(transaction.amount)}
          receiverUserName={transaction.receiverUserName}
        />
      ))}
    </div>
  );
};

export default LastTransList;
