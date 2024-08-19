import React from 'react'
import Image from 'next/image'
export default function InvestmentCard(props:any){
  return (
    <div className="flex bg-white rounded-3xl mb-2 w-[90%] md:w-full h-[100px] mx-auto p-2">
      <div className='my-auto w-[15%] md:w-[10%] h-auto'>
        <Image src={props.icon} alt='' width={1} height={1} className='my-auto size-[25px] md:size-[50px] rounded-full'  />
      </div>
      <div className='my-auto w-[50%] md:w-[30%]'>
        <h1>{props.name}</h1>
        <div className='flex flex-row gap-1 flex-wrap'>
            {props.type.map((item:any) => (
              <p className='text-bg-gray font-[400] text-blue-900 text-[14px] md:text-[16px] text-wrap w-[100px] ml-0' key={item}>{item}</p>
            ))}
        </div>
      </div>
      <div className='my-auto w-[30%] hidden lg:table-cell'>
        <h1 className='text-[16px] md:text-[20px]' >{props.investmentValue}</h1>
        <p className='text-blue-800 text-[10px] md:text-[16px]'>Envestment Value</p>
      </div>
      <div className='my-auto w-[40%] md:w-[30%] pl-16 md:pl-0'>
        <h1 className='text-emerald-700'>{props.returnValue}</h1>
        <p className='text-blue-800 hidden md:table-cell my-auto'>Return Value</p>
      </div>
    </div>
  )
}

