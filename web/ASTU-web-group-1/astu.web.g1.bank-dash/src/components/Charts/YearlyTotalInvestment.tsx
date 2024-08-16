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

const YearlyTotalInvestment = () => {
  // Define the chart data with type annotations
  const data: ChartData<"line"> = {
    labels: ["2016", "2017", "2018", "2019", "2020", "2021"],
    datasets: [
      {
        data: [65, 69, 80, 81, 56, 79],
        borderColor: "#FCAA0B",
        borderWidth: 2,
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
    <div className="w-full md:w-1/2">
      <h1 className="text-[#333B69] text-20px py-2 font-semibold">
        Yearly Total Investment
      </h1>
      <div className="bg-white p-6 rounded-3xl">
        <Line data={data} options={options} className="w-full" />
      </div>
    </div>
  );
};

export default YearlyTotalInvestment;
