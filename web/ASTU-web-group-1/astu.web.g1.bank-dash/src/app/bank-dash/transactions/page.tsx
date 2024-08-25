import React from 'react';
import MyCard from '@/components/MyCard/MyCard';
import RecentTransactionTable from '@/components/RecentTransactionTable/RecentTransactionTable';
import { Plus } from 'lucide-react';
import MyExpence from '@/components/Charts/MyExpence';
import StoreProvider from '@/providers/StoreProvider';
import MyCardLists from '@/components/MyCard/MyCardLists';
import Link from 'next/link';

interface idParamsType {
  params: {
    transactionId: string;
  };
}
export default function page({ params }: idParamsType) {
  return (
    <div className='flex flex-col gap-5'>
      <div className='flex flex-col lg:flex-row gap-6'>
        <div className='w-full lg:flex '>
          <div className='lg:w-2/3 md:pr-3 xl:pr-5 flex-shrink'>
            {/* <div className="w-full"> */}
            <div className='flex justify-between'>
              <p className='text-[#333B69] pb-2 font-semibold'>My Card</p>
              <p className='text-[#333B69] pb-2 font-semibold'>+ Add</p>
            </div>
            <div className='flex overflow-x-auto scrollbar-none space-x-2'>
              <StoreProvider>
                <MyCardLists />
              </StoreProvider>
              <div className='w-[295px] h-[175px] bg-gray-200 rounded-3xl justify-center items-center flex flex-shrink-0'>
              <Link href='/bank-dash/credit-card'>
                <Plus size={32} />
              </Link>
              </div>
            </div>
            {/* </div> */}
          </div>
          <MyExpence />
        </div>
      </div>
      <div>
        <StoreProvider>
          <RecentTransactionTable />
        </StoreProvider>
      </div>
    </div>
  );
}
