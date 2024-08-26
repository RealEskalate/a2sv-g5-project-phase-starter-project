"use client";
import { useRef, useEffect } from "react";
import { Chart, ChartData, ChartOptions } from "chart.js/auto";
import LineChartSkeleton from "./LineChartSkeleton";
import { useSession } from "next-auth/react";
import { useGetMonthInvestmentHistoryQuery } from "@/lib/service/TransactionService";
import ErrorImage from "../Error/ErrorImage";

interface CustomCanvasElement extends HTMLCanvasElement {
  chart?: Chart;
}

function MonthlyRevenueChart() {
  const chartRef = useRef<CustomCanvasElement | null>(null);
  const { data: session } = useSession();
  const accessToken = session?.user.accessToken!;
  const { data, isError, isLoading } = useGetMonthInvestmentHistoryQuery(accessToken);

  useEffect(() => {
    if (chartRef.current) {
      if (chartRef.current.chart) {
        chartRef.current.chart.destroy();
      }

      const context = chartRef.current.getContext("2d");

      if (context) {
        const chartItem = context.canvas;

        // Sort the data from oldest to newest
        const sortedMonthlyRevenue = data?.data.monthlyRevenue.slice().sort((a:any, b: any) => 
          new Date(a.time.split('/').reverse().join('-')).getTime() - new Date(b.time.split('/').reverse().join('-')).getTime()
        );

        const labels = sortedMonthlyRevenue?.map((item: { time: string }) => item.time);
        const values = sortedMonthlyRevenue?.map((item: { value: number }) => item.value);

        const chartData: ChartData<"line"> = {
          labels: labels,
          datasets: [
            {
              label: "Monthly Revenue",
              data: values,
              fill: false,
              borderColor: "#16DBCC",
              borderWidth: 4,
              tension: 0.4,
              pointRadius: 0,
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
  }, [data]); // Add data as a dependency

  if (isLoading){
    return(
      <LineChartSkeleton />
    );
  } 

  if (isError) {
    return  <div className="text-gray-500 border rounded-[22px] bg-gray-200 p-5 w-full h-auto animate-pulse"><ErrorImage /></div>;
  }

  return (
    <div className="text-gray-500 border rounded-[22px] bg-white p-5">
      <canvas ref={chartRef} />
    </div>
  );
}

export default MonthlyRevenueChart;
