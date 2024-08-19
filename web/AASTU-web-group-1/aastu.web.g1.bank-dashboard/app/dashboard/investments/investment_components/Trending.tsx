import React from 'react';
import TrendingTable from './TrendingTable';
export default function Trending() {
    return (
        <div className='w-[95%] md:w-[25%] mx-auto md:mx-0'>
          <h1  className="my-4 font-[600] text-[22px] text-[#333B69] ml-5 md:ml-0" >Trending Stock</h1>
          <TrendingTable />
        </div>
      )
}
