"use client";

import React, { useState, useEffect } from 'react';
import { useSelector } from 'react-redux';
import { CartesianGrid, Line, LineChart, Tooltip, XAxis, YAxis, ResponsiveContainer } from "recharts";
import { Card, CardContent } from "../../@/components/ui/card";
import { RootState } from '@/app/redux/store'; // Adjust the import path as necessary
import Shimmer1 from "../Accounts/shimmer";
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
        <div className="custom-tooltip bg-white dark:bg-gray-800 p-2 border border-gray-300 dark:border-gray-600 rounded">
          <p className="label text-black dark:text-white">{`${label} : $${value.toLocaleString()}`}</p>
        </div>
      );
    }
  }
  return null;
};

const Shimmer = () => (
  <div className="shimmer-wrapper">
    <div className="shimmer"></div>
  </div>
);

export default function MonthlyRevenueChart({ data }: MonthlyRevenueChartProps) {
  const darkMode = useSelector((state: RootState) => state.theme.darkMode);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    // Simulate loading delay
    const timer = setTimeout(() => setLoading(false), 2000);
    return () => clearTimeout(timer);
  }, []);

  if (loading) {
    return (
      <Card style={{ height: '100%' }} className={darkMode ? 'bg-gray-800 text-white' : 'bg-white text-black'}>
        <CardContent style={{ height: '100%' }}>
          <Shimmer1 />
        </CardContent>
      </Card>
    );
  }

  return (
    <Card style={{ height: '100%' }} className={darkMode ? 'bg-gray-800 text-white' : 'bg-white text-black'}>
      <CardContent style={{ height: '100%' }}>
        <div className='pt-6' style={{ width: '100%', height: '100%' }}>
          <ResponsiveContainer width="100%" height="100%">
            <LineChart
              data={data}
              margin={{
                top: 5, right: 20, left: 5, bottom: 5,
              }}
            >
              <CartesianGrid strokeDasharray="3 3" stroke={darkMode ? '#444' : '#ccc'} />
              <XAxis dataKey="time" stroke={darkMode ? '#ccc' : '#000'} />
              <YAxis tickFormatter={(value) => `$${value}`} stroke={darkMode ? '#ccc' : '#000'} />
              <Tooltip content={<CustomTooltip />} />
              <Line type="monotone" dataKey="value" stroke={darkMode ? '#82ca9d' : '#82ca9d'} strokeWidth={3} dot={false} />
            </LineChart>
          </ResponsiveContainer>
        </div>
      </CardContent>
    </Card>
  );
}