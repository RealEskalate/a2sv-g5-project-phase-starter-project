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
      data: [25, 25, 25, 25], // Adjust these values as per your requirement
      backgroundColor: ["#22E2B8", "#FF678E", "#FFBF47", "#4E88FF"],
      hoverBackgroundColor: ["#22E2B8", "#FF678E", "#FFBF47", "#4E88FF"],
    },
  ],
};

const options = {
  responsive: true,
  plugins: {
    legend: {
      position: "bottom" as const,
      rectancle: "round" as const,
    },
  },
};

const DoughnutChart = () => {
  return (
    <div className="w-auto h-72 md:h-[240px] flex justify-between items-center">
      <Doughnut data={data} options={options} />
    </div>
  );
};

export default DoughnutChart;
