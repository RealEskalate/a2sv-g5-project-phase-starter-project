"use client";
import React, { useRef, useEffect } from "react";

import {
  Chart,
  LineController,
  ArcElement,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Filler,
  Tooltip,
  Legend,
  ScriptableContext,
} from "chart.js";

interface GRAPHDATA {
  balanceHistory: number[];
}

const BalanceHistoryLineGraph = ({ balanceHistory }: GRAPHDATA) => {
  const chartRef = useRef<HTMLCanvasElement>(null);
  const chartInstanceRef = useRef<Chart | null>(null);
  useEffect(() => {
    if (chartRef.current && !chartInstanceRef.current) {
      const context = chartRef.current.getContext("2d");

      if (context) {
        chartInstanceRef.current = new Chart(context, {
          type: "line",
          data: {
            labels: ["Sat", "Sun", "Mon", "Tue", "Wed", "Thu", "Fri"],
            datasets: [
              {
                label: "",
                data: balanceHistory, // Sample data
                borderWidth: 2,
                pointBackgroundColor: "#1D8CF8",
                pointBorderColor: "#ffffff",
                pointBorderWidth: 2,
                pointRadius: 0, // Remove the points
                tension: 0.45, // Smooth curve
                

                backgroundColor: (context: ScriptableContext<"line">) => {
                  const ctx = context.chart.ctx;
                  const gradient = ctx.createLinearGradient(0, 0, 0, 250);
                  gradient.addColorStop(0, "rgba(91,56,237,0.45)");
                  gradient.addColorStop(1, "rgba(91,56,237,0.0)");
                  return gradient;
                }, //background gradient color
                borderColor: "#1814F3", // Line color
                fill: true,
              },
            ],
          },
          options: {
            responsive: true,
            maintainAspectRatio: false,

            scales: {
              x: {
                type: "category",
                grid: {
                  display: false,
                },
                border: {
                  display: false,
                },
              },
              y: {
                beginAtZero: true,
                grid: {
                  lineWidth: 0.4,
                },
                border: {
                  display: false,
                },
              },
            },
            layout: {
              padding: {
                left: 20,
                right: 20,
                top: 30,
                bottom: 20,
              },
            },
            plugins: {
              datalabels: {
                display: false,
              },
              legend: {
                display: false,
                position: "top",
                align: "end",

                labels: {
                  usePointStyle: true,
                  pointStyle: "circle",
                  color: "#718EBF",
                  font: {
                    size: 12, // Font size
                  },
                  padding: 20,
                  boxHeight: 12,
                },
              },
            },
          },
        });
      }

      return () => {
        if (chartInstanceRef.current) {
          chartInstanceRef.current.destroy();
          chartInstanceRef.current = null;
        }
      };
    }
  }, [balanceHistory]);
  return (
    // <div className=" flex flex-grow mobile:w-3/5 max-mobile:w-full max-mobile:h-52 mobile:h-80 bg-white rounded-3xl">
    <div className="max-mobile:h-[223px] max-mobile:w-[325px] mobile:w-[423px] mobile:h-[220px] tablet:h-[276px] tablet:w-[635px]">
      <canvas className=" bg-white rounded-3xl w-full h-full " ref={chartRef} />
    </div>
  );
};

export default BalanceHistoryLineGraph;
