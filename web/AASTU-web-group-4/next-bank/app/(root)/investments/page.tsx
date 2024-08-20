import React from 'react'
import LineChartNoBg from '@/components/LineChartNoBg'
import LineChartStright from '@/components/LineChartStraight'
import {colors} from '@/constants/index'
import { FaSackDollar } from "react-icons/fa6";
import { AiOutlineRetweet } from "react-icons/ai";
import { GiTakeMyMoney } from "react-icons/gi";
import  MyInvestment from '@/components/MyInvestment'
import TrendingStock from '@/components/TrendingStock';
import { getTrendingCompanies } from '@/services/companygetch';


  const data = [
    {
      icon: "/icons/apple_store.png",
      color: "bg-red-100 ",
      colortext: colors.textblack,
      category: "E-commerce, marketplace",
      categorycolor: colors.textgray,
      name: "Apple Store",
      amount: "54000",
      percentage: "1.6%",
    },
    {
      icon: "/icons/Google_store.png",
      color: "bg-blue-100",
      colortext: colors.textblack,
      category: "E-commerce, marketplace",
      categorycolor: colors.textgray,
      name: "Google Store",
      amount: "25000",
      percentage: "2.23%",
    },
    {
      icon: "/icons/tesla.png",
      color: "bg-yellow-100",
      colortext: colors.textblack,
      category: "E-commerce, marketplace",
      categorycolor: colors.textgray,
      name: "Tesla Store",
      amount: "95000",
      percentage: "2.23%",
    },

  ];
  const trendingdata = [
    { slNo: '01.', name: 'Nokia', price: '$940', return: '+2%' },
    { slNo: '02.', name: 'Apple', price: '$1500', return: '+5%' },
    { slNo: '03.', name: 'Google', price: '$2500', return: '-3%' },
    { slNo: '04.', name: 'Amazon', price: '$3000', return: '+4%' },
    { slNo: '05.', name: 'Microsoft', price: '$2000', return: '-6%' },
  ];

const Investments = async () => {

  const fetch = async () => {
   
      try {
        const trendingcomp = await getTrendingCompanies();
        return trendingcomp
       
      } catch (error) {
        console.error('Login Error:', error);
      }
    };

   const  Trendingcomp = fetch()

  return (
    <div className={` ${colors.graybg}   flex  flex-col  lg:gap-5 lg:ml-64 lg:pr-6 xl:pr-10 `}>

      <div className='flex flex-col items-center px-6 pt-10 gap-4 lg:flex-row ' >
        <div className='flex gap-3 w-[80%] bg-white justify-center items-center py-3  rounded-xl'>
          <div className='bg-cyan-100 w-[50px] h-[50px] flex items-center justify-center rounded-full  '>
            <FaSackDollar className='text-cyan-500 h-[25px] w-[20px] '/>
          </div>
          <div>
            <p className={`${colors.textgray} font-normal text-[12px]`}>Total Invested Amount</p>
            <p className={`${colors.textblack} font-semibold text-[16px] `}>$10000</p>
          </div>
        </div>

        <div className='flex gap-3 w-[80%] bg-white justify-center items-center py-3  rounded-xl'>
          <div className='bg-pink-100 w-[50px] h-[50px] flex items-center justify-center rounded-full  '>
            <GiTakeMyMoney className='text-pink-500 h-[25px] w-[20px] '/>
          </div>
          <div>
            <p className={`${colors.textgray} font-normal text-[12px]  `}>Number of Investments</p>
            <p className={`${colors.textblack} font-semibold text-[16px] `}>1809</p>
          </div>
        </div>

        <div className='flex gap-3 w-[80%] bg-white justify-center items-center py-3  rounded-xl'>
          <div className='bg-indigo-100 w-[50px] h-[50px] flex items-center justify-center rounded-full  '>
            <AiOutlineRetweet className='text-indigo-500 h-[25px] w-[20px] '/>
          </div>
          <div>
            <p className={`${colors.textgray} font-normal text-[12px] w-[132px] `}>Rate of Return</p>
            <p className={`${colors.textblack} font-semibold text-[16px] `}>+5.8%</p>
          </div>
        </div>

      </div>

      <div className='  flex flex-col py-5 px-6 gap-14 lg:grid lg:grid-cols-2 lg:gap-6  '>
        <div className='flex flex-col gap-3 lg:gap-4 xl:gap-5 '>
          <h2 className={`font-semibold text-[22px] ${colors.navbartext} `}>Yearly Total Investments</h2>
          <LineChartStright/>
        </div>
        <div className='flex  flex-col gap-3 lg:gap-4 xl:gap-5 '>
          <h2  className={`font-semibold text-[22px] ${colors.navbartext} `} >Monthly Revenue</h2>
          <LineChartNoBg/>
        </div>
      </div>

      <div className='flex flex-col lg:grid lg:grid-cols-5 '>
        <div className='px-6  lg:col-span-3 flex flex-col gap-5'>
          <h2 className={`font-semibold text-[22px] ${colors.navbartext} `} >My Investment</h2>
          {
            data.map((item,index)=>(
              <MyInvestment key={index} icon={item.icon} color={item.color} colortext={item.colortext} category={item.category} categorycolor={item.categorycolor} name={item.name} amount={item.amount} percentage={item.percentage}/>
            ))
          }
        </div>
        <div className='lg:col-span-2' >

       <div className='p-6 lg:p-0'>
       <h2 className="text-lg font-semibold text-gray-800 mb-3">Trending Stock</h2>

         <TrendingStock items = {trendingdata}/>
       </div>
        </div>
      </div>

    </div>
  )
}

export default Investments