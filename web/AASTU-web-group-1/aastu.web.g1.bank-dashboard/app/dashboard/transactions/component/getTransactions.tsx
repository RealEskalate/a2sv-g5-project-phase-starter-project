import { TransactionProps } from "@/types";

export async function getallTransactions(page: number, size: number) {
  try {
    const accessToken =
      "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJlbW5ldC10ZXMiLCJpYXQiOjE3MjQwMDY5MzIsImV4cCI6MTcyNDA5MzMzMn0.I2q3aT6zWjY09lf5LrWGVmIMDvuQR1vKU2w3jze4iW02o-cQFIeQpi95yv-QTwSO";

    const res = await fetch(
      `https://bank-dashboard-6acc.onrender.com/transactions?page=${page}&size=${size}`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
        body: null,
      }
    );
    if (!res.ok) {
      throw new Error("faild to fetch data");
    }
    const transactions: TransactionProps = await res.json();
    console.log(transactions);

    return transactions.data;
  } catch (error) {
    console.log("An error occurred:", error);
    // alert("An unexpected error occurred. Please try again later.");
  }
}
