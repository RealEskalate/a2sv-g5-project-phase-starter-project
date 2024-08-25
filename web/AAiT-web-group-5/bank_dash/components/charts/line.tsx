"use client";
import React, { useMemo } from "react";
import EChartsReactCore from "echarts-for-react/lib/core";
import { LineChart } from "echarts/charts";
import {
  GridComponent,
  LegendComponent,
  TooltipComponent,
  TitleComponent,
  ToolboxComponent,
} from "echarts/components";
import * as echarts from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";

// Register required components
echarts.use([
  LineChart,
  GridComponent,
  LegendComponent,
  TooltipComponent,
  TitleComponent,
  ToolboxComponent,
  CanvasRenderer,
]);

const GradientStackedAreaChart = () => {
  const chartOptions = useMemo(() => {
    return {
      color: ["#80FFA5", "#00DDFF", "#37A2FF", "#FF0087", "#FFBF00"],
      //   title: {
      //     text: "Gradient Stacked Area Chart",
      //   },
      tooltip: {
        trigger: "axis",
        axisPointer: {
          type: "cross",
          label: {
            backgroundColor: "#6a7985",
          },
        },
      },
      visualMap: [
        {
          show: true,
          type: "continuous",
          seriesIndex: 0,
          min: 0,
          max: 400,
        },
      ],
      //   legend: {
      //     data: ["Line 1"],
      //   },
      //   toolbox: {
      //     feature: {
      //       saveAsImage: {},
      //     },
      //   },
      grid: {
        left: "3%",
        right: "4%",
        bottom: "3%",
        containLabel: true,
      },
      xAxis: [
        {
          type: "category",
          boundaryGap: false,
          data: [
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
          ],
        },
      ],
      yAxis: [
        {
          type: "value",
        },
      ],
      series: [
        {
          name: "Line 1",
          type: "line",
          stack: "Total",
          smooth: true,
          lineStyle: {
            width: 2,
            color: "#13117a",
          },
          showSymbol: false,
          areaStyle: {
            opacity: 0.8,
            color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
              {
                offset: 0,
                color: "#9e9ae4",
              },
              {
                offset: 1,
                color: "#ffffff",
              },
            ]),
          },
          emphasis: {
            focus: "series",
          },
          data: [120, 232, 151, 264, 200, 340, 250, 230, 180, 300, 400, 300],
        },
      ],
    };
  }, []);

  return (
    <div className="flex flex-col gap-3 p-4 rounded-lg w-full">
      <h2 className="text-lg font-semibold text-gray-700">Weekly Activity</h2>
      <div className="bg-white rounded-2xl w-full">
        <EChartsReactCore
          echarts={echarts}
          option={chartOptions}
          notMerge={true}
          lazyUpdate={true}
          style={{ height: "300px", width: "100%" }}
        />
      </div>
    </div>
  );
};

export default GradientStackedAreaChart;
