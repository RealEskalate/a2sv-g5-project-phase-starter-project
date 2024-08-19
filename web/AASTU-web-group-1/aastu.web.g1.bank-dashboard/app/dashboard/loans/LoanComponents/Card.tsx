import React from 'react'
import Image from "next/image";
export default function Card(props:any){
  return (
    <div className='h-[90px] w-[280px] md:w-[380px] bg-white rounded-3xl m-2 flex justify-around'>
            <Image src={props.icon} alt='' width={1} height={1} className='my-auto h-[60%] w-[20%] p-3 md:p-4 size-[50px] text-indigo-700 bg-[#577bb813] rounded-full'  />
            <div className='pl-3 my-auto w-2/3'>
              <h1 className='text-bg-gray font-[400]  text-[12px] md:text-[17px] text-[#718EBF]'>{props.name}</h1>
              <p className='font-[500] text-[12px] md:text-[14px] md:font-[600]'>{props.description}</p>
            </div>
    </div>
  )
}

