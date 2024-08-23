import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { getToken } from "next-auth/jwt";

export async function middleware(req: NextRequest) {
  const publicPaths = ["/login", "/signup"];
  const token = await getToken({ req, secret: process.env.NEXTAUTH_SECRET });

  // console.log('Token:', token);

  if (publicPaths.includes(req.nextUrl.pathname) || token) {
    return NextResponse.next();
  } else {
    const signInUrl = new URL("/login", req.url);
    return NextResponse.redirect(signInUrl);
  }
}

export const config = {
  matcher: ["/page-not-found"],
  // console.log("middleware")
};
