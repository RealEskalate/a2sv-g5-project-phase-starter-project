'use client';
import React from 'react';
import Linechart from './components/linechart'; 
import LineGraphWithDots from './components/dotchart';

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
      <h1> team 3 </h1>
      </div>
  );
};

export default Page;