import React from 'react';
import MyCard from '@/components/MyCard/MyCard';
import RecentTransactionTable from '@/components/RecentTransactionTable/RecentTransactionTable';
export default function page() {
  return (
    <div className="flex flex-col gap-5">
      <div className="flex flex-row gap-6">
        <div className="flex flex-col w-2/3 gap-4">
          <div className='flex justify-between'>
            <p>MyCard</p>
            <p>+ Add card</p>
          </div>
          <div className='flex flex-row justify-between'>
          <MyCard  />
          <MyCard />
          </div>
        </div>
        <div className="flex flex-col w-1/3 gap-4">
          <div className='flex justify-between'>
            <p>My Expense</p>
          </div>
          <div className='flex'>
          <MyCard  />
          </div>
        </div>
      </div>
      <div>
        <RecentTransactionTable />
      </div>
    </div>
  );
}
