"use client";

import { Area, AreaChart, CartesianGrid, Legend, XAxis, YAxis } from "recharts";
import { Card, CardContent } from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { useEffect, useState } from "react";
import { getRandomBalanceHistory } from "@/services/transactionfetch";
import { TbFileSad } from "react-icons/tb";
import { colors } from "@/constants/index";

// Mapping of numeric month values to month names
const monthNames = [
  "January",
  "February",
  "March",
  "April",
  "May",
  "June",
  "July",
  "August",
  "September",
  "October",
  "November",
  "December",
];

const chartConfig = {
  balance: {
    color: "#1814F3",
  },
} satisfies ChartConfig;

export default function LineChart() {
  const [lineChartData, setLineChartData] = useState([]);
  const [status, setStatus] = useState<
    "loading" | "success" | "error" | "nodata"
  >("loading");

  useEffect(() => {
    const fetchLineChart = async () => {
      setStatus("loading");
      try {
        const response = await getRandomBalanceHistory();

        if (response && response.data) {
          const mappedData = response.data.map(
            (item: { time: string; value: number }) => {
              const monthIndex = parseInt(item.time.split("-")[1]) - 1;
              return {
                month: monthNames[monthIndex],
                balance: item.value,
              };
            }
          );

          if (mappedData.length === 0) {
            setStatus("nodata");
          } else {
            setLineChartData(mappedData);
            setStatus("success");
          }
        } else {
          setStatus("error");
        }
      } catch (error) {
        console.error("Error fetching the random balance history: ", error);
        setStatus("error");
      }
    };
    fetchLineChart();
  }, []);

  if (status === "loading") {
    return (
      <div className="flex bg-white rounded-2xl flex-col py-4 items-start gap-3 w-[100%] dark:bg-dark text-gray-900 dark:text-white animate-pulse">
        <div className="w-full h-[300px] bg-gray-300 dark:bg-gray-600 rounded-2xl"></div>
      </div>
    );
  }

  if (status === "error") {
    return (
      <div className="flex flex-col items-center justify-center h-full text-red-500">
        <TbFileSad
          className={`text-gray-300 dark:text-[#993d4b] w-[400px] h-[70px] pb-2 block mx-auto`}
          strokeWidth={1}
        />

        <div>Error fetching the balance history</div>
      </div>
    );
  }

  if (status === "nodata") {
    return (
      <div className="flex flex-col items-center justify-center h-full text-gray-500 dark:text-gray-300">
        <TbFileSad className="text-blue-500 dark:text-white w-[70px] h-[70px]" />
        <div>No data available to display</div>
      </div>
    );
  }

  return (
    <Card className="w-[100%] rounded-2xl">
      <CardContent>
        <ChartContainer config={chartConfig} className="w-[100%]">
          <AreaChart accessibilityLayer data={lineChartData}>
            <defs>
              <linearGradient id="colorBalance" x1="0" y1="0" x2="0" y2="1">
                <stop
                  offset="0%"
                  stopColor={chartConfig.balance.color}
                  stopOpacity={0.2}
                />
                <stop
                  offset="100%"
                  stopColor={chartConfig.balance.color}
                  stopOpacity={0.05}
                />
              </linearGradient>
            </defs>
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="month"
              tickLine={false}
              axisLine={true}
              tickMargin={8}
              tickFormatter={(value) => value.slice(0, 3)} // Show only the first 3 characters of the month
            />
            <YAxis dataKey="balance" axisLine={true} />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="line" />}
            />
            <Area
              dataKey="balance"
              type="natural"
              fill="url(#colorBalance)"
              stroke={chartConfig.balance.color}
              fillOpacity={1}
              strokeWidth={3}
            />
            <Legend verticalAlign="top" height={36} />
          </AreaChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
