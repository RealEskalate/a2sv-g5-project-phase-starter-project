import {jwtDecode} from "jwt-decode";
import AuthService from "@/app/Services/api/authService";
import { getSession, signOut } from "next-auth/react"; 

interface DecodedToken {
  exp: number; 
}

export const checkAndRefreshToken = async (): Promise<string | null> => {

  const session = await getSession();
  

  if (!session || !session.accessToken || !session.refreshToken) {
    console.error("No session or access token found.");
    return null;
  }

  const decodedToken = jwtDecode<DecodedToken>(session.accessToken);
  const currentTime = Math.floor(Date.now() / 1000);
  console.log("DECODED TOKEN",decodedToken)

  if (decodedToken.exp < currentTime) {
    try {
      console.log("Token has expired, refreshing...",);
      const response:any = await AuthService.refreshToken(session.refreshToken);
      // console.log("Response",response)
      // console.log("RefreshToken",session.refreshToken)
      // console.log("BEFORE",session)

      if (response.success && response.data) {
        const newAccessToken = response.data.access_token; 
        const newRefreshToken = response.data.refresh_token
        
        // await updateSession({
        //   ...session,
        //   refreshToken: newRefreshToken,
        //   accessToken: newAccessToken,
        // });
        // console.log("AFTER",session)

        console.log("Token refreshed successfully.");
        return newAccessToken;
      } else {
        console.error("Failed to refresh token: Invalid response format.",);
        
        return null;
      }
    } catch (error) {
      console.error("Failed to refresh token:", error);
      return null;
    }
  } else {
    console.log("Token is still valid.");
    return session.accessToken;
  }
};
