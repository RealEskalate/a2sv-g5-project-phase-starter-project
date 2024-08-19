import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";
import { getToken } from "next-auth/jwt";

export async function middleware(req: NextRequest) {
  const { pathname } = req.nextUrl;

  if (pathname.startsWith("/auth") || pathname.startsWith("/api")) {
    return NextResponse.next();
  }

  const excludedPaths = [
    "/images/",
    "/icons/",
    "/_next/static/",
    "/_next/image/",
    "/favicon.ico",
];

if (excludedPaths.some(path => pathname.startsWith(path))) {
    return NextResponse.next();
}

  const token = await getToken({ req, secret: process.env.NEXTAUTH_SECRET });
  console.log("Token", token)

  if (!token) {
    const url = req.nextUrl.clone();
    url.pathname = "/auth/sign-in";
    return NextResponse.redirect(url);
  }

  return NextResponse.next();
}

export const config = {
    matcher: [
        "/((?!api|_next/static|_next/image|favicon.ico|auth).*)"
    ],
};