"use client";

import React from "react";
import dynamic from "next/dynamic";
const Chart = dynamic(() => import("react-apexcharts"), { ssr: false });

const chartConfig = {
  // height: 300,
  // width: 710,
  series: [
    {
      name: "Deposit",
      data: [450, 250, 350, 400, 100, 300, 450],
    },
    {
      name: "Withdraw",
      data: [100, 400, 150, 300, 450, 200, 350],
    },
  ],
  options: {
    chart: {
      toolbar: {
        show: false,
      },
    },
    colors: ["#1814F3", "#16DBCC", "#16DBCC"], // Blue for Deposit, Green for Withdraw
    plotOptions: {
      bar: {
        columnWidth: "30%",
        borderRadius: 4,
        barSpacing: 10,
      },
    },
    dataLabels: {
      enabled: false,
    },
    xaxis: {
      categories: ["Sat", "Sun", "Mon", "Tue", "Wed", "Thu", "Fri"],
      labels: {
        style: {
          colors: "#718EBF",
          fontSize: "12px",
          fontFamily: "inherit",
          fontWeight: 400,
        },
      },
    },
    yaxis: {
      labels: {
        style: {
          colors: "#718EBF",
          fontSize: "12px",
          fontFamily: "inherit",
          fontWeight: 400,
        },
      },
    },
    grid: {
      borderColor: "#dddddd",
      strokeDashArray: 5,
    },
    tooltip: {
      theme: "light",
    },
  },
};

const WeeklyActivity = () => {
  return (
    <div className="max-w-full max-h-[300px] h-[300px]  overflow-hidden  bg-white rounded-lg shadow-md">
      <Chart
        options={chartConfig.options}
        series={chartConfig.series}
        type="bar"
        width={"100%"}
        height={"100%"}
      />
    </div>
  );
};

export default WeeklyActivity;
