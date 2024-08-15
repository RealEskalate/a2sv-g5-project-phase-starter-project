"use client"

import { Pie, PieChart, Tooltip, Cell, Sector } from "recharts"

// Dummy data for the pie chart
const dummyChartData = [
  { browser: "Others", Expenses: 400, fill: "#FF6384" },
  { browser: "Transfer", Expenses: 300, fill: "#36A2EB" },
  { browser: "Shopping", Expenses: 200, fill: "#FFCE56" },
  { browser: "Services", Expenses: 100, fill: "#4BC0C0" },
]

// Define a custom label component
const CustomLabel = (props: any) => {
  const { cx, cy, midAngle, outerRadius, value, name } = props;
  const RADIAN = Math.PI / 180;
  const radius =  (outerRadius) / 2; // Position inside the sector

  // Calculate the position of the label
  const x = cx + radius * Math.cos(-RADIAN * midAngle);
  const y = cy + radius * Math.sin(-RADIAN * midAngle);

  return (
    <text
      x={x}
      y={y}
      fill="#fff"
      textAnchor="middle"

      className="text-xs font-semibold"
    >
      {name}
    </text>
  );
};

// Define a custom shape for the hover effect
const renderActiveShape = (props: any) => {
  const { cx, cy, midAngle, outerRadius, startAngle, endAngle, fill } = props;
  const RADIAN = Math.PI / 180;
  const sin = Math.sin(-RADIAN * midAngle);
  const cos = Math.cos(-RADIAN * midAngle);

  return (
    <g>
      <Sector
        cx={cx}
        cy={cy}
        startAngle={startAngle}
        endAngle={endAngle}
        outerRadius={outerRadius}
        fill={fill}
      />
      <Sector
        cx={cx}
        cy={cy}
        startAngle={startAngle}
        endAngle={endAngle}
        outerRadius={outerRadius}
        fill="black"
        stroke="none"
      />
    </g>
  );
};

// PieChart component
export function PieChartComponent() {
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
        >
        </Pie>
      </PieChart>
    </div>
  )
}
