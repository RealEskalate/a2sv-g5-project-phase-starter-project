'use client';
import React from 'react';
import { useSession } from 'next-auth/react';
import { useGetAllTransactionQuery } from '@/lib/service/TransactionService';
import RecentTransactionSkeleton from './RecentTransactionSkeleton';
import ErrorImage from '../Error/ErrorImage';

interface Props {
  description: string;
  date: string;
  amount: number;
  type: string;
  icon: string;
  receiverUserName: string;
}

const icons = [
  "/assets/recentTransaction/icon1.svg",
  "/assets/recentTransaction/icon2.svg",
  "/assets/recentTransaction/icon3.svg",
];

const recentlistitems = [
  {   
    transactionName: "Deposit from my",
    date: "28 January 2021", 
    amount: -880,
    isDeposited: false,
    icons: "/assets/recentTransaction/icon1.svg"
  },
  {
    transactionName: "Deposit Paypal",
    date: "28 January 2021",
    amount: 2500,
    isDeposited: true,
    icons: "/assets/recentTransaction/icon2.svg"
  },
  {
    transactionName: "Jemi Wilson",
    date: "28 January 2021",
    amount: 5400,
    isDeposited: true,
    icons: "/assets/recentTransaction/icon3.svg"
  }
];

const RecentTransaction = () => {
  const { data: session } = useSession();
  const accessToken = session?.user?.accessToken || "";
  const { data, isLoading, error } = useGetAllTransactionQuery(accessToken);

  if (isLoading) {
    return <RecentTransactionSkeleton />
  }

  if (error) {
   return ( <ErrorImage />)
  }


  let fetcheddata: Props[] = Array.isArray(data?.data.content) ? data.data.content : recentlistitems;

  if (fetcheddata.length > 3) {
    fetcheddata = fetcheddata.slice(-3);
  }

  console.log('Fetched data:', fetcheddata);
  console.log('Is array:', Array.isArray(fetcheddata));

  return (
    <div className='flex flex-col flex-initial flex-wrap gap-[10px] bg-white drop-shadow-xl font-medium rounded-[25px] p-[25px]'>
      {fetcheddata.length === 0 ? (
        <div>No recent transactions available</div>
      ) : (
        fetcheddata.map((value, index) => (
          <div key={index} className='flex items-center gap-3'>
            <img src={icons[index % icons.length]} alt='Icon' />
            <div className='flex flex-col gap-1'>
              <p className='text-[16px] text-[#232323] leading-[19.36px]'>
                {value.receiverUserName || recentlistitems[index].transactionName}
              </p>
              <p className='text-[15px] leading-[18.36px] text-[#718EBF]'>
                {value.date}
              </p>
            </div>
            <p className={`text-lg ml-auto ${value.amount >= 0 ? 'text-[#41D4A8]' : 'text-[#FF4B4A]'}`}>
              {value.amount >= 0 ? `+${value.amount}` : `${value.amount}`}
            </p>
          </div>
        ))
      )}
    </div>
  );
};

export default RecentTransaction;
