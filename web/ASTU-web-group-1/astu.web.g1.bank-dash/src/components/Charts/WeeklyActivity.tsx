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
        borderRadius: 6,
        barThickness: 23,
        barPercentage: 0.5,
      },
      {
        label: 'Withdraw',
        data: datavalues2,
        backgroundColor: '#16dbcc',
        borderRadius: 6,
        barThickness: 23,
        barPercentage: 0.5,
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
          boxWidth: 23,
          boxHeight: 23,
          usePointStyle: true,
          pointStyle: 'circle',
        },
      },
      tooltip: {
        enabled: true, // Keep tooltips if needed
      },
      datalabels: {
        display: false, // Disable data labels if using chartjs-plugin-datalabels
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
          stepSize: 5, // Set the y-axis increments to 5
        },
      },
    },
  };

  return (
    <div className='w-full md:w-8/12 md:me-6'>
      <h1 className='text-[#333B69] pb-3 font-semibold'>Weekly Activity</h1>
      <div className='bg-white rounded-3xl md:px-5 md:py-6 p-2 w-full'>
        <div style={{ position: 'relative', height: '270px', width: '100%' }}>
          <Bar data={data} options={options} />
        </div>
      </div>
    </div>
  );
};

export default WeeklyActivity;
