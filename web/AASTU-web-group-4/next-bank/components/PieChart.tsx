'use client'
import { Pie } from 'react-chartjs-2';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import React from 'react';

ChartJS.register(ArcElement, Tooltip, Legend);

const PieChart = () => {
    const dummyData = {
        labels: ["Category 1", "Category 2", "Category 3", "Category 4"],
        datasets: [
          {
            label: "Distribution",
            data: [300, 250, 200, 150],
            backgroundColor: [
              "rgba(24, 20, 243, 0.2)", // Light blue with 20% opacity
              "rgba(10, 200, 243, 0.2)", // Light cyan with 20% opacity
              "rgba(243, 134, 10, 0.2)", // Light orange with 20% opacity
              "rgba(243, 10, 134, 0.2)", // Light pink with 20% opacity
            ],
            borderColor: [
              "#1814F3", // Original blue color
              "#0AC8F3", // Original cyan color
              "#F3860A", // Original orange color
              "#F30A86", // Original pink color
            ],
            borderWidth: 1,
          },
        ],
      };
      
      const options = {
        responsive: true,
        plugins: {
          legend: {
            position: "top" as const,
          },
        },
      };
  return <Pie data={dummyData} options={options} />;
};

export default PieChart;
