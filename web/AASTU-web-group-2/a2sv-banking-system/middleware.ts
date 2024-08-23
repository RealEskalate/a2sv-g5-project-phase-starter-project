import { NextRequest, NextResponse } from "next/server";

export async function middleware(request: NextRequest) {
  const sessionToken = request.cookies.get("next-auth.session-token")?.value;

  console.log("Session Token:", sessionToken);

  if (!sessionToken) {
    // Redirect to sign-in if no session token
    return NextResponse.redirect(new URL('/api/auth/signin', request.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: [
    '/', 
    '/dashboard',
    '/transaction', 
    '/accounts', 
    '/creditCards', 
    '/investments', 
    '/loans', 
    '/bankingServices', 
    '/bankingSettings'
  ],
};
