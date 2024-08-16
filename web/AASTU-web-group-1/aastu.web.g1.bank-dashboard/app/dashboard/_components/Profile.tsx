import React from 'react'
import Image from 'next/image';
interface Props{
    image:string,
    name:string,
    job:string
}
export const Profile = ({image,name,job}:Props) => {
  return (
    <div className=" min-w-[55px]">
      <div className="inline bg-transparent rounded-xl ">
        <Image
          src={image}
          alt={`profile picture`}
          className="inline-block !rounded-full object-cover object-center ml-2"
          width={50}
          height={50}
        />
        <div className='block'>
          <h4 className="font-inter font-normal text-[12px] min-w-[55px] text-center">{name}</h4>
          <h4 className="font-inter font-normal text-[12px] text-[#718EBF] text-center">{job}</h4>
        </div>
      </div>
    </div>
  );
}
