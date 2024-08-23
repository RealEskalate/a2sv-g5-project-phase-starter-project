"use client";

import React from "react";
import dynamic from "next/dynamic";
const Chart = dynamic(() => import("react-apexcharts"), { ssr: false });

const chartConfig = {
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
    <div className="px-2 pb-0 mx-auto  bg-white rounded-my-card-radius shadow-md">
      <Chart
        options={chartConfig.options}
        series={chartConfig.series}
        type="bar"
        height={300}
        width={710}
      />
    </div>
  );
};

export default WeeklyActivity;
