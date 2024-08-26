import AccountInformation from '@/components/AccountInformation/AccountInformation';
import DebiteAndCredit from '@/components/Charts/DebiteAndCredit';
import InvoiceSent from '@/components/InvoiceSent/InvoiceSent';
import LastTransaction from '@/components/Transaction/LastTransaction';
import React from 'react';
import StoreProvider from '@/providers/StoreProvider';
import SingleCard from './SingleCard';
import Link from 'next/link';

export default function page() {
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
            <p className='text-[#333B69] pb-2 font-semibold'>
              <Link href='/credit-card'>See All</Link>
            </p>
          </div>
          <div>
            <StoreProvider>
              <SingleCard />
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
