"use client";

import { useState, useEffect } from "react";


import CreditCard from "../_components/Credit_Card";
import { ExpenseChart } from "./component/ExpenseChart";
import { ExpenseTable } from "./component/ExpenseTable";
import { CardDetails, TransactionContent, TransactionData } from "@/types";
import {
  getallTransactions,
  getCreditCards,
  getExpenses,
  getIncomes,
} from "@/lib/api";
import { Loading } from "../_components/Loading";
import { useUser } from "@/contexts/UserContext";
import { TransactionCards } from "./component/transactionCards";
import { TransactionTable } from "./component/TransactionTable";

const Transactions = () => {
  const { isDarkMode } = useUser();

  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try{}
      finally{ setLoading(false);}
       
    };
    fetchData();
  }, []);

  

  

  if (loading) {
    return <Loading />;
  }

  return (
    <div
      className={`space-y-5 p-5 ${
        isDarkMode ? "bg-gray-700 text-gray-200" : "bg-[#F5F7FA] text-gray-900"
      }`}
    >
      {/* first row */}
      <TransactionCards/>
          {/* second row */}
     <TransactionTable/>
    </div>
  );
};

export default Transactions;
