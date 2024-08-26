'use client';
import { FC, useEffect, useState } from "react";
import Image from "next/image";
import { signOut } from "next-auth/react";
import Cookies from "js-cookie";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { DropdownMenu, DropdownMenuContent, DropdownMenuGroup, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from "@/components/ui/dropdown-menu";
import { UserCircleIcon, Cog6ToothIcon, QuestionMarkCircleIcon, ArrowLeftStartOnRectangleIcon } from "@heroicons/react/24/outline";
import NotificationBell from './NotificationBell';
import { DialogDemo } from "./modal";
import { useNotifications } from '@/services/NotificationContext';
import { MdDoneAll } from 'react-icons/md'; 
import { UserData } from "@/types/index";
import { Dialog, DialogTrigger } from "@radix-ui/react-dialog";
import { currentuser } from "@/services/userupdate";
import ThemeSwitch from "./ThemeSwitch";

export function DropdownMenuDemo() {
  const router = useRouter();
  const handleSignOut = async () => {
    Cookies.remove("accessToken");
    router.push("/home"); // Redirect after sign-out
  };

  const [info, setinfo] = useState<UserData>();
  const [isSmallScreen, setIsSmallScreen] = useState<boolean>(false);
  const { notifications, markAllAsRead, unreadCount } = useNotifications();

  useEffect(() => {
    const fetch = async () => {
      try {
        const data = await currentuser();
        setinfo(data.data || []);
      } catch (error) {
        console.error("Error:", error);
      }
    };
    fetch();

    // Listen to window resize events to determine screen size
    const handleResize = () => {
      setIsSmallScreen(window.innerWidth < 768); // Set to true if screen width is less than 768px
    };
    
    handleResize(); // Set initial state
    window.addEventListener('resize', handleResize);
    return () => window.removeEventListener('resize', handleResize);
  }, []);

  return (
    <div className="">
      <Dialog>
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Image
              src="/Images/profilepic.jpeg"
              alt="User Profile"
              width={40}
              height={40}
              className="rounded-full aspect-square object-cover cursor-pointer"
            />
          </DropdownMenuTrigger>
          <DropdownMenuContent className="w-56">
            <DropdownMenuLabel className="text-xl text-gray-800 font-bold dark:text-white text-center">{info?.name}</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuGroup>
              {isSmallScreen && (
                <DropdownMenuItem>
                  <div className="flex gap-5 items-center">
                    <NotificationBell/>
                    <ThemeSwitch/>
                  </div>
                </DropdownMenuItem>
              )}
              <DialogDemo />
              <DialogTrigger asChild>
                <DropdownMenuItem>
                  <div className="flex gap-4 items-center">
                    <UserCircleIcon className="h-5 w-5" />
                    Profile
                  </div>
                </DropdownMenuItem>
              </DialogTrigger>
              <Link href="../setting" passHref>
                <DropdownMenuItem>
                  <div className="flex gap-4 items-center">
                    <Cog6ToothIcon className="h-5 w-5" />
                    Settings
                  </div>
                </DropdownMenuItem>
              </Link>
              <DropdownMenuItem>
                <div className="flex gap-4 items-center">
                  <QuestionMarkCircleIcon className="h-5 w-5" />
                  Help
                </div>
              </DropdownMenuItem>
            </DropdownMenuGroup>
            <DropdownMenuSeparator />
            <DropdownMenuItem onClick={handleSignOut}>
              <div className="flex gap-4 items-center">
                <ArrowLeftStartOnRectangleIcon className="h-5 w-5 text-red-600" />
                <p className="text-red-600">LogOut</p>
              </div>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
        <DialogDemo />
      </Dialog>
    </div>
  );
}
