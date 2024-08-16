"use client";
import React from "react";
import { Doughnut } from "react-chartjs-2";
import {
  Chart as ChartJS,
  ArcElement,
  Tooltip,
  Legend,
  CategoryScale,
  LinearScale,
} from "chart.js";

ChartJS.register(ArcElement, Tooltip, Legend, CategoryScale, LinearScale);

const data = {
  labels: ["MCP Bank", "ABM Bank", "BRC Bank", "DBL Bank"],
  datasets: [
    {
      data: [25, 25, 25, 25],
      backgroundColor: ["#22E2B8", "#FF678E", "#FFBF47", "#4E88FF"],
      hoverBackgroundColor: ["#22E2B8", "#FF678E", "#FFBF47", "#4E88FF"],
    },
  ],
};

const options = {
  responsive: true,
  maintainAspectRatio: false, 
  plugins: {
    legend: {
      position: "bottom" as const,
    },
  },
};

const DoughnutChart = () => {
  return (
    <div className="w-full h-full flex justify-center items-center">
      <div className="w-full h-80 md:h-full p-3">
        <Doughnut data={data} options={options} />
      </div>
    </div>
  );
};

export default DoughnutChart;
