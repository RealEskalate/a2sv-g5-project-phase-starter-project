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
    label: "Chrome",
    color: "#16DBCC",
  },
  safari: {
    label: "Safari",
    color: "#3464F3",
  },
  firefox: {
    label: "Firefox",
    color: "#FFB11F",
  },
  edge: {
    label: "Edge",
    color: "#FF82AC",
  },
} satisfies ChartConfig;

export default function Component() {
  return (
    <Card className="flex flex-col h-80 p-8 border rounded-3xl">
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
      </CardContent>
      <CardFooter className="flex-col gap-2 text-sm">
        <div className="flex items-center justify-center gap-8">
          <div>
            <div className="flex items-center gap-1 pb-3">
              <img src="/assets/ellipseB.svg" alt="" />
              <p className="text-[#718EBF] font-[15px]">DBL Bank</p>
            </div>
            <div className="flex items-center gap-1">
              <img src="/assets/ellipseG.svg" alt="" />
              <p className="text-[#718EBF] font-[15px]">ABM Bank</p>
            </div>
          </div>
          <div>
            <div className="flex items-center gap-1 pb-3">
              <img src="/assets/ellipseP.svg" alt="" />
              <p className="text-[#718EBF] font-[15px]">BRC Bank</p>
            </div>
            <div className="flex items-center gap-1">
              <img src="/assets/ellipseO.svg" alt="" />
              <p className="text-[#718EBF] font-[15px]">MCP Bank</p>
            </div>
          </div>
        </div>
      </CardFooter>
    </Card>
  );
}
