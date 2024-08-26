import { NextRequest, NextResponse } from "next/server";

export async function middleware(request: NextRequest) {
  // Access the cookies from the request
  const accessToken = request.cookies.get('accessToken');

  // // Redirect to /home if accessing the root path
  // if (request.nextUrl.pathname === '/') {
  //   return NextResponse.redirect(new URL('/home', request.url));
  // }

  // Redirect to /signin if the access token is missing
  if (!accessToken) {
    return NextResponse.redirect(new URL('/home', request.url));
  }

  // Allow the request to proceed if the access token exists
  return NextResponse.next();
}

// Define the paths where this middleware should be applied
export const config = {
  matcher: [
    '/', 
    '/transaction', 
    '/accounts', 
    '/credit-card', 
    '/transfer', 
    '/investments', 
    '/loans', 
    '/services', 
    '/setting'
  ],
};
