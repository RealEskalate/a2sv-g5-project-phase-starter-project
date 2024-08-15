"use client";
import React from "react";
import "chart.js/auto";
import { ChartData } from "chart.js";
import { Line } from "react-chartjs-2";

import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from "chart.js";

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);

const MonthlyRevenue = () => {
  // Define the chart data with type annotations
  const data: ChartData<"line"> = {
    labels: ["2016", "2017", "2018", "2019", "2020", "2021"],
    datasets: [
      {
        data: [6, 65, 69, 80, 81, 79],
        borderColor: "#16DBCC",
        borderWidth: 2,
        tension: 0.2,
      },
    ],
  };

  const options = {
    responsive: true,
    plugins: {
      legend: {
        display: false,
      },
      title: {
        display: false,
      },
    },
  };

  return (
    <div className="w-1/2 bg-white p-6 rounded-3xl">
      <Line data={data} options={options} className="w-full" />
    </div>
  );
};

export default MonthlyRevenue;
