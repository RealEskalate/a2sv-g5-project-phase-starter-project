import { authOptions } from './authOptions';
import NextAuth from 'next-auth/next';

export const handler = NextAuth(authOptions) ;

export { handler as GET, handler as POST };
