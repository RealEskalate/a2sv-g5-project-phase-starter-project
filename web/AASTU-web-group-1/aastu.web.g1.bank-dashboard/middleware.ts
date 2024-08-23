import { getToken } from "next-auth/jwt";
import type { NextRequest } from "next/server";
import { NextResponse } from "next/server";

export async function middleware(req: NextRequest) {
  const { pathname } = req.nextUrl;

  if (pathname.startsWith("/auth") || pathname.startsWith("/api") || pathname.startsWith("/landing")) {
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
  // const url = new URL("/path/to/another/page", req.nextUrl.origin);
  // url.searchParams.append("token", token);
  // return NextResponse.redirect(url);

  return NextResponse.next();
}

export const config = {
    matcher: [
        "/((?!api|_next/static|_next/image|favicon.ico|auth).*)"
    ],
};