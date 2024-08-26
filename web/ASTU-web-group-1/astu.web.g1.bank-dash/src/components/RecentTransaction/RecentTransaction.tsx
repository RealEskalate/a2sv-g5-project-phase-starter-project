'use client';
import React from 'react';
import RecentTransactionCard from './RecentTransactionCard';
import { useGetAllTransactionsQuery } from '@/lib/redux/api/transactionSlice';
import { Skeleton } from '../ui/skeleton';

const RecentTransaction = () => {
  const { data, isLoading } = useGetAllTransactionsQuery({ page: 0, size: 3 });
  if (isLoading) {
    return (
      <div className='w-full pb-3 '>
        <h1 className='text-[#333B69] pb-3 font-semibold'>Recent Transaction</h1>
        <div className='space-y-4 bg-white p-5 rounded-3xl'>
          <div className='flex'>
            <Skeleton className='h-9 w-9 rounded-xl bg-slate-200 flex-shrink-0 mr-3' />
            <Skeleton className='h-8 w-full bg-slate-200' />
          </div>
          <div className='flex'>
            <Skeleton className='h-9 w-9 rounded-xl bg-slate-200 flex-shrink-0 mr-3' />
            <Skeleton className='h-8 w-full bg-slate-200' />
          </div>
          <div className='flex'>
            <Skeleton className='h-9 w-9 rounded-xl bg-slate-200 flex-shrink-0 mr-3' />
            <Skeleton className='h-8 w-full bg-slate-200' />
          </div>
        </div>
      </div>
    );
  }

  if (data?.data.content.length === 0) {
    return (
      <div className='w-full h-full'>
        <h1 className='text-[#333B69] pb-3 font-semibold'>Recent Transaction</h1>
        <div className='w-full h-[85%] flex justify-center items-center bg-white rounded-3xl text-lg font-semibold'>
          No Transaction Found
        </div>
      </div>
    );
  }

  return (
    <div className='w-full'>
      <h1 className='text-[#333B69] pb-3 font-semibold'>Recent Transaction</h1>
      <div className='  bg-white  rounded-[15px] px-4 py-3'>
        <ul role='list' className=' '>
          {data?.data.content.map((transaction) => (
            <li className='py-1' key={transaction.transactionId}>
              <RecentTransactionCard
                TransactionName={transaction.description}
                calender={transaction.date}
                amount={transaction.amount}
                imageUrl='/assets/images/dollarCoin.png'
                sign={transaction.type}
              />
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default RecentTransaction;
