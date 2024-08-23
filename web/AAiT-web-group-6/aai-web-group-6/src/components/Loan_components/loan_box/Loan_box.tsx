import React from 'react'
import items from '@/constants/loan_constants/loan_box_items'

const Loan_box = () => {
  return (
    <div className='flex flex-wrap justify-between gap-4 md:gap-6'>
      {items.map((t) => (
        <div 
          key={t.id} 
          className="w-full sm:w-[255px] h-[120px] bg-white flex  gap-3 items-center justify-center rounded-[25px] md:w-[255px]"
        >
          <div 
            className='w-[70px] h-[70px] rounded-full flex justify-center items-center p-3' 
            style={{ backgroundColor: t.img_color }}
          >
            <img 
              className='h-[30px] w-[30px]' 
              src={t.img_src} 
              alt={t.img_src} 
            />
          </div>
          <div className='w-[116px] h-[51px]'>
            <p className='text-[16px] text-gray-300 whitespace-nowrap '>Personal Loans</p>
            <p className='text-[20px] font-bold'>$50,000</p>
          </div>
        </div>
      ))}
    </div>
  )
}

export default Loan_box
