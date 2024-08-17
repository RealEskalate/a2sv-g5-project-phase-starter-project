"use client";
import { useRef, useEffect } from "react";
import { Chart, ChartData, ChartOptions } from "chart.js/auto";

interface CustomCanvasElement extends HTMLCanvasElement {
  chart?: Chart;
}

function YearlyTotalInvestment() {
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
        const data: ChartData<"line"> = {
          labels: ["2016", "2017", "2018", "2019", "2020", "2021"],
          datasets: [
            {
              label: "Balance",
              data: [5000, 22000, 15000, 35000, 20000, 29000],
              fill: true, // Fill area under the line
              backgroundColor: "rgba(252, 170, 11,0)", // Color for the filled area
              borderColor: "rgba(252, 170, 11,1)", // Color of the line
              borderWidth: 4,
              pointRadius: 5,
              pointBackgroundColor: "rgba(255, 255, 255, 1)",
              pointBorderColor: "rgba(252, 170, 11,1)",
            },
          ],
        };

        const options: ChartOptions<"line"> = {
          responsive: true,
          plugins: {
            legend: {
              display: false,
              labels: {
                color: "rgb(113, 142, 191)", // Change legend label color
              },
            },
          },
          scales: {
            x: {
              stacked: false,
              ticks: {
                align: "end",
                autoSkip: true,
                color: "rgb(113, 142, 191)", // Change x-axis tick label color
              },
              grid: {
                display: false,
                tickBorderDash: [1, 1],
              },
              border: {
                display: false,
              },
            },
            y: {
              beginAtZero: true,
              ticks: {
                callback: function (value) {
                  return `$${value}`;
                },
                color: "rgb(113, 142, 191)", // Change y-axis tick label color
              },
              border: {
                display: false,
              },
            },
          },
        };

        const newChart = new Chart(chartItem, {
          type: "line",
          data,
          options,
        });

        chartRef.current.chart = newChart;
      }
    }
  }, []);

  return (
    <div className="text-gray-500 border rounded-[22px] bg-white p-2  lg:w-[540px] lg:h-[282px] md:w-[359px] md:h-[226px] w-[325px] h-[225px]">
      <div>
        <div className="mt-8 expense-chart lg:mx-[20px] lg:w-[481px] lg:h-[228px] md:w-[321px] md:h-[190px] w-[283px] h-[157px]">
          <canvas ref={chartRef} />
        </div>
      </div>
    </div>
  );
}

export default YearlyTotalInvestment;
