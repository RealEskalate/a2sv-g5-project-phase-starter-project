import NextAuth from "next-auth";
import { options } from "./options";

export interface Userx {
  id: string;
  refreshToken: string;
  accessToken: string;
}

const handler = NextAuth(options);

export { handler as GET, handler as POST };
