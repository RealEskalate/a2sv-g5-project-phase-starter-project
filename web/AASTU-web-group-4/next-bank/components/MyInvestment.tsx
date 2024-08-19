import React from 'react'
import { colors , logo } from '@/constants';
import Image from 'next/image';
const MyInvestment = ({icon , color ,colortext , name ,category,categorycolor ,amount,percentage}:{
    icon:string;
    color:string;
    colortext : string;
    category:string;
    categorycolor:string
    name:string;
    amount:string;
    percentage:string;
}) => {
  return (
    <div className= "w-[100%] rounded-2xl">
      <div className={`${colors.white} rounded-2xl `}>
          <div className=''>
            <div className='flex p-2 justify-between lg:justify-evenly'>
              <div className=' flex gap-4 justify-center items-center'>
                  <div className={`${color} rounded-2xl flex items-center justify-center px-4 h-[60px]`}>
                      <Image
                    src={icon}
                    alt="next logo"
                    width={30}
                    height={10}
                    className="object-contain"
                  />
                  </div>
                  <div className=' '>
                      <p className='text-[14px] font-normal'> {name}</p>
                      <p className={`${categorycolor} text-wrap w-[100px] lg:w-auto text-[12px] font-normal text-start`}> {category}</p>
                  </div>
              </div>
              <div className='py-5 flex gap-6 items-center'>
                <div className=' hidden lg:block'>
                    <p className={`${colortext} text-[14px] font-normal `} >{amount}</p>
                    <p className={`${categorycolor} text-[12px] font-normal text-start`} >Investment Value</p>
                </div>
                  <div className='flex flex-col items-end'>
                      <p className={`${colortext}`}>
                         {
                                percentage.includes("-") ?
                                <span className='text-red-500 text-[14px] font-normal text-end lg:text-start'>{percentage}</span>:
                                <span className='text-green-500 text-end text-[14px] font-normal lg:text-start'>{percentage}</span>
                         }
                      
                         </p>
                        <p className={`${categorycolor} text-[12px] font-normal text-end`} >return value</p>
                  </div>
              </div>
            </div>
          </div>
      </div>
    </div>
  )
}

export default MyInvestment
