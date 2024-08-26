"use client";

import { useState, useEffect } from "react";
import { Bar, BarChart, CartesianGrid, XAxis, Cell, LabelList } from "recharts";
import { Card, CardContent } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

interface TransactionData {
  transactionId: string;
  type: string;
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string | null;
}

// Define all months of the year
const monthsOfYear = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;

export function BarChartComponent() {
  const [chartData, setChartData] = useState<{ month: string; desktop: number }[]>([]);
  const [activeIndex, setActiveIndex] = useState<number | null>(null);

  useEffect(() => {
    // Fetch data from the API
    const fetchData = async () => {
      try {
        const response = await fetch("https://bank-dashboard-rsf1.onrender.com/transactions?page=0");
        const data: TransactionData[] = await response.json();

        // Aggregate data by month
        const aggregatedData = monthsOfYear.map((month, index) => {
          const monthData = data.filter(item => {
            const itemMonth = new Date(item.date).getMonth(); // Extract month from date
            return itemMonth === index;
          });
          const totalAmount = monthData.reduce((sum, item) => sum + item.amount, 0);
          return { month, desktop: totalAmount + 100 }; // Example addition
        });

        setChartData(aggregatedData);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
  }, []);

  const handleMouseEnter = (index: number) => {
    setActiveIndex(index);
  };

  const handleMouseLeave = () => {
    setActiveIndex(null);
  };

  return (
    <Card>
      <CardContent>
        <ChartContainer config={chartConfig}>
          <BarChart data={chartData} onMouseLeave={handleMouseLeave}>
            <CartesianGrid vertical={false} horizontal={false} />
            <XAxis
              dataKey="month"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            
            <Bar dataKey="desktop" radius={10}>
              {chartData.map((entry, index) => (
                <Cell
                  key={`cell-${index}`}
                  fill={activeIndex === index ? "#12887E33" : "#EDF0F7"}
                  onMouseEnter={() => handleMouseEnter(index)}
                />
              ))}
              <LabelList
                dataKey="desktop"
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
        </ChartContainer>
      </CardContent>
    </Card>
  );
}

export default BarChartComponent;
