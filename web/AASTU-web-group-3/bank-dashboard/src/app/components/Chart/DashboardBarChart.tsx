"use client";

import { useState, useMemo } from "react";
import { Bar, BarChart, CartesianGrid, XAxis, Cell } from "recharts";

import { Card, CardContent } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

import { useGetExpenseTransactionsQuery } from "@/lib/redux/api/transactionsApi";

export default function DashboardBarChart() {
  const { data, isLoading, isError } = useGetExpenseTransactionsQuery({
    size: 10, // Adjust the size to fetch all relevant data
    page: 0,
  });

  // Generate the last seven days including today
  const sevenDays = useMemo(() => {
    const days = [];
    const now = new Date();

    for (let i = 6; i >= 0; i--) {
      const day = new Date(now.getFullYear(), now.getMonth(), now.getDate() - i);
      days.push(day.toLocaleString("default", { weekday: "short", day: "numeric" }));
    }

    return days;
  }, []);

  // Process the data
  const chartData = useMemo(() => {
    if (!data || isLoading || isError) return [];

    const dailyExpenses: Record<string, number> = {};

    // Iterate over data.data.content
    data.data.content.forEach((transaction) => {
      const day = new Date(transaction.date).toLocaleString("default", {
        weekday: "short",
        day: "numeric",
      });

      if (dailyExpenses[day]) {
        dailyExpenses[day] += transaction.amount;
      } else {
        dailyExpenses[day] = transaction.amount;
      }
    });

    return sevenDays.map((day) => ({
      day,
      expense: Math.abs(dailyExpenses[day] || 0), 
    }));
  }, [data, isLoading, isError, sevenDays]);

  const [activeIndex, setActiveIndex] = useState<number | null>(null);

  const handleMouseOver = (index: number) => {
    setActiveIndex(index);
  };

  const handleMouseOut = () => {
    setActiveIndex(null);
  };

  return (
    <div className="flex flex-col w-full lg:w-1/3">
      <h1 className="font-semibold text-[#343C6A] h-20 flex items-center mx-2">
        My Expense
      </h1>
      <Card className="xl:w-[90%] mx-2 min-h-48">
        <CardContent>
          {isLoading ? (
            <p>Loading...</p>
          ) : isError ? (
            <p>Error loading data</p>
          ) : (
            <ChartContainer
              config={{ expense: { label: "Expenses", color: "#EDF0F7" } }}
              className="w-full px-1"
            >
              <BarChart accessibilityLayer data={chartData}>
                <CartesianGrid vertical={false} />
                <XAxis
                  dataKey="day"
                  tickLine={false}
                  tickMargin={10}
                  axisLine={false}
                />
                <ChartTooltip
                  cursor={false}
                  content={<ChartTooltipContent indicator="dashed" />}
                />
                <Bar dataKey="expense" radius={8}>
                  {chartData.map((entry, index) => (
                    <Cell
                      key={`cell-${index}`}
                      fill={
                        index === activeIndex
                          ? "#16DBCC" // Hover color
                          : "#EDF0F7" // Default color
                      }
                      onMouseOver={() => handleMouseOver(index)}
                      onMouseOut={handleMouseOut}
                    />
                  ))}
                </Bar>
              </BarChart>
            </ChartContainer>
          )}
        </CardContent>
      </Card>
    </div>
  );
}
