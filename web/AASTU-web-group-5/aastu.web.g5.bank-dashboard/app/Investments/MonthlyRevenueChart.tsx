"use client";

import React from 'react';
import { CartesianGrid, Line, LineChart, Tooltip, XAxis, YAxis, ResponsiveContainer } from "recharts";
import { Card, CardContent } from "../../@/components/ui/card";

// Define the shape of the data prop
interface RevenueData {
  time: string;  // Assuming 'time' represents the month
  value: number; // The revenue value
}

interface MonthlyRevenueChartProps {
  data: RevenueData[];  // Prop for passing revenue data from parent
}

const CustomTooltip = ({ active, payload, label }: { active?: boolean; payload?: any; label?: any; }) => {
  if (active && payload && payload.length) {
    const value = payload[0]?.value;
    if (value !== undefined) {
      return (
        <div className="custom-tooltip bg-white p-2 border border-gray-300 rounded">
          <p className="label">{`${label} : $${value.toLocaleString()}`}</p>
        </div>
      );
    }
  }
  return null;
};

export default function MonthlyRevenueChart({ data }: MonthlyRevenueChartProps) {
  return (
    <Card style={{ height: '100%' }}>
      <CardContent style={{ height: '100%' }}>
        <div className='pt-6' style={{ width: '100%', height: '100%' }}>
          <ResponsiveContainer width="100%" height="100%">
            <LineChart
              data={data}
              margin={{
                top: 5, right: 20, left: 5, bottom: 5,
              }}
            >
              <CartesianGrid strokeDasharray="3 3" />
              <XAxis dataKey="time" />
              <YAxis tickFormatter={(value) => `$${value}`} />
              <Tooltip content={<CustomTooltip />} />
              <Line type="monotone" dataKey="value" stroke="#82ca9d" strokeWidth={3} dot={false} />
            </LineChart>
          </ResponsiveContainer>
        </div>
      </CardContent>
    </Card>
  );
}
