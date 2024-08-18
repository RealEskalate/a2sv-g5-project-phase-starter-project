import React from 'react';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';

const TinyLineGraphWithDots = () => {
  // Dummy data for the graph
  const data = [
    { year: '2016', value: 16000 },
    { year: '2017', value: 40000 },
    { year: '2018', value: 20000 },
    { year: '2019', value: 30000 },
    { year: '2020', value: 20000 },
    { year: '2021', value: 50000 },
  ];

  return (
    
    <div className=' w-full md:w-[40%]'>
      <h1 className="text-[22px] font-bold leading-[26.63px] text-[rgba(51,59,105,1)] text-left px-4 py-4">
        Monthly Revenue
      </h1>
    <div className="w-full  h-[300px] bg-white rounded-2xl p-4 shadow-lg">
      <ResponsiveContainer width="100%" height="100%">
        <LineChart
          data={data}
          margin={{ top: 20, right: 20, bottom: 20, left: 10 }} // Add padding around the chart
        >
          <CartesianGrid strokeDasharray="3" stroke="#DFE5EE" />
          <XAxis
            dataKey="year"
            stroke="#718EBF"
            tick={{ fill: '#718EBF' }}
            axisLine={false} // Hide axis line
            tickLine={false} // Hide tick lines
            />
          <YAxis
            tickFormatter={(value) => `$${value}`}
            stroke="#718EBF"
            tick={{ fill: '#718EBF' }}
            axisLine={false} // Hide axis line
            tickLine={false} // Hide tick lines
            />
          <Tooltip formatter={(value) => `$${value}`} />
          <Line
            type="monotone"
            dataKey="value"
            stroke="#16DBCC"
            strokeWidth={4}
            dot={{ stroke: '#16DBCC', strokeWidth: 3, r: 0, fill: '#fff' }} // Customize dots
            />
        </LineChart>
      </ResponsiveContainer>
    </div>
            </div>
  );
};

export default TinyLineGraphWithDots;
