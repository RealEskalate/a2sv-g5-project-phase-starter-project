"use client";
import {
  CartesianGrid,
  Area,
  AreaChart,
  Tooltip,
  XAxis,
  YAxis,
  ResponsiveContainer,
  ReferenceLine,
} from 'recharts';

const data = [
  { name: 'Jan', totalIncome: 323 },
  { name: 'Feb', totalIncome: 424 },
  { name: 'Mar', totalIncome: 515 },
  { name: 'Apr', totalIncome: 602 },
  { name: 'May', totalIncome: 518 },
  { name: 'Jun', totalIncome: 951 },
  { name: 'Jul', totalIncome: 615 },
  { name: 'Aug', totalIncome: 815 },
  { name: 'Sep', totalIncome: 715 },
  { name: 'Oct', totalIncome: 915 },
  { name: 'Nov', totalIncome: 715 },
  { name: 'Dec', totalIncome: 515 },
];
const BalanceHistory = () => {
  return (
    <div className='w-full md:w-7/12'>
      <h1 className='text-[#333B69] pb-3 font-semibold'>Balance History</h1>
      <div className=' bg-white px-3 py-5 rounded-3xl'>
        <ResponsiveContainer width={'100%'} height={240}>
          <AreaChart data={data} margin={{ top: 10, right: 30, left: 0, bottom: 0 }}>
            <defs>
              <linearGradient id='color' x1='0' y1='0' x2='0' y2='1'>
                <stop offset='5%' stopColor='#1814f3' stopOpacity={0.8} />
                <stop offset='95%' stopColor='#1814f3' stopOpacity={0} />
              </linearGradient>
            </defs>
            <XAxis dataKey='name' tick={{ fontSize: 10 }} />
            <YAxis tick={{ fontSize: 10 }} />
            <CartesianGrid strokeDasharray='3 3' />
            <Tooltip />
            <ReferenceLine x='Page C' stroke='green' label='Min PAGE' />
            <ReferenceLine y={4000} label='Max' stroke='red' strokeDasharray='3 3' />
            <Area
              type='monotone'
              dataKey='totalIncome'
              stroke='#1814f3'
              strokeWidth={3}
              fillOpacity={1}
              fill='url(#color)'
            />
          </AreaChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
};

export default BalanceHistory;
