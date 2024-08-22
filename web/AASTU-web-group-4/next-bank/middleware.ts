import { NextRequest, NextResponse } from "next/server";

export async function middleware(request: NextRequest) {
  // Access the cookies from the request
  const accessToken = request.cookies.get('accessToken');
  const refreshToken = request.cookies.get('refreshToken');

  // Log the cookies for debugging purposes
  // console.log('Request Cookies:', { accessToken, refreshToken });

  // Check if the access token exists
  if (!accessToken) {
    // Redirect to the sign-in page if the access token is missing
    return NextResponse.redirect(new URL('/signin', request.url));
  }

  // Optionally, you can add further validation for the token (e.g., expiration check)

  // Allow the request to proceed if the access token exists
  return NextResponse.next();
}

// Define the paths where this middleware should be applied
export const config = {
  matcher: [
    
  ],
};
