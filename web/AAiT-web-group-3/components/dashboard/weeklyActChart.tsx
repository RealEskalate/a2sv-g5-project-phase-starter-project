"use client";

import { Bar, BarChart, CartesianGrid, XAxis, YAxis } from "recharts";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

const chartData = [
  { month: "January", desktop: 186, mobile: 80 },
  { month: "February", desktop: 305, mobile: 200 },
  { month: "March", desktop: 237, mobile: 120 },
  { month: "April", desktop: 73, mobile: 190 },
  { month: "May", desktop: 209, mobile: 130 },
  { month: "June", desktop: 214, mobile: 140 },
  { month: "June", desktop: 214, mobile: 140 },
];

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "hsl(var(--chart-1))",
  },
  mobile: {
    label: "Mobile",
    color: "hsl(var(--chart-2))",
  },
} satisfies ChartConfig;

const WeeklyActivityChart = () => {
  return (
    <ChartContainer
      config={chartConfig}
      className="bg-white rounded-2xl h-[260px]"
    >
      <BarChart accessibilityLayer data={chartData}>
        <CartesianGrid vertical={false} horizontal={true} stroke="#F3F3F5" />
        <XAxis
          dataKey="month"
          tickLine={false}
          tickMargin={10}
          width={3}
          axisLine={false}
          tickFormatter={(value) => value.slice(0, 3)}
        />
        <YAxis
          tickLine={false}
          axisLine={false}
          tickMargin={10}
          tickCount={7}
          tickFormatter={(value) => `${value}`}
        />
        <ChartTooltip
          cursor={false}
          content={<ChartTooltipContent indicator="dashed" />}
        />
        <Bar dataKey="desktop" fill="#1814F3" radius={4} strokeWidth={2} />
        <Bar dataKey="mobile" fill="#16DBCC" radius={4} strokeWidth={2} />
      </BarChart>
    </ChartContainer>
  );
};

export default WeeklyActivityChart;
