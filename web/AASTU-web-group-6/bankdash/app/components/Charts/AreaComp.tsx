"use client";

import { Area, AreaChart, CartesianGrid, XAxis } from "recharts";
import { Card, CardContent } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

const chartData = [
  { month: "Jul", desktop: 20 },
  { month: "Aug", desktop: 80 },
  { month: "Sep", desktop: 390 },
  { month: "Oct", desktop: 500 },
  { month: "Nov", desktop: 150 },
  { month: "Dev", desktop: 400 },
  { month: "January", desktop: 86 },
];

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;

export function AreaComp() {
  return (
    <Card className="w-full border-0 shadow-none bg-transparent ">
      <CardContent>
        <ChartContainer
          config={chartConfig}
          className="barHeight w-full min-h-[100px] h-56"
        >
          <AreaChart
            accessibilityLayer
            data={chartData}
            margin={{
              left: 12,
              right: 12,
            }}
          >
            <defs>
              <linearGradient id="gradient" x1="0%" y1="100%" x2="0%" y2="0%">
                <stop offset="0%" stopColor="rgba(45, 96, 255, 0)" />
                <stop offset="90%" stopColor="rgba(45, 96, 255, 0.5)" />
              </linearGradient>
            </defs>

            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="month"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              tickFormatter={(value) => value.slice(0, 4)}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dot" />}
            />

            <Area
              dataKey="desktop"
              type="natural"
              fill="url(#gradient)" // Use the gradient fill
              stroke="#1814F3" // Stroke color
              strokeWidth={3} // Stroke width
              stackId="a"
            />
          </AreaChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
