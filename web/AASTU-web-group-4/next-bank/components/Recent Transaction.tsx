import React from 'react'
import { colors , logo } from '@/constants';
import Image from 'next/image';
const RecentTransaction = ({icon , color ,colortext , name , date ,amount}:{
    icon:string;
    color:string;
    colortext : string;
    name:string;
    date:string;
    amount:string
}) => {
  return (
    <div className= "">
      <div className={`${colors.white} `}>
          <div className=''>
            <div className='flex  p-2 justify-between w-screen md:w-auto'>
              <div className=' flex gap-4'>
                  <div className={`${color} rounded-full flex items-center justify-center p-4 h-20 w-20  `}>
                      <Image
                    src={icon}
                    alt="next logo"
                    width={150}
                    height={50}
                    className="object-contain"
                  />
                  </div>
                  <div className='my-2'>
                      <p className='text-lg font-medium'> {name}</p>
                      <p className={`${colors.textgray} text-sm text-start`}> {date}</p>
                  </div>
              </div>
              <div className='my-5'>
                  <p className={`${colortext}`}> ${amount}</p>
              </div>
            </div>
          </div>
      </div>
    </div>
  )
}

export default RecentTransaction
