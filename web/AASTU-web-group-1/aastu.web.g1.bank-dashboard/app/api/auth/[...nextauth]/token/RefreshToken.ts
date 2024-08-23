// import { refreshAccessToken } from "@/lib/auth";
// import {jwtDecode, JwtPayload } from "jwt-decode";
// import { getSession } from "next-auth/react";

// type Data = {
//   accessToken: string;
//   data: string;
//   refreshToken: string;
// };

// type SessionDataType = {
//   user: Data;
// };

// export default async function Refresh(): Promise<string> {
//   const session = (await getSession()) as SessionDataType | null
//   console.log("session", session)
//   if (!session?.user?.accessToken || !session?.user?.refreshToken) {
//     ("Fall back")
//     return ""
//   }
//   const accessToken = session.user.accessToken;
//   const refreshTokenValue = session.user.refreshToken;

//   const decodedToken = jwtDecode<JwtPayload>(accessToken);
//   const currentTime = Date.now() / 1000;
//   const expiry = decodedToken.exp;

//   if (expiry && expiry < currentTime) {
//     console.log(refreshTokenValue)

//     const newAccessToken = await refreshAccessToken(refreshTokenValue) as string;
//     console.log("New AT generated", newAccessToken)
//     return newAccessToken;
//   } else {
//     return accessToken;
//   }
// }
