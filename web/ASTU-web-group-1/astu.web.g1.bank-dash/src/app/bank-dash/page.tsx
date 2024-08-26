import BalanceHistory from '@/components/Charts/BalanceHistory';
import DebiteAndCredit from '@/components/Charts/DebiteAndCredit';
import ExpenseStatistics from '@/components/Charts/ExpenseStatistics';
import WeeklyActivity from '@/components/Charts/WeeklyActivity';
import MyCard from '@/components/MyCard/MyCard';
import QuickTransfer from '@/components/QuickTransfer/QuickTransfer';
import RecentTransaction from '@/components/RecentTransaction/RecentTransaction';
import { Plus } from 'lucide-react';
import StoreProvider from '@/providers/StoreProvider';
import MyCardLists from '@/components/MyCard/MyCardLists';
import Link from 'next/link';

export default function Home() {
  return (
    <>
      <div className='w-full lg:flex '>
        <div className='lg:w-2/3 md:pr-3 xl:pr-5 flex-shrink'>
          <div className='w-full'>
            <div className='flex justify-between'>
              <p className='text-[#333B69] pb-3 font-semibold'>My Cards</p>
              <p className='text-[#333B69] pb-3 font-semibold'>
                <Link href='bank-dash/credit-card'>See All</Link>{' '}
              </p>
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
          </div>
        </div>
        <div className='lg:w-1/3 w-full'>
          <StoreProvider>
            <RecentTransaction />
          </StoreProvider>
        </div>
      </div>
      <div className='lg:flex my-5'>
        <WeeklyActivity />
        <ExpenseStatistics />
      </div>
      <div className='lg:flex justify-between w-full lg:space-x-5'>
        <div className='lg:w-5/12'>
          <QuickTransfer />
        </div>
        <BalanceHistory />
      </div>
    </>
  );
}
