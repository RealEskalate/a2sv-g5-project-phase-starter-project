"use client";

import { TrendingUp } from "lucide-react";
import { Bar, BarChart, CartesianGrid, XAxis, YAxis } from "recharts";

import {
  Card,
  CardContent,
  CardDescription,
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
  { day: "Fri", desktop: 186, mobile: 80 },
  { day: "Sat", desktop: 305, mobile: 200 },
  { day: "Sun", desktop: 237, mobile: 120 },
  { day: "Mon", desktop: 73, mobile: 190 },
  { day: "Tue", desktop: 209, mobile: 130 },
  { day: "Wed", desktop: 214, mobile: 140 },
];

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "#1814F3",
  },
  mobile: {
    label: "Mobile",
    color: "#16DBCC",
  },
} satisfies ChartConfig;

export function BarChartComponent() {
  return (
    <Card className="w-full lg:w-[800px] mb-8 lg:mt-9 dark:bg-darkComponent">
      <CardHeader>
        <CardTitle>Weekly Activity</CardTitle>
        {/* <CardDescription>January - June 2024</CardDescription> */}
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig}>
          <BarChart accessibilityLayer data={chartData} barSize={20}>
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="day"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={(value) => value}
            />
            <YAxis
              tickCount={6} 
              domain={[0, 500]} 
              ticks={[0, 100, 200, 300, 400, 500]} 
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Bar dataKey="desktop" fill="var(--color-desktop)" radius={15} />
            <Bar dataKey="mobile" fill="var(--color-mobile)" radius={15} />
          </BarChart>
        </ChartContainer>
      </CardContent>
      <CardFooter className="flex-col items-start gap-2 text-sm">
        {/* <div className="flex gap-2 font-medium leading-none">
          Trending up by 5.2% this day <TrendingUp className="h-4 w-4" />
        </div>
        <div className="leading-none text-muted-foreground">
          Showing total visitors for the last 6 days
        </div> */}
      </CardFooter>
    </Card>
  );
}
