import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';
import { getToken } from 'next-auth/jwt';

export async function middleware(req: NextRequest) {
//   const publicPaths = ['/sign_in', '/sign_up'];
//   const token = await getToken({ req, secret: process.env.NEXTAUTH_SECRET });

//   // console.log('Token:', token);

//   if (publicPaths.includes(req.nextUrl.pathname) || token) {
//     return NextResponse.next();
//   } else {
//     const signInUrl = new URL('/sign_in', req.url);
//     return NextResponse.redirect(signInUrl);
//   }
// }

// export const config = {
//   matcher: ['/'],
// console.log("middleware")
};
