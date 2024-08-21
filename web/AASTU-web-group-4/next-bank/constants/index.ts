import BusinessLoans from "@/public/icons/BusinessLoans";
import CorporateLoans from "@/public/icons/CorporateLoans";
import CustomLoans from "@/public/icons/CustomLoans";
import PersonalLoans from "@/public/icons/PersonalLoans";

import {
  FaHome,
  FaExchangeAlt,
  FaWallet,
  FaChartLine,
  FaCreditCard,
  FaMoneyCheckAlt,
  FaCog,
  FaGift,
  FaUserShield,
} from "react-icons/fa";



// Define a mapping for your loan types
export const loanCardMapping = [
  {
    type: 'Personal Loan',
    icon: PersonalLoans,
    title: 'Personal Loans',
    descriptionKey: 'personalLoan',
  },
  {
    type: 'Corporate Loan',
    icon: CorporateLoans,
    title: 'Corporate Loans',
    descriptionKey: 'corporateLoan',
  },
  {
    type: 'Business Loan',
    icon: BusinessLoans,
    title: 'Business Loans',
    descriptionKey: 'businessLoan',
  },
];




// src/constants/index.ts

export const ITEMS_PER_PAGE = 5;

export const TABLE_HEADERS = [
  'Description',
  'Transaction ID',
  'Type',
  'Card',
  'Date',
  'Amount',
  'Receipt',
];


// src/constants/index.ts

export const transactionsData = [
  // Existing data
  {
    id: '1',
    description: 'Spotify Subscription',
    type: 'expense',
    date: '12 Jan, 10:30 PM',
    amount: '-$9.99',
  },
  {
    id: '2',
    description: 'Freelance Payment',
    type: 'income',
    date: '15 Jan, 02:00 PM',
    amount: '+$300.00',
  },
  {
    id: '3',
    description: 'Uber Ride',
    type: 'expense',
    date: '16 Jan, 08:45 AM',
    amount: '-$15.00',
  },
  {
    id: '4',
    description: 'Salary',
    type: 'income',
    date: '20 Jan, 09:00 AM',
    amount: '+$1500.00',
  },
  {
    id: '5',
    description: 'Online Course',
    type: 'expense',
    date: '22 Jan, 06:30 PM',
    amount: '-$50.00',
  },
  // Additional data for testing pagination
  {
    id: '6',
    description: 'Grocery Shopping',
    type: 'expense',
    date: '24 Jan, 01:15 PM',
    amount: '-$80.00',
  },
  {
    id: '7',
    description: 'Freelance Project Payment',
    type: 'income',
    date: '25 Jan, 03:30 PM',
    amount: '+$500.00',
  },
  {
    id: '8',
    description: 'Restaurant Bill',
    type: 'expense',
    date: '26 Jan, 08:00 PM',
    amount: '-$40.00',
  },
  {
    id: '9',
    description: 'Monthly Subscription',
    type: 'expense',
    date: '27 Jan, 07:45 AM',
    amount: '-$12.99',
  },
  {
    id: '10',
    description: 'New Laptop',
    type: 'expense',
    date: '28 Jan, 10:00 AM',
    amount: '-$1200.00',
  },
  {
    id: '11',
    description: 'Investment Earnings',
    type: 'income',
    date: '29 Jan, 11:30 AM',
    amount: '+$200.00',
  },
  {
    id: '12',
    description: 'Rent Payment',
    type: 'expense',
    date: '30 Jan, 02:00 PM',
    amount: '-$1000.00',
  },
  {
    id: '13',
    description: 'Utility Bill',
    type: 'expense',
    date: '31 Jan, 04:15 PM',
    amount: '-$150.00',
  },
  {
    id: '14',
    description: 'Gift',
    type: 'expense',
    date: '01 Feb, 09:00 AM',
    amount: '-$75.00',
  },
  {
    id: '15',
    description: 'Bonus',
    type: 'income',
    date: '02 Feb, 05:00 PM',
    amount: '+$250.00',
  },
  {
    id: '16',
    description: 'Concert Tickets',
    type: 'expense',
    date: '03 Feb, 07:30 PM',
    amount: '-$120.00',
  },
  {
    id: '17',
    description: 'Book Purchase',
    type: 'expense',
    date: '04 Feb, 03:00 PM',
    amount: '-$20.00',
  },
  {
    id: '18',
    description: 'Side Job Payment',
    type: 'income',
    date: '05 Feb, 01:30 PM',
    amount: '+$150.00',
  },
  {
    id: '19',
    description: 'Online Shopping',
    type: 'expense',
    date: '06 Feb, 08:45 PM',
    amount: '-$90.00',
  },
  {
    id: '20',
    description: 'Side Gig Earnings',
    type: 'income',
    date: '07 Feb, 11:00 AM',
    amount: '+$400.00',
  },
];



export const sidebarLinks = [
  { id:1, route: "/", label: "Dashboard", Icon: FaHome },
  { id:2, route: "/transaction", label: "Transaction", Icon: FaExchangeAlt },
  { id:3, route: "/accounts", label: "Accounts", Icon: FaWallet },
  { id:4, route: "/investments", label: "Investments", Icon: FaChartLine },
  { id:4, route: "/credit-card", label: "Credit Card", Icon: FaCreditCard },
  { id:6, route: "/loans", label: "Loans", Icon: FaMoneyCheckAlt },
  { id:7, route: "/services", label: "Services", Icon: FaCog },
  { id:8, route: "/transfer", label: "Transfer", Icon: FaGift },
  { id:9, route: "/setting", label: "Settings", Icon: FaUserShield },
];

export const user = {
  name: "John Doe",
  email: "john.doe@example.com",
  profileImage: "/path/to/profile-image.jpg", // Use a valid path for the profile image
};

export const creditcardstyles = [{
  iconwhite : "/icons/chip.png",
  icongray : "w-6 h-6"

}]

export const colors = {
  blue: 'bg-[#1814F3]',
  white: 'bg-white',
  navbartext:' text-[#343C6A]',
  black: 'bg-[#000000]',
  textwhite: 'text-white',
  textblack: 'text-[#232323]',
  textred: 'text-[#FF4B4A]',
  textgreen: 'text-[#41D4A8]',
  green :'bg-[#16DBCC]',
  red : 'bg-[#FF82AC]',
  textgray : 'text-[#718EBF]',
  gradientcard:'linear-gradient(to bottom, #4C49ED,#0A06F4)',
  gradientchart:'linear-gradient(to bottom, #2D60FF,#2D60FF)',
  textblue : "text-[#1814F3]",
  lightblue : "bg-gray-100",
  lightorange : 'bg-[#FFF5D9]',
  lightpurple : 'bg-[#E7EDFF]',
  lightgreen : 'bg-[#DCFAF8]',
  graybg : 'bg-[#F5F7FA]'



};

export const logo = {
  icon : "/icons/logo.png",
  RT1 : '/icons/RT1.png',
  RT2 : '/icons/RT2.png',
  RT3 : '/icons/RT3.png',
  transfer : '/icons/transfer2.png',
}

export const textColors = {
  textWhite: 'text-white',
  textDimmed: 'text-[#718EBF]',
}