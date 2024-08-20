"use client";

import React from 'react';
import { Line, LineChart, CartesianGrid, XAxis, YAxis, Tooltip, TooltipProps, ResponsiveContainer } from "recharts";
import { Card, CardContent } from "../../@/components/ui/card";

const yearlyInvestmentData = [
  { year: "2020", totalInvestment: "$10,000" },
  { year: "2021", totalInvestment: "$12,000" },
  { year: "2022", totalInvestment: "$9,000" },
  { year: "2023", totalInvestment: "$15,000" },
  { year: "2024", totalInvestment: "$13,000" },
].map(item => ({
  ...item,
  totalInvestment: Number(item.totalInvestment.replace(/[^0-9.-]+/g, "")),
}));

console.log("Yearly Investment Data:", yearlyInvestmentData);

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

export default function YearlyInvestmentChart() {
  return (
    <Card style={{ height: '100%' }}>
      <CardContent style={{ height: '100%' }}>
        <div className="pt-6" style={{ width: '100%', height: '100%' }}>
          <ResponsiveContainer width="100%" height="100%">
            <LineChart
              margin={{
                top: 5, right: 30, left: 20, bottom: 5,
              }}
              data={yearlyInvestmentData}
            >
              <CartesianGrid strokeDasharray="3 3" />
              <XAxis 
                dataKey="year" 
                axisLine={false} 
                tickFormatter={(value) => `${value}`} 
                tickMargin={10} // Add margin between axis line and tick text
              />
              <YAxis 
                axisLine={false} 
                tickFormatter={(value) => `$${value.toLocaleString()}`} 
                tickMargin={10} // Add margin between axis line and tick text
              />
              <Tooltip content={<CustomTooltip />} />
              <Line type="linear" dataKey="totalInvestment" stroke="#ff7300" strokeWidth={3} />
            </LineChart>
          </ResponsiveContainer>
        </div>
      </CardContent>
    </Card>
  );
}