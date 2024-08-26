"use client";

import { LabelList, Pie, PieChart } from "recharts";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
  ChartConfig,
} from "@/components/ui/chart";

import { useGetRandomBalanceHistoryQuery } from "@/lib/redux/api/transactionsApi";

// Define the type for the data item
interface DataItem {
  time: string;
  value: number;
}

// Define the type for pie chart data
interface PieChartData {
  browser: keyof typeof chartConfig;
  visitors: number;
  fill: string;
}

// Define the chart configuration
const chartConfig = {
  chrome: {
    label: "Transfer",
    color: "#343C6A",
  },
  safari: {
    label: "Shopping",
    color: "#FC7900",
  },
  firefox: {
    label: "Services",
    color: "#1814F3",
  },
  edge: {
    label: "Others",
    color: "#FA00FF",
  },
} satisfies ChartConfig;

export function PieChartComponent() {
  const { data, isLoading, error } = useGetRandomBalanceHistoryQuery({
    monthsBeforeFirstTransaction: 4,
  });

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading data</div>;

  const pieChartData: PieChartData[] = data?.data.map(
    (item: DataItem, index: number) => {
      const browsers: Array<keyof typeof chartConfig> = [
        "chrome",
        "safari",
        "firefox",
        "edge",
      ];
      const browser = browsers[index % browsers.length];

      return {
        browser,
        visitors: item.value,
        fill: chartConfig[browser].color,
      };
    }
  );

  return (
    <Card className="w-full h-full lg:h-fit lg:rounded-3xl">
      <CardContent className="flex-1 pb-0 lg:h-[400px]">
        <ChartContainer
          config={chartConfig}
          className="mx-auto aspect-square   lg:h-[400px] lg:w-full"
        >
          <PieChart>
            <ChartTooltip
              content={<ChartTooltipContent nameKey="visitors" hideLabel />}
            />
            <Pie data={pieChartData} dataKey="visitors">
              <LabelList
                dataKey="browser"
                className="fill-background"
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
