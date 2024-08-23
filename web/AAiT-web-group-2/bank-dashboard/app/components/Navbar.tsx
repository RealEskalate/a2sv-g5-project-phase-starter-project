"use client"
import React from 'react'
import Link from 'next/link'
import { usePathname } from 'next/navigation'
// import { FiMenu } from "react-icons/fi";



const Navbar = ({onMenuClick}: {onMenuClick: () => void}) => {
  const path = usePathname();
  let title = path.split("/")[1] || ''; // Use an empty string if there is no second segment

  if (title) {
    title = title[0].toUpperCase().concat(title.slice(1));
  }

  return (
    <div className="flex justify-between items-center  h-24 px-3 ">
      <div className='text-2xl sm:hidden '>
        {/* <FiMenu onClick={onMenuClick} /> */}

      </div>

      <div className='flex gap-[50px] items-center'>
        <p className='text-wrap font-[600] text-[28px] text-[#343C6A]'>{path.slice(1, 2).toUpperCase() + path.slice(2)}</p>
      </div>
      <div className='flex gap-2  md:gap-2 lg:gap-5 items-center'>
        <div className=' hidden  sm:flex sm:gap-1  md:gap-2  lg:gap-3 '>
          <label className='flex items-center gap-3 h-[40px] bg-[#F5F7FA] rounded-3xl cursor-pointer px-5'>
              <img src="/search.png" alt="" />
              <input className="bg-inherit w-20 md:w-full p-1 focus:outline-none " type="search" placeholder='search for something' />
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
