'use client';
import { useSession } from "next-auth/react";
import Investments from "./Investments/page";
import Signin from "./auth/signin/page";

export default function Home() {
  const { status } = useSession();

  return (
    <div className="p-4">
      {status === "authenticated" ? <Investments /> : <Signin />}
    </div>
  );
}
