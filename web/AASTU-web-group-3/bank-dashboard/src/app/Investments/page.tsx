'use client';
import React, { useEffect, useState } from 'react';
import LineGraphWithDots from '../components/investment/dotchart';
import CurveGraph from '../components/investment/curvegraph';
import Image from 'next/image';
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "@/lib/redux/store";
import { useGetRandomInvestmentDataQuery } from "@/lib/redux/api/settingApi";
import { setSetting,setError,setLoading } from "@/lib/redux/slices/settingSlice";
import {apple,google,tesla,returnValue,nameinvestment,totalinvestment} from '@/../../public/Icons'
import { RandomInvestmentData } from '@/lib/redux/types/setting';
interface Investment {
  index: number;
  title: string;
  value: number;
  icon:string;
}

interface Datas {
  titles: string;
  jobtitle: string[];
  price: number;
  investmentvalue: string;
  percent: number;
  returnval: string;
  icon:string;
}


interface tables {
  no: number;
  name: string;
  price: number;
  return: number;
}


const investor: Datas[] = [
  {icon:apple, titles: 'Apple Store', jobtitle: ['E-commerce', 'Marketplace'], price: 54000, investmentvalue: 'investment value', percent: 16, returnval: 'Return Value' },
  {icon:google, titles: 'Samsung Mobile', jobtitle: ['E-commerce', 'Marketplace'], price: 25300, investmentvalue: 'investment value', percent: -4, returnval: 'Return Value' },
  {icon:tesla, titles: 'Tesla Motors', jobtitle: ['Electric Vehicles'], price: 8200, investmentvalue: 'investmentvalue', percent: 25, returnval: 'Return Value' }
];

const tabledata: tables[] = [
  { no: 1, name: "Trivago", price: 520, return: +5 },
  { no: 2, name: "Canon", price: 480, return: +10 },
  { no: 3, name: "Uber Food", price: 350, return: -3 },
  { no: 4, name: "Nokia", price: 940, return: +2 },
  { no: 5, name: "Tiktok", price: 670, return: -12 }
];

const Page: React.FC = () => {
  const [RandInvestment , setRandInvestment] = useState<RandomInvestmentData>()
  
  const { data:investment, isLoading, isError } = useGetRandomInvestmentDataQuery({years:6,months:5});
  useEffect (()=>{
    setRandInvestment(investment)
  },[investment])
  
  console.log(RandInvestment?.data)
  if (isLoading || RandInvestment === undefined) return <div>Loading...</div>;

  const datacorner: Investment[] = [
    {icon: totalinvestment, index: 0, title: 'Total Invested Amount', value:Math.floor(RandInvestment?.data?.totalInvestment)},
    {icon: nameinvestment, index: 1, title: 'Number of Investments', value:2000 },
    {icon: returnValue, index: 2, title: 'Rate of Return', value:Math.floor(RandInvestment?.data?.rateOfReturn) }
  ];
  if (isError) return <div>{isError}</div>;




  return (
    <div className="dark:bg-[#0f1a2b] dark:text-lightText">
      <div className="justify-between grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 mb-2 py-8 px-8 w-full">
        {datacorner.map((card) => (
          <div
            key={card.index}
            className="flex justify-evenly h-[85px] rounded-2xl shadow-xl bg-white dark:bg-darkComponent p-4"
          >
            <div className="flex items-center">
              <div className="bg-[#DCFAF8] rounded-full p-4 mr-3">
                <Image
                  width={30}
                  height={30}
                  className="text-gray-500 dark:text-gray-300"
                  src={card.icon}
                  alt="Image Icon"
                />
              </div>
              <div>
                <h3 className="text-gray-500 dark:text-gray-200 text-sm">{card.title}</h3>
                <p className="text-black dark:text-gray-100 text-2xl font-bold">
                  ${card.value}
                </p>
              </div>
            </div>
          </div>
        ))}
      </div>
  
      <div className="gap-6 max-w-full flex flex-col md:flex-row px-9 py-10">
        <LineGraphWithDots TotInvestment={RandInvestment?.data?.yearlyTotalInvestment} />
        <CurveGraph MonthlyRevenue={RandInvestment?.data?.monthlyRevenue} />
      </div>
  
      <div className="flex flex-col md:flex-row w-full">
        <div className="w-full md:w-[55%] px-8">
          <h1 className="text-[22px] font-bold leading-[26.63px] text-[rgba(51,59,105,1)] dark:text-gray-100 text-left px-4 py-4">
            My Investment
          </h1>
          {investor.map((item, index) => (
            <div
              key={index}
              className="flex justify-between items-center p-6 min-h-[40px] bg-white dark:bg-darkComponent shadow-xl rounded-2xl mb-3"
            >
              <div className="flex items-center space-x-4 w-1/2">
                <div className="bg-[#FFE0EB] rounded-xl p-3">
                  <Image
                    src={item.icon}
                    alt={item.titles}
                    width={30}
                    height={30}
                    className="object-cover rounded-lg"
                  />
                </div>
                <div className="w-3/4">
                  <h2 className="text-sm font-medium text-gray-800 dark:text-gray-200">
                    {item.titles}
                  </h2>
                  <p className="text-sm font-light text-gray-500 dark:text-gray-400">
                    {item.jobtitle.join(', ')}
                  </p>
                </div>
              </div>
  
              <div className="flex items-center w-1/3 gap-3">
                <div className="flex flex-col gap-2">
                  <span className="hidden sm:block text-lg font-semibold text-gray-800 dark:text-gray-200">
                    ${item.price}
                  </span>
                  <span className="hidden sm:block text-xs font-light text-gray-500 dark:text-gray-400">
                    {item.investmentvalue}
                  </span>
                </div>
              </div>
  
              <div className="flex flex-col w-1/5">
                <span
                  className={`text-sm font-semibold ${
                    item.percent > 0 ? 'text-green-500 dark:text-green-400' : 'text-red-500 dark:text-red-400'
                  }`}
                >
                  {item.percent}%
                </span>
                <span className="hidden sm:block text-xs font-light text-gray-500 dark:text-gray-400">
                  {item.returnval}
                </span>
              </div>
            </div>
          ))}
        </div>
  
        <div className="w-11/12 md:w-[40%] py-1">
          <h1 className="text-[22px] font-bold leading-[26.63px] text-[rgba(51,59,105,1)] dark:text-gray-100 text-left px-4 py-4">
            Trending Stock
          </h1>
          <div className="flex w-full flex-col bg-white dark:bg-darkComponent shadow-lg rounded-2xl p-4 py-10 px-5">
            <div className="flex font-bold text-[16px] text-[#718EBF] dark:text-gray-300 gap-4">
              <span className="w-1/5">SL No</span>
              <span className="w-1/3">Name</span>
              <span className="w-1/4">Price</span>
              <span className="w-1/5">Return</span>
            </div>
            <div className="border-b border-[#D6DDEB] dark:border-gray-600 my-2"></div>
  
            {tabledata.map((item, index) => (
              <div
                key={index}
                className="flex text-base font-normal text-[#232323] dark:text-gray-200 gap-4 p-3"
              >
                <span className="w-1/5">{item.no}</span>
                <span className="w-1/3">{item.name}</span>
                <span className="w-1/4">${item.price}</span>
                <span
                  className={`w-1/5 ${
                    item.return > 0 ? 'text-[#16DBAA] dark:text-green-400' : 'text-[#FE5C73] dark:text-red-400'
                  }`}
                >
                  {item.return}%
                </span>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
  
};

export default Page;
