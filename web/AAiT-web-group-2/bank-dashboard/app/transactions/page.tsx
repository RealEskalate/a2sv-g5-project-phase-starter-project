import React from 'react'
import MyCards from '../components/MyCards'
import MyExpense from '../components/MyExpense'
import ChipCard from '@/public/ChipCard'
import RecentTransactions from '../components/RecentTransactions'

const page = () => {
  return (
    <div className='bg-background px-8  '> 
        <div className=' flex gap-[30px] justify-between  items-center mb-6 text-[#343C6A] max-sm:block'>
           <div className='sm:w-[550px]  overflow-x-scroll  flex gap-5   items-center '>
                <div className=''>
                    <p className='font-[600] text-[22px] py-4'>My Cards</p>
                    <MyCards />
                </div>
                <div className=''>
                    <p className='font-[600] text-[17px] py-4 text-end '>+ Add Cards</p>
                    <MyCards />
                </div>
           </div>
            <div className=''>
                <p className='font-[600] text-[22px] py-4'>My Expense</p>
                <MyExpense />
            </div>
        </div>

        <div>
            <p className="font-[600] pl-2 py-5 text-lg text-[#343C6A]"> Recent Transactions</p>
            <div className="flex pl-2 gap-10 max-sm:gap-7 text-[#718EBF]">
                <p className="pb-2 border-b-2 border-transparent group hover:text-[#1814F3] hover:font-bold hover:border-[#1814F3]">All Transactions</p>
                <p className="pb-2 border-b-2 border-transparent group hover:text-[#1814F3] hover:font-bold hover:border-[#1814F3]">Income</p>
                <p className="pb-2 border-b-2 border-transparent group hover:text-[#1814F3] hover:font-bold hover:border-[#1814F3]">Expense</p>
            </div>

            <RecentTransactions />
        </div>
              
    </div>
    
  )
}

export default page
