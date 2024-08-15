"use client";
import {
  LabelList,
  Pie,
  PieChart,
  Tooltip,
  ResponsiveContainer,
} from "recharts";
import {
  Card,
  CardContent,
 
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
} from "@/components/ui/chart";

const chartData = [
  { expense: "Other", value: 30, fill: "#1814F3" },
  { expense: "Investment", value: 20, fill: "#343C6A" },
  { expense: "Bill Expense", value: 15, fill: "#FC7900" },
  { expense: "Entertainment", value: 35, fill: "#FA00FF" },
];

const chartConfig = {
  value: {
    label: "Value",
  },
  other: {
    label: "Other",
    color: "#1814F3",
  },
  investment: {
    label: "Investment",
    color: "#343C6A",
  },
  billExpense: {
    label: "Bill Expense",
    color: "#FC7900",
  },
  entertainment: {
    label: "Entertainment",
    color: "#FA00FF",
  },
} satisfies ChartConfig;

export function Piechart() {
  return (
    <Card >
      <CardContent >
        <ChartContainer
          config={chartConfig}
          className="flex" // Increased the size for better visibility
        >
          <ResponsiveContainer width="100%" height={300}>
<PieChart>
            <Tooltip />
            <Pie
              data={chartData}
              dataKey="value"
              nameKey="expense"
              outerRadius="80%"
              fill="#8884d8"
              label
              labelLine
              paddingAngle={10}

            >
              <LabelList
                dataKey="expense"
                position="inside"
                fontSize={12} 
                formatter={(value: string) => value}
                stroke="#fff"
                className="font-inter font-light"
              />
            </Pie>
          </PieChart>
          </ResponsiveContainer>
          
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
