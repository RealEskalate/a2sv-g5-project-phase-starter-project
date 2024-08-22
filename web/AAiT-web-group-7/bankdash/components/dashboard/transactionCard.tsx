import React from 'react'
import Image from 'next/image';
import img1 from '@/public/dash1.svg';
import img2 from '@/public/dash2.svg';
import img3 from '@/public/dash3.svg';

const TransactionCard = () => {
  return (
    <div className='flex flex-col gap-4 bg-white p-5 rounded-3xl text-sm '>
      <div className="flex items-center gap-6">
        <Image src={img1} alt='deposit from card' className='w-12 h-12' />
        <span className='flex flex-col items-start gap-2 justify-center'>
            <p>Deposit from my Card</p>
            <p className='text-[#718EBF] text-sm'>28 January 2023</p>
        </span>
        <p className='text-[#FF4B4A]'>-$850</p>
      </div>
      <div className="flex gap-6 items-center">
        <Image src={img2} alt='deposit from card' className='w-12 h-12' />
        <span className='flex flex-col items-start justify-center gap-2'>
            <p>Deposit PayPal</p>
            <p className='text-[#718EBF] text-sm'>28 January 2023</p>
        </span>
        <p className='text-[#41d4a8] ml-8'>$2,500</p>
      </div>
      <div className="flex gap-6 items-center">
        <Image src={img3} alt='deposit from card' className='w-12 h-12' />
        <span className='flex flex-col items-start gap-2 justify-center'>
            <p>Deposit from my Card</p>
            <p className='text-[#718EBF] text-sm'>28 January 2023</p>
        </span>
        <p className='text-[#41d4a8]'>$5,400</p>
      </div>
    </div>
  )
}

export default TransactionCard
