import CredentialsProvider from "next-auth/providers/credentials";
import { JWT } from 'next-auth/jwt';
import { NextAuthOptions, Session } from 'next-auth';
import { login } from "@/lib/api/authenticationController";
interface User {
  refresh_token: string;
  accessToken: string;
  data: string;
}

interface MyToken extends JWT {
  refresh_token?: string;
  accessToken?: string;
  data?: string;
}

// interface MySession extends Session {
//   user: User;
// }

export const options: NextAuthOptions = {
  session: {
    strategy: "jwt", // Use JWT for session strategy
  },


  providers: [
    CredentialsProvider({
      name: "credentials",
      credentials: {
        userName: { label: "Username", type: "string" },
        password: { label: "Password", type: "string" }
      },
    
      async authorize(credentials, req) {
        try {
          const response = await login(credentials as { userName: string; password: string });
          if (response.success) {
            console.log("successful Response");
            return response.data; // Return the user data object
          } else {
            console.log("DIDN'T GET A SUCCESSFUL RESPONSE");
            return null;
          }
        } catch (error) {
          console.error('Authorization error:', error);
          return null;
        }
      },
    })
    
  ],

  pages:{
    signIn: '/api/auth/signin'
    // signUp: '/api/auth/signup',
  },
  callbacks: {
    // Store the user information in the JWT token
    async jwt({ token, user }: any) {
      if (user) {
        token.accessToken = user.accessToken;
        token.data = user.data;     // Assuming the user object has a 'name' field
        token.refresh_token = user.refresh_token
      }
      return token;
    },
    // Make custom user data available in the session
    async session({ session, token }: any) {
      session.user.accessToken = token.accessToken;
      session.user.data = token.data;
      session.user.refresh_token = token.refresh_token;
      return session;
    },
  },
};
