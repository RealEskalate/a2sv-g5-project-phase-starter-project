import { NextAuthOptions, Session as NextAuthSession } from 'next-auth';
import NextAuth from 'next-auth';
import CredentialsProvider from 'next-auth/providers/credentials';

export const authOptions: NextAuthOptions = {
  providers: [
    CredentialsProvider({
      name: 'Credentials',
      credentials: {
        username: { label: 'Username', type: 'username', placeholder: 'Enter Your Email Address' },
        password: { label: 'Password', type: 'password' },
      },
      async authorize(credentials) {
        // console.log('data from credentials', credentials);
        const userName = credentials?.username;
        const password = credentials?.password;
        const res = await fetch(`https://bank-dashboard-6acc.onrender.com/auth/login`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ userName, password }),
        });
        // console.log('res is', res);
        const user = await res.json();
        // console.log('first user is ', user);
        if (res.status === 200) {
          // console.log('authorize response is ', user);
          return {
            id: user.data.id,
            email: user.data.email,
            accessToken: user.data.access_token,
            refreshToken: user.data.refresh_token,
            profileComplete: user.data.profileComplete,
            message: user.message,
            success: user.success,
            name: user.data.name,
            role: user.data.role,
          };
        } else {
          return null;
        }
      },
    }),
  ],
  secret: process.env.NEXTAUTH_SECRET as string,
  callbacks: {
    async signIn({ user, account, profile, email, credentials }) {
      // console.log('sign in with user', user);
      return true;
    },
    async jwt({ token, user }) {
      // console.log('first token is ', token, 'user in jwt is ', user);
      if (user) {
        // token = { ...token, ...user };
        token.accessToken = user.accessToken;
        token.refreshToken = user.refreshToken;
      }
      // console.log('new token is ', token);
      // token.name = user;
      return token;
    },
    async session({ session, token }) {
      // session.user = token;
      session.accessToken = token.accessToken;
      session.refreshToken = token.refreshToken;
      // console.log('first session is ', session, 'token in session is ', token);
      return session;
    },
    async redirect({ url, baseUrl }) {
      return baseUrl;
    },
  },
};

export default NextAuth(authOptions);
