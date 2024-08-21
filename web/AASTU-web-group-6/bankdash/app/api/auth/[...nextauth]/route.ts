import NextAuth from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import AuthService from "@/app/Services/api/authService";

interface User {
  refreshToken: string;
  accessToken: string;
}

const handler = NextAuth({
  providers: [
    CredentialsProvider({
      name: "Credentials",
      // type: "credentials",
      credentials: {
        userName: { label: "UserName", type: "text", placeholder: "jondoe" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials) {
        // console.log("Authorize called with credentials:", credentials);

        if (!credentials) {
          return null;
        }

        try {
          const response = await AuthService.login(credentials);
          if (response.success) {

            const data: any = response.data;

            const userData: User = {
              refreshToken: data.refresh_token,
              accessToken: data.access_token,
            };

            return userData;
          } else {
            return null;
          }
        } catch (error) {
          console.error("Login error:", error);
          return null;
        }
      },
    }),
  ],
  callbacks: {
    async jwt({ token, user }) {
      if (user) {
        token.accessToken = user.accessToken;
        token.refreshToken = user.refreshToken;
      }
      return token;
    },
    async session({ session, token }) {
      session.accessToken = token.accessToken;
      session.refreshToken = token.refreshToken;
      return session;
    },
  },
  pages: {
    signIn:"/login",
    error: "/error",
  },
});

export { handler as GET, handler as POST };
