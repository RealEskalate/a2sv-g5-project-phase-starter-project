"use client";

import { Bar, BarChart, CartesianGrid, Legend, XAxis, YAxis } from "recharts";

import {
  Card,
  CardContent,
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
import { useGetRandomBalanceHistoryQuery } from "@/lib/redux/api/transactionsApi";

const chartConfig = {
  withdraw: {
    label: "Withdraw",
    color: "#1814F3",
  },
  deposit: {
    label: "Deposit",
    color: "#16DBCC",
  },
} satisfies ChartConfig;

export function BarChartComponent() {
  const { data, isLoading, error } = useGetRandomBalanceHistoryQuery({
    monthsBeforeFirstTransaction: 14,
  });

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading data</div>;

  // Extract and transform data for the chart
  const rawData = data?.data || [];
  const totalEntries = 14;
  const half = Math.ceil(totalEntries / 2);

  const slicedData = rawData.slice(0, totalEntries);

  const weekdays = ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"];

  const withdrawData = slicedData
    .slice(0, half)
    .map((item: { time: string; value: number }, index: number) => ({
      day: weekdays[index % 7],
      withdraw: item.value,
      deposit: 0,
    }));

  const depositData = slicedData
    .slice(half)
    .map((item: { time: string; value: number }) => ({
      day: new Date(item.time).toLocaleString("default", { weekday: "short" }),
      withdraw: 0,
      deposit: item.value,
    }));

  //format the data
  const chartData = withdrawData.map(
    (withdrawItem: { time: string; value: number }, index: number) => ({
      ...withdrawItem,
      deposit: depositData[index]?.deposit || 0,
    })
  );

  return (
    <Card className="w-full h-auto lg:ml-4 lg:rounded-3xl  dark:bg-darkComponent">
      <CardContent className="p-0 lg:h-auto ">
        <ChartContainer config={chartConfig} className="lg:h-[400px]">
          <BarChart data={chartData} barSize={20}>
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="day"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
            />
            <YAxis
              tickCount={6}
              domain={[0, 5000]}
              ticks={[0, 1000, 2000, 3000, 4000, 5000]}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Bar
              dataKey="withdraw"
              fill={chartConfig.withdraw.color}
              radius={15}
            />
            <Bar
              dataKey="deposit"
              fill={chartConfig.deposit.color}
              radius={15}
            />

            <Legend
              verticalAlign="top"
              align="right"
              height={36}
              width={500}
              iconType="circle"
              iconSize={15}
              payload={[
                {
                  value: chartConfig.deposit.label,
                  type: "circle",
                  color: chartConfig.deposit.color,
                },
                {
                  value: chartConfig.withdraw.label,
                  type: "circle",
                  color: chartConfig.withdraw.color,
                },
              ]}
            />
          </BarChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
