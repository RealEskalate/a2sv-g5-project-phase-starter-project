"use client";

import * as React from "react";
import { Label, Pie, PieChart, Sector } from "recharts";
import { PieSectorDataItem } from "recharts/types/polar/Pie";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartStyle,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
const desktopData = [
  { month: "ABM Bank", desktop: 186, fill: "#16DBCC" },
  { month: "DBL Bank", desktop: 305, fill: "#4C78FF" },
  { month: "MCP Bank", desktop: 237, fill: "#FFBB38" },
  { month: "BRC Bank", desktop: 173, fill: "#FF82AC" },
];

const chartConfig = {
  visitors: {
    label: "Visitors",
  },
  desktop: {
    label: "Desktop",
  },
  mobile: {
    label: "Mobile",
  },
  january: {
    label: "January",
    color: "hsl(var(--chart-1))",
  },
  february: {
    label: "February",
    color: "hsl(var(--chart-2))",
  },
  march: {
    label: "March",
    color: "hsl(var(--chart-3))",
  },
  april: {
    label: "April",
    color: "hsl(var(--chart-4))",
  },
  may: {
    label: "May",
    color: "hsl(var(--chart-5))",
  },
} satisfies ChartConfig;

export function PieChartPage() {
  const id = "pie-interactive";
  const [activeMonth, setActiveMonth] = React.useState(desktopData[0].month);

  const activeIndex = React.useMemo(
    () => desktopData.findIndex((item) => item.month === activeMonth),
    [activeMonth]
  );
  const months = React.useMemo(() => desktopData.map((item) => item.month), []);

  return (
    <div className="flex flex-col gap-5 bg-white rounded-2xl py-3 justify-center px-4 items-center shadow-sm">
      <Card
        data-chart={id}
        className="flex flex-col w-52 shadow-none bg-transparent border-none "
      >
        <ChartStyle id={id} config={chartConfig} />

        <CardContent className="flex flex-1 justify-center p-0">
          <ChartContainer
            id={id}
            config={chartConfig}
            className="mx-auto aspect-square w-full max-w-[300px]"
          >
            <PieChart className="">
              <ChartTooltip
                cursor={false}
                content={<ChartTooltipContent hideLabel />}
              />
              <Pie
                data={desktopData}
                dataKey="desktop"
                nameKey="month"
                innerRadius={50}
                strokeWidth={5}
                activeIndex={[0, 1, 2, 3]}
                activeShape={({
                  outerRadius = 0,
                  ...props
                }: PieSectorDataItem) => (
                  <g>
                    <Sector {...props} outerRadius={outerRadius + 10} />
                    <Sector
                      {...props}
                      outerRadius={outerRadius + 25}
                      innerRadius={outerRadius + 12}
                    />
                  </g>
                )}
              >
                <Label
                  content={({ viewBox }) => {
                    if (viewBox && "cx" in viewBox && "cy" in viewBox) {
                      return (
                        <text
                          x={viewBox.cx}
                          y={viewBox.cy}
                          textAnchor="middle"
                          dominantBaseline="middle"
                        >
                          <tspan
                            x={viewBox.cx}
                            y={viewBox.cy}
                            className="fill-foreground text-3xl font-bold"
                          >
                            {desktopData[activeIndex].desktop.toLocaleString()}
                          </tspan>
                          <tspan
                            x={viewBox.cx}
                            y={(viewBox.cy || 0) + 24}
                            className="fill-muted-foreground"
                          >
                            Visitors
                          </tspan>
                        </text>
                      );
                    }
                  }}
                />
              </Pie>
            </PieChart>
          </ChartContainer>
        </CardContent>
      </Card>
      <div className="space-y-3">
        <ChartLabel
          colors={["#4C78FF", "#FF82AC"]}
          labels={["DBL Bank", "BRC Bank"]}
        />
        <ChartLabel
          colors={["#16DBCC", "#FFBB38"]}
          labels={["ABM Bank", "MCP Bank"]}
        />
      </div>
    </div>
  );
}

export const ChartLabel = ({
  colors,
  labels,
}: {
  colors: string[];
  labels: string[];
}) => {
  return (
    <div className="flex gap-5">
      {colors.map((color, index) => (
        <div key={index} className="flex justify-center gap-3">
          <div
            className="w-3 h-3 rounded-full"
            style={{ backgroundColor: color }}
          ></div>
          <p className="text-[#718EBF] text-xs font-medium w-20">
            {labels[index]}
          </p>
        </div>
      ))}
    </div>
  );
};

export default PieChartPage;
