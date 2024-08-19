import { UserResponse } from "@/types";

export async function getCurrentUser() {
  const accessToken =
    "eeyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJlbW5ldC10ZXMiLCJpYXQiOjE3MjQwNDg3NjAsImV4cCI6MTcyNDEzNTE2MH0.baqrlqraepMSM7YMMdUKSUd2j_Z3ui7hyQjvw8b-ENDP9cly77sngGLsVvC3lpC-";
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
