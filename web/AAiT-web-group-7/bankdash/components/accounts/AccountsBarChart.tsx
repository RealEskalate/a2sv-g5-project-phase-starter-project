'use client';
import {
  BarChart,
  Bar,
  XAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from 'recharts';

const data = [
  { name: 'Sat', Debit: 500, Credit: 100 },
  { name: 'Sun', Debit: 300, Credit: 200 },
  { name: 'Mon', Debit: 200, Credit: 300 },
  { name: 'Tue', Debit: 400, Credit: 200 },
  { name: 'Wed', Debit: 100, Credit: 300 },
  { name: 'Thu', Debit: 400, Credit: 200 },
  { name: 'Fri', Debit: 300, Credit: 300 },
];

const AccountsBarChartComponent = () => {
  return (
    <ResponsiveContainer width="100%" height={300} className='bg-white rounded-3xl'>
      <BarChart
        data={data}
        margin={{
          top: 20,
          right: 30,
          left: 20,
          bottom: 5,
        }}
        barCategoryGap='20%'
      > 
        <CartesianGrid strokeDasharray="3 3" horizontal={false} vertical={false} />
        <XAxis dataKey="name" />
        <Tooltip />
        <Bar dataKey="Debit" fill="#1A16F3" radius={[10, 10, 10, 10]}/>
        <Bar dataKey="Credit" fill="#FCAA0B" radius={[10, 10,10, 10]} />
        <Legend verticalAlign='top' align='right' />
      </BarChart>
    </ResponsiveContainer>
  );
};

export default AccountsBarChartComponent;
