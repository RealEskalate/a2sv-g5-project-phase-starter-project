"use client";

import {
  Bar,
  BarChart,
  CartesianGrid,
  LabelList,
  Legend,
  XAxis,
} from "recharts";
import { Card, CardContent } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { TransactionContent } from "@/types";
import { useEffect, useState } from "react";
import { useUser } from "@/contexts/UserContext";

interface ExpenseChartProps {
  expenses: TransactionContent[];
}

export function ExpenseChart({ expenses }: ExpenseChartProps) {
  const { isDarkMode } = useUser();
  const [chartData, setChartData] = useState<
    { month: string; expense: number }[]
  >([]);

  useEffect(() => {
    const monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun"];
    const newChartData = expenses.map((expense, index) => ({
      month: monthNames[index],
      expense: Math.round(expense.amount),
    }));
    setChartData(newChartData);
  }, [expenses]);

  const chartConfig = {
    desktop: {
      label: "expense",
      color: isDarkMode ? "#16DBCC" : "#16DBCC", // Adjust as needed for dark mode
    },
  } satisfies ChartConfig;

  return (
    <Card
      className={`${
        isDarkMode ? "bg-gray-800 " : "bg-white"
      } border-none py-4`}
    >
      <CardContent
        className=""
        style={{
          backgroundColor: isDarkMode ? "#1f2937" : "#ffffff", // Card background color
          borderColor: isDarkMode ? "#333333" : "#dddddd", // Card border color
        }}
      >
        <ChartContainer
          config={chartConfig}
          className={`h-40 w-full rounded-xl`}
        >
          <BarChart
            data={chartData}
            margin={{
              top: 20,
              right: 20,
              left: 0,
              bottom: 0,
            }}
          >
            <CartesianGrid
              stroke={isDarkMode ? "#444" : "#ddd"}
              vertical={false}
            />
            <XAxis
              dataKey="month"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={(value) => value.slice(0, 3)}
              stroke={isDarkMode ? "#ddd" : "#333"} // XAxis color
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent hideLabel />}
            />
            <Bar
              dataKey="expense"
              fill={isDarkMode ? "#1e3a8a" : "#EDF0F7"} // Bar color
              radius={8}
              barSize={30}
              className={`hover:fill-${isDarkMode ? "#16DBCC" : "#16DBCC"}`}
            >
              <LabelList
                position="top"
                offset={12}
                className={`fill-${isDarkMode ? "white" : "black"}`} // Label color
                fontSize={12}
              />
            </Bar>
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
