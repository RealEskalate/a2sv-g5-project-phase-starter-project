"use client";

import { TrendingUp } from "lucide-react";
import { Area, AreaChart, CartesianGrid, XAxis, YAxis } from "recharts";

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
  { month: "2016", desktop: 5000 },
  { month: "2017", desktop: 22000 },
  { month: "2018", desktop: 17000 },
  { month: "2019", desktop: 35000 },
  { month: "2020", desktop: 21000 },
  { month: "2021", desktop: 29000 },
];

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;

export default function LineChartComp() {
  return (
    <div>
      <Card className="py-3 rounded-3xl" >
        <CardContent>
          <ChartContainer config={chartConfig} className="">
            <AreaChart
              accessibilityLayer
              data={chartData}
              margin={{
                top: 18,
              }}
            >
              <CartesianGrid vertical={false} strokeDasharray="3 3" strokeWidth={2}/>
              <XAxis
                dataKey="month"
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
                dataKey="desktop"
                type="linear"
                fill="white"
                fillOpacity={0.1}
                stroke="#EDA10D"
                strokeWidth={2.5}
                dot={{
                  fill: "white",
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
    </div>
  );
}
