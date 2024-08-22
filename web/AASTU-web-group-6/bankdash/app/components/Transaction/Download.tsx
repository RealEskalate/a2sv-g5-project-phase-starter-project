import React from "react";
import axios from "axios";
import { TransactionType } from "@/types/TransactionValue";

const Download = ({
  transactionId,
  transaction,
  accessToken,
}: {
  transactionId: string;
  transaction: TransactionType;
  accessToken: string;
}) => {
  const handleDownload = async () => {
    try {
      const headers = {
        Authorization: `Bearer ${accessToken}`,
        "Content-Type": "application/json",
      };

      // Fetch the transaction data using Axios with headers
      const response = await axios.get(
        `https://bank-dashboard-o9tl.onrender.com/transactions/${transactionId}`,
        {
          headers,
        }
      );

      // Create a simple text-based receipt
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

      // Create a Blob with the receipt content
      const blob = new Blob([receipt], { type: "text/plain" });

      // Create a temporary URL for the Blob
      const url = window.URL.createObjectURL(blob);

      // Create a temporary anchor element and trigger the download
      const link = document.createElement("a");
      link.href = url;
      link.download = `receipt-${transactionId}.txt`;
      document.body.appendChild(link);
      link.click();

      // Clean up
      document.body.removeChild(link);
      window.URL.revokeObjectURL(url);
    } catch (error) {
      console.error("Error downloading receipt:", error);
      alert("Failed to download receipt. Please try again.");
    }
  };
  return (
    <button className="table-button" onClick={() => handleDownload()}>
      Download
    </button>
  );
};

export default Download;
