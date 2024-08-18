'use client';
import React from 'react';
import Linechart from '../components/linechart'; 
import LineGraphWithDots from '../components/dotchart';

interface Investment {
  index: number;
  title: string;
  value: number;
}

interface Datas {
  titles: string;
  jobtitle: string[];
  price: number;
  investmentvalue: string;
  percent: number;
  returnval: string;
}

interface tables {
  no: number;
  name: string;
  price: number;
  return: number;
}

const datacorner: Investment[] = [
  { index: 0, title: 'Total Invested Amount', value: 150000 },
  { index: 1, title: 'Number of Investments', value: 1250 },
  { index: 2, title: 'Rate of Return', value: 5.69 }
];

const investor: Datas[] = [
  { titles: 'Apple Store', jobtitle: ['E-commerce', 'Marketplace'], price: 54000, investmentvalue: 'investment value', percent: 16, returnval: 'Return Value' },
  { titles: 'Samsung Mobile', jobtitle: ['E-commerce', 'Marketplace'], price: 25300, investmentvalue: 'investment value', percent: -4, returnval: 'Return Value' },
  { titles: 'Tesla Motors', jobtitle: ['Electric Vehicles'], price: 8200, investmentvalue: 'investmentvalue', percent: 25, returnval: 'Return Value' }
];

const tabledata: tables[] = [
  { no: 1, name: "Trivago", price: 520, return: +5 },
  { no: 2, name: "Canon", price: 480, return: +10 },
  { no: 3, name: "Uber Food", price: 350, return: -3 },
  { no: 4, name: "Nokia", price: 940, return: +2 },
  { no: 5, name: "Tiktok", price: 670, return: -12 }
];

const Page: React.FC = () => {
  return (
    <div>
      <div className="justify-between grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-2 py-8 px-8 w-auto">
        {datacorner.map((card) => (
          <div key={card.index} className="flex justify-evenly h-[85px] rounded-2xl shadow-xl bg-white p-4">
            <div className="flex items-center">
              <img className="w-12 h-12 text-gray-500 mr-4" src={`/images/${card.index + 11}.png`} alt="Image Icon" />
              <div>
                <h3 className="text-gray-500 text-sm">{card.title}</h3>
                <p className="text-black text-2xl font-bold">${card.value}</p>
              </div>
            </div>
          </div>
        ))}
      </div>

      <div className="gap-6 md:gap-32 max-w-full flex flex-col md:flex-row px-9 py-10">
        

        <LineGraphWithDots />
        <Linechart />
      </div>
      <div className='flex flex-col md:flex-row w-full'>

      <div className='w-full md:w-2/3 px-8'>
      <h1 className="text-[22px] font-bold leading-[26.63px] text-[rgba(51,59,105,1)] text-left px-4 py-4">
        My Investment
      </h1>
        {investor.map((item, index) => (
          <div key={index} className="flex justify-between items-center p-6 min-h-[40px] bg-white shadow-xl rounded-2xl mb-3">
            <div className="flex items-center space-x-4 w-4/2">
              <img 
                src={`/images/${index + 14}.png`} 
                alt={item.titles} 
                className="w-16 h-16 object-cover rounded-lg"
              />
              <div>
                <h2 className="text-xl font-medium text-gray-800 items-start">{item.titles}</h2>
                <p className="text-sm font-light text-gray-500">
                  {item.jobtitle.join(', ')}
                </p>
              </div>
            </div> 

            <div className="flex items-center px-20 w-1/3 gap-">
              <div className="flex flex-col gap-2 place-items-center">
                <span className="hidden sm:block text-2xl font-semibold text-gray-800">${item.price}</span>
                <span className="hidden sm:block text-sm font-light text-gray-500">{item.investmentvalue}</span>
              </div>
            </div>

            <div className="flex flex-col px-10 w-1/3 items-end">
              <span className={`text-sm font-semibold place-items-center ${item.percent > 0 ? 'text-green-500' : 'text-red-500'}`}>
                {item.percent}%
              </span>
              <span className="hidden sm:block text-sm font-light text-gray-500">{item.returnval}</span>
            </div>
          </div>
        ))}
      </div>

  <div className="w-full md:w-[40%] py-1 px-10">
  <h1 className="text-[22px] font-bold leading-[26.63px] text-[rgba(51,59,105,1)] text-left px-4 py-4">
        Trending Stock
      </h1>
  <div className="h-full flex flex-col bg-white shadow-lg rounded-2xl p-4 py-10 px-5">
    <div className="flex justify-between font-bold text-[16px] text-[#718EBF] gap-8">
      <span className="flex-1">SL No</span>
      <span className="flex-1">Name</span>
      <span className="flex-1">Price</span>
      <span className="flex-1">Return</span>
    </div>
    <div className="border-b border-[#D6DDEB] my-2"></div>

    {tabledata.map((item, index) => (
      <div key={index} className="flex justify-between text-[16px] font-normal text-[#232323] gap-8 p-3">
        <span className="flex-1">{item.no}</span>
        <span className="flex-1">{item.name}</span>
        <span className="flex-1">${item.price}</span>
        <span className= {`flex-1 ${item.return>0 ? `text-[#16DBAA]`:`text-[#FE5C73]`} `} >{item.return}%</span>

      </div>
    ))}
  </div>
</div>

      </div>
    </div>
  );
};

export default Page;