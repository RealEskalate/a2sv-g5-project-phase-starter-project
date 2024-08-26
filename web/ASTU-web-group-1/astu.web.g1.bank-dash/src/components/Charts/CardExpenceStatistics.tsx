'use client';
import React from 'react';
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Doughnut } from 'react-chartjs-2';

// Register the necessary components with Chart.js
ChartJS.register(ArcElement, Tooltip, Legend);

const data = {
  plugins: {
    legend: {
      display: false,
    },
  },
  labels: ['Red', 'Blue', 'Yellow'],
  datasets: [
    {
      label: 'My First Dataset',
      data: [10, 10, 10, 10],
      backgroundColor: ['#4C78FF', '#16DBCC', '#FF82AC', '#FFBB38'],
      hoverOffset: 4,
    },
  ],
};

const options = {
  plugins: {
    legend: {
      display: false,
    },
  },
  responsive: true,
  maintainAspectRatio: false, // Allows the chart to be responsive
};

const CardExpenceStatistics: React.FC = () => {
  return (
    <div className='w-full'>
      <h1 className='text-[#333B69] pb-3 font-semibold'>Card and Expense Statistics</h1>
      <div className='bg-white p-4 rounded-3xl shadow-md'>
        <div className='flex justify-center mb-4'>
          <div className='w-full max-w-md h-[250px]'>
            <Doughnut data={data} options={options} />
          </div>
        </div>
        <div className='flex flex-row justify-around mt-4'>
          <div className='space-y-4 text-center md:text-left'>
            <h1 className='flex items-center justify-center md:justify-start'>
              <p className='w-4 h-4 rounded-full bg-[#4C78FF] me-2'></p> DBL Bank
            </h1>
            <h1 className='flex items-center justify-center md:justify-start'>
              <p className='w-4 h-4 rounded-full bg-[#16DBCC] me-2'></p> AMB Bank
            </h1>
          </div>
          <div className='space-y-4 text-center md:text-left'>
            <h1 className='flex items-center justify-center md:justify-start'>
              <p className='w-4 h-4 rounded-full bg-[#FF82AC] me-2'></p> RBC Bank
            </h1>
            <h1 className='flex items-center justify-center md:justify-start'>
              <p className='w-4 h-4 rounded-full bg-[#FFBB38] me-2'></p> MCP Bank
            </h1>
          </div>
        </div>
      </div>
    </div>
  );
};

export default CardExpenceStatistics;
