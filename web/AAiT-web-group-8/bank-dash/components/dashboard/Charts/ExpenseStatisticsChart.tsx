'use client'
import { LabelProps } from '@/types/index.';
import { PieChart, Pie, Cell, ResponsiveContainer } from 'recharts';

const data = [
  { name: 'Transfer', value: 30, color: '#333b76' },
  { name: 'Service', value: 15, color: '#FF7F00' },
  { name: 'Shopping', value: 20, color: '#FF00FF' },
  { name: 'Others', value: 35, color: '#0000FF' },
];

const RADIAN = Math.PI / 180;
const renderCustomizedLabel = ({
  cx,
  cy,
  midAngle,
  innerRadius,
  outerRadius,
  percent,
  index,
}: LabelProps) => {
  const radius = innerRadius + (outerRadius - innerRadius) * 0.5;
  const x = cx + radius * Math.cos(-midAngle * RADIAN);
  const y = cy + radius * Math.sin(-midAngle * RADIAN);

  return (
    <text
      x={x}
      y={y}
      fill="white"
      textAnchor="middle"
      dominantBaseline="central"
      fontSize={13}
      fontWeight={500}
    >
      {`${data[index].value}% ${data[index].name}`}
    </text>
  );
};

const ExpenseStatisticsChart = () => {
  return (
    <ResponsiveContainer width="100%" height={210} className="bg-white rounded-xl">
      <PieChart>
        <Pie
          data={data}
          cx="50%"
          cy="50%"
          labelLine={false}
          label={renderCustomizedLabel}
          outerRadius={90}
          innerRadius={4}
          fill="#8884d8"
          dataKey="value"
        >
          {data.map((entry, index) => (
            <Cell key={`cell-${index}`} fill={entry.color} />
          ))}
        </Pie>
      </PieChart>
    </ResponsiveContainer>
  );
};

export default ExpenseStatisticsChart;
