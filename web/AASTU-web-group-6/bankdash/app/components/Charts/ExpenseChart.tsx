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
  { browser: "chrome", visitors: 275, fill: "var(--color-chrome)" },
  { browser: "safari", visitors: 200, fill: "var(--color-safari)" },
  { browser: "firefox", visitors: 187, fill: "var(--color-firefox)" },
  { browser: "edge", visitors: 173, fill: "var(--color-edge)" },
];

const chartConfig = {
  visitors: {
    label: "Visitors",
  },
  chrome: {
    label: "ABM Bank",
    color: "#16DBCC",
  },
  safari: {
    label: "DBL Bank",
    color: "#3464F3",
  },
  firefox: {
    label: "MCP Bank",
    color: "#FFB11F",
  },
  edge: {
    label: "BRC Bank",
    color: "#FF82AC",
  },
} satisfies ChartConfig;

export default function Component() {
  return (
    <div className="flex w-full flex-col items-center justify-center gap-[-10px] p-8 bg-white dark:bg-[#232328] rounded-3xl shadow-sm">
      <Card className="w-full border-0 shadow-none bg-transparent">
        {" "}
        <ChartContainer
          config={chartConfig}
          className="we mx-auto aspect-square max-h-[250px] lg:h-64"
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
              innerRadius={30}
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
      </Card>
      <div className="flex items-center justify-center gap-8">
        <div>
          <div className="flex items-center gap-2 pb-3">
            <div className="circle w-4 h-4 rounded-full bg-[#4C78FF]"></div>
            <p className="text-[#718EBF] font-[15px]">DBL Bank</p>
          </div>
          <div className="flex items-center gap-2">
            <div className="circle w-4 h-4 rounded-full bg-[#16DBCC]"></div>
            <p className="text-[#718EBF] font-[15px]">ABM Bank</p>
          </div>
        </div>
        <div>
          <div className="flex items-center gap-2 pb-3">
            <div className="circle w-4 h-4 rounded-full bg-[#FF82AC]"></div>
            <p className="text-[#718EBF] font-[15px]">BRC Bank</p>
          </div>
          <div className="flex items-center gap-2">
            <div className="circle w-4 h-4 rounded-full bg-[#FFBB38]"></div>
            <p className="text-[#718EBF] font-[15px]">MCP Bank</p>
          </div>
        </div>
      </div>
    </div>
  );
}
