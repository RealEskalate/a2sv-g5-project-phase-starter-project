'use client';

import React, { useRef, useEffect } from 'react';
import { Chart, ArcElement, Tooltip, Legend } from 'chart.js';
import ChartDataLabels from 'chartjs-plugin-datalabels';
import PieChartSkeleton from './PieChartSkeleton';

Chart.register(ArcElement, Tooltip, Legend);

const data = {
  labels: ['Category A', 'Category B', 'Category C', 'Category D'],
  datasets: [
    {
      data: [25, 25, 25, 25],
      backgroundColor: ['#343C6A', '#FC7900', '#1814F3', '#FA00FF'],
      borderWidth: 5,
      hoverBorderWidth: 2,
    },
  ],
};

const options = {
  plugins: {
    tooltip: {
      callbacks: {
        label: function (tooltipItem: { dataIndex: any; }) {
          const dataIndex = tooltipItem.dataIndex;
          const label = data.labels ? data.labels[dataIndex] : '';
          const value = data.datasets[0].data[dataIndex];
          const total = data.datasets[0].data.reduce(
            (acc, curr) => acc + (typeof curr === 'number' ? curr : 0),
            0
          );
          const percentage = ((value / total) * 100).toFixed(2);
          return `${label} ${percentage}% (${value})`;
        },
      },
    },
    legend: {
      display: false,
    },
    datalabels: {
      color: '#fff',
      display: true,
      formatter: (value: any, context: { chart: { data: { labels: { [x: string]: any; }; }; }; dataIndex: string | number; }) => {
        const label = context.chart.data.labels[context.dataIndex];
        return `${label}: ${value}`;
      },
    },
  },
  layout: {
    padding: 10,
  },
  elements: {
    arc: {
      borderRadius: (context: { dataIndex: number; }) => {
        // Define different radii for different slices
        const radii = [5, 10, 15, 20];
        return radii[context.dataIndex] || 0;
      },
    },
  },
};

const ExpenseStatisticsChart: React.FC = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null);

  useEffect(() => {
    if (canvasRef.current) {
      const ctx = canvasRef.current.getContext('2d');
      if (ctx) {
        const newChart = new Chart(ctx, {
          type: 'pie',
          data,
          options: {
            ...options,
            plugins: {
              ...options.plugins,
              datalabels: {
                ...options.plugins.datalabels,
                formatter: (value: any, context: any) => {
                  const label = context.chart.data.labels[context.dataIndex];
                  return `      ${value}\n${label}`;
                },
              },
            },
          },
          plugins: [ChartDataLabels],
        });

        return () => {
          newChart.destroy();
        };
      }
    }
  }, []);

  return (
    <div className='bg-white rounded-3xl lg:h-[322px] h-[261px] flex justify-center items-center'>
      <canvas ref={canvasRef} />
    </div>
  );
};

export default ExpenseStatisticsChart;
