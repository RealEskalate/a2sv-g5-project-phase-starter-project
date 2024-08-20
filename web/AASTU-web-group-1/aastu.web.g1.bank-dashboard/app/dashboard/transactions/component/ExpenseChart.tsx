"use client";

import { Bar, BarChart, CartesianGrid, LabelList, Legend, XAxis } from "recharts";

import {
  Card,
  CardContent,
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { TransactionData } from "@/types";
import { useEffect, useState } from "react";

interface ExpenseChartProps {
  expenses: TransactionData[];
}

export function ExpenseChart({expenses}: ExpenseChartProps) {
 
  const [chartData, setChartData] = useState<{ month: string; expense: number }[]>([]);
  const monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun"];

  useEffect(() => {
    const newChartData = expenses.map((expense, index) => ({
      month: monthNames[index],
      expense: Math.round(expense.amount),
    }));
    setChartData(newChartData);
  }, [expenses]);
 
  

  const chartConfig = {
    desktop: {
      label: "expense",
      color: "#16DBCC",
    },
  } satisfies ChartConfig;
  return (
    <Card>
      <CardContent>
        <ChartContainer config={chartConfig} className="h-40 w-full">
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
