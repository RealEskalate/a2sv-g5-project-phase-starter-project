"use client";
import React, { useRef, useEffect } from "react";
import { Chart } from "chart.js/auto";

const MyExpenseGraph = () => {
  const chartRef = useRef<HTMLCanvasElement>(null);
  const chartInstanceRef = useRef<Chart | null>(null);
  useEffect(() => {
    if (chartRef.current && !chartInstanceRef.current) {
      const context = chartRef.current.getContext("2d");

      if (context) {
        chartInstanceRef.current = new Chart(context, {
          type: "bar",
          data: {
            labels: [ "Mar", "Apr", "May", "Jun", "July", "Aug"],
            datasets: [
              {
                label: "Deposit",
                data: [380, 300, 575, 200, 425, 350, 305],
                backgroundColor: "#EDF0F7",
                hoverBackgroundColor: "#16DBCC",
                borderRadius: 10,
                borderSkipped: false,
                barPercentage: 0.8,
                categoryPercentage: 0.8,
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
    <div className="max-mobile:h-[214px] max-mobile:w-[325px] mobile:w-[231px] mobile:h-[170px] tablet:h-[225px] tablet:w-[350px]">
      <canvas className=" bg-white rounded-3xl w-full h-full " ref={chartRef} /></div>
      

   
      
    
  );
};

export default MyExpenseGraph;
