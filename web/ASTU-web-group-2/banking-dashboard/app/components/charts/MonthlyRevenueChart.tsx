"use client";
import { useRef, useEffect } from "react";
import { Chart, ChartData, ChartOptions } from "chart.js/auto";

interface CustomCanvasElement extends HTMLCanvasElement {
  chart?: Chart;
}

// Sample raw data with multiple entries per month
const rawData = [
  { month: "2016", balance: 200 },
  { month: "2016", balance: 0 },
  { month: "2016", balance: 200 },
  { month: "2017", balance: 0 },
  { month: "2017", balance: 200 },
  { month: "2017", balance: 0 },
  { month: "2018", balance: 1900 },
  { month: "2018", balance: 1800 },
  { month: "2019", balance: 1700 },
  { month: "2019", balance: 1600 },
  { month: "2019", balance: 2100 },
  { month: "2019", balance: 20 },
  { month: "2020", balance: 1800 },
  { month: "2020", balance: 1700 },
  { month: "2020", balance: 2400 },
  { month: "2021", balance: 2300 },
  { month: "2021", balance: 2200 },
  { month: "2021", balance: 0 },
];

// Function to aggregate data by month
const aggregateData = (data: { month: string; balance: number }[]) => {
  const monthMap: { [key: string]: { total: number; count: number } } = {};

  data.forEach((item) => {
    if (!monthMap[item.month]) {
      monthMap[item.month] = { total: 0, count: 0 };
    }
    monthMap[item.month].total += item.balance;
    monthMap[item.month].count += 1;
  });

  return Object.keys(monthMap).map((month) => ({
    month,
    balance: monthMap[month].total / monthMap[month].count,
  }));
};

function MonthlyRevenueChart() {
  const chartRef = useRef<CustomCanvasElement | null>(null);

  useEffect(() => {
    if (chartRef.current) {
      // Destroy the chart instance if it already exists
      if (chartRef.current.chart) {
        chartRef.current.chart.destroy();
      }

      const context = chartRef.current.getContext("2d");

      if (context) {
        const chartItem = context.canvas;

        // Aggregate the raw data
        const aggregatedData = aggregateData(rawData);

        // Extract labels and data for the chart
        const labels = aggregatedData.map((item) => item.month);
        const data = aggregatedData.map((item) => item.balance);

        const chartData: ChartData<"line"> = {
          labels: labels,
          datasets: [
            {
              label: "Average Balance",
              data: data,
              fill: false, // Remove background fill
              borderColor: "#16DBCC", // Change color of the line
              borderWidth: 4,
              tension: 0.4, // Set tension to curve the line
              pointRadius: 0, // Remove points
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
              grid: {
                display: false, 
              },
            },
            y: {
              beginAtZero: true,
              ticks: {
                callback: function (value) {
                    return `$${value}`;
                  },
                color: "rgb(113, 142, 191)",
                font: {
                  size: 12,
                },
              },
              grid: {
                color: "rgba(0, 0, 0, 0.1)",
             
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
  }, []);

  return (
    <div className="lg:w-[540px] lg:h-[276px] md:w-[423px] md:h-[232px] w-[325px] h-[254px]">
      <div className="rounded-3xl bg-white p-5 lg:w-[540px] lg:h-[276px] md:w-[423px] md:h-[200px] w-full h-[195px]">
        <div className="balance-history-chart text-[#718EBF] lg:w-[500px] lg:h-[177px] md:w-[347px] md:h-[147px] w-full h-[19px]">
          <div className="">
            <canvas ref={chartRef} />
          </div>
        </div>
      </div>
    </div>
  );
}

export default MonthlyRevenueChart;
