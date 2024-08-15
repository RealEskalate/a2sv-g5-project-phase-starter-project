"use client";
import { useRef, useEffect } from "react";
import { Chart } from "chart.js/auto";

export interface ChartRef extends HTMLCanvasElement {
  chart?: Chart;
}

function MyExpenseChart() {
  const chartRef = useRef<ChartRef>(null);

  useEffect(() => {
    const currentChartRef = chartRef.current;

    if (currentChartRef) {
      if (currentChartRef.chart) {
        currentChartRef.chart.destroy();
      }

      const context = currentChartRef.getContext("2d");

      if (context) {
        const barColors = Array(12).fill("#EDF0F7");
        barColors[4] = "#16DBCC";

        const newChart = new Chart(context, {
          type: "bar",
          data: {
            labels: ["Jan", "Feb", "Mar", "Ap", "May", "Jun"],
            datasets: [
              {
                label: "Expenses",
                data: [200, 150, 300, 250, 220, 180],
                backgroundColor: barColors,
                borderColor: "rgba(54, 162, 235, 1)",
                borderWidth: 0,
                borderRadius: 10,
                barThickness: 34,
                maxBarThickness: 40,
                categoryPercentage: 0.5,
                barPercentage: 0.8,
              },
            ],
          },
          options: {
            responsive: true,
            plugins: {
              legend: {
                display: false,
              },

              tooltip: {
                enabled: true,
                callbacks: {
                  label: function (context) {
                    return ` $${context.parsed.y}`;
                  },
                },
              },
            },
            scales: {
              x: {
                stacked: false,
                ticks: {
                  align: "end",
                  autoSkip: true,
                },
                grid: {
                  display: false,
                },
                border: {
                  display: false,
                },
              },
              y: {
                display: false,
                beginAtZero: true,
                grid: {
                  display: false,
                },
              },
            },
          },
        });

        currentChartRef.chart = newChart;
      }
    }
  }, []);

  return (
    <div className="text-gray-500 border rounded-[22px] bg-white lg:w-[370px] px-[10px] lg:h-[225px] md:w-[251px] md:h-[170px] w-fit h-[214px]">
      <div>
        <div className="mt-8 expense-chart lg:w-[350px] lg:h-[225px] md:w-[231px] md:h-[170px] w-[325px] h-[214px]">
          <canvas ref={chartRef} />
        </div>
      </div>
    </div>
  );
}

export default MyExpenseChart;
