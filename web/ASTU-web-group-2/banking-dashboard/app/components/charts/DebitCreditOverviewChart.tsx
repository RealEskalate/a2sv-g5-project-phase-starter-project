"use client";
import { useRef, useEffect } from "react";
import { Chart } from "chart.js/auto";

export interface ChartRef extends HTMLCanvasElement {
  chart?: Chart;
}

function DebitCreditOverviewChart() {
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
                label: "Debit",
                data: [50, 25, 35, 45, 55, 43, 23],
                backgroundColor: "#1A16F3",
                borderColor: "rgba(54, 162, 235, 0)",
                borderWidth: 3,
                borderRadius: 10,
              
                barThickness: 25,
                categoryPercentage: 0.8, // Adjusts space around groups of bars
                barPercentage: 0.9, // A
             
              },
              
              {
                label: "Credit",
                data: [45, 20, 30, 40, 50, 48, 45],
                backgroundColor: "#FCAA0B",
                borderColor: "rgba(255, 99, 132, 0)",
                borderWidth: 3,
                // borderRadius: { topLeft: 5, topRight: 5, bottomLeft: 5, bottomRight: 5 },
               
                borderRadius: 10,
              
                barThickness: 25,
                categoryPercentage: 0.8, // Adjusts space around groups of bars
                barPercentage: 0.9,  // A
                // maxBarThickness: 20,
              
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
                border: {
                  display: false, // Hide the axis line
                },
              },
              y: {
                display : false,
                beginAtZero: true,
                grid : {
                  display : false,
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
      <div className="lg:w-[730px] lg:h-[411px] md:w-[487px] md:h-[299px] h-[254]">
        <div className="text-gray-500 rounded-[22px] bg-white lg:w-[730px] lg:h-[411px] md:w-[487px] md:h-[300px] h-[234] w-[350px]">
          <div className="flex flex-row justify-between lg:w-[680px] md:w-[400px]   w-[325px]">
            <div className="flex flex-col mx-5 mt-5">
              <div className="text-black text-[9px] md:text-[12px] lg:text-[14px] font-light">
                $7,560 Debited & $5,420 Credited in this Week
              </div>
            </div>
            <div className="flex flex-row mx-2 mt-5">
              <div className="w-[12px] h-[12px]  mx-2 mt-[6px] border rounded-full bg-[#1A16F3]"></div>
              <div className="">Debit</div>
            </div>
            <div className="flex flex-row mx-2 mt-5">
              <div className="w-[12px] h-[12px] mx-2 mt-[6px] border rounded-full bg-[#FCAA0B]"></div>
              <div className="">Credit</div>
            </div>
          </div>
          <div>
            <div className="weekly-activity-chart lg:w-[670px]  lg:h-[309px] md:w-[417px] md:h-[240px] w-[325px] h-[204px] lg:mx-[30px] md:mx-[35px] mx-[12px]">
              <canvas ref={chartRef} />
            </div>
          </div>
        </div>
      </div>
  
  );
}

export default DebitCreditOverviewChart;
