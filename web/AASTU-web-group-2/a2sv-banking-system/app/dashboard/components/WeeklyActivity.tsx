"use client";

import { TrendingUp } from "lucide-react";
import { Bar, BarChart, CartesianGrid, XAxis, YAxis } from "recharts";
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

const chartData = [
  { month: "Saturday", desktop: 186, mobile: 80 },
  { month: "Sunday", desktop: 305, mobile: 200 },
  { month: "Monday", desktop: 237, mobile: 120 },
  { month: "Tuesday", desktop: 73, mobile: 190 },
  { month: "Wednesday", desktop: 209, mobile: 130 },
  { month: "Thursday", desktop: 214, mobile: 140 },
  { month: "Friday", desktop: 214, mobile: 140 },
];

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "#1814f3", // Color for desktop bars
  },
  mobile: {
    label: "Mobile",
    color: "#16dbcc", // Color for mobile bars
  },
} satisfies ChartConfig;

export function WeeklyActivity() {
  return (
    <Card className="my-4 mx-4 rounded-3xl flex-grow">
      <CardHeader>
        <CardTitle className="text-[#343C6A] font-bold text-xl md:hidden">
          Weekly Activity
        </CardTitle>
      </CardHeader>
      <CardContent className="flex justify-center items-center">
        <ChartContainer
          config={chartConfig}
          className="w-full h-60 max-w-[800-px]" // Adjust the width to ensure proper fitting
        >
          <BarChart
            data={chartData}
            width={800}
            height={400}
            margin={{
              top: 10, right: 30, left: 0, bottom: 0, // Adjust margins if needed
            }}
          >
            <CartesianGrid
              vertical={false}
              strokeDasharray="none" // Remove dashed lines for solid lines
              stroke="#E0E0E0" // Lighter grey color for the grid lines
              strokeWidth={0.5} // Thinner lines for a lighter appearance
            />
            <YAxis
              tickCount={6}
              tickFormatter={(value) => value}
              domain={[0, 500]}
              interval={0}
              tickLine={false}
              axisLine={false}
              tickMargin={10}
            />
            <XAxis
              dataKey="month"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Bar dataKey="desktop" fill={chartConfig.desktop.color} radius={4} />
            <Bar dataKey="mobile" fill={chartConfig.mobile.color} radius={4} />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
