import React from 'react';
import ProfileCard from '../ProfileCard/ProfileCard';
import Telegram from '../../../public/assets/icons/telegram-icon.svg';
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from '@/components/ui/carousel';

const QuickTransfer = () => {
  return (
    <div className='flex flex-col gap-[12px] md:gap-[16px] lg:gap-[20px]'>
      <h1 className='text-navy text-18px lg:text-22px'>Quick Transfer</h1>
      <div className='flex flex-col lg:gap-5 rounded-3xl  md:w-[320px]  w-[325px] py-5 px-6 bg-white'>
        <div className='px-10 md:px-6'>
          <Carousel className='w-full max-w-sm'>
            <CarouselContent className='-ml-1'>
              {Array.from({ length: 5 }).map((_, index) => (
                <CarouselItem
                  key={index}
                  className='pl-1 mb-1.5 md:basis-1/3 basis-1/3 lg:basis-1/3'
                >
                  <div className='p-1'>
                    <ProfileCard />
                  </div>
                </CarouselItem>
              ))}
            </CarouselContent>
            <CarouselPrevious />
            <CarouselNext />
          </Carousel>
        </div>
        <div className='flex flex-row gap-[12px] lg:gap-[35px] w-full text-12px lg:text-16px'>
          <p className='text-blue-steel flex whitespace-nowrap items-center'>Write amount</p>
          <div className='flex relative flex-row items-center h-10 rounded-full'>
            <div className='relative flex-1 h-full'>
              <input
                type='number'
                placeholder='552.50'
                className='bg-[#EDF1F7] rounded-full pl-[10px] pr-[100px] lg:pl-[20px] lg:pr-[120px] w-full border-none h-full'
              />
            </div>
            <button className='bg-blue-bright absolute right-0 rounded-full px-[21px] lg:px-[24px] gap-[9px] lg:gap-[15px] h-full text-white flex items-center justify-center'>
              <p className='block'>send</p>
              <Telegram />
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default QuickTransfer;
