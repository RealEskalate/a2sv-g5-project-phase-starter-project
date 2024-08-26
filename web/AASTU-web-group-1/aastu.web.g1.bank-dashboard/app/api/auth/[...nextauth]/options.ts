import NextAuth, { NextAuthOptions } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import { jwtDecode, JwtPayload } from "jwt-decode";
import { signInWithCredentials } from "@/lib/auth";

interface returnData {
  access_token: string;
  refresh_token: string;
  expires_in: number;
}

async function refreshAccessTokenWithRetry(
  refresh_token: string,
  maxRetries: number = 3,
  delay: number = 1000
): Promise<returnData | null> {
  let attempt = 0;

  while (attempt < maxRetries) {
    try {
      console.log(`Attempt ${attempt + 1} to refresh token.`);
      const refreshedTokens = await refreshAccessToken(refresh_token);

      if (refreshedTokens) {
        console.log("Successfully refreshed tokens:", refreshedTokens);
        return refreshedTokens;
      }
    } catch (error) {
      console.error(`Error refreshing token on attempt ${attempt + 1}:`, error);
    }

    attempt++;
    await new Promise((resolve) => setTimeout(resolve, delay)); // Delay before retrying
  }

  console.error("Failed to refresh tokens after maximum retries.");
  return null;
}

async function refreshAccessToken(
  refresh_token: string
): Promise<returnData | null> {
  console.log("Refresh Token Entering:", refresh_token);

  try {
    const url = `${process.env.NEXT_PUBLIC_BASE_URL}/auth/refresh_token`;
    console.log("Request URL:", url);

    const response = await fetch(url, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${refresh_token}`,
        "Content-Type": "application/json",
      },
    });

    console.log("Response Status:", response.status);
    const refreshedTokens = await response.json();
    console.log("Refreshed Tokens:", refreshedTokens);

    if (!response.ok) {
      throw refreshedTokens;
    }

    if (
      refreshedTokens.data &&
      refreshedTokens.data.access_token &&
      refreshedTokens.data.refresh_token
    ) {
      console.log("Returning refreshed tokens");
      return {
        access_token: refreshedTokens.data.access_token,
        refresh_token: refreshedTokens.data.refresh_token,
        expires_in: refreshedTokens.data.expires_in,
      };
    } else {
      throw new Error("Missing access or refresh token in response data");
    }
  } catch (error) {
    console.error("Error refreshing token:", error);
    throw error; // Rethrow error to allow retry logic to handle it
  }
}

export const authOptions: NextAuthOptions = {
  providers: [
    CredentialsProvider({
      name: "credentials",
      credentials: {
        userName: { label: "Username", type: "text" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials) {
        console.log("Authorize called with credentials:", credentials);
        try {
          if (!credentials) throw new Error("No credentials provided");

          const response = await signInWithCredentials(credentials);
          console.log("Sign-In Response:", response);

          if (response.success) {
            console.log("Authorization successful, returning user:", {
              id: response.data.id,
              name: response.data.name,
              email: response.data.email,
              accessToken: response.data.access_token,
              refreshToken: response.data.refresh_token,
            });
            return {
              id: response.data.id,
              name: response.data.name,
              email: response.data.email,
              accessToken: response.data.access_token,
              refreshToken: response.data.refresh_token,
            };
          } else {
            throw new Error(response.message || "Failed to sign in");
          }
        } catch (error) {
          console.error("Authorize error:", error);
          return null;
        }
      },
    }),
  ],
  callbacks: {
    async jwt({ token, user }) {
      console.log("JWT callback - Token:", token, "User:", user);

      // Initial sign in
      if (user) {
        token.accessToken = user.accessToken;
        token.refreshToken = user.refreshToken;

        // Decode the access token to get the expiry time
        const decodedToken = jwtDecode<JwtPayload>(user.accessToken);
        console.log("Decoded Access Token:", decodedToken);

        token.accessTokenExpires = decodedToken.exp
          ? decodedToken.exp * 1000
          : Date.now() + 60 * 60 * 1000; // Fallback to 1 hour
        console.log("Access Token Expires:", token.accessTokenExpires);

        return token;
      }

      // Return previous token if the access token has not expired yet
      if (Date.now() < (token.accessTokenExpires as number)) {
        console.log("Access token is still valid.");
        return token;
      }

      // Access token has expired, try to refresh it
      console.log("Access token has expired, refreshing...");
      const refreshedTokens = await refreshAccessTokenWithRetry(
        token.refreshToken as string
      );
      if (refreshedTokens) {
        console.log("Refreshed Tokens:", refreshedTokens);
        token.refreshToken = refreshedTokens.refresh_token;
        token.accessToken = refreshedTokens.access_token;

        // Decode the new access token to get the new expiry time
        const decodedNewToken = jwtDecode<JwtPayload>(
          refreshedTokens.access_token
        );
        console.log("Decoded New Access Token:", decodedNewToken);

        token.accessTokenExpires = decodedNewToken.exp
          ? decodedNewToken.exp * 1000
          : Date.now() + refreshedTokens.expires_in * 1000;

        console.log("New Access Token Expires:", token.accessTokenExpires);
      } else {
        console.log("Failed to refresh tokens.");
      }

      return token;
    },

    async session({ session, token }) {
      console.log("Session callback - Session:", session, "Token:", token);
      if (token) {
        session.user.accessToken = token.accessToken;
        session.user.refreshToken = token.refreshToken;
      }
      console.log("Updated Session:", session);
      return session;
    },
  },
  pages: {
    signIn: "/auth/sign-in",
  },
  secret: process.env.NEXTAUTH_SECRET,
};
