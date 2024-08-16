import React from 'react';
import { Bar } from 'react-chartjs-2';
import { Chart as ChartJS, BarElement, CategoryScale, LinearScale, Tooltip, Legend, TooltipItem } from 'chart.js';

ChartJS.register(BarElement, CategoryScale, LinearScale, Tooltip, Legend);

const BarChart: React.FC = () => {
  const data = {
    labels: ['Aug', 'Sep', 'Oct', 'Nov', 'Dec', 'Jan'],
    datasets: [
      {
        label: '',
        data: [5000, 8000, 6000, 3000, 12500, 7000],
        backgroundColor: [
          'rgba(0, 0, 0, 0.05)',
          'rgba(0, 0, 0, 0.05)',
          'rgba(0, 0, 0, 0.05)',
          'rgba(0, 0, 0, 0.05)',
          'rgba(0, 204, 204, 0.8)',
          'rgba(0, 0, 0, 0.05)',
        ],
        borderRadius: 10,
        borderSkipped: false,
        barPercentage: 0.6,
      },
    ],
  };

  const options = {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        display: false,
      },
      tooltip: {
        callbacks: {
          label: function (tooltipItem: TooltipItem<'bar'>) {
            const value = tooltipItem.raw as number;
            return `$${value.toLocaleString()}`;
          },
        },
      },
    },
    scales: {
      x: {
        grid: {
          display: false,
        },
        ticks: {
          color: 'rgba(0, 0, 0, 0.5)',
        },
      },
      y: {
        beginAtZero: true,
        display: false,
      },
    },
  };

  return (
    <div className="bg-white p-4 rounded-lg shadow-md w-full max-w-xs md:max-w-md lg:max-w-lg xl:max-w-xl h-64">
      <Bar data={data} options={options} />
      <div className="text-right text-gray-600 font-medium mt-2">$12,500</div>
    </div>
  );
};

export default BarChart;
