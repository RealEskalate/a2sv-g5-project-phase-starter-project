"use client";

import React from 'react';
import { Line, LineChart, CartesianGrid, XAxis, YAxis, Tooltip, ResponsiveContainer } from "recharts";
import { Card, CardContent } from "../../@/components/ui/card";

interface YearlyInvestmentData {
  time: string;
  value: number;
}

interface ChartCardInvestProps {
  data: YearlyInvestmentData[];
}

const CustomTooltip = ({ active, payload, label }: TooltipProps<number, string>) => {
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

export default function ChartCard_Invest({ data }: ChartCardInvestProps) {
  return (
    <Card style={{ height: '100%' }}>
      <CardContent style={{ height: '100%' }}>
        <div className="pt-6" style={{ width: '100%', height: '100%' }}>
          <ResponsiveContainer width="100%" height="100%">
            <LineChart
              margin={{
                top: 5, right: 30, left: 20, bottom: 5,
              }}
              data={data}
            >
              <CartesianGrid strokeDasharray="3 3" />
              <XAxis 
                dataKey="time" 
                axisLine={false} 
                tickMargin={10} 
              />
              <YAxis 
                axisLine={false} 
                tickFormatter={(value) => `$${value.toLocaleString()}`} 
                tickMargin={10} 
              />
              <Tooltip content={<CustomTooltip />} />
              <Line type="linear" dataKey="value" stroke="#ff7300" strokeWidth={3} />
            </LineChart>
          </ResponsiveContainer>
        </div>
      </CardContent>
    </Card>
  );
}
