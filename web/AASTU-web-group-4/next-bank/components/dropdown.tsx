import { Button } from "@/components/ui/button"
import { currentuser } from "@/services/userupdate";
import {FC, useState } from "react";
import { signOut } from 'next-auth/react';
import Cookies from 'js-cookie';
import { useRouter } from 'next/navigation';
import Link from "next/link";

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
} from "@/components/ui/dropdown-menu"
import Image from 'next/image';

import {UserCircleIcon , PencilSquareIcon , Cog6ToothIcon , QuestionMarkCircleIcon , ArrowLeftStartOnRectangleIcon } from "@heroicons/react/24/outline";

const modalStyle = {
  animation: 'fadeIn 0.5s ease-in-out',
};



export function DropdownMenuDemo() {
 
  const router = useRouter();
  const handleSignOut = async () => {
    Cookies.remove('accessToken');
     // Call signOut to handle session termination
    router.push('/signin'); // Redirect after sign-out
  };
  const [name , setname ] = useState("")

const fetch = async () => {

    try {
      const data = await currentuser();
      setname(data.data.name)
    } catch (error) {
      console.error("Error:", error);
  }
}
  return (
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
        <DropdownMenuLabel>{name}</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuGroup>
          <DropdownMenuItem>
          <div className="flex  gap-4 items-center">
            <UserCircleIcon className="h-5 w-5"/>
              Profile
          </div>
            {/* <DropdownMenuShortcut>⇧⌘P</DropdownMenuShortcut> */}
          </DropdownMenuItem>
          <Link href="../setting" passHref>
          <DropdownMenuItem>
          <div className="flex  gap-4 items-center">
            <Cog6ToothIcon className="h-5 w-5"/>
              Settings
          </div>
            {/* <DropdownMenuShortcut>⌘S</DropdownMenuShortcut> */}
          </DropdownMenuItem> </Link>
          <DropdownMenuItem>
          <div className="flex  gap-4 items-center">
            <QuestionMarkCircleIcon className="h-5 w-5"/>
              Help
          </div>
            {/* <DropdownMenuShortcut>⌘K</DropdownMenuShortcut> */}
          </DropdownMenuItem>
        </DropdownMenuGroup>
        <DropdownMenuSeparator />
        <DropdownMenuItem onClick={handleSignOut}>
        <div className="flex  gap-4 items-center">
            <ArrowLeftStartOnRectangleIcon className="h-5 w-5 text-red-600"/>
              <p className="text-red-600">LogOut</p>
          </div>
          {/* <DropdownMenuShortcut>⇧⌘Q</DropdownMenuShortcut> */}
        </DropdownMenuItem>
        
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
