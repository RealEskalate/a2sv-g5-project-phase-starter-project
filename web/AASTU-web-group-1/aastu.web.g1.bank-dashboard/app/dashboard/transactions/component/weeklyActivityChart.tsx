"use client";

import {
  Bar,
  BarChart,
  CartesianGrid,
  XAxis,
  YAxis,
  Tooltip,
  Legend,
} from "recharts";
import { useState, useEffect, useMemo } from "react";
import { Card, CardContent } from "@/components/ui/card";
import { ChartContainer, ChartTooltip, ChartTooltipContent } from "@/components/ui/chart";
import { TransactionContent } from "@/types";
import { useUser } from "@/contexts/UserContext";

interface BarchartProps {
  weeklyDeposit: TransactionContent[];
  weeklyWithdraw: TransactionContent[];
}

export function Barchart({ weeklyDeposit, weeklyWithdraw }: BarchartProps) {
  const { isDarkMode } = useUser();
  const [chartData, setChartData] = useState<
    { day: string; Deposite: number; Withdraw: number }[]
  >([]);

  const chartConfig = {
    Deposite: {
      label: "Deposite",
      color: isDarkMode ? "#4A90E2" : "#1814F3", // Light and dark mode colors
    },
    Withdraw: {
      label: "Withdraw",
      color: isDarkMode ? "#50E3C2" : "#16DBCC", // Light and dark mode colors
    },
  };

  const weekofdays = useMemo(
    () => ["Sat", "Sun", "Mon", "Tue", "Wed", "Thu", "Fri"],
    []
  );

  const income = useMemo(
    () => weeklyDeposit.map((deposit) => deposit.amount),
    [weeklyDeposit]
  );
  const Withdraw = useMemo(
    () => weeklyWithdraw.map((withdraw) => withdraw.amount),
    [weeklyWithdraw]
  );

  useEffect(() => {
    const newChartData: { day: string; Deposite: number; Withdraw: number }[] =
      [];
    weekofdays.map((day, index) => {
      newChartData.push({
        day: day,
        Deposite: income[index] || 0,
        Withdraw: Withdraw[index] || 0,
      });
    });
    setChartData(newChartData);
  }, [weeklyDeposit, weeklyWithdraw, weekofdays, income, Withdraw]);

  return (
    <Card
      className={` ${
        isDarkMode ? "bg-gray-800 border-none " : "bg-white"
      } py-5`}
    >
      <CardContent className="p-0 ">
        <ChartContainer
          config={chartConfig}
          className="h-64 w-full"
        >
          <BarChart data={chartData}>
            <CartesianGrid
              vertical={false}
              // strokeDasharray="3 3"
              stroke={isDarkMode ? "bg-gray-800" : "#dddddd"}
            />
            <XAxis
              dataKey="day"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={(value) => value.slice(0, 3)}
              tick={{ fill: isDarkMode ? "#cccccc" : "#666666" }} // XAxis tick color
            />
            <YAxis
              tick={{ fill: isDarkMode ? "#cccccc" : "#666666" }} // YAxis tick color
            />
            <ChartTooltip
              cursor={false}
               content={<ChartTooltipContent indicator="dashed" />}
            />
            <Legend
              wrapperStyle={{
                color: isDarkMode ? "#cccccc" : "#333333", // Legend text color
              }}
            />

            <Bar
              dataKey="Deposite"
              fill={chartConfig.Deposite.color}
              radius={8}
              barSize={10}
            />
            <Bar
              dataKey="Withdraw"
              fill={chartConfig.Withdraw.color}
              radius={8}
              barSize={10}
            />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
