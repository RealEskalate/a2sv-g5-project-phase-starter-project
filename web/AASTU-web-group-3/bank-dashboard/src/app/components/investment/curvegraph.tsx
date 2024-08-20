'use client';
import React, { PureComponent } from 'react';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';

const data = [
  {
    name: '2016',
    pv: 15000,
    amt: 10000,
  },
  {
    name: '2017',
    pv: 30000,
    amt: 20000,
  },
  {
    name: '2018',
    pv: 9000,
    amt: 30000,
  },
  {
    name: '2019',
    pv: 20000,
    amt: 40000,
  },
  {
    name: '2020',
    pv: 24000,
    amt: 50000,
  },
  {
    name: '2021',
    pv: 45000,
    amt: 60000,
  },
];

export default class CurveGraph extends PureComponent {
  render() {
    return (
    <div className='w-full md:w-1/2'>
      <h1 className="text-[22px] font-bold leading-[26.63px] text-[rgba(51,59,105,1)] text-left px-4 py-4">
      Monthly Revenue
      </h1>
   
        <div className='rounded-3xl h-[300px] bg-white shadow-md p-8'>

<ResponsiveContainer width= {400} >
        <LineChart
          data={data}
          margin={{
              top: 5,
              right: 25,
              left: 5,
              bottom: 5,
            }}
            >
          <CartesianGrid strokeDasharray="3" vertical={false} />
          <XAxis
            dataKey="name"
            axisLine={false}
            tickLine={false}
            domain={[0, 'auto']}
            padding={{ left: 10, right: 30 }} 
            tickMargin={4}
            stroke="#718EBF"
            />
          <YAxis
            
            tick={{ dx: -5, dy: 15 }}
            dataKey="amt"
            axisLine={false}
            tickLine={false}
            padding={{ top: 20, bottom: 20 }}
            tickMargin={3}
            stroke="#718EBF"
            tickCount={5} 
            />
          <Tooltip />
          <Line
            type="monotone"
            dataKey="pv"
            dot={false}
            fontWeight={400}
            stroke="#16DBCC"
            strokeWidth={3}
            />
        </LineChart>
      </ResponsiveContainer>
      </div>
            </div>
    );
  }
}
