"use client";
import { useRef, useEffect } from "react";
import { Chart } from "chart.js/auto";
import { useGetAllTransactionQuery } from "@/lib/service/TransactionService"; 

export interface ChartRef extends HTMLCanvasElement {
  chart?: Chart;
}

function WeeklyActivityChart() {
  const chartRef = useRef<ChartRef>(null);
  const { data, isError, isLoading } = useGetAllTransactionQuery("acessToken");

  const processDataForChart = (transactions: any[]) => {
    const daysOfWeek = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
    const deposits = new Array(7).fill(5);
    const withdrawals = new Array(7).fill(5);

    transactions.forEach((transaction) => {
      const date = new Date(transaction.date);
      const dayIndex = date.getDay(); // Sunday is 0, Saturday is 6

      if (transaction.type === "deposit") {
        deposits[dayIndex] += transaction.amount+2;
      } else if (transaction.type === "shopping") { // Assuming "shopping" means withdrawal
        withdrawals[dayIndex] += transaction.amount;
      }
    });

    return { deposits, withdrawals, daysOfWeek };
  };

  useEffect(() => {
    const currentChartRef = chartRef.current;

    if (currentChartRef && data?.success) {
      if (currentChartRef.chart) {
        currentChartRef.chart.destroy();
      }

      const context = currentChartRef.getContext("2d");

      if (context) {
        const { deposits, withdrawals, daysOfWeek } = processDataForChart(data.data);

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
                maxBarThickness: 20,
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
                maxBarThickness: 20,
                categoryPercentage: 0.6,
                barPercentage: 0.7,
              },
            ],
          },
          options: {
            responsive: true,
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
  }, [data]);

  if (isLoading) {
    return <div className="text-gray-500 border rounded-[22px] bg-white lg:w-[730px] px-5 lg:h-[367px] md:w-[520px] md:h-[299px] w-fit h-[254]">Loading...</div>;
  }

  if (isError) {
    return <div className="text-gray-500 border rounded-[22px] bg-white lg:w-[730px] px-5 lg:h-[367px] md:w-[520px] md:h-[299px] w-fit h-[254]">Ooops! error loading your Activities.</div>;
  }

  return (
    <div className="">
      <div className="text-gray-500 border rounded-[22px] bg-white">
        <div className="flex flex-row justify-end lg:w-[650px] md:w-[467px]  w-[325px]">
          <div className="flex flex-row mx-5 mt-5">
            <div className="w-[12px] h-[12px]  mx-2 mt-[6px] border rounded-full bg-[#16DBCC]"></div>
            <div className="">Deposits</div>
          </div>
          <div className="flex flex-row mx-5 mt-5">
            <div className="w-[12px] h-[12px] mx-2 mt-[6px] border rounded-full bg-blue-700"></div>
            <div className="">Withdrawals</div>
          </div>
        </div>
        <div>
          <div className="weekly-activity-chart md:ml-5 lg:w-[667px] lg:h-[310px] md:w-[487px] md:h-[204px] w-[325px] h-[204px] ml-0">
            <canvas ref={chartRef} />
          </div>
        </div>
      </div>
    </div>
  );
}

export default WeeklyActivityChart;
