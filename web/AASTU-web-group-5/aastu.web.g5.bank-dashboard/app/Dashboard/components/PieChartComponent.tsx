'use client';

import React, { useState, useEffect } from 'react';
import { Pie, PieChart, Tooltip, Cell, Sector } from 'recharts';
import axios from 'axios';
import { useSession } from 'next-auth/react';

interface Transaction {
  type: string;
  amount: number;
}

const typeColors: { [key: string]: string } = {
  transfer: "#36A2EB",
  shopping: "#FFCE56",
  deposit: "#4BC0C0",
  service: "#FF6384",
  other: "#C0C0C0"
};

export function PieChartComponent() {
  const { data: session } = useSession();
  const user = session?.user as { accessToken?: string; refreshToken?: string };
  
  const [data, setData] = useState<Transaction[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [activeIndex, setActiveIndex] = useState<number | 0>(0);
  
  useEffect(() => {
    const fetchData = async () => {
      if (user?.accessToken) {
        try {
          const response = await axios.get('https://bank-dashboard-rsf1.onrender.com/transactions/expenses?page=0&size=7', {
            headers: {
              Authorization: `Bearer ${user.accessToken}`,
            },
          });

          const fetchedData = response.data.data.content;
          setData(fetchedData);
          setError(null); // Clear any previous errors
        } catch (error) {
          if (error.response?.status === 401) {
            // Token might be expired, try to refresh it
            try {
              const refreshResponse = await axios.post('https://bank-dashboard-rsf1.onrender.com/auth/refresh_token', {}, {
                headers: {
                  "Content-Type": "application/json",
                  Authorization: `Bearer ${user.refreshToken}`,
                },
              });

              const refreshedTokens = refreshResponse.data.data;
              const newAccessToken = refreshedTokens.access_token;

              // Retry the original request with the new access token
              const retryResponse = await axios.get('https://bank-dashboard-rsf1.onrender.com/transactions/expenses?page=0&size=7', {
                headers: {
                  Authorization: `Bearer ${newAccessToken}`,
                },
              });

              const retriedData = retryResponse.data.data.content;
              setData(retriedData);
            } catch (refreshError) {
              console.error("Failed to refresh access token:", refreshError);
              setError("Failed to refresh access token. Please log in again.");
            }
          } else {
            console.error("Failed to fetch data:", error);
            setError("Failed to fetch data. Please check the console for more details.");
          }
        }
      }
    };

    fetchData();
  }, [user]);

  // Process data
  const processedData = data.reduce((acc: any[], item: Transaction) => {
    const type = ['service', 'transfer', 'shopping'].includes(item.type?.toLowerCase())
      ? item.type.toLowerCase()
      : 'other';

    const existing = acc.find((entry) => entry.name === type);

    if (existing) {
      existing.value += item.amount;
    } else {
      acc.push({
        name: type,
        value: item.amount,
        fill: typeColors[type],
      });
    }

    return acc;
  }, []);

  const totalValue = processedData.reduce((sum, entry) => sum + entry.value, 0);
  const threshold = 100; // Example threshold for "Other"
  const filteredData = processedData.filter(entry => entry.value >= threshold);

  const othersValue = totalValue - filteredData.reduce((sum, entry) => sum + entry.value, 0);
  
  if (othersValue > 0) {
    filteredData.push({
      name: 'other',
      value: othersValue,
      fill: typeColors['other']
    });
  }

  const CustomLabel = (props: any) => {
    const { cx, cy, midAngle, outerRadius, value, name } = props;
    const RADIAN = Math.PI / 180;
    const radius = outerRadius / 2;

    const x = cx + radius * Math.cos(-RADIAN * midAngle);
    const y = cy + radius * Math.sin(-RADIAN * midAngle);

    const percentage = ((value / totalValue) * 100).toFixed(1);

    return (
      <text
        x={x}
        y={y}
        className="text-xs font-semibold text-gray-800 dark:text-gray-200"
        textAnchor="middle"
        dominantBaseline="central"
      >
        <tspan x={x} dy="-1.2em">{name}</tspan>
        <tspan x={x} dy="1.2em">{percentage}%</tspan>
      </text>
    );
  };

  const renderActiveShape = (props: any) => {
    const { cx, cy, midAngle, outerRadius, startAngle, endAngle, fill } = props;
    const RADIAN = Math.PI / 180;
    const hoverRadius = outerRadius + 15; // Adjust the value as needed

    return (
      <g>
        <Sector
          cx={cx}
          cy={cy}
          startAngle={startAngle}
          endAngle={endAngle}
          outerRadius={hoverRadius}
          fill={fill}
          className="transition-all duration-200 dark:fill-opacity-80"
        />
        <Sector
          cx={cx}
          cy={cy}
          startAngle={startAngle}
          endAngle={endAngle}
          outerRadius={outerRadius}
          fill={fill}
          stroke="none"
        />
      </g>
    );
  };

  const onPieEnter = (data: any, index: number) => {
    setActiveIndex(index);
  };

  const onPieLeave = () => {
    setActiveIndex(0);
  };

  if (error) return <p className="text-red-500 dark:text-red-300">Error: {error}</p>;

  return (
    <div className="flex justify-center items-center h-screen bg-gray-50 dark:bg-gray-900 transition-colors duration-300">
      <PieChart width={300} height={300}>
        <Pie
          data={filteredData}
          dataKey="value"
          nameKey="name"
          cx="50%"
          cy="50%"
          outerRadius={100}
          fill="#8884d8"
          activeShape={renderActiveShape}
          label={CustomLabel}
          stroke="#fff"
          strokeWidth={3}
          onMouseEnter={onPieEnter}
          onMouseLeave={onPieLeave}
          activeIndex={activeIndex}
        >
          {filteredData.map((entry, index) => (
            <Cell key={`cell-${index}`} fill={entry.fill} />
          ))}
        </Pie>
        <Tooltip />
      </PieChart>
    </div>
  );
}

export default PieChartComponent;
