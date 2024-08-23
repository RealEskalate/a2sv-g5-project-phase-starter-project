"use client"

import React, { useState } from 'react'; // Import React and useState
import { Pie, PieChart, Tooltip, Cell, Sector } from "recharts"

// Dummy data for the pie chart
const dummyChartData = [
  { browser: "Others", Expenses: 350, fill: "#FF6384" },
  { browser: "Transfer", Expenses: 300, fill: "#36A2EB" },
  { browser: "Shopping", Expenses: 200, fill: "#FFCE56" },
  { browser: "Services", Expenses: 150, fill: "#4BC0C0" },
]

<<<<<<< HEAD
  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get('https://bank-dashboard-o9tl.onrender.com/transactions/expenses?page=0&size=10', {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        });

        // Extract the content array from the response
        const fetchedData = response.data.data.content;

        // Set the content array as the data
        setData(fetchedData);
      } catch (error) {
        console.error("Failed to fetch data:", error);
        setError("Failed to fetch data. Please check the console for more details.");
      }
    };
=======
// Calculate the total sum of expenses
const totalExpenses = dummyChartData.reduce((sum, entry) => sum + entry.Expenses, 0);
>>>>>>> aastu.web.g5.yetnayet.transactions

// Define a custom label component
const CustomLabel = (props: any) => {
  const { cx, cy, midAngle, outerRadius, value, name } = props;
  const RADIAN = Math.PI / 180;
  const radius = (outerRadius) / 2; // Position inside the sector

  // Calculate the position of the label
  const x = cx + radius * Math.cos(-RADIAN * midAngle);
  const y = cy + radius * Math.sin(-RADIAN * midAngle);

  // Calculate the percentage
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

// Define a custom shape for the hover effect
const renderActiveShape = (props: any) => {
  const { cx, cy, midAngle, outerRadius, startAngle, endAngle, fill } = props;
  const RADIAN = Math.PI / 180;

  // Increase outer radius for the hover effect
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

// PieChart component
export function PieChartComponent() {
  const [activeIndex, setActiveIndex] = useState<number | 0>(0);

  const onPieEnter = (data: any, index: number) => {
    setActiveIndex(index);
  };

<<<<<<< HEAD
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
  }, []);

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
=======
  const onPieLeave = () => {
    setActiveIndex(0);
>>>>>>> aastu.web.g5.yetnayet.transactions
  };

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
          activeShape={renderActiveShape} // Apply the hover effect
          label={CustomLabel} // Use the custom label component
          stroke="#fff" // Set the stroke color to white
          strokeWidth={3} // Set the stroke width for the border
          onMouseEnter={onPieEnter} // Handle mouse enter
          onMouseLeave={onPieLeave} // Handle mouse leave
          activeIndex={activeIndex} // Apply hover effect based on activeIndex
        >
          {dummyChartData.map((entry, index) => (
            <Cell key={`cell-${index}`} fill={entry.fill} />
          ))}
        </Pie>
      </PieChart>
    </div>
  )
}
