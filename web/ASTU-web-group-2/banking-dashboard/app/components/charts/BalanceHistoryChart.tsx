"use client";
import { useGetBalanceHistoryQuery } from "@/lib/service/TransactionService";
import { useRef, useEffect } from "react";
import { Chart, ChartData, ChartOptions } from "chart.js/auto";

const generateMonths = () => {
  const months = [];
  const currentDate = new Date();
  
  for (let i = 6; i >= 0; i--) {
    const date = new Date(currentDate.getFullYear(), currentDate.getMonth() - i, 1);
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
  const { data, isError, isLoading } = useGetBalanceHistoryQuery("eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJzYW1pdGVzdCIsImlhdCI6MTcyNDA1NDU0MiwiZXhwIjoxNzI0MTQwOTQyfQ.Mg31Fn623uFEgDrc8gkAB8u7GgkZqA3U4fTfznz107LTXc8jD6t5G8q-bd1jOWxT");
  const chartRef = useRef<ChartRef | null>(null);

  useEffect(() => {
    if (chartRef.current && data?.success) {
      // Destroy the chart instance if it already exists
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
        const gradient = context.createLinearGradient(0, 0, 0, chartItem.height);
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
    return <div className="rounded-3xl bg-white p-5 lg:w-[635px] lg:h-[276px] md:w-[423px] md:h-[200px] w-[325px] h-[223px]">Loading...</div>;
  }

  if (isError) {
    return <div className="rounded-3xl bg-white p-5 lg:w-[635px] lg:h-[276px] md:w-[423px] md:h-[200px] w-[325px] h-[223px]">Ooops! Error loading balance history.</div>;
  }

  return (
    <div className="lg:w-[635px] lg:h-[276px] md:w-[423px] md:h-[232px] w-[325px] h-[254px]">
      <div className="rounded-3xl bg-white p-5 lg:w-[635px] lg:h-[276px] md:w-[423px] md:h-[200px] w-[325px] h-[223px]">
        <div className="balance-history-chart text-[#718EBF] lg:w-[500px] lg:h-[177px] md:w-[347px] md:h-[147px] w-[289px] h-[190px]">
          <div className="">
            <canvas ref={chartRef} />
          </div>
        </div>
      </div>
    </div>
  );
}

export default BalanceHistoryChart;
