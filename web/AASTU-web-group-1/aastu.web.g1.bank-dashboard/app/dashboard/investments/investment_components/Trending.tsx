'use client'
import React from 'react';
import TrendingTable from './TrendingTable';
import { useUser } from '@/contexts/UserContext';
export default function Trending() {
  const { isDarkMode } = useUser();
    return (
        <div className='w[95%] md:w-[25%] mx-auto md:mx-0 md:min-w-[450px]'>
          <h1  className={`my-4 font-[600] text-[22px] ${isDarkMode ? "text-white":"text-[#333B69]"} ml-5 md:ml-0`} >Trending Stock</h1>
          <TrendingTable />
        </div>
      )
}
