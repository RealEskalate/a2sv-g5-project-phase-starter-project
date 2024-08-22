// auth/api/[...nextauth]/route.ts
import NextAuth from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import axios from 'axios';

const authOptions = {
  providers: [
    CredentialsProvider({
      name: 'Credentials',
      credentials: {
        userName: { label: 'Username', type: 'text' },
        password: { label: 'Password', type: 'password' }
      },
      authorize: async (credentials) => {
        try {
          const response = await axios.post('https://bank-dashboard-1tst.onrender.com/auth/login', {
            userName: credentials.userName,
            password: credentials.password
          });
          const user = response.data.data;
          if (user && user.access_token) {
            return {
              ...user,
              accessToken: user.access_token,
              refreshToken: user.refresh_token
            };
          } else {
            return null;
          }
        } catch (error) {
          console.error('Error during authorization:', error);
          return null;
        }
      }
    })
  ],
  pages: {
    signIn: '/auth/signin',
    error: '/auth/error'
  }
};

export default NextAuth(authOptions);