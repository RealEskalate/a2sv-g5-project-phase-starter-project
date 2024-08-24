import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { getToken } from "next-auth/jwt";

export async function middleware(req: NextRequest) {
  const token = await getToken({ req, secret: process.env.NEXTAUTH_SECRET });

  if (token) {
    return NextResponse.next();
  } else {
    const signInUrl = new URL("/login", req.url);
    return NextResponse.redirect(signInUrl);
  }
}

export const config = {
  matcher: ["/settings"], // Protect all routes except login and signup
};
