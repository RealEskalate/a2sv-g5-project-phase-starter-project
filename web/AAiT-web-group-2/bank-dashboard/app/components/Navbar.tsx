import React from 'react'
import Link from 'next/link'

const Navbar = () => {
  return (
    <div className="flex justify-between items-center px-10 h-[100px]">
      <div className='flex gap-[50px] items-center'>
        <div className="flex gap-2 items-center text-[#343C6A] font-extrabold text-[25px]"><img src="/images/logo.png" alt="" />BankDash.</div>
        <p className='font-[600] text-[28px] text-[#343C6A]'>Overview</p>
      </div>
      <div className='flex gap-5 items-center'>
        <label className='flex items-center gap-3 h-[50px] bg-[#F5F7FA] rounded-3xl cursor-pointer px-5'>
            <img src="/images/search.png" alt="" />
            <input className="bg-inherit p-1 focus:outline-none " type="search" placeholder='search for something' />
        </label>
        <Link  href="/settings" className="h-[50px] w-[50px] flex items-center justify-center rounded-full bg-[#F5F7FA]">
            <img className='w-6 h-6'  src="/images/settings.png" alt="" />
        </Link>
        <Link  href='/notifications' className="h-[50px] w-[50px] flex items-center justify-center rounded-full bg-[#F5F7FA]">
            <img className='w-6 h-6'  src="/images/notification.png" alt="" />
        </Link>
        <Link href='/settings/editProfile'>        
          <img src="/images/profile.png" alt="" />
        </Link>
      </div>
    </div>
  )
}

export default Navbar
