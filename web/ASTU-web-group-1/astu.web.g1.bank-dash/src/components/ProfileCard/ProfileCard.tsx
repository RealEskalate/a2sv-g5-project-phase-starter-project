import React from 'react';
import Avatar from '../Avatar/Avatar';

const ProfileCard = () => {
  return (
    <div className='flex flex-col gap-4 justify-center items-center max-w-[100px] text-11px lg:mx-2'>
      <Avatar />
      <div className='flex flex-col gap-0.5 justify-center w-full'>
        <h1 className='text-gray-dark truncate whitespace-nowrap overflow-hidden text-overflow-ellipsis font-medium text-lg'>
          Chala olani
        </h1>
        <p className='text-center text-blue-steel text-md'>CEO</p>
      </div>
    </div>
  );
};

export default ProfileCard;
