import React from "react";
import { Area, AreaChart, CartesianGrid, XAxis, YAxis } from "recharts";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

import { useGetRandomBalanceHistoryQuery } from "@/lib/redux/api/transactionsApi";
import { Slash } from "lucide-react";

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;

export function AreaChartComponent() {
  // Fetching data from the API
  const { data, isLoading, error } = useGetRandomBalanceHistoryQuery({
    monthsBeforeFirstTransaction: 7,
  });

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading data</div>;

  // Extracting and transforming the data for the chart
  const chartData = data?.data.map((item: { time: string; value: number }) => {
    const month = new Date(item.time).toLocaleString("default", {
      month: "short",
    });
    return {
      month, // Updated to show the abbreviated month
      desktop: item.value,
    };
  });

  return (
    <div className="Balance-History dark:bg-darkComponent">
      <ChartContainer config={chartConfig}>
        <AreaChart
          data={chartData}
          margin={{
            top: 10,
            left: 0,
            right: 30,
            bottom: 10,
          }}
        >
          <defs>
            <linearGradient id="gradient" x1="0" y1="0" x2="0" y2="1">
              <stop offset="0%" stopColor="#DFE7FF" stopOpacity={1} />
              <stop offset="100%" stopColor="#DFE7FF" stopOpacity={0.4} />
            </linearGradient>
          </defs>
          <CartesianGrid />
          <XAxis
            dataKey="month" // Matches the transformed 'month' key
            tickLine={false}
            axisLine={false}
            tickMargin={8}
            tickFormatter={(value) => value.slice(0, 3)} // If you want the month abbreviation
          />
          <YAxis
            domain={[0, 8000]} // Adjust the domain based on your data range
            tickCount={5}
            tickLine={false}
            axisLine={false}
            tickMargin={8}
            interval={0}
          />
          <ChartTooltip
            cursor={true}
            content={<ChartTooltipContent indicator="dot" />}
          />
          <Area
            dataKey="desktop" // Matches the transformed 'desktop' key
            type="natural"
            fill="url(#gradient)"
            fillOpacity={1}
            stroke="#1814F3"
            strokeWidth={5}
            stackId="a"
          />
        </AreaChart>
      </ChartContainer>
    </div>
  );
}
