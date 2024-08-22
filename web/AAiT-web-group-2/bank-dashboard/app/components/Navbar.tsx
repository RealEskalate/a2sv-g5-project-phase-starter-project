"use client"
import React from 'react'
import Link from 'next/link'
import { usePathname } from 'next/navigation'



const Navbar = () => {
  const path = usePathname()
  return (
    <div className="flex justify-between items-center  h-24 px-3 ">
      <div className='flex gap-[50px] items-center'>
        <p className='font-[600] text-[28px] text-[#343C6A]'>{path}</p>
      </div>
      <div className='flex gap-5 items-center'>
        <div className='flex hidden  sm:flex md:gap-1 lg:gap-5'>
          <label className='flex items-center gap-3 h-[40px] bg-[#F5F7FA] rounded-3xl cursor-pointer px-5'>
              <img src="/search.png" alt="" />
              <input className="bg-inherit p-1 focus:outline-none " type="search" placeholder='search for something' />
          </label>
          <Link  href="/settings" className="h-[40px] w-[40px] flex items-center justify-center rounded-full bg-[#F5F7FA]">
              <img className='w-6 h-6'  src="/settings.png" alt="" />
          </Link>
          <Link  href='/notifications' className="h-[40px] w-[40px] flex items-center justify-center rounded-full bg-[#F5F7FA]">
              <img className='w-6 h-6'  src="/notification.png" alt="" />
          </Link>
        </div>
        <Link className='w-11' href='/settings/editProfile'>        
          <img src="/profile.png" alt="" />
        </Link>
      </div>
    </div>
  )
}

export default Navbar
