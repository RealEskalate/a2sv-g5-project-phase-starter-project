'use client'
import { HeadersTitle } from '@/constants/Layout/lay_constant';
import Image from 'next/image';
import { usePathname } from 'next/navigation';
import React from 'react'

type Props = {}

export default function Header({ }: Props) {
    const currentPath = usePathname();
    return (
        <div className='flex justify-between items-center h-[100px] px-[30px] md:px-[40px]'>
            <div className="headerTitle">
                <h1 className='text-[28px] text-[#343C6A]'>{HeadersTitle(currentPath)}</h1>
            </div>
            <div className="headerRight flex gap-x-[35px] items-center ">
                <div className='md:flex gap-x-[30px] hidden '>
                    <div className='flex gap-x-4 pl-5 bg-[#F5F7FA] rounded-[40px]  py-4 max-w-[255px] overflow-hidden' >
                        <Image src='/icons/search.svg' alt="Search icon" width={20} height={20}  />
                        <input type="text" placeholder='Search for something' className='outline-none bg-[#F5F7FA] max-w-fit overflow-hidden' />
                    </div>
                    <div className="setting w-[50px] h-[50px] rounded-full bg-[#F5F7FA] flex justify-center items-center">
                        <Image src='/icons/header_settings.svg' alt="Setting icon" width={25} height={25} />
                    </div>
                    <div className="notifiacation w-[50px] h-[50px] rounded-full bg-[#F5F7FA] flex justify-center items-center">
                        <Image src='/icons/notification.svg' alt="Setting icon" width={25} height={25} />
                    </div>
                </div>
                <div className="profile flex-shrink-0">
                    <Image src='/images/profile.png' alt="Setting icon" width={60} height={60} className='rounded-full' />
                </div>
            </div>
        </div>
    )
}