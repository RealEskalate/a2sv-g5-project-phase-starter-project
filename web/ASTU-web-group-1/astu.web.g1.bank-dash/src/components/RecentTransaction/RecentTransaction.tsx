'use client';
import React from 'react';
import RecentTransactionCard from './RecentTransactionCard';
import { useGetAllTransactionsQuery } from '@/lib/redux/slices/transactionSlice';

const RecentTransaction = () => {
  const { data, isLoading } = useGetAllTransactionsQuery({ page: 0, size: 3 });
  console.log(data, 'data transaction');

  return (
    <div className='w-full'>
      <h1 className='text-[#333B69] pb-3 font-semibold'>Recent Transaction</h1>
      <div className='  bg-white  rounded-[15px] px-4 py-3'>
        <ul role='list' className=' '>
          {/* <li className='py-2'>
            <RecentTransactionCard
              TransactionName='Deposit from my'
              calender='28 January 2021'
              amount={850}
              imageUrl='/assets/images/deposit.png'
              moneyColor='#FF4B4A'
              sign='-'
            />
          </li>
          <li className='py-2'>
            <RecentTransactionCard
              TransactionName='Depoist Paypal'
              calender='25 January 2021'
              amount={2500}
              imageUrl='/assets/images/paypal.png'
              moneyColor='#41D4A8'
              sign='+'
            />
          </li>
          <li className='py-2'>
            <RecentTransactionCard
              TransactionName='Jemi Wilson'
              calender='21 January 2021'
              amount={5400}
              imageUrl='/assets/images/dollarCoin.png'
              moneyColor='#41D4A8'
              sign='+'
            />
          </li> */}
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
