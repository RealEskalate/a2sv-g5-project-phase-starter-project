"use client";

import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Pie, PieChart, Tooltip, Cell, Sector } from "recharts";

const PieChartComponent: React.FC = () => {
  const [data, setData] = useState<any[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [activeIndex, setActiveIndex] = useState<number | undefined>(undefined);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('https://bank-dashboard-6acc.onrender.com/transactions/expenses?page=0&size=10', {
          headers: {
            Authorization: `Bearer ${process.env.NEXT_PUBLIC_ACCESS_TOKEN}`,
          },
        });
        setData(response.data.data);
      } catch (error) {
        console.error("Failed to fetch data:", error);
        setError("Failed to fetch data. Please check the console for more details.");
      }
    };

    fetchData();
  }, []);

  // Map transaction types to colors
  const typeColors: Record<string, string> = {
    service: "#FC7900",
    transfer: "#343C6A",
    shopping: "#FA00FF",
    other: "#1814F3", // Default color for "other"
  };

  // Process data
  const processedData = data.reduce((acc: any[], item: any) => {
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
  }, [] as any[]);

  // Handle "Other" category directly within the processed data
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
        fill="#fff"
        textAnchor="middle"
        className="text-xs font-semibold"
        dominantBaseline="central"
      >
        <tspan x={x} dy="-1.2em">{name}</tspan>
        <tspan x={x} dy="1.2em">{percentage}%</tspan>
      </text>
    );
  };

  const renderActiveShape = (props: any) => {
    const { cx, cy, midAngle, outerRadius, startAngle, endAngle, fill } = props;
    const hoverRadius = outerRadius + 15;

    return (
      <g>
        <Sector
          cx={cx}
          cy={cy}
          startAngle={startAngle}
          endAngle={endAngle}
          outerRadius={hoverRadius}
          fill={fill}
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

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div className="flex justify-center items-center h-screen">
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
          onMouseEnter={(_, index) => setActiveIndex(index)}
          onMouseLeave={() => setActiveIndex(undefined)}
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
};

export default PieChartComponent;
