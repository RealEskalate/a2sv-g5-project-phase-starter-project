"use client";
import React from "react";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
import { Doughnut } from "react-chartjs-2";

// Register the necessary components with Chart.js
ChartJS.register(ArcElement, Tooltip, Legend);

const data = {
  plugins: {
    legend: {
      display: false,
    },
  },
  labels: ["Red", "Blue", "Yellow"],
  datasets: [
    {
      label: "My First Dataset",
      data: [10, 10, 10, 10],
      backgroundColor: ["#4C78FF", "#16DBCC", "#FF82AC", "#FFBB38"],
      hoverOffset: 4,
    },
  ],
};

const config = {
  type: "doughnut" as const,
  data: data,
};

const options = {
  plugins: {
    legend: {
      display: false,
    },
  },
  responsive: true,
};

const CardExpenceStatistics: React.FC = () => {
  return (
    <div>
      <h1 className="text-[#333B69] pb-2 font-semibold">
        Card and Expence Statistics
      </h1>
      <div className="w-full bg-white p-6 rounded-3xl">
        <div className="mx-14 md:mx-32 lg:mx-12 xl:mx-16 mt-2 flex justify-center">
          <Doughnut data={data} options={options} />
        </div>
        <div className="flex justify-around mt-4">
          <div className="space-y-4">
            <h1 className="flex items-center">
              {" "}
              <p className="w-4 h-4 rounded-full bg-[#4C78FF] me-3"></p> DBL
              Bank
            </h1>
            <h1 className="flex items-center">
              {" "}
              <p className="w-4 h-4 rounded-full bg-[#16DBCC] me-3"></p> AMB
              Bank
            </h1>
          </div>
          <div className="space-y-4">
            <h1 className="flex items-center">
              {" "}
              <p className="w-4 h-4 rounded-full bg-[#FF82AC] me-3"></p> RBC
              Bank
            </h1>
            <h1 className="flex items-center">
              {" "}
              <p className="w-4 h-4 rounded-full bg-[#FFBB38] me-3"></p> MCP
              Bank
            </h1>
          </div>
        </div>
      </div>
    </div>
  );
};

export default CardExpenceStatistics;
