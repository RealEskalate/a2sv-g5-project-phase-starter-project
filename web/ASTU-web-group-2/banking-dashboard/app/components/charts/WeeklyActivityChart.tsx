"use client";
import { useRef, useEffect } from "react";
import { Chart } from "chart.js/auto";
import { useGetAllTransactionQuery } from "@/lib/service/TransactionService";
import { useSession } from "next-auth/react";
import BarChartSkeleton from "./BarChartSkeleton";
import ErrorImage from "../Error/ErrorImage";

export interface ChartRef extends HTMLCanvasElement {
  chart?: Chart;
}

function WeeklyActivityChart() {
  const chartRef = useRef<ChartRef>(null);
  const { data: session, status } = useSession();
  const accessToken = session?.user.accessToken!;
  const { data: res, isError, isLoading } = useGetAllTransactionQuery(accessToken);

  const processDataForChart = (transactions: any[]) => {
    const daysOfWeek = [];
    let deposits = new Array(7).fill(0);
    let withdrawals = new Array(7).fill(0);

    const today = new Date();
    const todayIndex = today.getDay();
    
    // Get the date 7 days ago
    const sevenDaysAgo = new Date();
    sevenDaysAgo.setDate(today.getDate() - 7);

    for (let i = 0; i < 7; i++) {
        const dayIndex = (todayIndex - i + 7) % 7;
        const day = new Date(today);
        day.setDate(today.getDate() - i);
        const dayLabel = day.toLocaleString('en-US', { weekday: 'short' });
        daysOfWeek.unshift(dayLabel);
    }

    transactions.forEach((transaction) => {
        const date = new Date(transaction.date);

        if (date >= sevenDaysAgo && date <= today) {
            const dayIndex = (todayIndex - date.getDay() + 7) % 7;

            if (transaction.type.toLowerCase() === "deposit") {
                deposits[dayIndex] += transaction.amount;
            } else {
                withdrawals[dayIndex] += transaction.amount;
            }
        }
    });
    deposits = deposits.filter(deposit => deposit <= 1000)
    withdrawals = withdrawals.filter(withdrawal => withdrawal <= 1000)
    

    return { deposits, withdrawals, daysOfWeek };
};




  useEffect(() => {
    const currentChartRef = chartRef.current;

    if (currentChartRef && res?.success) {
      if (currentChartRef.chart) {
        currentChartRef.chart.destroy();
      }

      const context = currentChartRef.getContext("2d");

      if (context) {
        const { deposits, withdrawals, daysOfWeek } = processDataForChart(
          res.data.content
        );

        const newChart = new Chart(context, {
          type: "bar",
          data: {
            labels: daysOfWeek,
            datasets: [
              {
                label: "Deposits",
                data: deposits,
                backgroundColor: "rgb(24, 20, 243)",
                borderColor: "rgba(54, 162, 235, 1)",
                borderWidth: 0,
                borderRadius: 50,
                barThickness: 10,
                maxBarThickness: 10,
                categoryPercentage: 0.6,
                barPercentage: 0.7,
              },
              {
                data: [null, null, null, null, null],
                backgroundColor: "rgba(0, 0, 0, 0)",
                borderColor: "rgba(0, 0, 0, 0)",
                borderWidth: 0,
                barThickness: 0,
                maxBarThickness: 0,
                categoryPercentage: 0.1,
                barPercentage: 0.1,
              },
              {
                label: "Withdrawals",
                data: withdrawals,
                backgroundColor: "rgb(22, 219, 204)",
                borderColor: "rgba(255, 99, 132, 1)",
                borderWidth: 0,
                borderRadius: 50,
                barThickness: 10,
                maxBarThickness: 10,
                categoryPercentage: 0.6,
                barPercentage: 0.7,
              },
            ],
          },
          options: {
            responsive: true,
            maintainAspectRatio: false,
            plugins: {
              legend: {
                display: false,
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
              },
              y: {
                beginAtZero: true,
              },
            },
          },
        });

        currentChartRef.chart = newChart;
      }
    }
  }, [res]);

  if (isLoading) {
    return (
      <BarChartSkeleton />
    );
  }

  if (isError) {
    return (
     <ErrorImage />
    );
  }

  return (
    <div className="bg-white rounded-[22px] lg:h-[322px] h-[261px] ">
      <div className="flex flex-row justify-end gap-2">
        <div className="flex flex-row mx-5 mt-5 gap-1">
          <div className="w-[12px] h-[12px] mt-[6px] border rounded-full bg-[#16DBCC]"></div>
          <div className="">Deposit</div>
        </div>
        <div className="flex flex-row mx-5 mt-5 gap-1">
          <div className="w-[12px] h-[12px] mt-[6px] border rounded-full bg-blue-700"></div>
          <div className="">Withdraw</div>
        </div>
      </div>
      <div className="h-[75%] mx-5 mb-5">
        <canvas ref={chartRef} className="" /> 
      </div>
    </div>
  );
}

export default WeeklyActivityChart;
