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
    <div className='flex flex-col'>
      <p className='text-[#333B69] pb-3 font-semibold'>Quick Transfer</p>
      <div className='flex flex-col lg:gap-5 rounded-3xl  md:w-[320px] lg:w-full  w-[320px] py-5 px-6 bg-white'>
        <div className='px-10 md:px-6'>
          <Carousel className='w-full max-w-sm'>
            <CarouselContent className='-ml-1'>
              {Array.from({ length: 6 }).map((_, index) => (
                <CarouselItem
                  key={index}
                  className='pl-1 mb-1.5 md:basis-1/3 basis-1/3 lg:basis-1/4'
                >
                  {/* <div className='p-1'> */}
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
        <div className='flex flex-row  w-full text-15px'>
          <p className='text-blue-steel flex whitespace-nowrap items-center mr-2'>Write amount</p>
          <div className='flex relative flex-row items-center h-10 rounded-full'>
            <div className='relative flex-1 h-full'>
              <input
                type='number'
                placeholder='552.50'
                className='bg-[#EDF1F7] rounded-full pl-[10px] pr-[100px] w-full border-none h-full outline-none'
              />
            </div>
            <button className='bg-blue-bright absolute right-0 rounded-full px-5  h-full text-white flex items-center justify-center '>
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
