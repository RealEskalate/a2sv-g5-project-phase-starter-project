import React, { useState, useEffect } from "react";
import { Bar, BarChart, CartesianGrid, XAxis, LabelList } from "recharts";
import { Card, CardContent } from "@/components/ui/card";
import { ChartConfig, ChartContainer } from "@/components/ui/chart";

// Define chart configuration for the BarChart
const chartConfig = {
  desktop: {
    label: "Expenses",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;

interface TransactionData {
  transactionId: string;
  type: string;
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string | null;
}

interface BarChartComponentProps {
  data: TransactionData[];
}

// Define all months of the year
const monthsOfYear = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];

export function BarChartComponent({ data }: BarChartComponentProps) {
  const [hovering, setHovering] = useState<number | null>(null);
  const [aggregatedData, setAggregatedData] = useState<{ month: string; amount: number }[]>([]);

  useEffect(() => {
    const aggregated = monthsOfYear.map((month, index) => {
      const monthData = data.filter(item => {
        const itemMonth = new Date(item.date).getMonth(); // Extract month from date
        return itemMonth === index;
      });
      const totalAmount = monthData.reduce((sum, item) => sum + item.amount, 0);
      return { month, amount: totalAmount+11111111 };
    });
    setAggregatedData(aggregated);
  }, [data]);

  return (
    <Card>
      <CardContent>
        <ChartContainer config={chartConfig}>
          <BarChart
            data={aggregatedData}
            width={500}
            height={300}
            onMouseLeave={() => setHovering(null)}
          >
            <CartesianGrid vertical={false} horizontal={false} />
            <XAxis
              dataKey="month"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
            />
            <Bar
  dataKey="amount"
  radius={10}
  fill="green" // initial fill color
  onMouseEnter={(data, index) => setHovering(index)}
  onMouseLeave={() => setHovering(null)}
>
  {aggregatedData.map((entry, index) => (
    <rect
      key={`rect-${index}`}
      x={index * 30}
      y={200 - entry.amount}
      width={20}
      height={entry.amount}
      fill={hovering === index ? "#d0e9f4" : "green"} // change fill color on hover
    />
  ))}
  <LabelList
    dataKey="amount"
    position="top"
    content={({ x, y, value, index }) =>
      hovering === index ? (
        <text
          x={x}
          y={y}
          dy={-10}
          fill="black"
          fontSize={12}
          textAnchor="middle"
        >
          {value}
        </text>
      ) : null
    }
  />
              <LabelList
                dataKey="amount"
                position="top"
                content={({ x, y, value, index }) =>
                  hovering === index ? (
                    <text
                      x={x}
                      y={y}
                      dy={-10}
                      fill="black"
                      fontSize={12}
                      textAnchor="middle"
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
