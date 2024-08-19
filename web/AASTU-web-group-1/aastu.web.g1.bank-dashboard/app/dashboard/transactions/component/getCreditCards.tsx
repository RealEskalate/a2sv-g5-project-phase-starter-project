import { CardDetails } from "@/types";

export async function getCreditCards() {
  const accessToken =
    "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJlbW5ldC10ZXMiLCJpYXQiOjE3MjQwNDg3NjAsImV4cCI6MTcyNDEzNTE2MH0.baqrlqraepMSM7YMMdUKSUd2j_Z3ui7hyQjvw8b-ENDP9cly77sngGLsVvC3lpC-";
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
