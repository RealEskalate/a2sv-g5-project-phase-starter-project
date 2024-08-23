"use client";
import React from "react";
import { useSession } from "next-auth/react";
import { TransactionType } from "@/types/TransactionValue";

const Download = ({
  transactionId,
  transaction,
}: {
  transactionId: string;
  transaction: TransactionType;
}) => {
  const { data: session } = useSession();
  const accessToken = session?.accessToken as string;

  const handleDownload = async () => {
    if (!accessToken) {
      alert("You are not authenticated. Please log in.");
      return;
    }

    try {
      const response = await fetch(
        `https://bank-dashboard-rsf1.onrender.com/transactions/${transactionId}`,
        {
          headers: {
            Authorization: `Bearer ${accessToken}`,
            "Content-Type": "application/json",
          },
        }
      );

      if (!response.ok) {
        throw new Error("Failed to fetch transaction data");
      }

      const data = await response.json();

      const receipt = `
        Transaction Receipt
        -------------------
        Transaction ID: ${transaction.transactionId}
        Description: ${transaction.description}
        Type: ${transaction.type}
        Amount: $${Math.abs(transaction.amount)}
        Date: ${transaction.date}
        -------------------
      `;

      const blob = new Blob([receipt], { type: "text/plain" });
      const url = window.URL.createObjectURL(blob);

      const link = document.createElement("a");
      link.href = url;
      link.download = `receipt-${transactionId}.txt`;
      document.body.appendChild(link);
      link.click();

      document.body.removeChild(link);
      window.URL.revokeObjectURL(url);
    } catch (error) {
      console.error("Error downloading receipt:", error);
      alert("Failed to download receipt. Please try again.");
    }
  };

  return (
    <button className="table-button" onClick={handleDownload}>
      Download
    </button>
  );
};

export default Download;
