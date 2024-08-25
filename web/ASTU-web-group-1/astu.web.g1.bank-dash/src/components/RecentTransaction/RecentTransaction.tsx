'use client';
import React from 'react';
import RecentTransactionCard from './RecentTransactionCard';
import { useGetAllTransactionsQuery } from '@/lib/redux/slices/transactionSlice';
import { Skeleton } from '../ui/skeleton';

const RecentTransaction = () => {
  const { data, isLoading } = useGetAllTransactionsQuery({ page: 0, size: 3 });
  if (12) {
    return (
      <div className='w-full'>
        <Skeleton className='h-8 w-52 bg-slate-200' />
        hello world
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
