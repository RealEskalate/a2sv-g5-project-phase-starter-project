import Credentials from "next-auth/providers/credentials";
import CredentialsProvider from "next-auth/providers/credentials";

interface User {
  id: string | null;
  name: string;
  email: string;
  accessToken: string | null;
  refreshToken: string | null;
  profileStatus?: string; 
}


export const options = {
  providers: [
    CredentialsProvider({
        name: "Credentials",
        credentials: {
        userName: { label: 'User Name', type: 'text' },
        password: { label: 'Password', type: 'password' },
      },
      
      async authorize(credentials: { userName: string; password: string }) {
        console.log("first")
        const response = await fetch('https://bank-aait-web-group-1.onrender.com/auth/login', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            userName: credentials.userName,
            password: credentials.password,
          }),
        });

        const data = await response.json();
        console.log("Login Response Data:", data); 

        if (data.success && data.data) {
          const user:User = {
            id: data.data.id || null,
            name: credentials.userName,
            email: credentials.userName,
            accessToken: data.data.access_token,
            refreshToken: data.data.refresh_token,
          };
          return user;
  }
      },
    }),
  ],
  callbacks: {
    async jwt({ token, user }: { token: any, user: any }) {
      if (user) {
        token.accessToken = user.accessToken;
        token.profileStatus = user.profileStatus || undefined;
      }
      console.log("JWT Token:", token.accessToken);
      return token;
    },
    async session({ session, token }: { session: any, token: any }) {
      session.user.accessToken = token.accessToken;
      session.user.profileStatus = token.profileStatus;
      return session;
    },
  },
  pages: {
    signIn: '/auth/login',  
  },
};