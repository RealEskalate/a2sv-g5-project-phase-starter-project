"use client";
import React, { useMemo, useRef } from "react";
import { SxProps, useMediaQuery, useTheme } from "@mui/material";
import EChartsReactCore from "echarts-for-react/lib/core";
import { BarChart, BarSeriesOption } from "echarts/charts";
import {
  GridComponent,
  GridComponentOption,
  LegendComponent,
  TitleComponent,
  TooltipComponentOption,
} from "echarts/components";
import * as echarts from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import { LegendComponentOption } from "echarts";
// import { useBreakpoints } from "providers/useBreakpoints"; // Adjust this path as per your implementation

// Compose the ECharts option type with necessary components
export type ECOption = echarts.ComposeOption<
  | BarSeriesOption
  | TooltipComponentOption
  | GridComponentOption
  | LegendComponentOption
>;

// Register necessary ECharts components
echarts.use([
  BarChart,
  LegendComponent,
  CanvasRenderer,
  GridComponent,
  TitleComponent,
]);

export type TransactionDataType = {
  day: string;
  deposit: number;
  withdraw: number;
}[];

export const transactionData: TransactionDataType = [
  { day: "Sat", deposit: 420, withdraw: 220 },
  { day: "Sun", deposit: 332, withdraw: 132 },
  { day: "Mon", deposit: 301, withdraw: 251 },
  { day: "Tue", deposit: 334, withdraw: 334 },
  { day: "Wed", deposit: 490, withdraw: 390 },
  { day: "Thu", deposit: 160, withdraw: 230 },
  { day: "Fri", deposit: 320, withdraw: 320 },
];

const BarStat: React.FC = () => {
  const theme = useTheme();
  const chartRef = useRef<EChartsReactCore | null>(null);
  //   const { up } = useBreakpoints(); // Custom hook for breakpoints, adjust this as per your setup
  const upSm = useMediaQuery(theme.breakpoints.up("sm"));
  const upMd = useMediaQuery(theme.breakpoints.up("md"));

  const barWidth = upMd ? 14 : upSm ? 9 : 6;
  const barSpacing = upMd ? 0.8 : upSm ? 0.6 : 0.4;

  const chartOptions: ECOption = useMemo(() => {
    const xAxisData = transactionData.map((item) => item.day);
    const depositData = transactionData.map((item) => item.deposit);
    const withdrawData = transactionData.map((item) => item.withdraw);

    return {
      xAxis: {
        axisLabel: {
          padding: 10,
          baseline: "top",
          color: theme.palette.primary.light,
          fontSize: 13,
        },
        axisLine: { show: false },
        axisTick: { show: false },
        type: "category",
        data: xAxisData,
      },
      yAxis: {
        axisLabel: { color: theme.palette.primary.light, fontSize: 20 },
        axisLine: { show: false },
        axisTick: { show: false },
        splitLine: {
          lineStyle: {
            color: theme.palette.secondary.contrastText,
          },
        },
        type: "value",
      },
      grid: {
        left: "2%",
        top: "15%",
        right: "2%",
        bottom: "5%",
        containLabel: true,
      },
      tooltip: {
        trigger: "item",
        formatter: "{b}: ${c}",
        // backgroundColor: theme.palette.neutral.dark,
        textStyle: { color: theme.palette.secondary.contrastText },
        borderWidth: 0,
        padding: 10,
      },
      legend: {
        data: [
          { name: "Deposit", icon: "circle" },
          { name: "Withdraw", icon: "circle" },
        ],
        itemGap: 33,
        itemHeight: 16,
        textStyle: {
          color: theme.palette.primary.light,
        },
        right: -2,
        zLevel: 10,
      },
      series: [
        {
          data: depositData,
          type: "bar",
          stack: "1",
          name: "Deposit",
          barWidth: barWidth,
          itemStyle: {
            borderRadius: 30,
          },
          color: "#1814F3",
          emphasis: {
            itemStyle: { color: theme.palette.primary.dark },
          },
          barGap: barSpacing,
          animationDuration: 500,
        },
        {
          data: withdrawData,
          type: "bar",
          stack: "2",
          name: "Withdraw",
          barWidth: barWidth,
          itemStyle: {
            borderRadius: 30,
          },
          color: theme.palette.success.light,
          animationDuration: 500,
        },
      ],
    };
  }, [theme, upMd, upSm]);

  return (
    <div className="flex flex-col gap-3 p-4 rounded-lg w-full">
      <h2 className="text-lg font-semibold text-gray-700">Weekly Activity</h2>
      <div className="bg-white p-5 rounded-2xl">
        <EChartsReactCore
          ref={chartRef}
          echarts={echarts}
          option={chartOptions}
          notMerge={true}
          lazyUpdate={true}
          style={{ height: "400px", width: "100%" }}
        />
      </div>
    </div>
  );
};

export default BarStat;
