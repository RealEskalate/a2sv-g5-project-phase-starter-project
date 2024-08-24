import { AuthOptions } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import { JWT } from "next-auth/jwt";
console.log(process.env.NEXTAUTH_SECRET, "enviroment");
async function refreshAccessToken(token: JWT) {
  console.log("authOptions", token);

  try {
    const res = await fetch(
      "https://bank-dashboard-rsf1.onrender.com/auth/refresh_token",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token.refreshToken}`,
        },
      }
    );

    const refreshedTokens = await res.json();
    console.log(refreshedTokens, "refreshedTokens");
    if (!res.ok) {
      throw refreshedTokens;
    }
    console.log();
    return {
      ...token,
      accessToken: refreshedTokens.access_token,
      accessTokenExpires: Date.now() + 10 * 60 * 1000,
      refreshToken: refreshedTokens.refresh_token,
    };
  } catch (error) {
    console.error("Failed to refresh access token", error);

    return {
      ...token,
      error: "RefreshAccessTokenError",
    };
  }
}

const authOptions: AuthOptions = {
  providers: [
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        userName: { label: "Username", type: "text", placeholder: "" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials) {
        const res = await fetch(
          "https://bank-dashboard-rsf1.onrender.com/auth/login",
          {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              userName: credentials?.userName,
              password: credentials?.password,
            }),
          }
        );

        const user = await res.json();

        if (user && user.data) {
          return {
            id: "f",
            username: credentials?.userName,
            accessToken: user.data.access_token,
            refreshToken: user.data.refresh_token,
            accessTokenExpires: Date.now() + 10 * 60 * 1000,
          };
        } else {
          return null;
        }
      },
    }),
  ],
  secret: process.env.NEXTAUTH_SECRET,
  callbacks: {
    async jwt({ token, user }: { token: JWT; user: any }) {
      console.log(user, "initial signin ");
      if (user) {
        return {
          accessToken: user.accessToken,
          refreshToken: user.refreshToken,
          accessTokenExpires: Date.now() + 10 * 60 * 1000,
          username: user.username,
        };
      }
      console.log(
        (Date.now() - (token.accessTokenExpires as number)) / 600,
        "time111"
      );

      if (Date.now() < (token.accessTokenExpires as number)) {
        return token;
      }

      console.log(
        (Date.now() - (token.accessTokenExpires as number)) / 600,
        "timeleft"
      );
      return refreshAccessToken(token);
    },
    async session({ session, token }: { session: any; token: JWT }) {
      session.user.name = token.username;
      session.user.accessToken = token.accessToken;
      session.user.refreshToken = token.refreshToken;
      session.error = token.error;

      return session;
    },
  },
};

export { authOptions };
