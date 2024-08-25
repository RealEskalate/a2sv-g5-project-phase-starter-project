'use client';
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from 'recharts';

const data = [
  { name: 'Sat', Deposit: 500, Withdraw: 100 },
  { name: 'Sun', Deposit: 300, Withdraw: 200 },
  { name: 'Mon', Deposit: 200, Withdraw: 300 },
  { name: 'Tue', Deposit: 400, Withdraw: 200 },
  { name: 'Wed', Deposit: 100, Withdraw: 300 },
  { name: 'Thu', Deposit: 400, Withdraw: 200 },
  { name: 'Fri', Deposit: 300, Withdraw: 300 },
];

const WeeklyActivityChart = () => {
  return (
    <ResponsiveContainer width="100%" height={280} className='bg-white rounded-2xl'>
      <BarChart
        data={data}
        margin={{
          top: 20,
          right: 20,
          left: 20,
          bottom: 5,
        }}
        barCategoryGap='20%'
      > 
        <CartesianGrid strokeDasharray="0" horizontal={true} vertical={false} />
        <XAxis dataKey="name" />
        <YAxis domain={[0, 500]} tickCount={6} />
        <Tooltip />
        <Bar dataKey="Deposit" fill="#0000FF" radius={[10, 10, 0, 0]} barSize={15} />
        <Bar dataKey="Withdraw" fill="#00E0E0" radius={[10, 10, 0, 0]} barSize={15} />
        <Legend verticalAlign='top' align='right' />
      </BarChart>
    </ResponsiveContainer>
  );
};

export default WeeklyActivityChart;
