import NextAuth from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import axios from "axios";

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
            return response.data;  // Return the user object if authentication is successful
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
        token.accessToken = user.access_token;  // Set access token in JWT token
      }
      return token;
    },
    async session({ session, token }) {
      session.accessToken = token.accessToken;  // Attach access token to session
      return session;
    },
  },
});

export { handler as GET, handler as POST };
