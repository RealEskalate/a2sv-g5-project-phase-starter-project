"use client";

import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Pie, PieChart, Tooltip, Cell, Sector } from "recharts";

// Define state for data and errors
const PieChartComponent: React.FC = () => {
  const [data, setData] = useState<any[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [activeIndex, setActiveIndex] = useState(0);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('https://bank-dashboard-6acc.onrender.com/transactions/expenses', {
          headers: {
            Authorization: `Bearer your_bearer_token_here`, // Replace with your token
          },
        });
        setData(response.data.data); // Adjust based on API response
      } catch (error) {
        console.error("Failed to fetch data:", error);
        setError("Failed to fetch data. Please check the console for more details.");
      }
    };

    fetchData();
  }, []);

  // Define dummy data to fallback on if needed
  const dummyChartData = data.length === 0 ? [
    { browser: "Others", Expenses: 350, fill: "#FF6384" },
    { browser: "Transfer", Expenses: 300, fill: "#36A2EB" },
    { browser: "Shopping", Expenses: 200, fill: "#FFCE56" },
    { browser: "Services", Expenses: 150, fill: "#4BC0C0" },
  ] : data.map((item: any) => ({
    browser: item.type, // or other relevant property
    Expenses: item.amount,
    fill: "#8884d8", // Customize as needed
  }));

  const totalExpenses = dummyChartData.reduce((sum, entry) => sum + entry.Expenses, 0);

  const CustomLabel = (props: any) => {
    const { cx, cy, midAngle, outerRadius, value, name } = props;
    const RADIAN = Math.PI / 180;
    const radius = (outerRadius) / 2;

    const x = cx + radius * Math.cos(-RADIAN * midAngle);
    const y = cy + radius * Math.sin(-RADIAN * midAngle);

    const percentage = ((value / totalExpenses) * 100);

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
        <tspan x={x} dy="1.2em">{percentage.toFixed(1)}%</tspan>
      </text>
    );
  };

  const renderActiveShape = (props: any) => {
    const { cx, cy, midAngle, outerRadius, startAngle, endAngle, fill } = props;
    const RADIAN = Math.PI / 180;
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
          data={dummyChartData}
          dataKey="Expenses"
          nameKey="browser"
          cx="50%"
          cy="50%"
          outerRadius={100}
          fill="#8884d8"
          activeShape={renderActiveShape}
          label={CustomLabel}
          stroke="#fff"
          strokeWidth={3}
          onMouseEnter={(_, index) => setActiveIndex(index)}
          onMouseLeave={() => setActiveIndex(0)}
          activeIndex={activeIndex}
        >
          {dummyChartData.map((entry, index) => (
            <Cell key={`cell-${index}`} fill={entry.fill} />
          ))}
        </Pie>
      </PieChart>
    </div>
  );
};

export default PieChartComponent;
