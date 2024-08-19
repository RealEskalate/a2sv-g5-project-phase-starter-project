'use client';
import React, { PureComponent } from 'react';
import { BarChart, Bar, XAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';

const data = [
  {
    name: 'Sat',
    Debit: 4000,
    Credit: 2400,
  },
  {
    name: 'Sun',
    Debit: 3000,
    Credit: 1398,
  },
  {
    name: 'Mon',
    Debit: 2000,
    Credit: 9800,
  },
  {
    name: 'Tue',
    Debit: 2780,
    Credit: 3908,
  },
  {
    name: 'Wed',
    Debit: 1890,
    Credit: 4800,
  },
  {
    name: 'Thu',
    Debit: 2390,
    Credit: 3800,
  },
  {
    name: 'Fri',
    Debit: 3490,
    Credit: 4300,
  },
];


export default class LineChart extends PureComponent {
  render() {
    return (
      <div className="w-full h-80">
        <div className="flex justify-between mb-2">
          {/* Left text */}
          
          <Legend layout="horizontal" verticalAlign="top" align="right" />
        </div>
        
        <ResponsiveContainer>
          <BarChart
            width={80}
            height={120}
            data={data}
            margin={{
              top: 1,
              right: 30,
              left: 0,
              bottom: 5,
            }}
          >
            
             <XAxis dataKey="name" axisLine={false} tickLine={false} />
            
            <Bar dataKey="Debit" fill="#1A16F3" radius={[10, 10, 10, 10]} />
            <Bar dataKey="Credit" fill="#FCAA0B" radius={[10, 10, 10, 10]} />
          </BarChart>
        </ResponsiveContainer>
        </div>
    );
  }
}
