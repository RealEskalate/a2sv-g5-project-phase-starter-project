'use client'
import React from 'react';
import { Bar } from 'react-chartjs-2';
import { Chart as ChartJS, CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend } from 'chart.js';

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

const BarChart = () => {
    const dummyData = {
        labels: ["2023-10", "2023-11", "2023-12", "2024-01", "2024-02", "2024-03"],
        datasets: [
          {
            label: "Monthly Balance",
            data: [1500, 1750, 1600, 1800, 1700, 1900],
            backgroundColor: "rgba(24, 20, 243, 0.2)", // 20% opacity
            borderColor: "#1814F3", // Original color for the border
            borderWidth: 1,
            borderRadius: 10, // Rounded bars
            barThickness: 10, // Bar thickness
          },
          {
            label: "Balance",
            data: [1500, 1750, 1600, 1800, 1700, 1900],
            backgroundColor: "rgba(10, 200, 243, 0.2)", // 20% opacity
            borderColor: "#1814F3", // Original color for the border
            borderWidth: 1,
            borderRadius: 10, // Rounded bars
            barThickness: 10, // Bar thickness
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
        scales: {
          x: {
            stacked: false,
          },
          y: {
            beginAtZero: true,
          },
        },
      };

    return <Bar data={dummyData} options={options} />;
};

export default BarChart;
