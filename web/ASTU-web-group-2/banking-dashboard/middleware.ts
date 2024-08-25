import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';
import { getToken } from 'next-auth/jwt';
import { isTokenExpired } from './utils/authUtils'; // Adjust the path to where your utility function is located

export async function middleware(req: NextRequest) {
  const sessionToken = await getToken({ req, secret: process.env.AUTH_SECRET });
  const { pathname } = req.nextUrl;

  // Check if the token is expired
  if (sessionToken && isTokenExpired(sessionToken.accessToken)) {
    return NextResponse.redirect(new URL('/', req.url));
  }

  // Redirect authenticated users from `/` to `/dashboard`
  if (sessionToken && pathname === '/') {
    return NextResponse.redirect(new URL('/dashboard', req.url));
  }

  // Redirect unauthenticated users to `/signup` if they try to access routes other than `/`, `/signup`, or `/login`
  if (!sessionToken && pathname !== '/' && pathname !== '/signup' && pathname !== '/login') {
    return NextResponse.redirect(new URL('/', req.url));
  }

  // Redirect users trying to access `/login` to `/`
  if (pathname === '/login') {
    return NextResponse.redirect(new URL('/', req.url));
  }

  // Handle non-existent routes by redirecting to a 404 page
  if (pathname !== '/' && pathname !== '/signup' && pathname !== '/login' && !pathname.startsWith('/api') && !pathname.startsWith('/static') && !pathname.startsWith('/favicon.ico')) {
    return NextResponse.redirect(new URL('/404', req.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: [
    '/',
    '/signup',
    '/login',
    '/((?!api|static|favicon.ico).*)' // Matches all routes except APIs and static files
  ],
};
