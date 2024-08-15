import React from 'react';
import InvestmentItem from '@/components/InvestmentItems/InvestmentItem';
import InvestmentList from '@/components/InvestmentList/InvestmentList';
import TrendingList from '@/components/TrendingStock/trendingList';

export default function page() {
    return (
    <div>
      <InvestmentItem/>
      <div className='border h-[282px]'>

      </div>
      <div className="flex flex-col md:flex-row">
      <InvestmentList />
      <TrendingList />
      </div>
  </div>
  );
}
