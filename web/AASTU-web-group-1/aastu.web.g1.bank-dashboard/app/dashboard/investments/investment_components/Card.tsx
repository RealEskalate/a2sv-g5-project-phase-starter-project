import React from 'react'
import Image from 'next/image'
export default function Card(props:any){
  return (
    <div className='flex px-2 pl-5 bg-white h-[90px] md:h-[100px] w-[90%] md:w-[400px] rounded-3xl mx-auto md:mx-0'>
        <Image src={props.icon} alt='' width={1} height={1} className='my-auto h-[50%] w-[20%] p-3 size-[50px] text-indigo-700 bg-[#577bb813] rounded-full'  />
        <div className='pl-4 my-auto w-[70%]'>
            <h1 className='text-bg-gray font-[400]  text-[14px] md:text-[16px] text-[#718EBF]'>{props.name}</h1>
            <p className='font-[500] text-[12px] md:text-[14px] md:font-[600] pl-1  '>{props.description}</p>
        </div>
    </div>
  )
}
 