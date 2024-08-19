"use client";
import { Bar, BarChart, CartesianGrid, Legend, XAxis } from "recharts";

import { Card, CardContent } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
const chartData = [
  { day: "Monday", debit: 186, credit: 80 },
  { day: "Tuesday", debit: 305, credit: 200 },
  { day: "Wednesday", debit: 237, credit: 120 },
  { day: "Thursday", debit: 73, credit: 190 },
  { day: "Friday", debit: 209, credit: 130 },
  { day: "Saturday", debit: 214, credit: 140 },
  { day: "Sunday", debit: 214, credit: 140 },
];

const chartConfig = {
  debit: {
    label: "Debit",
    color: "#1814F3", // Debit color
  },
  credit: {
    label: "Credit",
    color: "#FC7900", // Credit color
  },
} satisfies ChartConfig;

const CustomLegend = (props: any) => {
  const { payload } = props;

  return (
    <ul className="flex justify-center gap-4">
      {payload.map((entry: any, index: number) => (
        <li key={`item-${index}`} className="flex items-center gap-2">
          <div
            style={{ backgroundColor: entry.color }}
            className="w-4 h-4 rounded-sm"
          />
          <span className="text-sm font-medium text-[#718EBF]">
            {entry.value.charAt(0).toUpperCase() + entry.value.slice(1)}
          </span>
        </li>
      ))}
    </ul>
  );
};

export default function BarChartForAccounts() {
  return (
    <Card>
      <CardContent>
        <ChartContainer config={chartConfig}>
          <BarChart accessibilityLayer data={chartData}>
            <Legend
              verticalAlign="top"
              align="center"
              wrapperStyle={{ paddingBottom: "20px" }}
              content={<CustomLegend />}
            />
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="day"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Bar dataKey="debit" fill="var(--color-debit)" radius={4} />
            <Bar dataKey="credit" fill="var(--color-credit)" radius={4} />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
