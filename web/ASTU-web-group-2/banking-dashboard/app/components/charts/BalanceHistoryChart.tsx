"use client";
import { useGetBalanceHistoryQuery } from "@/lib/service/TransactionService";
import { useRef, useEffect } from "react";
import { Chart, ChartData, ChartOptions } from "chart.js/auto";
import { useSession } from "next-auth/react";
import BarChartSkeleton from "./BarChartSkeleton";
import ErrorImage from "../Error/ErrorImage";

const generateMonths = () => {
  const months = [];
  const currentDate = new Date();

  for (let i = 6; i >= 0; i--) {
    const date = new Date(
      currentDate.getFullYear(),
      currentDate.getMonth() - i,
      1
    );
    const month = date.toLocaleString("default", { month: "short" });
    months.push(`${month}`);
  }

  return months;
};

const aggregateData = (data: { time: string; value: number }[]) => {
  const monthMap: { [key: string]: number } = {};

  const months = generateMonths();

  months.forEach((month) => {
    monthMap[month] = 0;
  });

  data.forEach((item) => {
    const [year, month] = item.time.split("-"); // Extract year and month from 'time'
    const date = new Date(parseInt(year), parseInt(month) - 1);
    const shortMonth = date.toLocaleString("default", { month: "short" });

    if (monthMap[shortMonth] !== undefined) {
      monthMap[shortMonth] += item.value;
    }
  });

  // Convert the monthMap into an array of objects for the chart
  const aggregatedData = months.map((month) => ({
    month,
    balance: monthMap[month],
  }));

  return aggregatedData;
};

export interface ChartRef extends HTMLCanvasElement {
  chart?: Chart;
}

function BalanceHistoryChart() {
  const { data: session, status } = useSession();
  const accessToken = session?.user.accessToken!;
  const { data, isError, isLoading } = useGetBalanceHistoryQuery(accessToken);
  const chartRef = useRef<ChartRef | null>(null);

  useEffect(() => {
    if (chartRef.current && data?.success) {
   
      if (chartRef.current.chart) {
        chartRef.current.chart.destroy();
      }

      const context = chartRef.current.getContext("2d");

      if (context) {
        const chartItem = context.canvas;

        // Extract the 'data' array from the API response
        const apiData = data.data;

        // Aggregate the data from the API and fill missing months
        const aggregatedData = aggregateData(apiData);

        // Extract labels and data for the chart
        const labels = aggregatedData.map((item) => item.month);
        const chartDataPoints = aggregatedData.map((item) => item.balance);

        // Create the linear gradient
        const gradient = context.createLinearGradient(
          0,
          0,
          0,
          chartItem.height
        );
        gradient.addColorStop(0, "rgba(45, 96, 255, 0.25)");
        gradient.addColorStop(1, "rgba(45, 96, 255, 0)");

        const chartData: ChartData<"line"> = {
          labels: labels,
          datasets: [
            {
              label: "Average Balance",
              data: chartDataPoints,
              fill: true, // Fill area under the line
              backgroundColor: gradient, // Use the gradient as the fill
              borderColor: "rgba(45, 96, 255, 1)", // Color of the line
              borderWidth: 2,
              tension: 0.4, // Set tension to curve the line
              pointRadius: 0, // Remove points
              pointBackgroundColor: "rgba(75, 192, 192, 1)",
              pointBorderColor: "#fff",
            },
          ],
        };

        const options: ChartOptions<"line"> = {
          responsive: true,
          plugins: {
            legend: {
              display: false,
            },
          },
          scales: {
            x: {
              ticks: {
                color: "rgb(113, 142, 191)",
                font: {
                  size: 12,
                },
              },
            },
            y: {
              beginAtZero: true,
              ticks: {
                color: "rgb(113, 142, 191)",
                font: {
                  size: 12,
                },
              },
            },
          },
        };

        const newChart = new Chart(chartItem, {
          type: "line",
          data: chartData,
          options,
        });

        chartRef.current.chart = newChart;
      }
    }
  }, [data]);

  if (isLoading) {
    return (
      <div className="w-full h-auto bg-gray-300 rounded-lg flex items-center justify-center animate-pulse">
     
      <svg width="100%" height="100%" viewBox="0 0 200 100">
        <polyline
          fill="none"
          stroke="#aaaaaa"
          strokeWidth="2"
          points="10,80 40,60 70,65 100,45 130,50 160,30 190,40"
        />
        <circle cx="10" cy="80" r="3" fill="#aaaaaa" />
        <circle cx="40" cy="60" r="3" fill="#aaaaaa" />
        <circle cx="70" cy="65" r="3" fill="#aaaaaa" />
        <circle cx="100" cy="45" r="3" fill="#aaaaaa" />
        <circle cx="130" cy="50" r="3" fill="#aaaaaa" />
        <circle cx="160" cy="30" r="3" fill="#aaaaaa" />
        <circle cx="190" cy="40" r="3" fill="#aaaaaa" />
      </svg>
    </div>
    );
  }

  if (isError) {
    return (
      <ErrorImage />
    );
  }

  return (
        <div className="balance-history-chart text-[#718EBF] rounded-3xl bg-white  w-full md:h-[300px] h-[250] pt-5 flex justify-center items-center">
            <canvas ref={chartRef} className=" w-full" />
        </div>
  );
}

export default BalanceHistoryChart;
