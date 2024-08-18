'use client';
import AccountInformation from '@/components/AccountInformation/AccountInformation';
import DebiteAndCredit from '@/components/Charts/DebiteAndCredit';
import MonthlyRevenue from '@/components/Charts/MonthlyRevenue';
import InvoiceSent from '@/components/InvoiceSent/InvoiceSent';
import MyCard from '@/components/MyCard/MyCard';
import LastTransaction from '@/components/Transaction/LastTransaction';
import { getServerSession } from 'next-auth';
import { getSession } from 'next-auth/react';
import React from 'react';
import { authOptions } from '../../api/auth/[...nextauth]/options';
import { ApiProvider } from '@reduxjs/toolkit/query/react';
import StoreProvider from '@/providers/StoreProvider';

export default async function page() {
  // const session = await getServerSession(authOptions);
  // console.log(session, 'session is from accounts page');
  return (
    <>
      <AccountInformation />
      <div className=' min-[890px]:flex min-[890px]:space-x-4 lg:space-x-10 mb-5'>
        <LastTransaction />
        <div className='mb-5'>
          <div className='flex justify-between'>
            <h1 className='text-[#333B69] pb-2 font-semibold'>My Card</h1>
            <p className='text-[#333B69] pb-2 font-semibold'>See All</p>
          </div>
          <div>
            <StoreProvider>
              <MyCard />
            </StoreProvider>
          </div>
        </div>
      </div>
      <div className='min-[890px]:flex min-[890px]:space-x-4 lg:space-x-10'>
        <DebiteAndCredit />
        <InvoiceSent />
      </div>
    </>
  );
}
