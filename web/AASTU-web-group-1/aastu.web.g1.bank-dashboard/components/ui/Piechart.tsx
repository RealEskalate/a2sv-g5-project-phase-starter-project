"use client";

import { TrendingUp } from "lucide-react";
import { Label, Pie, PieChart, Sector } from "recharts";
import { PieSectorDataItem } from "recharts/types/polar/Pie";

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
  { browser: "BRC Bank", visitors: 275, fill: "#1EC6B8" },
  { browser: "ABM Bank", visitors: 200, fill: "#3464F3" },
  { browser: "DBL Bank", visitors: 187, fill: "#FFB11F" },
  { browser: "MCP Bank", visitors: 173, fill: "#FF6195" },
];

const chartConfig = {
  visitors: {
    label: "Visitors",
  },
  BRC_Bank: {
    label: "BRC Bank",
    color: "#1EC6B8",
  },
  ABM_Bank: {
    label: "ABM Bank",
    color: "#3464F3",
  },
  DBL_Bank: {
    label: "DBL Bank",
    color: "#FFB11F",
  },
  MCP_Bank: {
    label: "MCP Bank",
    color: "#FF6195",
  },
} satisfies ChartConfig;

export function Donut() {
  return (
    <Card className="flex flex-col">
      <CardHeader className="items-center pb-0"></CardHeader>
      <CardContent className="flex-1 pb-0">
        <ChartContainer
          config={chartConfig}
          className="mx-auto aspect-square max-h-[250px]"
        >
          <PieChart>
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent hideLabel />}
            />
            <Pie
              data={chartData}
              dataKey="visitors"
              nameKey="browser"
              innerRadius={60}
              strokeWidth={5}
              activeIndex={0}
              activeShape={({
                outerRadius = 0,
                ...props
              }: PieSectorDataItem) => (
                <Sector {...props} outerRadius={outerRadius + 10} />
              )}
            
            />
          </PieChart>
        </ChartContainer>
      </CardContent>
      <CardFooter className="flex flex-col gap-2 text-sm">
        {chartData.map((entry, index) => (
          <div key={index} className="flex items-center gap-2">
            <span
              className="inline-block h-4 w-4 rounded-full"
              style={{ backgroundColor: entry.fill }}
            />
            <span className="font-medium leading-none text-[#718EBF]">
              {entry.browser}
            </span>
          </div>
        ))}
      </CardFooter>
    </Card>
  );
}
