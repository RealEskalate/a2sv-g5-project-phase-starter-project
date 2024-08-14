import { Button } from "@/components/ui/button";
import Link from "next/link";
import React from "react";

const Home = () => {
  return (
    <div className="flex flex-col gap-2 items-center justify-center h-[100vh]">
      <Link href="/dashboard">
        <Button>Go to Dashboard</Button>
      </Link>

      <Link href="/auth/sign-in">
        <Button>Go to Auth</Button>
      </Link>
    </div>
  );
};

export default Home;
