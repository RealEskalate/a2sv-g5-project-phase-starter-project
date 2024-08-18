import { TransactionProps } from "./ExpenseTable";

export async function getExpenses(page: number, size: number) {
  try {
    const accessToken =
      "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJlbW5ldC10ZXMiLCJpYXQiOjE3MjM5NjM4NjQsImV4cCI6MTcyNDA1MDI2NH0.vR8GlJLHI7X9_aISaO4jwuoGayo1Kyo61o0Qc0TsTDBJowGNQ5V1juj88rkvDOO1";

    const res = await fetch(
      `https://bank-dashboard-6acc.onrender.com/transactions/expenses?page=${page}&size=${size}`,
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
    const expenses: TransactionProps = await res.json();
    console.log(expenses);

    return expenses.data;
  } catch (error) {
    console.log("An error occurred:", error);
    // alert("An unexpected error occurred. Please try again later.");
  }
}
