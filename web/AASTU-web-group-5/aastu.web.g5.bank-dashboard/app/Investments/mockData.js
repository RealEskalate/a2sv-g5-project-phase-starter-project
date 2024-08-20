// mockData.js
import Apple from '@/public/assests/icon/Investments/Group245.png'
import Google from '@/public/assests/icon/Investments/Group875.png'
import Tesla from '@/public/assests/icon/Investments/Rectangle65.png'
export const investmentsData = [
    {
      id: 1,
      name: "Tech Innovators Fund",
      service: "Equity Investment",
      value: "$10,000",
      return: 12.5,
      image:Apple ,
    },
    {
      id: 2,
      name: "Green Energy Ventures",
      service: "Mutual Fund",
      value: "$8,500",
      return: 10.2,
      image: Google,
    },
    {
      id: 3,
      name: "Real Estate Trust",
      service: "Real Estate",
      value: "$15,000",
      return: 8.3,
      image: Tesla,
    },
    {
      id: 4,
      name: "Healthcare Fund",
      service: "ETF",
      value: "$9,200",
      return: 14.1,
      image: "/assets/images/investment4.png",
    },
    {
      id: 5,
      name: "Blockchain Ventures",
      service: "Crypto Investment",
      value: "$5,000",
      return: 20.7,
      image: "/assets/images/investment5.png",
    },
    // Add more investments as needed
  ];
  
  export const tradingStockData = [
    {
      id: 1,
      name: "Apple Inc.",
      price: "$150",
      return: 6.5,
    },
    {
      id: 2,
      name: "Tesla Inc.",
      price: "$680",
      return: 7.3,
    },
    {
      id: 3,
      name: "Amazon Inc.",
      price: "$3200",
      return: 5.2,
    },
    {
      id: 4,
      name: "Google LLC",
      price: "$2700",
      return: 4.8,
    },
    {
      id: 5,
      name: "Microsoft Corp.",
      price: "$290",
      return: 6.9,
    },
    // Add more trading stocks as needed
  ];
  
  export const investmentOverview = {
    totalAmount: 50000,
    numberOfInvestments: 25,
    rateOfReturn: 11.5,
  };
  
  export const yearlyInvestmentData = [
    { year: "2020", value: 10000 },
    { year: "2021", value: 15000 },
    { year: "2022", value: 20000 },
    { year: "2023", value: 25000 },
  ];
  
  export const monthlyRevenueData = [
    { month: "Jan", value: 1000 },
    { month: "Feb", value: 1200 },
    { month: "Mar", value: 1400 },
    { month: "Apr", value: 1300 },
    { month: "May", value: 1600 },
    { month: "Jun", value: 1700 },
    { month: "Jul", value: 1800 },
    { month: "Aug", value: 1500 },
    { month: "Sep", value: 1900 },
    { month: "Oct", value: 2000 },
    { month: "Nov", value: 2100 },
    { month: "Dec", value: 2200 },
  ];
  