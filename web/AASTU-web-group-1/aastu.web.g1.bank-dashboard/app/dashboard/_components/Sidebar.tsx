"use client";
import React, { useEffect, useState } from "react";
import { sidebarLinks } from "@/constants";
import Link from "next/link";
import Image from "next/image";
import { usePathname } from "next/navigation";
import { cn } from "@/lib/utils";
import { useUser } from "@/contexts/UserContext";

import { getSession, signOut } from "next-auth/react";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import ky from "ky";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";

const Sidebar = () => {
  const [profileUrl, setProfileUrl] = useState("");
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const { isDarkMode } = useUser(); // Get the dark mode state

  useEffect(() => {
    const fetchUser = async () => {
      const session = await getSession();
      const accessToken = session?.user.accessToken;
      setIsLoading(true);
      if (!accessToken) {
        throw new Error("No access token found");
      }

      try {
        const res: any = await ky(
          `${process.env.NEXT_PUBLIC_BASE_URL}/user/current`,
          {
            headers: {
              Authorization: `Bearer ${accessToken}`,
            },
            timeout: 10000,
          }
        ).json();

        setProfileUrl(res.data.profilePicture);
        setName(res.data.name);
        setEmail(res.data.email);
      } catch (error) {
        console.error("Failed to fetch user:", error);
      } finally {
        setIsLoading(false);
      }
    };

    fetchUser();
  }, []);

  const pathname = usePathname();

  return (
    <div
      className={cn(
        "sticky left-0 top-0 h-screen border-r pt-4 max-md:hidden sm:p-2 xl:p-4 2xl:w-[300px] flex flex-col justify-between",
        isDarkMode
          ? "bg-gray-800 border-gray-700 text-white"
          : "bg-white border-gray-200 text-black"
      )}
    >
      <div>
        <div className="flex items-center gap-2 p-3 pb-8">
          <Image src="/icons/logo.png" width={25} height={25} alt="logo" />
          <h1
            className={cn(
              "font-[900] text-[1.5rem]",
              isDarkMode ? "text-white" : "text-primaryBlack"
            )}
          >
            BankDash.
          </h1>
        </div>
        <div className="flex flex-col gap-2">
          {sidebarLinks.map((link, index) => {
            const isActive =
              pathname === link.route ||
              pathname.startsWith(`dashboard/${link.route}/`);
            return (
              <Link
                href={link.route}
                key={link.title}
                className={cn(
                  "flex gap-6 items-center py-1 md:p-3 2xl:px-4 pl-0 justify-center xl:justify-start",
                  {
                    "bg-nav-focus": isActive && !isDarkMode,
                    "": isActive && isDarkMode,
                    "text-gray-400": !isActive && isDarkMode,
                  }
                )}
              >
                <div
                  className={`${
                    isActive ? "visible" : "hidden"
                  } flex w-6 h-[55px] rounded-[32px] bg-[#1814F3] absolute left-[-20px]`}
                ></div>
                <Image
                  src={link.icon}
                  alt={link.title}
                  width={20}
                  height={20}
                  className={cn({
                    "filter-custom-blue": isActive,
                    "filter-custom-white": !isActive && isDarkMode,
                  })}
                />
                <p
                  className={cn("text-sm font-semibold", {
                    "text-primaryBlue": isActive && !isDarkMode,
                    "text-blue-500": isActive && isDarkMode,
                    "text-[#B1B1B1]": !isActive && !isDarkMode,
                  })}
                >
                  {link.title}
                </p>
              </Link>
            );
          })}
        </div>
      </div>
      <div
        className={cn(
          "cursor-pointer rounded-full p-1 w-[70%] transition-all duration-500",
          isDarkMode ? "hover:bg-gray-800" : "hover:bg-neutral-200"
        )}
      >
        <Popover>
          <PopoverTrigger className="flex items-center gap-4">
            <Avatar>
              <AvatarImage src={profileUrl} />
              <AvatarFallback>CN</AvatarFallback>
            </Avatar>

            {isLoading ? (
              "Loading..."
            ) : (
              <div className="flex gap-4">
                <div className="space-y-1 text-left">
                  <h4 className="text-sm font-medium">{name}</h4>
                </div>
              </div>
            )}
          </PopoverTrigger>

          <PopoverContent
            className={cn(
              "w-80 p-2",
              isDarkMode ? "bg-gray-800 text-white" : "bg-white text-black"
            )}
          >
            <div className="mt-4 space-y-2">
              <Link
                href="/dashboard/setting/"
                className={cn(
                  "flex items-center gap-2 rounded-md px-3 py-2 text-sm font-medium",
                  isDarkMode ? "hover:bg-gray-700" : "hover:bg-muted"
                )}
                prefetch={false}
              >
                <FilePenIcon className="h-4 w-4" />
                Edit Profile
              </Link>
              <button
                onClick={() => signOut({ callbackUrl: "/auth/sign-in" })}
                className={cn(
                  "flex items-center gap-2 rounded-md px-3 py-2 text-sm font-medium",
                  isDarkMode ? "hover:bg-gray-700 text-white" : "hover:bg-muted"
                )}
              >
                <LogOutIcon className="h-4 w-4" />
                Log Out
              </button>
            </div>
          </PopoverContent>
        </Popover>
      </div>
    </div>
  );
};

export default Sidebar;

function FilePenIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M12 22h6a2 2 0 0 0 2-2V7l-5-5H6a2 2 0 0 0-2 2v10" />
      <path d="M14 2v4a2 2 0 0 0 2 2h4" />
      <path d="M10.4 12.6a2 2 0 1 1 3 3L8 21l-4 1 1-4Z" />
    </svg>
  );
}

function LogOutIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4" />
      <polyline points="16 17 21 12 16 7" />
      <line x1="21" x2="9" y1="12" y2="12" />
    </svg>
  );
}
