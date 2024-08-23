import { NextRequest, NextResponse } from "next/server";
import Refresh from "./app/api/auth/[...nextauth]/token/RefreshToken";

export async function middleware(request: NextRequest) {
  // Check if the session token exists in the cookies
  const sessionToken = request.cookies.get("next-auth.session-token")?.value;

  // If no session token, try to refresh it
  if (!sessionToken) {
    const newAccessToken = await Refresh();
    
    // Debugging: Log the new access token
    console.log("New Access Token:", newAccessToken);
    
    // If no new access token, redirect to the sign-in page
    if (!newAccessToken) {
      return NextResponse.redirect(new URL('/api/auth/signin', request.url));
    }

    // If a new access token is obtained, proceed and set it in the cookies
    const response = NextResponse.next();
    response.cookies.set("accessToken", newAccessToken);
    return response;
  }

  // If a session token exists, allow the request to proceed
  return NextResponse.next();
}

// Define the paths where this middleware should be applied
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
