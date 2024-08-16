'use client';
import React from 'react';
import { Bar } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
  ChartOptions,
} from 'chart.js';

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

const WeeklyActivity = () => {
  // Define the chart data with type annotations
  const datavalues1 = [12, 19, 3, 5, 2, 3];
  const datavalues2 = [12, 8, 10, 5, 8, 4];
  const label = ['Red', 'Blue', 'Yellow', 'Green', 'Purple', 'Orange'];
  const data = {
    labels: label,
    datasets: [
      {
        label: 'Deposite',
        data: datavalues1,
        backgroundColor: 'rgba(75, 10, 192)',
        borderRadius: 6,
        barThickness: 30,
        barPercentage: 0.5,
      },
      {
        label: 'Withdraw',
        data: datavalues2,
        backgroundColor: '#16dbcc',
        borderRadius: 6,
        barThickness: 30,
        barPercentage: 0.5,
      },
    ],
  };

  const options: ChartOptions<'bar'> = {
    responsive: true,
    plugins: {
      legend: {
        position: 'top' as const,
        align: 'end' as const,
        labels: {
          boxWidth: 20,
          boxHeight: 20,
          usePointStyle: true,
          pointStyle: 'circle',
        },
      },
    },
    scales: {
      x: {
        stacked: false,
        ticks: {
          autoSkip: false,
          maxRotation: 0,
          padding: 10,
        },
      },
      y: {
        beginAtZero: true,
      },
    },
  };

  return (
    <div className='w-full md:w-8/12 md:me-6'>
      <h1 className='text-[#333B69] pb-3 font-semibold'>Weekly Activity</h1>
      <div className='bg-white rounded-3xl p-10'>
        <Bar data={data} options={options} height={350} width={1000} />
      </div>
    </div>
  );
};

export default WeeklyActivity;
