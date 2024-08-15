"use client";

import { LabelList, Pie, PieChart } from "recharts";

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

const chartData = [
  { browser: "Services", visitors: 15, fill: "#FC7900", outerRadius: 130 },
  { browser: "Transfer", visitors: 30, fill: "#343C6A", outerRadius: 140 },
  { browser: "Shopping", visitors: 20, fill: "#FA00FF", outerRadius: 150 },
  { browser: "Other", visitors: 35, fill: "#1814F3", outerRadius: 160 },
];

const totalVisitors = 100;

const chartConfig = {
  Services: {
    label: `Services`,
    color: "#FC7900",
  },
  Transfer: {
    label: `Transfer`,
    color: "#343C6A",
  },
  Shopping: {
    label: `Shopping`,
    color: "#FA00FF",
  },
  Other: {
    label: `Other`,
    color: "#1814F3",
  },
} satisfies ChartConfig;

export function PieComp() {
  return (
    <Card className="w-full flex flex-col border-0 shadow-none bg-transparent">
      <CardContent className="flex-1 pb-0">
        <ChartContainer config={chartConfig} className="mx-auto aspect-square">
          <PieChart className="h-80">
            <ChartTooltip
              content={<ChartTooltipContent nameKey="visitors" hideLabel />}
            />
            <Pie
              data={chartData}
              dataKey="visitors"
              innerRadius={2} // Adjust the inner radius of the pie
              paddingAngle={5} // Add padding between each section
              // Use a fixed radius for the pie and handle slice rendering manually
              outerRadius={150} // Set a fixed maximum outer radius
            >
              {}
              <LabelList
                dataKey="browser"
                className="flex flex-col fill-background font-bold text-[13px] text-wrap"
                stroke="none"
                fontSize={12}
                formatter={(value: keyof typeof chartConfig) =>
                  chartConfig[value]?.label
                }
              />
            </Pie>
          </PieChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
