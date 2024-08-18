import { CardDetails } from "@/types";

export async function getCreditCards() {
  const accessToken =
    "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJlbW5ldC10ZXMiLCJpYXQiOjE3MjQwMDY5MzIsImV4cCI6MTcyNDA5MzMzMn0.I2q3aT6zWjY09lf5LrWGVmIMDvuQR1vKU2w3jze4iW02o-cQFIeQpi95yv-QTwSO";
  try {
    const res = await fetch(
      `https://bank-dashboard-6acc.onrender.com/cards`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
        body: null,
      }
    );
    if (!res.ok) {
      throw new Error("failed to get data");
    }
    const cards: CardDetails[]= await res.json();
    return cards;
  } catch (error) {
    console.error("An error occurred:", error);
    alert("An unexpected error occurred. Please try again later.");
  }
}
