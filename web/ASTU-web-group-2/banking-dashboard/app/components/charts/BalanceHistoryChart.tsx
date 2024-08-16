"use client";
import { useRef, useEffect } from "react";
import { Chart, ChartData, ChartOptions } from "chart.js/auto";

interface CustomCanvasElement extends HTMLCanvasElement {
  chart?: Chart;
}

// Sample raw data with multiple entries per month
const rawData = [
  { month: "January", balance: 200 },
  { month: "January", balance: 0 },
  { month: "January", balance: 200 },
  { month: "January", balance: 0 },
  { month: "January", balance: 200 },
  { month: "January", balance: 0 },
  { month: "February", balance: 1900 },
  { month: "February", balance: 1800 },
  { month: "March", balance: 1700 },
  { month: "March", balance: 1600 },
  { month: "April", balance: 2100 },
  { month: "April", balance: 20 },
  { month: "May", balance: 1800 },
  { month: "May", balance: 1700 },
  { month: "June", balance: 2400 },
  { month: "June", balance: 2300 },
  { month: "July", balance: 2200 },
  { month: "July", balance: 0 },
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

function BalanceHistoryChart() {
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
              data: data,
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
              //add text color to axis labels
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
  }, []);

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

