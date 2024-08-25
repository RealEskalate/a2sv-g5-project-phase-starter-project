import CardSettings from '@/components/CardSettings/CardSettings';
import StoreProvider from '@/providers/StoreProvider';
import { Plus } from 'lucide-react';
import React from 'react';
import MyCardLists from '@/components/MyCard/MyCardLists';
import CardSkeleton from '@/components/AllSkeletons/CardSkeleton/CardSkeleton';
import CardAndExpenceStatisticsSkeleton from '@/components/AllSkeletons/CardAndExpenseStatistics/CardAndExpenseStatistics';
import CardListCardSkeleton from '@/components/AllSkeletons/CardListSkeleton/CardListSkeleton';
import { Skeleton } from '@/components/ui/skeleton';
import AddNewCardSkeleton from '@/components/AllSkeletons/AddNewCardSkeleton/AddnewCardSkeleton';

export default function page() {
  return (
    <div className='space-y-5 '>
      <div>
        <Skeleton className='text-[#333B69] pb-3 bg-slate-200 mb-2' />

        <div className='flex overflow-x-scroll space-x-2 scroll whitespace-nowrap scroll-smooth lg:flex lg:space-x-3  '>
          {[...Array(4)].map((_, index) => (
            <CardSkeleton key={index} />
          ))}
        </div>
      </div>

      <div className='space-y-5 lg:space-y-0 lg:flex lg:gap-6'>
        <div className=' w-full lg:w-4/12'>
          <CardAndExpenceStatisticsSkeleton />
        </div>
        <div className='w-full lg:w-2/3  '>
          <Skeleton className='bg-slate-200 w-48 h-6 mb-2 ' />

          {[...Array(4)].map((_, index) => (
            <CardListCardSkeleton key={index} />
          ))}
        </div>
      </div>
      <div className='space-y-5 md:space-y-0 md:flex md:gap-6'>
        <AddNewCardSkeleton />
      </div>
    </div>
  );
}
