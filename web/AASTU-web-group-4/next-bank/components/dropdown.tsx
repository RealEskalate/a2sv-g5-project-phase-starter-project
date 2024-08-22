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
const Modal: FC<{ message: string }> = ({ message }) => (
  <div style={modalStyle} className="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-50 z-50">
    <div className="bg-white p-6 rounded-lg shadow-lg">
      <h2 className="text-lg font-semibold">{message}</h2>
    </div>
  </div>
);


export function DropdownMenuDemo() {
  const [showModal, setShowModal] = useState(false);

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
                width={30}
                height={30}
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
        {showModal && (
        <Modal
          message="We're sad to see you go! Come back again!"
          onClose={() => setShowModal(false)} // Optional close function
        />
      )}
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
