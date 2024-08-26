'use client'
import * as React from "react";
import { Label, Legend, Pie, PieChart, Sector } from "recharts";
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

const desktopData = [
  { month: "january", desktop: 186, fill: "#FF6384" },  // Red
  { month: "february", desktop: 305, fill: "#36A2EB" }, // Blue
  { month: "march", desktop: 237, fill: "#FFCE56" },    // Yellow
  { month: "april", desktop: 173, fill: "#4BC0C0" },    // Teal
  { month: "may", desktop: 209, fill: "#9966FF" },      // Purple
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
    color: "#FF6384",  // Red
  },
  february: {
    label: "February",
    color: "#36A2EB",  // Blue
  },
  march: {
    label: "March",
    color: "#FFCE56",  // Yellow
  },
  april: {
    label: "April",
    color: "#4BC0C0",  // Teal
  },
  may: {
    label: "May",
    color: "#9966FF",  // Purple
  },
} satisfies ChartConfig;

export default function Component() {
  const id = "pie-interactive";
  const [activeIndex, setActiveIndex] = React.useState(0);

  return (
    <Card data-chart={id} className="flex flex-col rounded-3xl">
      <ChartStyle id={id} config={chartConfig} />
      <CardContent className="flex flex-1 justify-center pb-0">
        <ChartContainer
          id={id}
          config={chartConfig}
          className="mx-auto aspect-square w-full max-w-[300px]"
        >
          <PieChart>
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent hideLabel />}
            />
            <Pie
              data={desktopData}
              dataKey="desktop"
              nameKey="month"
              innerRadius={60}
              strokeWidth={5}
              activeIndex={activeIndex}
              onMouseEnter={(_, index) => setActiveIndex(index)}
              className="dark:text-white"
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
                        className="dark:text-white]"
                      >
                        <tspan
                          x={viewBox.cx}
                          y={viewBox.cy}
                          className="dark:text-yellow-500 fill-foreground text-3xl font-bold"
                        >
                          {desktopData[activeIndex].desktop.toLocaleString()}
                        </tspan>
                        <tspan
                          x={viewBox.cx}
                          y={(viewBox.cy || 0) + 24}
                          className="dark:text-yellow-500 fill-muted-foreground"
                        >
                          expense
                        </tspan>
                      </text>
                    );
                  }
                }}
              />
            </Pie>
            <Legend
              layout="horizontal"
              verticalAlign="bottom"
              align="center"
            />
          </PieChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}