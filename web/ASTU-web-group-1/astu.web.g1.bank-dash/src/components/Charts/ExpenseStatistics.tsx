'use client';
import React from 'react';

import { PolarArea } from 'react-chartjs-2';

import {
  Chart as ChartJS,
  RadialLinearScale,
  ArcElement,
  Tooltip,
  Legend,
  CategoryScale,
  LinearScale,
  ChartOptions,
} from 'chart.js';

import ChartDataLabels from 'chartjs-plugin-datalabels';

ChartJS.register(
  Tooltip,
  Legend,
  ChartDataLabels,
  ArcElement,
  RadialLinearScale,
  CategoryScale,
  LinearScale,
);

const vals = [18, 16, 18, 24];
const labels = ['Transfer', 'Services', 'Otheres', 'Shopping'];
let index = 0;

const roundToDecimal = (num: number): number => {
  const factor = Math.pow(10, 2);
  return Math.round(num * factor) / factor;
};

let valsInPercentage = vals.map((val) => (val * 100) / vals.reduce((a, b) => a + b, 0));
valsInPercentage = valsInPercentage.map((val) => roundToDecimal(val));

const data = {
  labels: labels,
  datasets: [
    {
      label: 'Dataset 1',
      data: valsInPercentage,
      backgroundColor: ['#343C6A', '#FC7900', '#1814F3', '#FA00FF'],
      borderColor: ['white', 'white', 'white', 'white'],
      borderWidth: 5,
    },
  ],
};

const options: ChartOptions<'polarArea'> = {
  responsive: true,
  plugins: {
    legend: {
      display: false,
    },
    tooltip: {
      callbacks: {
        label: (context) => {
          const label = context.label || '';
          const value = context.raw || 0;
          return `${label}: ${value}`;
        },
      },
    },
    datalabels: {
      display: true,
      formatter: (value) => {
        // return `${labels[index++]} ${value}%`;
        return `${value}%`;
      },
      color: 'white',
      font: {
        size: 12,
        weight: 'bold',
      },
    },
  },

  scales: {
    r: {
      grid: {
        display: false,
      },
      ticks: {
        display: false,
      },
    },
  },
};

const ExpenseStatistics: React.FC = () => {
  return (
    <div className='w-full md:w-4/12'>
      <h1 className='text-[#333B69] pb-3 font-semibold'>Expense Statistics</h1>
      <div className='bg-white rounded-3xl flex justify-center items-center'>
        <PolarArea data={data} options={options} />
      </div>
    </div>
  );
};

export default ExpenseStatistics;
