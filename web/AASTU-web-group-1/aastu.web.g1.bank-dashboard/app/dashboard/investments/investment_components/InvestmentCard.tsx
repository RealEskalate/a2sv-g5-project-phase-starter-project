import React from 'react'
import Image from 'next/image'
export default function InvestmentCard(props:any){
  return (
    <div className="flex  bg-white rounded-3xl mb-2 mt-4 justify-evenly w-full h-[100px]">
      <div className='my-auto'>
        <Image src={props.icon} alt='' width={1} height={1} className='my-auto size-[50px] text-indigo-700 bg-[#577bb813] rounded-full'  />
      </div>
      <div className='my-auto'>
        <h1>{props.name}</h1>
        <div className='flex flex-row gap-1'>
            {props.type.map((item:any) => (
              <p className='text-bg-gray font-[400] text-blue-900 text-[14px] md:text-[16px]' key={item}>{item}</p>
            ))}
        </div>
      </div>
      <div className='my-auto'>
        <h1>{props.investmentValue}</h1>
        <p className='text-blue-800'>Envestment Value</p>
      </div>
      <div className='my-auto'>
        <h1 className='text-emerald-700'>{props.returnValue}</h1>
        <p className='text-blue-800'>Return Value</p>
      </div>
    </div>
  )
}

