import Image from 'next/image';
import React from 'react';

const MyCard = () => {
  return (
    <div>
      <div className='w-[280px] h-[175px] bg-grad-end text-white  border rounded-3xl flex flex-col justify-between'>
        <div className='flex flex-col  px-4  pt-4 h-full'>
          <div className='flex justify-between '>
            <div className='flex flex-col'>
              <span className='text-12px'>Balance</span>
              <span className='text-14px font-semibold'>$5,756</span>
            </div>
            <div className=''>
              <Image src='/chip-card-white.svg' alt='chip_card' width={30} height={30} />
            </div>
          </div>
          <div className='flex h-full pb-3'>
            <div className='flex flex-1 flex-col justify-center'>
              <span className='text-12px'>CARD HOLDER</span>
              <span className='text-14px font-semibold'>Eddy Cusuma</span>
            </div>
            <div className='flex flex-1 flex-col justify-around items-center'>
              <div className='flex flex-col'>
                <span className='text-12px'>VALID THRU</span>
                <span className='text-14px  font-semibold'>12/22</span>
              </div>
            </div>
          </div>
        </div>
        <div className='rounded-b-3xl flex justify-between px-6 py-3 bg-gradient-to-b from-grad-start to-grad-end'>
          <div className='flex items-center'>
            <span className='text-16px font-text-navy'>3778 **** **** 1234</span>
          </div>
          <div className='flex pr-[15px]'>
            <div className='flex bg-[#9199AF] h-8 w-8 rounded-full opacity-50'></div>
            <div className='flex bg-[#9199AF] h-8 w-8 rounded-full opacity-50 -mx-[15px]'></div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default MyCard;
