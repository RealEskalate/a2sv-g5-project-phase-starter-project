"use client";
import { useEffect } from "react";
import { useRouter } from "next/navigation";

import { getServerSession } from "next-auth";
import Image from "next/image";
import { authOptions } from "../api/auth/[...nextauth]/options";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <main className="flex min-h-screen w-full font-inter">
      {children}
      <div className="flex h-screen w-full sticky top-0 items-center justify-end bg-sky-1 max-lg:hidden">
        <Image
          src="/icons/AuthPage.jpg "
          alt="Auth image"
          width={500}
          height={500}
          className="object-contain"
        />
      </div>
    </main>
  );
}
