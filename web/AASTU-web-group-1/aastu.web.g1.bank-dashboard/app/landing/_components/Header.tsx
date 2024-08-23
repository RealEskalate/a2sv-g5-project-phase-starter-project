import { Button } from "@/components/ui/button";
import { useUser } from "@/contexts/UserContext";
import { cn } from "@/lib/utils";
import Image from "next/image";
import Link from "next/link";
import React from "react";

const Header = () => {
  return (
    <div className="flex justify-between sticky top-0 px-12 py-6 items-center backdrop-filter backdrop-blur-lg bg-white/50  z-10">
      <div className="flex items-center justify-center">
        <div className="flex items-center gap-2">
          <Image src="/icons/logo.png" width={25} height={25} alt="logo" />
          <h1 className={"font-[900] text-[1.5rem] text-primaryBlack"}>
            BankDash.
          </h1>
        </div>
      </div>

      <div className="flex gap-8 font-normal text-lg">
        <Link href="/dashboard">Dashboard</Link>
        <Link href="/dashboard">Team</Link>
        <Link href="/dashboard">Community</Link>
        <Link href="/dashboard">About</Link>
      </div>

      <div className="flex gap-8">
        <Button className="rounded-full" variant="outline">
          <Link href="/auth/sign-in">Sign In</Link>
        </Button>
        <Button className="px-8 rounded-full bg-[#343C6A]">
          <Link href="/auth/sign-up">Open Account</Link>
        </Button>
      </div>
    </div>
  );
};

export default Header;
