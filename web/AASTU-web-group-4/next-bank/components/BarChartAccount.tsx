// import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts';

// const data = [
//   {
//     "name": "Page A",
//     "uv": 4000,
//   },
//   {
//     "name": "Page B",
//     "uv": 3000,
//   },
//   {
//     "name": "Page C",
//     "uv": 2000,
//   },
//   {
//     "name": "Page D",
//     "uv": 2780,
//   },
//   {
//     "name": "Page E",
//     "uv": 1890,
//   },
//   {
//     "name": "Page F",
//     "uv": 2390,
//   },
//   {
//     "name": "Page G",
//     "uv": 3490,
//   }
// ];

// const MyBarChart = () => (
//   <BarChart width={730} height={250} data={data}>
//     <CartesianGrid strokeDasharray="3 3" />
//     <XAxis dataKey="name" />
//     <YAxis />
//     <Tooltip />
//     <Legend />
//     <Bar dataKey="uv" fill="#82ca9d" />
//   </BarChart>
// );

// export default MyBarChart;

import React from "react";
import { Bar } from "react-chartjs-2";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from "chart.js";

// Register the components we will be using from Chart.js
ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

interface BarChartProps {
  labels: string[];
  data: number[];
  title: string;
}

const BarChart: React.FC<BarChartProps> = ({ labels, data, title }) => {
  const chartData = {
    labels: labels,
    datasets: [
      {
        label: title,
        data: data,
        backgroundColor: "rgba(75, 192, 192, 0.2)",
        borderColor: "rgba(75, 192, 192, 1)",
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
      title: {
        display: true,
        text: title,
      },
    },
    scales: {
      y: {
        beginAtZero: true,
      },
    },
  };

  return <Bar data={chartData} options={options} />;
};

export default BarChart;
