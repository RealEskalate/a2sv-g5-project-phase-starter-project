import NextAuth from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import axios from "axios";
import LoginResponseValue from "@/types/LoginResponseValue";

// Define User type based on what NextAuth expects
interface User {
  id: string;
  name?: string;
  email?: string;
  accessToken: string;
  // other fields as required
}

const handler = NextAuth({
  providers: [
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        userName: { label: "UserName", type: "text", placeholder: "jondoe" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials, req) {
        if (!credentials) {
          return null; // Return null if credentials are undefined
        }

        const { userName, password } = credentials;

        try {
          const response = await axios.post(
            "https://bank-dashboard-6acc.onrender.com/auth/login",
            { userName, password },
            {
              headers: {
                "Content-Type": "application/json",
              },
            }
          );

          if (response.status === 200) {
            const data: LoginResponseValue = response.data;

            // Ensure the returned object conforms to the User type
            const user: User = {
              id: data.data.access_token,  // Assuming access_token is used as ID, adjust as needed
              name: data.message,  // Adjust as needed based on response structure
              accessToken: data.data.access_token,
              // Populate other fields as necessary
            };

            return user;  // Return the user object
          } else {
            return null;  // Return null if authentication fails
          }
        } catch (err) {
          console.error(err);
          return null; // Return null in case of any error during the authentication process
        }
      },
    }),
  ],
  callbacks: {
    async jwt({ token, user }) {
      if (user) {
        token.accessToken = user.accessToken;  // Set access token in JWT token
      }
      return token;
    },
    async session({ session, token }) {
      session.accessToken = (token as { accessToken?: string }).accessToken; // Attach access token to session
      return session;
    },
  },
});

export { handler as GET, handler as POST };
