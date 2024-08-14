import React from 'react';
import Hamburger from '../../../public/assets/icons/hamburger-icon.svg';
import Settings from '../../../public/assets/icons/setting-icon.svg';
import Notification from '../../../public/assets/icons/notification-icon.svg';
import Avatar from '../Avatar/Avatar';

export default function NavBar() {
  return (
    <div className='w-full p-5 bg-white h-[80px] flex justify-between px-5 items-center'>
      <button type='button' className='items-center rounded-lg sm:hidden font-bold'>
        <Hamburger className='w-7 h-7 object-cover' />
      </button>
      <p className=''>active</p>
      <div className='flex space-x-2 items-center'>
        <div className='hidden md:flex '>
          <input
            type='text'
            placeholder='Search for something'
            className='bg-slate-100 px-3 pl-5 rounded-full py-2 outline-none'
          />
        </div>
        <div className=' hidden md:flex w-11 h-11 bg-slate-100 rounded-full justify-center items-center text-blue-steel'>
          <Settings className='w-6 h-6' />
        </div>
        <div className='hidden md:flex w-11 h-11 bg-slate-100 rounded-full justify-center items-center text-red-400'>
          <Notification className='w-6 h-6' />
        </div>
        <div className='w-11 h-11 object-contain object-center rounded-full overflow-hidden'>
          <Avatar />
        </div>
      </div>
    </div>
  );
}
