import React from 'react';
import Image from 'next/image'

interface InvoiceItem {
    image: string;
    entity: string;
    time: string;
    amount: string;
  }


const Invoice = ({item}: {item: InvoiceItem}) => {
  return (
    <div className='flex item-center justify-around'>
      <div className="flex items-center gap-8 text-sm">
        <Image src={item.image} alt={item.entity} className='text-bold w-8 h-8' />
        <div className="flex flex-col items-start justify-center gap-2">
            <span className='text-[#B1B1B1]'>{item.entity}</span>
            <span className='text-[#718EBF]'>{item.time} ago</span>
        </div>
      </div>
      <span className='text-[#718EBF]'>${item.amount}</span>
    </div>
  )
}

export default Invoice 