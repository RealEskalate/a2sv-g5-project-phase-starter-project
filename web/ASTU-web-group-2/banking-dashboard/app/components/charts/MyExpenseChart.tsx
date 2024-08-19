"use client";
import { useGetExpensesQuery } from "@/lib/service/TransactionService";
import { useRef, useEffect } from "react";
import { Chart } from "chart.js/auto";

export interface ChartRef extends HTMLCanvasElement {
  chart?: Chart;
}

function MyExpenseChart() {
  const chartRef = useRef<ChartRef>(null);
  const access = process.env.ACCESS_TOKEN;
  const { data, isError, isLoading } = useGetExpensesQuery(
    "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJzYW1pdGVzdCIsImlhdCI6MTcyMzkwMzIwNywiZXhwIjoxNzIzOTg5NjA3fQ.9oD43mrMrIADSyodH11ija2vLg3U8LVMsEPd7kvSZaVAHbpgItEr02YPkcSFKhCK"
  );

  useEffect(() => {
    if (isLoading || isError) return;

    const currentChartRef = chartRef.current;

    if (currentChartRef && data?.success) {
      if (currentChartRef.chart) {
        currentChartRef.chart.destroy();
      }

      const context = currentChartRef.getContext("2d");

      if (context) {
        // Group expenses by month and sum up the amounts
        const expensesByMonth: { [key: string]: number } = {};

        data.data.forEach((expense: { date: string; amount: number; }) => {
          const month = expense.date.substring(0, 7); // e.g., "2024-08"
          if (!expensesByMonth[month]) {
            expensesByMonth[month] = 0;
          }
          expensesByMonth[month] += expense.amount;
        });

        // Get the last 6 months (including the current month)
        const now = new Date();
        const lastSixMonths = Array.from({ length: 6 }, (_, i) => {
          const date = new Date(now.getFullYear(), now.getMonth() - i, 1);
          return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, "0")}`;
        }).reverse();

        // Extract labels and summed amounts from the grouped data
        const monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];
        const labels = lastSixMonths.map((month) => {
          const [year, monthNumber] = month.split("-");
          return monthNames[parseInt(monthNumber, 10) - 1];
        });
        const amounts = lastSixMonths.map((month) => expensesByMonth[month] || 0);
        const custommonth = [5,3,1,5,6,7]
        const customData = custommonth.concat(amounts)

        // Get the current month
        const currentMonth = now.getMonth(); // 0-based index (Jan = 0)

        // Set bar colors and highlight the current month
        const barColors = Array(labels.length).fill("#EDF0F7");
        const currentMonthIndex = labels.findIndex(label => monthNames[currentMonth] === label);
        if (currentMonthIndex !== -1) {
          barColors[currentMonthIndex] = "#16DBCC";
        }

        const newChart = new Chart(context, {
          type: "bar",
          data: {
            labels: labels,
            datasets: [
              {
                label: "Expenses",
                data: customData,
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
  }, [data, isLoading, isError]);

  if (isLoading) {
    return (
      <div className="flex justify-center items-center h-[225px] w-[350px] bg-white drop-shadow-xl font-medium rounded-[25px]">
        <p>Loading chart data...</p>
      </div>
    );
  }

  if (isError) {
    return (
      <div className="flex justify-center items-center h-[225px] w-[350px] bg-white drop-shadow-xl font-medium rounded-[25px]">
        <p>Failed to load data. Please try again later.</p>
      </div>
    );
  }

  return (
    <div className="flex flex-col flex-initial flex-wrap gap-[10px] bg-white drop-shadow-xl font-medium rounded-[25px] h-[225px] pt-[45px] w-[350px]">
      <canvas ref={chartRef} />
    </div>
  );
}

export default MyExpenseChart;
