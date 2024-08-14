import React from 'react';

const RecentTransaction = () => {
  return (
    <div className='w-[350px] h-[282px]'>
      <h1 className='mb-[40px] text-[22px] text-[#343C6A] font-semibold'>Recent Transaction</h1>
      <div className='bg-white drop-shadow-xl font-medium rounded-[25px] p-[25px]'>
        <div className='flex flex-col flex-wrap gap-3'>
          <div className='flex items-center gap-2'>
            <img src='/assets/recentTransaction/p1.png' alt='Icon 1' className='w-[55px] h-[55px]' />
            <div className='flex flex-col basis-1/2 g-1 ml-2'>
              <p className='text-[16px] leading-[19.36px]'>Deposit from my</p>
              <p className='text-[#718EBF] text-[15px] leading-[18.15px]'>28 January 2021</p>
            </div>
            <p className='text-[#FF4B4A] ml-auto '>-$850</p>
          </div>
          <div className='flex items-center gap-2 '>
            <img src='/assets/recentTransaction/p2.png' alt='Icon 2' className='w-[55px] h-[55px]' />
            <div className='flex flex-col flex-wrap basis-1/2  gap-1 ml-2'>
              <p className='text-[16px] leading-[19.36px]'>Deposit Paypal</p>
              <p className='text-[#718EBF] text-[15px] leading-[18.15px]'>28 January 2021</p>
            </div>
            <p className=' text-[#41D4A8] ml-auto'>+$2,5000</p>
          </div>
          <div className='flex items-center gap-2'>
            <img src='/assets/recentTransaction/p3.png' alt='Icon 3' className='w-[55px] h-[55px]' />
            <div className='flex flex-col  flex-wrap basis-1/2  gap-1 ml-2'>
              <p className='text-[16px] leading-[19.36px]'>Jemi Wilson</p>
              <p className='text-[#718EBF] text-[15px] leading-[18.15px]'>28 January 2021</p>
            </div>
            <p className='flex flex-wrap  text-[#41D4A8] ml-auto'>+$5,4000</p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default RecentTransaction;
