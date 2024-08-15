import { NextAuthOptions, User as NextAuthUser } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";

interface Credentials {
  userName: string;
  password: string;
}

interface User extends NextAuthUser {
  data: {
    access_token: string;
    refresh_token: string;
  };
  id: string;
}

function isUser(user: unknown): user is User {
  return (
    typeof user === "object" &&
    user !== null &&
    "data" in user &&
    typeof (user as User).data === "object" &&
    "access_token" in (user as User).data &&
    "refresh_token" in (user as User).data
  );
}

export const options: NextAuthOptions = {
  providers: [
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        userName: {
          label: "Username",
          type: "text",
          placeholder: "Enter email address",
        },
        password: {
          label: "Password",
          type: "password",
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
          const user: User = {
            ...result,
            id: result.data.userId,
          };
          return user;
        }

        return null;
      },
    }),
  ],
  pages: {
    signIn: '/auth/signin',
  },
  session: {
    strategy: 'jwt',
  },
  callbacks: {
    async jwt({ token, user }) {
      if (isUser(user)) {
        token.accessToken = user.data.access_token;
        token.refreshToken = user.data.refresh_token;
      }
      return token;
    },
    async session({ session, token }) {
      session.accessToken = token.accessToken as string;
      session.refreshToken = token.refreshToken as string;
      return session;
    },
  },
};
