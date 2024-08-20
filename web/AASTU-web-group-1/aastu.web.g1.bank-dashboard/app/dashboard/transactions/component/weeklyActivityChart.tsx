"use client";

import {
  Bar,
  BarChart,
  CartesianGrid,
  XAxis,
  YAxis,
  Legend,
  Tooltip,
  ResponsiveContainer,
} from "recharts";
import { useState, useEffect, useMemo } from "react";
import { Card, CardContent } from "@/components/ui/card";
import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { TransactionData } from "@/types";

interface BarchartProps {
  weeklyDeposit: TransactionData[];
  weeklyWithdraw: TransactionData[];
}

export function Barchart({ weeklyDeposit, weeklyWithdraw }: BarchartProps) {
  const [chartData, setChartData] = useState<
    { day: string; Deposite: number; Withdraw: number }[]
  >([]);
  const newChartData: { day: string; Deposite: number; Withdraw: number }[] = [];
  const chartConfig = {
    Deposite: {
      label: "Deposite",
      color: "#1814F3",
    },
    Withdraw: {
      label: "Withdraw",
      color: "#16DBCC",
    },
  };
    const weekofdays =useMemo(()=> ["Sat", "Sun", "Mon", "Tue", "Wed", "Thu", "Fri"],[])
 const income = useMemo(
   () => weeklyDeposit.map((deposit) => deposit.amount),
   [weeklyDeposit]
 );
 const Withdraw = useMemo(
   () => weeklyWithdraw.map((withdraw) => withdraw.amount),
   [weeklyWithdraw]
 );
  weeklyDeposit.map((deposit, index) => {
    income.push(deposit.amount);
  });
  weeklyWithdraw.map((withdraw, index) => {
    Withdraw.push(withdraw.amount);
  });
  weekofdays.map((day, index) => {
    newChartData.push({
      day: day,
      Deposite: income[index],
      Withdraw: Withdraw[index],
    });
  });

 

  useEffect(() => {
    const newChartData: { day: string; Deposite: number; Withdraw: number }[] = [];
    weekofdays.map((day, index) => {
      newChartData.push({
        day: day,
        Deposite: income[index],
        Withdraw: Withdraw[index],
      });
    });
    setChartData(newChartData);
  }, [weeklyDeposit, weeklyWithdraw, weekofdays, income, Withdraw]);
  
  return (
    <Card >
      <CardContent className="p-0">
        <ChartContainer config={chartConfig} className="h-60 w-full">
        
            <BarChart data={chartData}>
              <CartesianGrid vertical={false} strokeDasharray="3 3" />
              <XAxis
                dataKey="day"
                tickLine={false}
                tickMargin={10}
                axisLine={false}
                tickFormatter={(value) => value.slice(0, 3)}
              />
              <YAxis />
              <Tooltip content={<ChartTooltipContent indicator="dashed" />} />
              <Legend />

              <Bar
                dataKey="Deposite"
                fill={chartConfig.Deposite.color}
                radius={10}
                barSize={10}
              />
              <Bar
                dataKey="Withdraw"
                fill={chartConfig.Withdraw.color}
                radius={10}
                barSize={10}
              />
            </BarChart>
        
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
