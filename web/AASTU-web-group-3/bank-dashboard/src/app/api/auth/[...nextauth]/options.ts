import { NextAuthOptions } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";

interface Credentials {
  userName: string;
  password: string;
}

interface User {
  data: {
    access_token: string;
    refresh_token: string;
  }
}

export const options: NextAuthOptions = {
  providers: [
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        userName: {
          label: "Username ",
          type: "text",
          placeholder: "Enter email address",
        },
        password: {
          label: "Password",
          type: "Password",
          placeholder: "Enter password",
        },
      },
      async authorize(credentials: Credentials | undefined): Promise<User | null> {
        if (!credentials) return null;

        const res = await fetch("https://bank-dashboard-6acc.onrender.com/auth/login", {
          method: "POST",
          body: JSON.stringify(credentials),
          headers: { "Content-Type": "application/json" },
        });

        if (!res.ok) return null;

        const result = await res.json();
        
        if (result.success && result.data) {
          const user: User = result;
          return user;
        }

        return null;
      },
    }),
  ],
  pages:{
      signIn: '/auth/signin',

  },
  session: {
    strategy: 'jwt',
  },
  callbacks: {
    async jwt({ token, user }) {
      console.log('User', user)
      if (user) {
        token.accessToken = user.data.access_token;
        token.refreshToken = user.data.refresh_token;
      }
      return token;
    },
    async session({ session, token }) {
      session.accessToken = token.accessToken;
      session.refreshToken = token.refreshToken;
      return session;
    },
  },
};
