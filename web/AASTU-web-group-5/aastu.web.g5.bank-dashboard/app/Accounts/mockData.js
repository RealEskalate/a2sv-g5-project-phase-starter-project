export const balance = 1200;
export const income = 3000;
export const expense = 1800;
export const netBalance = 1400;
import User2 from '@/public/assests/icon/Accounts/user2.png'
import Apple from '@/public/assests/icon/Accounts/Apple.png'
import Playstation from '@/public/assests/icon/Accounts/playstation1.png'
import Rectangle from '@/public/assests/icon/Accounts/Rectangle65.png';
import Spotify from '@/public/assests/icon/Accounts/Group328.png';
import Mobile_Service from '@/public/assests/icon/Accounts/Group327.png';
export const transaction = [
  {
    image:Spotify,
    name: 'Amazon Purchase',
    date: '2024-08-14',
    type: 'Debit',
    number: '**** 1234',
    status: 'Completed',
    amount: '-$120.00'
  },
  {
    image:Mobile_Service,
    name: 'Salary Credit',
    date: '2024-08-13',
    type: 'Credit',
    number: '**** 5678',
    status: 'Completed',
    amount: '+$3000.00'
  },
  {
    image:  User2,
    name: 'Emmily wilson',
    date: '2024-08-12',
    type: 'Debit',
    number: '**** 9101',
    status: 'Completed',
    amount: '-$15.00'
  },
  {
    image:User2,
    name: 'Grocery Shopping',
    date: '2024-08-11',
    type: 'Debit',
    number: '**** 1121',
    status: 'Completed',
    amount: '-$85.00'
  },
  {
    image:User2,
    name: 'Electricity Bill',
    date: '2024-08-10',
    type: 'Debit',
    number: '**** 3141',
    status: 'Completed',
    amount: '-$100.00'
  }
];

// Invoices data
export const invoicesData = [
  {
    image: Apple,
    name: 'Apple store',
    date: '5h ago',
    amount: '$120.00',
  },
  {
   image:User2,
    name: 'Michael',
    date: '2 days ago',
    amount: '$100.00',
  },
    {
        image: Playstation,
        name: 'Playstation',
        date: '5 days ago',
        amount: '$50.00',
    },
    {
        image: Rectangle ,
        name: 'William',
        date: '10 days ago',
        amount: '$80.00',
    },
    {
        image: User2,
        name: 'Invoice #005',
        date: '2024-08-06',
        amount: '$120.00', 
    },
  // Add more invoice data if needed
];

export const dailyChartData = [
  { day: "Monday", desktop: 30, mobile: 20 },
  { day: "Tuesday", desktop: 50, mobile: 30 },
  { day: "Wednesday", desktop: 40, mobile: 25 },
  { day: "Thursday", desktop: 60, mobile: 35 },
  { day: "Friday", desktop: 70, mobile: 40 },
  { day: "Saturday", desktop: 80, mobile: 50 },
  { day: "Sunday", desktop: 90, mobile: 60 },
];
