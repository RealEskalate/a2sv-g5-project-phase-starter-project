"use client";
import React, { useRef, useEffect } from "react";
import { Chart } from "chart.js/auto";

const WeeklyActivity = () => {
  const chartRef = useRef<HTMLCanvasElement>(null);
  const chartInstanceRef = useRef<Chart | null>(null);
  useEffect(() => {
    if (chartRef.current && !chartInstanceRef.current) {
      const context = chartRef.current.getContext("2d");

      if (context) {
        chartInstanceRef.current = new Chart(context, {
          type: "bar",
          data: {
            labels: ["Sat", "Sun", "Mon", "Tue", "Wed", "Thu", "Fri"],
            datasets: [
              {
                label: "Deposit",
                data: [380, 300, 575, 200, 425, 350, 305],
                backgroundColor: "#1814F3",
                borderRadius: 100,
                borderSkipped: false,
                barPercentage: 0.6,
                categoryPercentage: 0.7,
              },
              {
                label: "Withdraw",
                data: [200, 260, 300, 400, 300, 200, 100],
                backgroundColor: "#16DBCC",
                borderRadius: 100,
                borderSkipped: false,
                barPercentage: 0.6,
                categoryPercentage: 0.7,
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
                top: 10,
                bottom: 10,
              },
            },
            plugins: {
              datalabels : {
                display : false,
              },
              legend: {
                display: true,
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
  }, []);
  return (
    // <div className=" flex flex-grow mobile:w-3/5 max-mobile:w-full max-mobile:h-52 mobile:h-80 bg-white rounded-3xl">
    <div className="w-full">
      <canvas className=" bg-white rounded-3xl w-full h-full " ref={chartRef} /></div>
      

   
      
    
  );
};

export default WeeklyActivity;
