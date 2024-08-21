"use client";

import { Area, AreaChart, CartesianGrid, XAxis } from "recharts";
import { Card, CardContent } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { BalanceType } from "@/app/Redux/slices/TransactionSlice";

const chartConfig = {
  desktop: {
    label: "Balance",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;

interface dataProps {
  data: BalanceType[];
}

export function AreaComp({ data }: dataProps) {
  const monthNames = [
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Nov",
    "Dec",
  ];

  // Transforming the API data
  const chartData = data.map((item) => {
    const [year, month] = item.time.split("-");
    return {
      month: monthNames[parseInt(month) - 1], // Convert month number to month name
      desktop: item.value, // Assign value to the desktop field
    };
  });

  console.log(chartData);
  return (
    <Card className="w-full border-0 shadow-none bg-transparent ">
      <CardContent className="p-0">
        <ChartContainer
          config={chartConfig}
          className=" w-full min-h-[100px] h-52"
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
