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
  const datavalues1 = [12, 19, 3, 5, 2, 3];
  const datavalues2 = [12, 8, 10, 5, 8, 4];
  const label = ['Red', 'Blue', 'Yellow', 'Green', 'Purple', 'Orange'];
  const data = {
    labels: label,
    datasets: [
      {
        label: 'Deposit',
        data: datavalues1,
        backgroundColor: 'rgba(80, 10, 192, 0.9)',
        borderRadius: 26,
        barThickness: 23,
        barPercentage: 0,
        categoryPercentage: 0,
      },
      {
        label: 'Spacer',
        data: new Array(datavalues1.length).fill(null),
        backgroundColor: 'transparent',
        barThickness: 5,
        barPercentage: 0.001,
        categoryPercentage: 0.001,
      },
      {
        label: 'Withdraw',
        data: datavalues2,
        backgroundColor: '#16dbcc',
        borderRadius: 26,
        barThickness: 26,
        barPercentage: 0,
        categoryPercentage: 0,
      },
    ],
  };

  const options: ChartOptions<'bar'> = {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        position: 'top' as const,
        align: 'end' as const,
        labels: {
          boxWidth: 10,
          boxHeight: 23,
          usePointStyle: true,
          pointStyle: 'circle',
        },
      },
      tooltip: {
        enabled: true,
      },
      datalabels: {
        display: false,
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
        grid: {
          display: false,
        },
      },
      y: {
        beginAtZero: true,
        ticks: {
          stepSize: 5,
        },
      },
    },
  };

  return (
    <div className='w-full md:w-8/12 md:me-6'>
      <h1 className='text-[#333B69] pb-3 font-semibold'>Weekly Activity</h1>
      <div className='bg-white rounded-3xl md:px-5 md:py-6 p-2 w-full'>
        <div className='w-full' style={{ height: '270px', maxWidth: '100%' }}>
          <Bar data={data} options={options} />
        </div>
      </div>
    </div>
  );
};

export default WeeklyActivity;
