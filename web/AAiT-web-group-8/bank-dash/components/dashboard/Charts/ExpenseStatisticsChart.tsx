'use client';

import { PieChart, Pie, Cell, ResponsiveContainer, Sector } from 'recharts';
import { useState } from 'react';
import { LabelProps } from '@/types/index.';

const data = [
  { name: 'Service', value: 15, color: '#FF7F00' },
  { name: 'Transfer', value: 30, color: '#333b76' },
  { name: 'Others', value: 35, color: '#FF00FF' },
  { name: 'Shopping', value: 20, color: '#0000FF' },
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
  const x = cx + radius * Math.cos(-midAngle * RADIAN) + 10; // Adjusted x coordinate
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
      <tspan x={x} dy="-0.8em">{`${data[index].value}%`}</tspan>
      <tspan x={x} dy="1.2em">{data[index].name}</tspan>
    </text>
  );
};

const renderActiveShape = (props: any) => {
  const {
    cx, cy, innerRadius, outerRadius, startAngle, endAngle, fill, payload, value,
  } = props;
  return (
    <g>
      <Sector
        cx={cx}
        cy={cy}
        innerRadius={innerRadius}
        outerRadius={outerRadius + 10}
        startAngle={startAngle}
        endAngle={endAngle}
        fill={fill}
      />
      <text x={cx} y={cy} dy={8} textAnchor="middle" fill={fill}>
        {`${value}%`}
      </text>
      <text x={cx} y={cy} dy={24} textAnchor="middle" fill={fill}>
        {payload.name}
      </text>
    </g>
  );
};

const ExpenseStatisticsChart = () => {
  const [activeIndex, setActiveIndex] = useState<number | null>(-1);

  const onPieEnter = (_: any, index: number) => {
    setActiveIndex(index);
  };

  const onPieLeave = () => {
    setActiveIndex(-1);
  };

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
          innerRadius={0}
          fill="#8884d8"
          dataKey="value"
          isAnimationActive={true}
          animationDuration={800}
          activeIndex={activeIndex}
          activeShape={renderActiveShape}
          onMouseEnter={onPieEnter}
          onMouseLeave={onPieLeave}
          paddingAngle={5}
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