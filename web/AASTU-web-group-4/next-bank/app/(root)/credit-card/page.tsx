import React from 'react'
import ResponsiveCreditCard from '@/components/CreditCard'
import { colors } from '@/constants'
import Component from '@/components/DoughnutChart'
import AddNewCard from '@/components/AddNewCard'
import CardSetting from '@/components/CardSetting'

const CreditCard = () => {
  const cards:string[] = [colors.blue, colors.white, colors.blue, colors.white]
  return (
    <div className='ml-40'>
      <div className="myCards w-[80%] mx-auto mt-4">
        <h1 className='text-[19px] mb-3 font-bold text-[#333B69]'>My Cards</h1>
        <div className="flex overflow-x-auto space-x-4 pr-3">
          {cards.map((bg, index) => (
            <span key={index} className='p-3'>
              <ResponsiveCreditCard backgroundColor={bg} />
            </span>
          ))}
        </div>
      </div>
      
      <div className="flex w-[80%] mx-auto">
        <div className="doughnutChart w-[30%] my-6">
          <h1 className='text-[19px] mb-3 font-bold text-[#333B69]'>Card Expense Statistics</h1>
          <Component />
        </div>
      </div>

      <div className="flex flex-col md:flex-row w-[80%] mx-auto mb-16">
        <div className='mb-2 sm:mb-0 md:mr-10'>
          <AddNewCard/>
        </div>

        <div>
          <CardSetting />
        </div>
      </div>
    </div>
  )
}

export default CreditCard