"use client"
import React from "react";
import { useSession } from "next-auth/react";
import { redirect } from "next/navigation";

const Page = () => {
  const { data: session } = useSession({
    required: true,
    onUnauthenticated() {
      redirect("/api/auth/signin?calbackUrl=/login");
    },
  });
  return (
    <>
      <h1 className="text-2xl">Transaction Page</h1>
    </>
  );
};

export default Page;
