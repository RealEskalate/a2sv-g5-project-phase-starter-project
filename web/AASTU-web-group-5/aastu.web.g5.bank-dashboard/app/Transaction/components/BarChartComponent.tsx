import React, { useState } from "react";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  LabelList,
  CartesianGrid,
  ResponsiveContainer,
  Cell,
} from "recharts";

interface ExpenseData {
  amount: number;
  month: string;
}

interface BarChartComponentProps {
  data: ExpenseData[];
}

const BarChartComponent: React.FC<BarChartComponentProps> = ({ data }) => {
  const [activeIndex, setActiveIndex] = useState<number | null>(null);

  // Default data structure for months with different default values
  const defaultData = [
    { month: "Jan", amount: 733 },
    { month: "Feb", amount: 1762 },
    { month: "Mar", amount: 500 },
    { month: "Apr", amount: 100 },
    { month: "May", amount: 1230 },
    { month: "Jun", amount: 1110 },
    { month: "Jul", amount: 12340 },
    { month: "Aug", amount: 2320 },
    { month: "Sep", amount: 1100 },
    { month: "Oct", amount: 1100 },
    { month: "Nov", amount: 1109 },
    { month: "Dec", amount: 1230 },
  ];

  // Merge default data with actual data
  const mergedData = defaultData.map((item) => {
    const found = data.find((d) => d.month === item.month);
    return { ...item, amount: found ? found.amount : item.amount }; // Use default value if not found
  });

  // Handle mouse enter and leave events
  const handleMouseEnter = (index: number) => {
    setActiveIndex(index);
  };

  const handleMouseLeave = () => {
    setActiveIndex(null);
  };

  return (
    <ResponsiveContainer width="100%" height={300}>
      <BarChart
        data={mergedData}
        onMouseLeave={handleMouseLeave}
        margin={{ top: 20, right: 30, left: 20, bottom: 5 }}
        
      >
 <CartesianGrid vertical={false} horizontal={false} />
         <XAxis dataKey="month" 
         tickLine={false}
         tickMargin={10}
         axisLine={false}
         tickFormatter={(value) => value.slice(0, 3)}
    />
        <Bar dataKey="amount" radius={10}>
              {mergedData.map((entry, index) => (
                <Cell
                  key={`cell-${index}`}
                  fill={activeIndex === index ? "#12887E33" : "#EDF0F7"}
                  onMouseEnter={() => handleMouseEnter(index)}
                />
              ))}
               <LabelList
                dataKey="amount"
                position="top"
                content={({ x, y, value, index }) =>
                  activeIndex === index ? (
                    <text
                      x={x}
                      y={y}
                      dy={-10}
                      fill="black"
                      fontSize={12}
                      textAnchor="top"
                    >
                      {value}
                    </text>
                  ) : null
                }
              />
            </Bar>
      </BarChart>
    </ResponsiveContainer>
  );
};

export default BarChartComponent;
