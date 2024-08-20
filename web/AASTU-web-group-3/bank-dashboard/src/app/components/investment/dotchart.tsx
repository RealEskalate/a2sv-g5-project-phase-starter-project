import React from 'react';
import { Line } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  LineElement,
  PointElement,
  LinearScale,
  Title,
  Tooltip,
  Legend,
  CategoryScale,
} from 'chart.js';

// Register necessary Chart.js components
ChartJS.register(
  LineElement,
  PointElement,
  LinearScale,
  CategoryScale,
  Title,
  Tooltip,
  Legend
);

const LineGraphWithDots = () => {
  // Dummy data for the graph
  const data = {
    labels: ['2016', '2017', '2018', '2019', '2020', '2021'],
    datasets: [
      {
        label: 'Dummy Data',
        data: [0, 40000, 20000, 30000, 10000, 50000], // Dummy dataset values
        borderColor: '#EDA10D', // Line color
        backgroundColor: 'transparent', // Ensure the line background is transparent
        borderWidth: 4,
        pointRadius: 6, // Dots size
        pointStyle: 'circle', // Circle dot shape
        pointBackgroundColor: '#fff', // Hollow middle (white to blend with chart background)
        pointBorderColor: '#EDA10D', // Dot border color
        pointBorderWidth: 4, // Bold dot border
      },
    ],
  };

  // Graph options and configurations
  const options = {
    responsive: true,
    plugins: {
      legend: {
        display: false, // Hide the dataset label
      },
      tooltip: {
        enabled: true, // Display tooltip on hover
      },
    },
    scales: {
      x: {
        display: true, // Show X-axis labels
        grid: {
          display: false, // Hide X-axis grid lines
        },
        ticks: {
          color: '#718EBF', // X-axis label color
          padding: 10, // Move X-axis labels down
        },
        border: {
          display: false, // Remove X-axis line
        },
      },
      y: {
        display: true, // Show Y-axis labels
        ticks: {
          color: '#718EBF', // Y-axis label color
          stepSize: 10000, // Control the number of ticks
          callback: (value:number) => `$${value}`, // Add $ symbol before Y-axis labels
          padding: 10, // Move Y-axis labels to the left
        },
        grid: {
          color: '#DFE5EE', // Custom horizontal gridline color
          drawOnChartArea: true, // Show gridlines within the chart area
          borderDash: [6, 6], // Dashed grid lines
        },
        border: {
          display: false, // Remove Y-axis line
        },
      },
    },
    elements: {
      point: {
        borderWidth: 4, // Bold border for hollow dots
        borderColor: '#EDA10D', // Dot border color
        backgroundColor: '#fff', // Make the inner part of the dots white (same as chart background)
      },
    },
  };

  return (
    <div className='w-full md:w-[45%]'>
      <h1 className="text-[22px] font-bold leading-[26.63px] text-[rgba(51,59,105,1)] text-left px-4 py-4">
        Yearly Total Investment
      </h1>
    <div className="flex justify-evenly h-[300px] bg-white rounded-2xl p-4 shadow-lg">
      <Line data={data}/>
    </div>
    </div>
  );
};

export default LineGraphWithDots;
