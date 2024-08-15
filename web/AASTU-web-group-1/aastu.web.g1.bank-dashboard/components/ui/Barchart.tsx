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
import { useState, useEffect } from "react";
import { Card, CardContent } from "@/components/ui/card";
import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

const initialChartData = [
  { day: "Sat", Deposite: 186, Withdraw: 80 },
  { day: "Sun", Deposite: 305, Withdraw: 200 },
  { day: "Mon", Deposite: 237, Withdraw: 120 },
  { day: "Tue", Deposite: 73, Withdraw: 190 },
  { day: "wed", Deposite: 209, Withdraw: 130 },
  { day: "Thu", Deposite: 214, Withdraw: 140 },
  { day: "Fri", Deposite: 214, Withdraw: 140 },
];

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

export function Barchart() {
  const [chartData, setChartData] = useState(initialChartData);
  const [chartOptions, setChartOptions] = useState({});

  useEffect(() => {
    setChartOptions({
      maintainAspectRatio: false,
      responsive: true,
      plugins: {
        legend: {
          position: "top",
          labels: {
            usePointStyle: true,
            pointStyle: "circle",
            boxWidth: 10,
          },
        },
      },
    });
  }, []);

  return (
    <Card>
      <CardContent>
        <ChartContainer config={chartConfig} className="w-full h-96">
          <ResponsiveContainer width="100%" height="100%">
            <BarChart data={chartData} >
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
                barSize={20}
              />
              <Bar
                dataKey="Withdraw"
                fill={chartConfig.Withdraw.color}
                radius={10}
                barSize={20}
              />
            </BarChart>
          </ResponsiveContainer>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
