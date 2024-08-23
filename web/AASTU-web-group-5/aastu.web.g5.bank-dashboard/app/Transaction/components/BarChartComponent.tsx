import React, { useState, useEffect } from "react";
import { BarChart, Bar, XAxis, CartesianGrid, Tooltip as RechartsTooltip } from "recharts";
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
  const [aggregatedData, setAggregatedData] = useState<{ month: string; amount: number }[]>([]);
  const [maxAmount, setMaxAmount] = useState<number>(0);
  const [highlightedIndex, setHighlightedIndex] = useState<number | null>(null);

  // Scaling factor to reduce the height of the tallest bar
  const scalingFactor = 0.9; // 90% of the chart height

  useEffect(() => {
    const aggregated = monthsOfYear.map((month, index) => {
      const monthData = data.filter(item => {
        const itemMonth = new Date(item.date).getMonth(); // Extract month from date
        return itemMonth === index;
      });
      const totalAmount = monthData.reduce((sum, item) => sum + item.amount, 0);
      return { month, amount: totalAmount + 100 }; // Added 100 to totalAmount as per your example
    });

    // Find the maximum amount
    const max = Math.max(...aggregated.map(d => d.amount));
    setMaxAmount(max);

    setAggregatedData(aggregated);
  }, [data]);

  return (
    <Card>
      <CardContent>
        <ChartContainer config={chartConfig}>
          <BarChart data={aggregatedData} width={500} height={300}>
            <CartesianGrid vertical={false} horizontal={false} />
            <XAxis dataKey="month" tickLine={false} tickMargin={10} axisLine={false} />
            <Bar
              dataKey="amount"
              radius={10}
              onMouseEnter={(data, index) => setHighlightedIndex(index)}
              onMouseLeave={() => setHighlightedIndex(null)}
            >
              {aggregatedData.map((entry, index) => {
                const isHighlighted = highlightedIndex === index;
                const barFill = isHighlighted ? "#green" : "#blue"; // Default gray for non-highlighted bars
                return (
                  <rect
                    key={`rect-${index}`}
                    x={index * (500 / aggregatedData.length)}
                    y={300 - ((entry.amount / maxAmount) * 300 * scalingFactor)}
                    width={(500 / aggregatedData.length) - 10}
                    height={((entry.amount / maxAmount) * 300 * scalingFactor)}
                    fill={barFill}
                    onMouseEnter={() => setHighlightedIndex(index)}
                    onMouseLeave={() => setHighlightedIndex(null)}
                  />
                );
              })}
            </Bar>
            <RechartsTooltip
              content={({ payload }) => {
                if (payload && payload.length) {
                  return (
                    <div className="custom-tooltip">
                      <p>{`Month: ${payload[0].payload.month}`}</p>
                      <p>{`Amount: $${payload[0].value}`}</p>
                    </div>
                  );
                }
                return null;
              }}
            />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}

export default BarChartComponent;
