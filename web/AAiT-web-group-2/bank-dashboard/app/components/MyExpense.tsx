'use client';

import React from "react";
import dynamic from "next/dynamic";
const Chart = dynamic(() => import("react-apexcharts"), { ssr: false });

const chartConfig = {
  type: "bar",
  height: 225,
  width: 350,
  series: [
    {
      name: "Earnings",
      data: [3000, 4000, 3200, 10000, 12500, 3000], // Example data
    },
  ],
  options: {
    chart: {
      toolbar: { show: false },
    },
    plotOptions: {
      bar: {
        borderRadius: 10,
        horizontal: false,
        columnWidth: "70%",
      },
    },
    colors: ["#16DBCC"],
    dataLabels: {
      enabled: false,
    },
    xaxis: {
      categories: ["Aug", "Sep", "Oct", "Nov", "Dec", "Jan"],
      labels: {
        style: {
          colors: "#718EBF",
          fontFamily: "inherit",
          fontSize: "12px",
          fontWeight: 400,
        },
      },
      axisBorder: {
        show: false, // Hides the bottom border line
      },
      axisTicks: {
        show: false, // Hides the bottom ticks
      },
    },
    yaxis: {
      show: false,
    },
    // tooltip: {
    //   enabled: true,
    //   y: {
    //     formatter: function (val) {
    //       return `$${val}`;
    //     },
    //   },
    // },
    grid: { show: false },
  },
};

const MyExpense = () => {
  return (
    <div className="px-2 h-[225px]  mx-auto w-[350px] max-md:w-[231px] max-md:h-[170px] rounded-my-card-radius drop-shadow-lg shadow-md overflow-hidden">
      <Chart options={chartConfig.options} series={chartConfig.series} type="bar" />
    </div>
  );
};

export default MyExpense;
