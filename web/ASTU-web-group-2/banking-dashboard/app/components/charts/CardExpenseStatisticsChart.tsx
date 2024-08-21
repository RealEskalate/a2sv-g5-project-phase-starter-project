"use client";
import { useRef, useEffect } from "react";
import {
  Chart,
  DoughnutController,
  ArcElement,
  Tooltip,
  Legend,
  ChartData,
  ChartOptions,
} from "chart.js";

Chart.register(DoughnutController, ArcElement, Tooltip, Legend);

const customRadiusPlugin = {
  id: "customRadius",
  beforeDraw(chart: {
    getDatasetMeta?: any;
    ctx?: any;
    chartArea?: any;
    data?: any;
  }) {
    const {
      ctx,
      chartArea: { left, top, right, bottom },
      data,
    } = chart;
    const totalWidth = right - left;
    const totalHeight = bottom - top;
    const radiusModifier = [1, 0.7, 0.8, 0.9];

    data.datasets.forEach((dataset: any, datasetIndex: any) => {
      const meta = chart.getDatasetMeta(datasetIndex);
      meta.data.forEach(
        (arc: { outerRadius: number }, arcIndex: number) => {
          const outerRadius = Math.min(totalWidth, totalHeight) / 2;
          arc.outerRadius = outerRadius * radiusModifier[arcIndex];
        }
      );
    });
  },
};

function CardExpenseStatisticsChart() {
  const chartRef = useRef<HTMLCanvasElement | null>(null);
  const chartInstanceRef = useRef<Chart<"doughnut"> | null>(null);

  useEffect(() => {
    if (chartRef.current) {
      const context = chartRef.current.getContext("2d");

      if (context) {
        if (chartInstanceRef.current) {
          chartInstanceRef.current.destroy();
        }

        const chartData: ChartData<"doughnut", number[], string> = {
          labels: ["DBL Bank", "ABM Bank", "BRC Bank", "MCP Bank"],
          datasets: [
            {
              data: [30, 20, 25, 25],
              backgroundColor: [
                "rgb(22, 219, 204)",
                "rgb(255, 130, 172)",
                "rgba(255, 187, 56, 1)",
                "rgb(76, 120, 255)",
              ],
              borderWidth: 0,
            },
            {
              data: [30, 20, 25, 25],
              backgroundColor: [
                "rgb(30, 198, 184)",
                "rgb(255, 97, 149)",
                "rgb(255, 177, 31)",
                "rgb(52, 100, 243)",
              ],
              borderWidth: 0,
            },
          ],
        };

        const options: ChartOptions<"doughnut"> = {
          responsive: true,
          maintainAspectRatio: false,
          plugins: {
            legend: {
              display: false,
            },
          },
          cutout: "30%",
        };

        const combinedChart = new Chart(context, {
          type: "doughnut",
          data: chartData,
          options: options,
          plugins: [customRadiusPlugin],
        });

        chartInstanceRef.current = combinedChart;
      }
    }

    return () => {
      if (chartInstanceRef.current) {
        chartInstanceRef.current.destroy();
      }
    };
  }, []);

  return (
    <div className="grid grid-cols-1 bg-white rounded-[28px] text-[#718EBF] h-[310px]">
      <div className="pt-6 pb-1 h-48">
        <canvas ref={chartRef}/>
      </div>
      <div className="grid grid-cols-[1fr_1fr] mx-4 pl-5 md:pl-1 md:mx-1 md:text-[15px] lg:mx-4 lg:pl-5">
        <div className="grid grid-cols-[1fr_6fr] m-1">
          <div className="w-4 h-4 bg-[#4C78FF] rounded-full mx-1 mt-1"></div>
          <div>
            <p>DBL Bank</p>
          </div>
        </div>
        <div className="grid grid-cols-[1fr_6fr] m-1 ">
          <div className="w-4 h-4 bg-red-400 rounded-full mx-1 mt-1"></div>
          <div>
            <p>BRC Bank</p>
          </div>
        </div>
        <div className="grid grid-cols-[1fr_6fr] m-1">
          <div className="w-4 h-4 bg-[#16DBCC] rounded-full mx-1 mt-1"></div>
          <div>
            <p>ABM Bank</p>
          </div>
        </div>
        <div className="grid grid-cols-[1fr_6fr] m-1">
          <div className="w-4 h-4 bg-[#FFBB38] rounded-full mx-1 mt-1"></div>
          <div>
            <p>MCP Bank</p>
          </div>
        </div>
      </div>
    </div>
  );
}

export default CardExpenseStatisticsChart;
