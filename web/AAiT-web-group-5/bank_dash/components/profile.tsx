import React from 'react'
import { Avatar, AvatarImage, AvatarFallback } from './ui/avatar'
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
  } from "@/components/ui/dropdown-menu"
  

export const Profile = () => {
  return (
    <div>

    <DropdownMenu>
        <DropdownMenuTrigger>
            <Avatar>
                <AvatarImage src="https://avatar.iran.liara.run/public/15" />
            </Avatar>
        </DropdownMenuTrigger>
        <DropdownMenuContent>
            <DropdownMenuLabel>My Account</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuItem className='text-red-500'>LogOut</DropdownMenuItem>
            
        </DropdownMenuContent>
    </DropdownMenu>

        

    </div>
  )
}
