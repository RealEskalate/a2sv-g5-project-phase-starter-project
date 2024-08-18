import { UserResponse } from "@/types";

export async function getCurrentUser() {
  const accessToken =
    "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJlbW5ldC10ZXMiLCJpYXQiOjE3MjM5NjM4NjQsImV4cCI6MTcyNDA1MDI2NH0.vR8GlJLHI7X9_aISaO4jwuoGayo1Kyo61o0Qc0TsTDBJowGNQ5V1juj88rkvDOO1";
    try {
      const res = await fetch(
        `https://bank-dashboard-6acc.onrender.com/user/current`,
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

      const currentUser: UserResponse = await res.json();
      return currentUser.data;
    } catch (error) {
      console.error("An error occurred:", error);
      alert("An unexpected error occurred. Please try again later.");
    }
 
}
