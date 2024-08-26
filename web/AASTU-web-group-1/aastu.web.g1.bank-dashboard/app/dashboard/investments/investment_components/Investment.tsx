import React from 'react'
import { investmentsArray } from '@/constants'
import InvestmentCard from './InvestmentCard'
import { useUser } from '@/contexts/UserContext'
export default function Investments(){
  const {isDarkMode} = useUser();
  return (
    <div className='w-full md:w-[50%] md:min-w-[600px]'>
      <h1 className={`my-4 font-[600] text-[22px] ${isDarkMode ? "text-white":"text-[#333B69]"}  ml-5 md:ml-0`} >My Investment</h1>
      <div className='flex flex-col gap-3 md:gap-1'>
        {investmentsArray.map((item, index) => (
          <InvestmentCard {...item} key={index}/>
        ))}
      </div>
    </div>
  )
}

