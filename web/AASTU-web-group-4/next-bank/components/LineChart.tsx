"use client";
import { Line } from "react-chartjs-2";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler
} from "chart.js";

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler
);

const LineChart = () => {


  const dummyData = {
    labels: ["2023-10", "2023-11", "2023-12", "2024-01", "2024-02", "2024-03"],
    datasets: [
      {
        label: "Balance History",
        data: [1500, 1750, 1600, 1800, 1700, 1900],
        borderColor: "#1814F3",
        backgroundColor: "rgba(24, 20, 243, 0.1)",
        borderWidth: 1,
        fill: true,
      },
    ],
  };
  


  return (
    <Line
      data={dummyData}
      options={{ responsive: true, plugins: { legend: { position: "top" } } }}
    />
  );
};

export default LineChart;
