import React from 'react';
import HomeIcon from '../../../public/assets/icons/home-icon.svg';

export default function SideBarItems({
  Icon,
  title,
  active,
}: {
  Icon: any;
  title: string;
  active: boolean;
}) {
  return (
    <div
      className={`${
        active ? 'text-blue-bright' : 'text-gray-light'
      } flex font-semibold items-center my-6 text-15px cursor-pointer`}
    >
      {active && <div className='w-1.5 h-10 bg-blue-bright rounded-r-[10px] absolute'></div>}
      <Icon className='w-[18px] h-[18px] mx-4 ml-6 object-cover' />
      <p className='font-inter text-sm'>{title}</p>
    </div>
  );
}
