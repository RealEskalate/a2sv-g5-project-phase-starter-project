"use client"
import SettingsPage from "@/components/ui/Settings/SettingsPage";
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
    <div className="bg-primary-color-50 h-full w-9/10 flex justify-center">
        <SettingsPage />
    </div>
  );
};

export default Page;
