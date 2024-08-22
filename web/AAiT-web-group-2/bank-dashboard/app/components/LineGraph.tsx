"use client";
import React from "react";
import { Line } from "react-chartjs-2";
import { Card } from "@mui/material";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Filler,
  Tooltip,
  Legend,
} from "chart.js";

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Filler,
  Tooltip,
  Legend
);

const data = {
  labels: ["January", "February", "March", "April", "May", "June", "July"],
  datasets: [
    {
      label: "",
      data: [75, 100, 85, 120, 90, 100, 90],
      fill: true,
      backgroundColor: (context: any) => {
        const ctx = context.chart.ctx;
        const gradient = ctx.createLinearGradient(0, 0, 0, 400);
        gradient.addColorStop(0, "rgba(0, 0, 255, 0.2)");
        gradient.addColorStop(1, "rgba(255, 255, 255, 0.2)");
        return gradient;
      },
      borderColor: "rgba(0, 0, 255, 1)",
      borderWidth: 2,
      pointRadius: 0,
      tension: 0.4,
    },
  ],
};

const options = {
  responsive: true,
  plugins: {
    legend: {
      display: false,
    },
    tooltip: {
      enabled: true,
    },
  },
  scales: {
    x: {
      grid: {
        display: true,
        color: "rgba(128, 128, 128, 0.2)", // Gray broken line grid
        borderDash: [5, 5],
      },
    },
    y: {
      grid: {
        display: true,
        color: "rgba(128, 128, 128, 0.2)", // Gray broken line grid
        borderDash: [5, 5],
      },
    },
  },
  maintainAspectRatio: false,
};

const LineGraph: React.FC = () => {
  return (
    <Card
      className="shadow-lg rounded-lg p-4"
      style={{ width: "400px", height: "200px" }}
    >
      <Line data={data} options={options} />
    </Card>
  );
};

export default LineGraph;
