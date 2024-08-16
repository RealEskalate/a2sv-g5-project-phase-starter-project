import React from 'react';
import Image from 'next/image';

const Avatar = () => {
  return (
    <img
      src='/assets/images/profilepicture.jpg'
      alt=''
      className='object-cover rounded-full w-[50px] h-[50px] md:w-[50px] md:h-[50px]'
    />
  );
};

export default Avatar;
