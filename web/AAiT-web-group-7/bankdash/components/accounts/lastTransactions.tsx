import React from 'react';
import Image from 'next/image'

interface Props {
    image: string,
    title: string,
    date: string,
    type: string,
    phone: string,
    status: string,
    amount: string
}


const LastTransactionsComp = ({item}: {item: Props}) => {
  return (
    <div className='text-sm grid grid-cols-7 p-3 gap-4'>
      <Image src={item.image} alt={item.type} className='col-span-1'/>
      <div className="flex flex-col items-start justify-center gap-2 col-span-2">
        <p className='font-semibold'>{item.title}</p>
        <p className='text-[#718EBF]'>{item.date}</p>
      </div>
      <span className='text-[#718EBF] col-span-1'>{item.type}</span>
      <span className='text-[#718EBF] col-span-1'>{item.phone}</span>
      <span className='text-[#718EBF] col-span-1'>{item.status}</span>
      <span className={`${Number(item.amount) < 0 ? "text-[#db3e3e]" : "text-[#3cd437]"} col-span-1`}>{item.amount}$</span>
    </div>
  )
}

export default LastTransactionsComp
