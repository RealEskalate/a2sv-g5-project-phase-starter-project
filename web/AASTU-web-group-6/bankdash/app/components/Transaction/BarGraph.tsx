"use client";

import { useState } from "react";
import { Bar, BarChart, CartesianGrid, XAxis, Cell } from "recharts";
import { DailyAmount } from "@/types/TransactionValue";

import { Card, CardContent } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

const chartConfig = {
  desktop: {
    label: "amount",
    color: "#EDF0F7",
  },
} satisfies ChartConfig;

const BarGraph = ({ chartData }: { chartData: DailyAmount[] }) => {
  const [activeIndex, setActiveIndex] = useState<number | null>(null);
  const handleMouseOver = (index: number) => {
    setActiveIndex(index);
  };

  const handleMouseOut = () => {
    setActiveIndex(null);
  };

  return (
    <div className="space-y-3">
      <Card className="rounded-[25px] dark:bg-[#232328]">
        <CardContent>
          <ChartContainer config={chartConfig} className="h-56 w-full">
            <BarChart data={chartData}>
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
                content={<ChartTooltipContent hideLabel />}
              />
              <Bar dataKey="amount" radius={8}>
                {chartData.map((entry, index) => (
                  <Cell
                    key={`cell-${index}`}
                    fill={
                      index === activeIndex
                        ? "#16DBCC" // Hover color
                        : "var(--color-desktop)" // Default color
                    }
                    onMouseOver={() => handleMouseOver(index)}
                    onMouseOut={handleMouseOut}
                  />
                ))}
              </Bar>
            </BarChart>
          </ChartContainer>
        </CardContent>
      </Card>
    </div>
  );
};

export default BarGraph;
