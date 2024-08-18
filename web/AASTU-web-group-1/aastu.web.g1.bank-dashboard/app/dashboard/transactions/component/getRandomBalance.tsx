import { RandomBalanceHistory } from "@/types";

export async function getRandomBalance() {
  try {
    const accessToken =
      "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJlbW5ldC10ZXMiLCJpYXQiOjE3MjQwMDY5MzIsImV4cCI6MTcyNDA5MzMzMn0.I2q3aT6zWjY09lf5LrWGVmIMDvuQR1vKU2w3jze4iW02o-cQFIeQpi95yv-QTwSO";
      const res = await fetch(
        `https://bank-dashboard-6acc.onrender.com/transactions/random-balance-history?monthsBeforeFirstTransaction=7`,
        {
            method:"GET",
            headers:{
                Authorization: `Bearer ${accessToken}`,
            },
            body:null,
            
        }
      );
      if(!res.ok){
          throw new Error("faild to fetch data");
      }
      const balanceHistory : RandomBalanceHistory= await res.json();
      return balanceHistory.data;
  }catch (error) {
    console.log("An error occurred:", error);
    // alert("An unexpected error occurred. Please try again later.");
  }}