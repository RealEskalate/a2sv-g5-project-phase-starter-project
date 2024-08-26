"use client";

import React, { useEffect, useState } from "react";
import { colors } from "@/constants";
import { getAllTransactionsss } from "@/services/transactionfetch";
import LifeInsuranceIcon from "@/public/icons/LifeInsuranceIcon";
import ShoppingIcon from "@/public/icons/ShoppingIcon";
import SavingAccountsIcon from "@/public/icons/SavingAccountsIcon";
import BusinessLoans from "@/public/icons/BusinessLoans";

import { TbFileSad } from "react-icons/tb";

const RecentTransaction = () => {
  const [recentTransaction, setRecentTransaction] = useState<any[]>([]);
  const [status, setStatus] = useState<"loading" | "success" | "error">(
    "loading"
  );

  useEffect(() => {
    const fetchRecentTransaction = async () => {
      setStatus("loading");
      try {
        const response = await getAllTransactionsss();
        setRecentTransaction(response.data.content);
        setStatus("success");
      } catch (error) {
        console.error("Error fetching the recent transactions: ", error);
        setStatus("error");
      }
    };
    fetchRecentTransaction();
  }, []);

  // Get the last 3 transactions
  const lastThreeTransactions = recentTransaction.slice(-3).reverse();

  if (status === "loading") {
    return (
      <div className="max-w-auto">
        {Array.from({ length: 2 }).map((_, index) => (
          <div
            key={index}
            className="flex justify-between p-4 mb-4 rounded-lg bg-gray-200 dark:bg-gray-700 animate-pulse"
          >
            <div className="flex gap-4 items-center">
              <div className="w-16 h-16 bg-gray-300 dark:bg-gray-600 rounded-full" />
              <div>
                <div className="w-24 h-4 bg-gray-300 dark:bg-gray-600 mb-2" />
                <div className="w-32 h-4 bg-gray-300 dark:bg-gray-600" />
              </div>
            </div>
            <div className="w-20 h-4 bg-gray-300 dark:bg-gray-600" />
          </div>
        ))}
      </div>
    );
  }

  if (status === "error") {
    return (
      <div className="flex flex-col items-center justify-center h-full text-red-500">
        <TbFileSad
          className={`text-gray-300 dark:text-[#993d4b] w-[400px] h-[70px] pb-2 block mx-auto`}
          strokeWidth={1}
        />
        <div>Error fetching the recent transactions</div>
      </div>
    );
  }

  return (
    <>
      {recentTransaction.length === 0 ? (
        <div className="flex flex-col items-center justify-center h-full text-white dark:text-blue-500">
        
        <div>No data to display</div>
      </div>
      ) : (
        <div className="max-w-auto">
          {lastThreeTransactions.map((transaction) => {
            let IconComponent = LifeInsuranceIcon;
            let amountTextColor = "text-green-500";
            let formattedAmount = `$${transaction.amount.toLocaleString()}`;

            if (transaction.amount > 1000000) {
              formattedAmount = `${formattedAmount.slice(
                0,
                formattedAmount.indexOf(",") + 6
              )}...`;
            }

            switch (transaction.type) {
              case "transfer":
                IconComponent = LifeInsuranceIcon;
                amountTextColor = "text-red-500";
                formattedAmount = `-${formattedAmount}`;
                break;
              case "shopping":
                IconComponent = ShoppingIcon;
                amountTextColor = "text-red-500";
                formattedAmount = `-${formattedAmount}`;
                break;
              case "service":
                IconComponent = SavingAccountsIcon;
                amountTextColor = "text-red-500";
                formattedAmount = `-${formattedAmount}`;
                break;
              case "deposit":
                IconComponent = BusinessLoans;
                amountTextColor = "text-green-500";
                formattedAmount = `+${formattedAmount}`;
                break;
              default:
                break;
            }

            return (
              <div
                key={transaction.transactionId}
                className={`${colors.white} max-w-auto rounded-lg dark:bg-dark text-gray-900 dark:text-white`}
              >
                <div className="flex p-2 justify-between items-center w-auto md:w-auto">
                  <div className="flex gap-4 items-center">
                    <div>
                      <IconComponent className="w-16 h-16" />
                    </div>
                    <div className="my-2">
                      <p className="text-lg font-medium">
                        {transaction.senderUserName}
                      </p>
                      <p className={`${colors.textgray} text-sm text-start`}>
                        {transaction.date}
                      </p>
                    </div>
                  </div>
                  <div className="my-5">
                    <p className={`${amountTextColor}`}>{formattedAmount}</p>
                  </div>
                </div>
              </div>
            );
          })}
        </div>
      )}
    </>
  );
};

export default RecentTransaction;
