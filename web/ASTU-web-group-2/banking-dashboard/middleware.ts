import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { getToken } from "next-auth/jwt";
import { isTokenExpired } from "./utils/authUtils"; // Adjust the path to where your utility function is located

export async function middleware(req: NextRequest) {
  const sessionToken = await getToken({ req, secret: process.env.AUTH_SECRET });
  console.log("sessionTokensessionToken", sessionToken);
  const { pathname } = req.nextUrl;

  // Redirect authenticated users from `/` to `/dashboard`
  if (sessionToken && pathname === "/") {
    return NextResponse.redirect(new URL("/dashboard", req.url));
  }

  // Redirect users trying to access `/login` to `/` if they are already authenticated
  if (sessionToken && pathname === "/login") {
    return NextResponse.redirect(new URL("/", req.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: [
    "/",
    "/signup",
    "/login",
    "/((?!api|static|favicon.ico).*)", // Matches all routes except APIs and static files
  ],
};
