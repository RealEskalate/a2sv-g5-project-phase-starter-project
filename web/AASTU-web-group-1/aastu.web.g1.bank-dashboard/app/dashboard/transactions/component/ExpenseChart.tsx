"use client";

import { TrendingUp } from "lucide-react";
import { Bar, BarChart, CartesianGrid, LabelList, XAxis } from "recharts";

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
  { month: "Aug", expense: 10000 },
  { month: "Sep", expense: 13000 },
  { month: "Oct", expense: 10000 },
  { month: "Nov", expense: 5000 },
  { month: "May", expense: 12500 },
  { month: "June", expense: 5000 },
];

const chartConfig = {
  desktop: {
    label: "expense",
    color: "#16DBCC",
  },
} satisfies ChartConfig;

export function ExpenseChart() {
  return (
    <Card>
      <CardContent>
        <ChartContainer config={chartConfig} className="h-36 w-full">
          <BarChart
            accessibilityLayer
            data={chartData}
            margin={{
              top: 20,
            }}
            width={10}
          >
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
              content={<ChartTooltipContent hideLabel />}
            />
            <Bar
              dataKey="expense"
              fill="#EDF0F7"
              radius={8}
              barSize={30}
              className="hover:fill-[#16DBCC]"
            >
              <LabelList
                position="top"
                offset={12}
                className="fill-foreground"
                fontSize={12}
              />
            </Bar>
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
