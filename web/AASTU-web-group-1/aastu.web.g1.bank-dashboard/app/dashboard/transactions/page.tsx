"use client";

import { useState, useEffect } from "react";
import { useUser } from "@/contexts/UserContext";
import { Loading } from "../_components/Loading";
import { TransactionCards } from "./component/transactionCards";
import { TransactionTable } from "./component/TransactionTable";

const Transactions = () => {
  const { isDarkMode } = useUser();
  const [dataFetched, setDataFetched] = useState(false);
  const [transactionCardsLoaded, setTransactionCardsLoaded] = useState(false);
  const [transactionTableLoaded, setTransactionTableLoaded] = useState(false);

  // Check if all data has been fetched
  useEffect(() => {
    if (transactionCardsLoaded && transactionTableLoaded) {
      setDataFetched(true);
    }
  }, [transactionCardsLoaded, transactionTableLoaded]);

  return (
    <div
      className={`min-h-screen ${
        dataFetched ? "p-10 space-y-5" : "flex justify-center items-center"
      } ${
        isDarkMode ? "bg-gray-700 text-gray-200" : "bg-[#F5F7FA] text-gray-900"
      }`}
    >
      {!dataFetched && <Loading />}

      {/* Render components in the background to trigger data fetching */}
      <div className={dataFetched ? "" : "hidden"}>
        {/* First row */}
        <TransactionCards
          onLoadingComplete={() => setTransactionCardsLoaded(true)}
        />
        {/* Second row */}
        <TransactionTable
          onLoadingComplete={() => setTransactionTableLoaded(true)}
        />
      </div>
    </div>
  );
};

export default Transactions;
