import {jwtDecode, JwtPayload } from "jwt-decode";
import { getSession } from "next-auth/react";
import { refreshToken } from "@/lib/api/authenticationController";
type Data = {
  access_token: string;
  data: string;
  refresh_token: string;
};

type SessionDataType = {
  user: Data;
};

export default async function Refresh(): Promise<string> {
  const session = (await getSession()) as SessionDataType | null
  if (!session?.user?.access_token || !session?.user?.refresh_token) {
    return ""
  }
  const accessToken = session.user.access_token;
  const refreshTokenValue = session.user.refresh_token;

  const decodedToken = jwtDecode<JwtPayload>(accessToken);
  const currentTime = Date.now() / 1000;
  const expiry = decodedToken.exp;
  console.log(accessToken, "access token");
  
  return accessToken
}
