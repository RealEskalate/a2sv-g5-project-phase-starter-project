import { jwtDecode, JwtPayload } from "jwt-decode";
import { getSession } from "next-auth/react";
import update, { AuthOptions } from "next-auth";
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
  const session = (await getSession()) as SessionDataType | null;
  if (!session?.user?.access_token || !session?.user?.refresh_token) {
    return "";
  }
  const accessToken = session.user.access_token;
  const refreshTokenValue = session.user.refresh_token;

  const decodedToken = jwtDecode<JwtPayload>(accessToken);
  const currentTime = Date.now() / 1000;
  const expiry = decodedToken.exp;

  if (expiry && expiry < currentTime) {
    const newTokens = await refreshToken(refreshTokenValue);

    // Update the session with the new tokens
    await update({
      // Specify the user property using the correct type
      user: {
        ...session.user,
        access_token: newTokens.access_token,
        refresh_token: newTokens.refresh_token,
      },
      providers: [], // Add an empty array or provide the required providers
    } as AuthOptions);

    return newTokens.access_token;
  } else {
    return accessToken;
  }
}
