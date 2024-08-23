"use client";

import { TrendingUp } from "lucide-react";
import { CartesianGrid, Line, LineChart, XAxis, YAxis } from "recharts";

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
interface YearlyInvestment {
  time: string;
  value: number;
}

interface MonthlyRevenue {
  time: string;
  value: number;
}

interface InvestmentData {
  totalInvestment: number;
  rateOfReturn: number;
  yearlyTotalInvestment: YearlyInvestment[];
  monthlyRevenue: MonthlyRevenue[];
}
interface YearlyInvestProps {
  data: InvestmentData | undefined;
}
function formatMonth(time: string): string {
  const [month] = time.split("/");
  const monthNames = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ];
  const monthIndex = parseInt(month, 10) - 1;
  return monthNames[monthIndex];
}

export function MonthlyRev({ data }: YearlyInvestProps) {
  const formchartData = data?.monthlyRevenue;
  const chartData = formchartData
    ?.map((item) => ({
      ...item,
      time: formatMonth(item.time),
    }))
    .slice()
    .reverse();

  const chartConfig = {
    desktop: {
      label: "Desktop",
      color: "hsl(var(--chart-1))",
    },
  } satisfies ChartConfig;
  return (
    <Card className=" rounded-3xl py-5 dark:bg-[#232328]">
      <CardContent>
        <ChartContainer config={chartConfig}>
          <LineChart
            accessibilityLayer
            data={chartData}
            margin={{
              left : 6,
              right: 12,
            }}
          >
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="time"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              padding={{ left: 20 }}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <YAxis
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              ticks={[0, 2000, 4000, 6000, 8000]}
              tickFormatter={(value) => `$${value.toLocaleString()}`}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent hideLabel />}
            />
            <Line
              dataKey="value"
              type="natural"
              stroke="#16DBCC"
              strokeWidth={3}
              dot={false}
            />
          </LineChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
