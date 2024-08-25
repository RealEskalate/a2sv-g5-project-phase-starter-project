"use client";
import React from "react";
import { Line } from "react-chartjs-2";
import {
  Chart as ChartJS,
  LineElement,
  CategoryScale,
  LinearScale,
  Title,
  Tooltip,
  Legend,
  PointElement,
} from "chart.js";

ChartJS.register(
  LineElement,
  CategoryScale,
  LinearScale,
  Title,
  Tooltip,
  PointElement
);

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
        label: "My Dataset",
        data: data.values,
        borderColor: "blue",
        backgroundColor: "rgba(0, 0, 0, 0)",
        fill: true,
        tension: 0.5,
        pointRadius: 0,
        Tooltip: {
          backgroundColor: "rgba(1, 0, 0, 0)",
        },
      },
    ],
  };

  const options = {
    responsive: true,
    plugins: {
      legend: {
        position: "top" as const,
      },
      tooltip: {
        callbacks: {
          label: (context: any) => {
            return `${context.dataset.label}: ${context.raw}`;
          },
        },
      },
    },
  };

  return (
    <div className="w-full max-w-md mx-auto my-10 bg-white shadow-lg rounded-3xl pt-10 pl-10 pb-10 pr-10">
      {" "}
    
      <Line data={chartData} options={options} />
    </div>
  );
};

export default LineGraph;
