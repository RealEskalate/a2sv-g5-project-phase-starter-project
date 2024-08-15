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
        active ? 'text-blue' : 'text-gray-light'
      } flex font-semibold items-center my-6 text-16px cursor-pointer`}
    >
      {active && <div className='w-1.5 h-10 bg-blue rounded-r-[10px] absolute'></div>}
      <Icon className='w-[21px] h-[21px] mx-5 ml-6 object-cover fon' />
      <p className='font-inter'>{title}</p>
    </div>
  );
}
