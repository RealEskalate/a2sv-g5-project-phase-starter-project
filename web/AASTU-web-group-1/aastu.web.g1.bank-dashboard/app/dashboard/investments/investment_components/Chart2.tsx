"use client";
import { TrendingUp } from "lucide-react";
import { CartesianGrid, Line, LineChart, XAxis, YAxis, Area, AreaChart } from "recharts";
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

import { useUser } from "@/contexts/UserContext";

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;

export default function Chart2(props:any) {
  const { isDarkMode } = useUser();
  return (
    <Card
    className={`${
      isDarkMode ? "bg-gray-800 border-none " : "bg-white shadow-xl "}  py-3 rounded-3xl md:min-w-[500px]`}
  >
    <CardContent>
      <ChartContainer config={chartConfig} className="">
        <AreaChart
          accessibilityLayer
          data={props.data}
          margin={{
            top: 20,
            left:0,
            bottom:0,
            right:15
          }}
        >
          <CartesianGrid
            vertical={false}
            strokeDasharray="3 3"
            strokeWidth={2}
          />
          <XAxis
            dataKey="time"
            tickLine={false}
            axisLine={false}
            tickMargin={8}
            tickFormatter={(value) => value.slice(0, 5)}
          />
          <YAxis
            tickLine={false}
            axisLine={false}
            tickFormatter={(value) => `$${value}`}
          />
          <ChartTooltip
            cursor={false}
            content={<ChartTooltipContent indicator="dot" hideLabel />}
          />
          <Area
            dataKey="value"
            type="linear"
            fill="white"
            fillOpacity={0}
            stroke="#008010"
            strokeWidth={2}
            dot={{
              fill: "#008010",
              fillOpacity: 1,
              r: 4,
            }}
            activeDot={{
              r: 4,
            }}
          />
        </AreaChart>
      </ChartContainer>
    </CardContent>
  </Card>
  );
}
