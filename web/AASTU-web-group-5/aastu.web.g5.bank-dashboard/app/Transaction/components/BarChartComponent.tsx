"use client";

import { useState } from "react";
import { Bar, BarChart, CartesianGrid, XAxis, Cell, LabelList } from "recharts";
import { Card, CardContent } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

const chartData = [
  { month: "January", desktop: 186 },
  { month: "February", desktop: 305 },
  { month: "March", desktop: 237 },
  { month: "April", desktop: 73 },
  { month: "May", desktop: 209 },
  { month: "June", desktop: 214 },
];

const chartConfig = {
  desktop: {
    label: "Desktop",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;

export function BarChartComponent() {
  const [activeIndex, setActiveIndex] = useState<number | null>(null);

  const handleMouseEnter = (index: number) => {
    setActiveIndex(index);
  };

  const handleMouseLeave = () => {
    setActiveIndex(null);
  };

  return (
    <Card>
      <CardContent>
        <ChartContainer config={chartConfig}>
          <BarChart data={chartData} onMouseLeave={handleMouseLeave}>
            <CartesianGrid vertical={false} horizontal={false} />
            <XAxis
              dataKey="month"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            
            <Bar dataKey="desktop" radius={10}>
              {chartData.map((entry, index) => (
                <Cell
                  key={`cell-${index}`}
                  fill={activeIndex === index ? "#12887E33" : "#EDF0F7"}
                  onMouseEnter={() => handleMouseEnter(index)}
                />
              ))}
              <LabelList
                dataKey="desktop"
                position="top"
                content={({ x, y, value, index }) =>
                  activeIndex === index ? (
                    <text
                      x={x}
                      y={y}
                      dy={-10}
                      fill="black"
                      fontSize={12}
                      textAnchor="top"
                    >
                      {value}
                    </text>
                  ) : null
                }
              />
            </Bar>
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
