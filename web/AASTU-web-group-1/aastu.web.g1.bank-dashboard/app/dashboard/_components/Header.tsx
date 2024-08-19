import React from "react";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import Image from "next/image";
import { signOut } from "next-auth/react";

import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import Link from "next/link";
import { Separator } from "@radix-ui/react-select";

const Header = ({ title }: { title: string }) => {
  return (
    <div className="bg-white max-md:hidden">
      <div className="flex justify-between px-10 py-4">
        <h1 className="text-3xl text-primaryBlack font-[600]">{title}</h1>

        <div className="flex gap-5 items-center">
          {/* Search */}
          <div className="flex gap-3 bg-[#F5F7FA] p-3 rounded-full">
            <Image
              src="/icons/Search.svg"
              width={20}
              height={20}
              alt="Search"
            />
            <input
              className="outline-none bg-[#F5F7FA]"
              type="text"
              placeholder="Search for something"
            />
          </div>

          {/* Settings */}
          <div className="bg-[#F5F7FA] p-2 rounded-full cursor-pointer">
            <Image
              src="/icons/Settings.svg"
              width={22}
              height={22}
              alt="Notif"
            />
          </div>

          <div className="bg-[#F5F7FA] p-2 rounded-full cursor-pointer">
            {/* Notif */}
            <Image
              src="/icons/Notification.svg"
              width={22}
              height={22}
              alt="Notif"
            />
          </div>
          {/* Avatar */}

          <Popover>
            <PopoverTrigger>
              <Avatar>
                <AvatarImage src="https://github.com/shadcn.png" />
                <AvatarFallback>CN</AvatarFallback>
              </Avatar>
            </PopoverTrigger>
            <PopoverContent className="w-80 p-4">
              <div className="flex items-center gap-4 border-b pb-4">
                <div className="space-y-1">
                  <h4 className="text-lg font-medium">John Doe</h4>
                  <p className="text-sm text-muted-foreground">
                    john@example.com
                  </p>
                </div>
              </div>
              <div className="mt-4 space-y-2">
                <Link
                  href="/dashboard/setting/"
                  className="flex items-center gap-2 rounded-md px-3 py-2 text-sm font-medium hover:bg-muted"
                  prefetch={false}
                >
                  <UserIcon className="h-4 w-4" />
                  View Profile
                </Link>
                <Link
                  href="/dashboard/setting/"
                  className="flex items-center gap-2 rounded-md px-3 py-2 text-sm font-medium hover:bg-muted"
                  prefetch={false}
                >
                  <FilePenIcon className="h-4 w-4" />
                  Edit Profile
                </Link>
                <Link
                  href="/dashboard/setting/"
                  className="flex items-center gap-2 rounded-md px-3 py-2 text-sm font-medium hover:bg-muted"
                  prefetch={false}
                >
                  <SettingsIcon className="h-4 w-4" />
                  Account Settings
                </Link>
                <Separator />
                <button
                  onClick={() => signOut({ callbackUrl: "/auth/sign-in" })}
                  className="flex items-center gap-2 rounded-md px-3 py-2 text-sm font-medium hover:bg-muted"
                >
                  <LogOutIcon className="h-4 w-4" />
                  Log Out
                </button>
              </div>
            </PopoverContent>
          </Popover>
        </div>
      </div>
    </div>
  );
};

export default Header;
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

function SettingsIcon(props: React.SVGProps<SVGSVGElement>) {
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
      <path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z" />
      <circle cx="12" cy="12" r="3" />
    </svg>
  );
}

function UserIcon(props: React.SVGProps<SVGSVGElement>) {
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
      <path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2" />
      <circle cx="12" cy="7" r="4" />
    </svg>
  );
}
