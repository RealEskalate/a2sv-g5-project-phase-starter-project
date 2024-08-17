"use client";
import React, { useState, useEffect } from "react";
import Last_trans from "./Last_trans";
import LastTransService from "@/app/Services/api/lastTransService";
interface LastTransData {
    transactionId: string;
    type: string;
    senderUserName: string;
    description: string;
    date: string;
    amount: number;
    receiverUserName: string;
  }
const LastTransList = () => {
  const [data, setData] = useState<LastTransData[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const getData = async () => {
      try {
        const accessToken = "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJtaWhyZXQiLCJpYXQiOjE3MjM4OTg5NjUsImV4cCI6MTcyMzk4NTM2NX0.xOmHULIGVJva2RHp9859jUC0zRjOzitYRFrcdbLwUkdYTsIIEtHJv8rz76nUgq0r";
        const expense = await LastTransService.getExpenseData(accessToken);
        const income = await LastTransService.getIncomeData(accessToken);
        const incomeData = (income ?? []).map((income: LastTransData) => ({
          ...income,
          amount: Math.abs(income.amount), 
        }));
        
        const expenseData = (expense ?? []).map((expense: LastTransData) => ({
          ...expense,
          amount: -Math.abs(expense.amount), 
        }));
        
        const combinedTransactions = [...incomeData, ...expenseData];
        
        combinedTransactions.sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime());
        
        console.log(combinedTransactions);
        console.log(expense)
        console.log(income)
        setData(combinedTransactions.slice(0, 3))
      } catch (error) {
        setError("Error fetching data");
        alert("Error Fetching data ");
      }
    };

    getData();
  }, []);

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
          amount={transaction.amount}
          receiverUserName={transaction.receiverUserName}
        />
      ))}
    </div>
  );
};

export default LastTransList;
