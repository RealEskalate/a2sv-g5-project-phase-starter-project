import React from 'react'
import items from '@/constants/loan_constants/loan_box_items'

const Loan_box = () => {
  return (
    <div className='bg-slate-100 p-20 flex gap-2'>
      {items.map((t) => (
        <div className= "w-max h-max bg-white flex  p-[15px] gap-4 items-center rounded-[20px]">
        <div className='w-max h-max bg-blue-300 rounded-[50%] flex justify-center items-center p-3'  style={{ backgroundColor: t.img_color }}>
          <img className='flex justify-center items-center h-[1.5rem]' src={t.img_src}/>
        </div>
        <div>
            <p className='text-[10pt] text-gray-300'>Personal Loans</p>
            <p className='text-[12pt] font-bold'> $50,000</p>
        </div>
    </div>
      ))}
         

    </div>
   
  )
}

export default Loan_box