"use client";
import { useRef, useEffect } from "react";
import { Chart } from "chart.js/auto";


export interface ChartRef extends HTMLCanvasElement {
  chart?: Chart;
}

function WeeklyActivityChart() {
  const chartRef = useRef<ChartRef>(null);

  useEffect(() => {
    const currentChartRef = chartRef.current;

    if (currentChartRef) {
      if (currentChartRef.chart) {
        currentChartRef.chart.destroy();
      }

      const context = currentChartRef.getContext("2d");

      if (context) {
        const newChart = new Chart(context, {
          type: "bar",
          data: {
            labels: ["Sat", "Sun", "Mon", "Tue", "Wed", "Thu", "Fri"],
            datasets: [
              {
                label: "Deposits",
                data: [50, 25, 35, 45, 55, 43, 23], 
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
                data: [45, 20, 30, 40, 50, 48, 45], 
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
  }, []);

  return (
    <div className="bg-[#F5F7FA] lg:w-[530px] lg:h-[300px] md:w-[487px] md:h-[299px] h-[254]">
      <div>
        <p className="font-[600] my-3 text-[18px] md:text-[18px] lg:text[22px]">
          Weekly Activity
        </p>
      </div>
      <div className="">
        <div className="text-gray-500 border rounded-[22px] bg-white lg:w-[530px] px-5 lg:h-[300px] md:w-[487px] md:h-[299px] w-auto h-[254]">
          <div className="flex flex-row justify-end lg:w-[450px] md:w-[467px]  w-[325px]">
            <div className="flex flex-row mx-5 mt-5">
              <div className="w-[12px] h-[12px]  mx-2 mt-[6px] border rounded-full bg-[#16DBCC]"></div>
              <div className="">Deposite</div>
            </div>
            <div className="flex flex-row mx-5 mt-5">
              <div className="w-[12px] h-[12px] mx-2 mt-[6px] border rounded-full bg-blue-700"></div>
              <div className="">Withdraw</div>
            </div>
          </div>
          <div>
            <div className="weekly-activity-chart lg:w-[730px] lg:h-[226px] md:w-[487px] md:h-[204px] w-[325px] h-[204px]">
              <canvas ref={chartRef} />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default WeeklyActivityChart;
