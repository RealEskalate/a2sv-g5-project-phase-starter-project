"use client";
import React, { useMemo } from "react";
import { useTheme } from "@mui/material";
import EChartsReactCore from "echarts-for-react/lib/core";
import { PieChart, PieSeriesOption } from "echarts/charts";
import {
  GridComponent,
  GridComponentOption,
  LegendComponent,
  TooltipComponentOption,
} from "echarts/components";
import * as echarts from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";

// Compose the ECharts option type with necessary components
export type ECOption = echarts.ComposeOption<
  PieSeriesOption | TooltipComponentOption | GridComponentOption
>;

// Register necessary ECharts components
echarts.use([PieChart, LegendComponent, CanvasRenderer, GridComponent]);

// Sample Expense Data
type ExpenseDataType = { value: number; name: string; selected: boolean }[];
const expenseData: ExpenseDataType = [
  { value: 26, name: "Service", selected: true },
  { value: 23, name: "Others", selected: true },
  { value: 26, name: "Shopping", selected: true },
  { value: 25, name: "Transfer", selected: true },
];

const PieStat = () => {
  const theme = useTheme();
  const { palette } = theme;

  // Memoize chart options for optimization
  const chartOptions: ECOption = useMemo(() => {
    return {
      backgroundColor: palette.common.white,
      tooltip: {
        trigger: "item",
      },
      color: ["#FC7900", "#1814F3", "#FA00FF", "#343C6A"],
      series: [
        {
          name: "Expense",
          type: "pie",
          selectedMode: "series",
          selectedOffset: 5,
          radius: "93%",
          center: ["50%", "50%"],
          roseType: "radius",
          avoidLabelOverlap: false,
          data: expenseData,
          startAngle: 45,
          label: {
            show: true,
            position: "inside",
            formatter: (params) => {
              return `{percent|${params.percent}%}\n{name|${params.name}}`;
            },
            rich: {
              percent: {
                fontSize: 16,
                fontWeight: "bold",
                color: palette.common.white,
              },
              name: {
                fontSize: 13,
                fontWeight: "bold",
                color: palette.common.white,
              },
            },
          },
          emphasis: {
            itemStyle: {
              borderColor: palette.common.white,
              borderWidth: 2,
            },
          },
          animationType: "expansion",
          animationEasing: "backOut",
          animationDuration: 1000,
        },
      ],
    };
  }, [palette]);

  return (
    <div className="flex flex-col gap-3 p-4 rounded-lg w-full">
      <h2 className="text-lg font-semibold text-center text-gray-700">
        Expense Statistics
      </h2>
      <div className="bg-white p-5 rounded-2xl">
        <EChartsReactCore
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

export default PieStat;
