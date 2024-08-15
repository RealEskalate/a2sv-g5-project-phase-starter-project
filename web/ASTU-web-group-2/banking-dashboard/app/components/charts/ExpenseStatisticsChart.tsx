'use client';

import React from 'react';
import {
  Chart as ChartJs,
  ArcElement,
  Tooltip,
  Legend,
  ChartOptions,
  ChartData,
} from 'chart.js';
import { Pie } from 'react-chartjs-2';
import ChartDataLabels from 'chartjs-plugin-datalabels';

ChartJs.register(ArcElement, Tooltip, Legend, ChartDataLabels);

const data: ChartData<'pie'> = {
  labels: ['Category A', 'Category B', 'Category C', 'Category D'],
  datasets: [
    {
      data: [30, 15, 35, 20],
      backgroundColor: [
        '#343C6A',
        '#FC7900',
        '#1814F3',
        '#FA00FF',
      ],
      borderWidth: 5,
      hoverBorderWidth: 2,
    },
  ],
};

const options: ChartOptions<'pie'> = {
  plugins: {
    datalabels: {
      color: '#fff',
      font: {
        weight: 'bold',
        size: 14,
      },
      formatter: (value: number, context: any) => {
        const total = context.chart._metasets[0].total;
        const percentage = ((value / total) *100).toFixed(2);
        return `   ${parseInt(percentage)}% 
${context.chart.data.labels[context.dataIndex]}`;
      },
    } ,
    tooltip: {
      callbacks: {
        label: function (tooltipItem) {
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
  },
  layout: {
    padding: 10,
  },
};

const ExpenseStatisticsChart: React.FC = () => {
  return (
    <div className='w-[350px] h-[350px] bg-white shadow-xl p-4 rounded-3xl'>
      <Pie data={data} options={options} />
    </div>
  );
};

export default ExpenseStatisticsChart;
