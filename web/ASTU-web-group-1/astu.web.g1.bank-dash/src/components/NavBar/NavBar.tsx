import React from 'react';
import Hamburger from '../../../public/assets/icons/hamburger-icon.svg';
import Settings from '../../../public/assets/icons/setting-icon.svg';
import Notification from '../../../public/assets/icons/notification-icon.svg';

export default function NavBar() {
  return (
    <div className='w-full p-5 bg-white h-[80px] flex justify-between px-5 items-center'>
      <button type='button' className='items-center rounded-lg sm:hidden font-bold'>
        <Hamburger className='w-7 h-7 object-cover' />
      </button>
      <p className=''>active</p>
      <div className='flex space-x-2'>
        <div className='w-11 h-11 bg-slate-50 rounded-full flex justify-center items-center text-blue-steel'>
          <Settings className='w-6 h-6' />
        </div>
        <div className='w-11 h-11 bg-slate-50 rounded-full flex justify-center items-center text-red-400'>
          <Notification className='w-6 h-6' />
        </div>
      </div>
    </div>
  );
}
