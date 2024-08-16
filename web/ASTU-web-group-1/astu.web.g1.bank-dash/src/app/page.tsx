<<<<<<< HEAD
import TrendingList from "@/components/TrendingStock/trendingList";
export default function Home() {
  return <div></div>;
=======
import BalanceHistory from '@/components/Charts/BalanceHistory';
import DebiteAndCredit from '@/components/Charts/DebiteAndCredit';
import ExpenseStatistics from '@/components/Charts/ExpenseStatistics';
import WeeklyActivity from '@/components/Charts/WeeklyActivity';
import MyCard from '@/components/MyCard/MyCard';
import QuickTransfer from '@/components/QuickTransfer/QuickTransfer';
import RecentTransaction from '@/components/RecentTransaction/RecentTransaction';
import TestNaol from '@/components/TestNaol';
import CardProvider from '@/providers/cardProvider';
import { Plus } from 'lucide-react';

export default function Home() {
  return (
    <>
      <div className='w-full lg:flex '>
        <div className='lg:w-2/3 md:pr-3 xl:pr-5 flex-shrink'>
          <div className='w-full'>
            <div className='flex justify-between'>
              <p className='text-[#333B69] pb-3 font-semibold'>My Card</p>
              <p className='text-[#333B69] pb-3 font-semibold'>See All</p>
            </div>
            <div className='flex  overflow-x-auto space-x-2'>
              <CardProvider>
                <MyCard />
                <MyCard />
              </CardProvider>
              <div className='w-[295px] h-[175px] bg-gray-200 rounded-3xl justify-center items-center flex flex-shrink-0'>
                <Plus size={32} />
              </div>
            </div>
          </div>
        </div>
        <div className='lg:w-1/3 w-full'>
          <RecentTransaction />
        </div>
      </div>
      <div className='md:flex my-5'>
        <WeeklyActivity />
        {/* <DebiteAndCredit /> */}
        <ExpenseStatistics />
      </div>
      <div className='md:flex justify-between'>
        <div className='w-5/12 pe-6'>
          <QuickTransfer />
        </div>
        <BalanceHistory />
      </div>
    </>
  );
>>>>>>> 715c07199bf7b1561bb67630c6c3611075825f77
}
