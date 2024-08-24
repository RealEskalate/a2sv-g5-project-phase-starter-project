// next-auth.d.ts
import NextAuth from 'next-auth';

declare module 'next-auth' {
  interface Session {
    user: {
      accessToken?: string;
    } & DefaultSession['user'];
  }

  interface User {
    accessToken?: string;
  }
}
