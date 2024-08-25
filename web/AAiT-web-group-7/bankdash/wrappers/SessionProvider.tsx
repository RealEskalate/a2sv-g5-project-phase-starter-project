"use client";

import { type ReactNode } from "react";
import { SessionProvider } from "next-auth/react";

interface props {
  children: ReactNode;
}

const AuthProvider: React.FC<props> = ({ children }) => {
  return <SessionProvider>{children}</SessionProvider>;
};

export default AuthProvider;
