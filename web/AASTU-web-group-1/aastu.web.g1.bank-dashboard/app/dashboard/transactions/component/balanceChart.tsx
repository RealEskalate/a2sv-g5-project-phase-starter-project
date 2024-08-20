"use client";

import { Card, CardContent } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { Area, AreaChart, CartesianGrid, Legend, XAxis, YAxis } from "recharts";

import { BalanceData } from "@/types";
import { useEffect, useState } from "react";
import { formatMonth } from "./utils";
import { useUser } from "@/contexts/UserContext";

const chartConfig = {
  value: {
    label: "Value",
    color: "#2D60FF",
  },
} satisfies ChartConfig;

export function BalanceAreachart({
  balanceHistory,
}: {
  balanceHistory: BalanceData[];
}) {
  const { isDarkMode } = useUser();
  const [chartData, setChartData] = useState<
    { month: string; value: number }[]
  >([]);

  useEffect(() => {
    const newChartData: { month: string; value: number }[] = [];
    balanceHistory.map((balance: BalanceData) => {
      newChartData.push({
        month: formatMonth(balance.time),
        value: Math.round(balance.value),
      });
    });
    setChartData(newChartData);
  }, [balanceHistory]);

  return (
    <Card className={`${isDarkMode ? "bg-gray-800 border-none" : "bg-white"} py-5`}>
      <CardContent className="px-0">
        <ChartContainer
          config={chartConfig}
          className={`h-40 w-full rounded-xl ${isDarkMode ? "bg-gray-800" : "bg-white"}`}
        >
          <AreaChart
            data={chartData}
            margin={{
              top: 10,
              right: 30,
              left: 0,
              bottom: 0,
            }}
          >
            <defs>
              <linearGradient id="colorValue" x1="0" y1="0" x2="0" y2="1">
                <stop
                  offset="5%"
                  stopColor={isDarkMode ? "#6EE7B7" : "#2D60FF"}
                  stopOpacity={0.8}
                />
                <stop
                  offset="95%"
                  stopColor={isDarkMode ? "#6EE7B7" : "#2D60FF"}
                  stopOpacity={0.01}
                />
              </linearGradient>
            </defs>
            <CartesianGrid
              strokeDasharray="3 3"
              stroke={isDarkMode ? "#4B5563" : "#E5E7EB"}
              vertical={false}
            />
            <XAxis
              dataKey="month"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              tickFormatter={(value) => value.slice(0, 3)}
              tick={{ fill: isDarkMode ? "#D1D5DB" : "#6B7280" }}
            />
            <YAxis
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              tick={{ fill: isDarkMode ? "#D1D5DB" : "#6B7280" }}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dot" />}
            />
            <Legend
              wrapperStyle={{
                color: isDarkMode ? "#D1D5DB" : "#6B7280",
              }}
            />
            <Area
              type="natural"
              dataKey="value"
              stroke={isDarkMode ? "#6EE7B7" : "#2D60FF"}
              fillOpacity={1}
              fill="url(#colorValue)"
            />
          </AreaChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
