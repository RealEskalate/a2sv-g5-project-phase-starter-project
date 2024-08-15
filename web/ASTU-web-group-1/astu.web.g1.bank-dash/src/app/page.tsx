import MyCard from '@/components/MyCard/MyCard';
import QuickTransfer from '@/components/QuickTransfer/QuickTransfer';
import RecentTransaction from '@/components/RecentTransaction/RecentTransaction';
import { Plus } from 'lucide-react';

export default function Home() {
  return (
    <>
      <div className='w-full lg:flex '>
        <div className='lg:w-2/3 md:pr-3 xl:pr-5 flex-shrink'>
          <div className='w-full'>
            <div className='flex justify-between'>
              <p className='text-[#333B69] pb-2 font-semibold'>My Card</p>
              <p className='text-[#333B69] pb-2 font-semibold'>See All</p>
            </div>
            <div className='flex  overflow-x-auto space-x-2'>
              <MyCard />
              <MyCard />
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
      <QuickTransfer />
    </>
  )
}
