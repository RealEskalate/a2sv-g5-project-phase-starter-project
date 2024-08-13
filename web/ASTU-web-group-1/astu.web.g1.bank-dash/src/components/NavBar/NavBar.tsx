import React from 'react';
import Hamburger from '../../../public/assets/icons/hamburger-icon.svg';

export default function NavBar() {
  return (
    <div className='w-full p-5 bg-white h-[80px] flex justify-between px-5 items-center'>
      <button type='button' className='items-center rounded-lg sm:hidden font-bold'>
        <Hamburger className='w-7 h-7 object-cover' />
      </button>
      <p className=''>url</p>
      <p>pp</p>
    </div>
  );
}
