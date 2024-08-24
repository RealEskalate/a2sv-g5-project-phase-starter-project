import AddNewCard from '@/components/AddNewCard/AddNewCard';
import CardList from '@/components/CardList/CardList';
import CardSettings from '@/components/CardSettings/CardSettings';
import CardExpenceStatistics from '@/components/Charts/CardExpenceStatistics';
import StoreProvider from '@/providers/StoreProvider';
import { Plus } from 'lucide-react';
import React from 'react';
import MyCardLists from '@/components/MyCard/MyCardLists';

export default function page() {
  return (
    <div className='space-y-5 '>
      <div>
        <h1 className='text-[#333B69] pb-3 font-semibold'>My Cards</h1>

        <div className='flex overflow-x-scroll scrollbar-none space-x-2 scroll whitespace-nowrap scroll-smooth lg:flex lg:space-x-4  '>
          <StoreProvider>
            <MyCardLists />
          </StoreProvider>
          <div className='w-[295px] h-[175px] bg-gray-200 rounded-3xl justify-center items-center flex flex-shrink-0'>
            <Plus size={32} />
          </div>
        </div>
      </div>

      <div className='space-y-5 lg:space-y-0 lg:flex lg:gap-6'>
        <div className=' w-full lg:w-4/12'>
          <CardExpenceStatistics />
        </div>
        <div className='w-full lg:w-2/3  '>
          <CardList />
        </div>
      </div>
      <div className='space-y-5 md:space-y-0 md:flex md:gap-6'>
        <AddNewCard />
        <CardSettings />
      </div>
    </div>
  );
}
