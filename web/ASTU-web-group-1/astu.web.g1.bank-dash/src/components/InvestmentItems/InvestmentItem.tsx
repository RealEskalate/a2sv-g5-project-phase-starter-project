'use client';
import Image from 'next/image';
import React from 'react';
import InvestmentCard from './investmentCard';
import { useGetInvestmentItemsQuery } from '@/lib/redux/slices/investmentSlice';
import { SkeletonCard } from '../AllSkeletons/investmentItemSkeleton/SkeletonCard';

const InvestmentItem = () => {
  const { data, isLoading } = useGetInvestmentItemsQuery({
    years: 2023,
    months: 5,
  });

  const Alldata = data?.data;

  // console.log(Alldata);
  if (isLoading) {
    return <SkeletonCard />;
  }
  const numberOfInvestments = Alldata?.yearlyTotalInvestment?.length;

  return (
    <div className='flex flex-col md:flex-row gap-2 md:gap-2 justify-evenly w-full'>
      <InvestmentCard
        image={'/assets/icons/moneyBag.svg'}
        name={'Total Invested Amount'}
        amount={`$${Alldata?.totalInvestment.toFixed(2)}`}
      />
      <InvestmentCard
        image={'/assets/icons/numberof-investment.svg'}
        name={'Number of Investments'}
        amount={`${numberOfInvestments}`}
      />
      <InvestmentCard
        image={'/assets/icons/rate-return.svg'}
        name={'Rate of Return'}
        amount={`+${Alldata?.rateOfReturn.toFixed(2)}%`}
      />
    </div>
  );
};

export default InvestmentItem;
