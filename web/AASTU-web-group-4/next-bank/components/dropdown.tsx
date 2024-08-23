'use client';
import { Button } from "@/components/ui/button";
import { currentuser } from "@/services/userupdate";
import { FC, useEffect, useState } from "react";
import { signOut } from "next-auth/react";
import Cookies from "js-cookie";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { DialogDemo } from "./modal";

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuPortal,
  DropdownMenuSeparator,
  DropdownMenuShortcut,
  DropdownMenuSub,
  DropdownMenuSubContent,
  DropdownMenuSubTrigger,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import Image from "next/image";

import {
  UserCircleIcon,
  PencilSquareIcon,
  Cog6ToothIcon,
  QuestionMarkCircleIcon,
  ArrowLeftStartOnRectangleIcon,
} from "@heroicons/react/24/outline";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "./ui/dialog";
import { Label } from "recharts";

const modalStyle = {
  animation: "fadeIn 0.5s ease-in-out",
};

export function DropdownMenuDemo() {
  const router = useRouter();
  const handleSignOut = async () => {
    Cookies.remove("accessToken");
    // Call signOut to handle session termination
    router.push("/signin"); // Redirect after sign-out
  };
  const [info, setinfo] = useState([]);

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
  } , []);

 
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
            <DropdownMenuLabel className="text-xl text-gray-800 font-bold  text-center">{info.name}</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuGroup>
              <DialogTrigger asChild  > 
                <DropdownMenuItem>
                  <div className="flex  gap-4 items-center">
                    <UserCircleIcon className="h-5 w-5"/>
                    Profile
                  </div>
      
                </DropdownMenuItem>
              </DialogTrigger>
              <Link href="../setting" passHref>
                <DropdownMenuItem>
                  <div className="flex  gap-4 items-center">
                    <Cog6ToothIcon className="h-5 w-5" />
                    Settings
                  </div>
      
                </DropdownMenuItem>{" "}
              </Link>
              <DropdownMenuItem>
                <div className="flex  gap-4 items-center">
                  <QuestionMarkCircleIcon className="h-5 w-5" />
                  Help
                </div>
      
              </DropdownMenuItem>
            </DropdownMenuGroup>
            <DropdownMenuSeparator />
            <DropdownMenuItem onClick={handleSignOut}>
              <div className="flex  gap-4 items-center">
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
