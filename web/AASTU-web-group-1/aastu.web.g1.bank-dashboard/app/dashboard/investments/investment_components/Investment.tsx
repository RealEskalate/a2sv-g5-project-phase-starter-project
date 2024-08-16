import React from 'react'
import { investmentsArray } from '@/constants'
import InvestmentCard from './InvestmentCard'
export default function Investments(){
  return (
    <div className='w-[65%]'>
      <h1>My Investment</h1>
      <div className='flex flex-col gap-3 md:gap-1'>
        {investmentsArray.map((item) => (
          <InvestmentCard {...item} key="item.name"/>
        ))}
      </div>
    </div>
  )
}

