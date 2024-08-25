"use client";
import React from "react";
import { Line } from "react-chartjs-2";
import {
  Chart as ChartJS,
  LineElement,
  CategoryScale,
  LinearScale,
  Tooltip,
  PointElement,
} from "chart.js";

ChartJS.register(LineElement, CategoryScale, LinearScale, Tooltip, PointElement);

interface ChartProps {
  data: {
    labels: string[];
    values: number[];
  };
}

const LineGraph: React.FC<ChartProps> = ({ data }) => {
  const chartData = {
    labels: data.labels,
    datasets: [
      {
        label: "Balance",
        data: data.values,
        borderColor: "blue",
        backgroundColor: "rgba(0, 0, 255, 0.2)", // Light blue shade for the filled area
        fill: true, // Enable fill to create the shaded area under the line
        tension: 0.4,
        pointRadius: 3,
      },
    ],
  };

  const options = {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        display: false, 
      },
      tooltip: {
        callbacks: {
          label: (context: any) => {
            return `${context.raw}`; 
          },
        },
      },
    },
    scales: {
      x: {
        grid: {
          display: false, 
        },
        ticks: {
          maxTicksLimit: 5,
        },
      },
      y: {
        grid: {
          drawBorder: false, 
          color: "rgba(0, 0, 0, 0.1)", 
        },
        ticks: {
          display: false, 
        },
      },
    },
  };

  return (
    <div className="w-full h-full">
      <Line data={chartData} options={options} />
    </div>
  );
};

export default LineGraph;