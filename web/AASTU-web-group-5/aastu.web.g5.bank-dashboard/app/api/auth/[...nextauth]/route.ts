<<<<<<< HEAD
import { authOptions } from './authOptions';
import NextAuth from 'next-auth/next';

export const handler = NextAuth(authOptions) ;

export { handler as GET, handler as POST };
=======
import NextAuth from "next-auth";
import { authOptions } from './authOptions';


const handler = NextAuth(authOptions);
export {handler as GET, handler as POST}
>>>>>>> be5493106c59d49db686b7cbdb78cfeda4a2b7a0
