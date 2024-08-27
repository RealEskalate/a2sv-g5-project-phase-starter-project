'use client';
import React from 'react';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';

interface MonthlyRevenueType {
  MonthlyRevenue: {
    time: string;
    value: number;
  }[];
}

const CurveGraph = ({ MonthlyRevenue }: MonthlyRevenueType) => {
  const chartData = MonthlyRevenue.map(item => ({
    time: item.time,  
    value: item.value,   
  })).reverse();

  return (
    <div className="w-full md:w-3/5">
    <h1 className="text-[22px] font-bold leading-[26.63px] text-[rgba(51,59,105,1)] dark:text-gray-100 text-left px-4 py-4">
      Monthly Revenue
    </h1>
  
    <div className="rounded-3xl h-[300px] bg-white dark:bg-darkComponent shadow-md p-8">
      <ResponsiveContainer className="w-full">
        <LineChart
          data={chartData}
          margin={{
            top: 5,
            right: 25,
            left: 5,
            bottom: 5,
          }}
        >
          <CartesianGrid strokeDasharray="3" vertical={false} stroke="#d3d3d3" />
          <XAxis
            dataKey="time"
            axisLine={false}
            tickLine={false}
            domain={[0, 'auto']}
            padding={{ left: 10, right: 30 }}
            tickMargin={4}
            stroke="#718EBF"
            tick={{ fill: "#718EBF", className: "dark:fill-gray-300" }}
          />
          <YAxis
            tick={{ dx: -5, dy: 15, fill: "#718EBF", className: "dark:fill-gray-300" }}
            dataKey="value"
            axisLine={false}
            tickLine={false}
            padding={{ top: 20, bottom: 20 }}
            tickMargin={3}
            stroke="#718EBF"
            tickCount={5}
          />
          <Tooltip contentStyle={{ backgroundColor: '#333', borderColor: '#333', color: '#fff' }} />
          <Line
            type="monotone"
            dataKey="value"
            dot={false}
            fontWeight={400}
            stroke="#16DBCC"
            strokeWidth={3}
            className="dark:stroke-[#16DBCC]"
          />
        </LineChart>
      </ResponsiveContainer>
    </div>
  </div>
  );  
};

export default CurveGraph;
