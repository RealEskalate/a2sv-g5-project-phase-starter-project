'use client'

import React from 'react'
import BankServiceList from './BankServiceList'
import LifeInsurance from './LifeInsurance'

const Page = () => {
  return (
    <div className='bg-gray-100'>
      <div className='px-4 py-2 sm:px-10'>
        <LifeInsurance />
      </div>
      <div className='p-2 sm:p-5'>
        <p className='text-xl sm:text-3xl font-sans text-blue-950 opacity-100 font-semibold'>
          Bank Service List
        </p>
      </div>
      <div className='flex flex-col gap-4'>
        <BankServiceList />
        <BankServiceList />
        <BankServiceList />
        <BankServiceList />
        <BankServiceList />
      </div>
    </div>
  )
}

export default Page
