"use client";

import { Bar, BarChart, CartesianGrid, XAxis } from "recharts";
import { Card, CardContent } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

const chartData = [
  { month: "Sat", desktop: 214, mobile: 140 },
  { month: "Sun", desktop: 214, mobile: 140 },
  { month: "Mon", desktop: 186, mobile: 80 },
  { month: "Tue", desktop: 305, mobile: 200 },
  { month: "Wed", desktop: 237, mobile: 120 },
  { month: "Thu", desktop: 73, mobile: 190 },
  { month: "Fri", desktop: 209, mobile: 130 },
];

const chartConfig = {
  desktop: {
    label: "Deposit",
    color: "hsl(var(--chart-1))",
  },
  mobile: {
    label: "Withdraw",
    color: "hsl(var(--chart-2))",
  },
} satisfies ChartConfig;

export default function BarComp() {
  return (
    <Card className="w-full border-0 shadow-none bg-transparent">
      <CardContent>
        <ChartContainer config={chartConfig} className="barHeight w-full h-52">
          <BarChart data={chartData}>
            <CartesianGrid vertical={false} />
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
            <Bar
              dataKey="desktop"
              fill="#1814F3" // Change to your desired color
              radius={4}
              barSize={12}
              // opacity={0.4}
            />
            <Bar
              dataKey="mobile"
              fill="#16DBCC" // Change to your desired color
              radius={4}
              barSize={10}
            />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
