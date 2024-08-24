import React, { useState, useEffect } from 'react';
import { Bar } from 'react-chartjs-2';
import { Chart as ChartJS, BarElement, CategoryScale, LinearScale, Tooltip, Legend, TooltipItem } from 'chart.js';
import { getBalanceHistory } from '@/lib/api/transactionController'; // Import the function

interface BalanceHistoryData {
  time: string;
  value: number;
}

ChartJS.register(BarElement, CategoryScale, LinearScale, Tooltip, Legend);

// Utility function to get the last six months' labels
const getLastSixMonthsLabels = (): string[] => {
  const monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];
  const currentMonth = new Date().getMonth(); // Get the current month (0-11)
  const labels = [];

  for (let i = 5; i >= 0; i--) {
    const monthIndex = (currentMonth - i + 12) % 12;
    labels.push(monthNames[monthIndex]);
  }

  return labels;
};

// Shimmer component with vertical bars
const Shimmer = () => {
  return (
    <div className="flex space-x-2 h-full">
      {/* Simulating 6 bars */}
      {[...Array(6)].map((_, index) => (
        <div key={index} className="w-1/6 bg-gray-300 rounded-lg shimmer h-100%"></div>
      ))}
    </div>
  );
};

const BarChart: React.FC<{ token: string }> = ({ token }) => {
  const [chartData, setChartData] = useState<number[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [labels, setLabels] = useState<string[]>(getLastSixMonthsLabels());

  useEffect(() => {
    const fetchBalanceHistory = async () => {
      try {
        const response = await getBalanceHistory(6, token); // Fetching data for the last 6 months
        const data = response.data;

        // Sort the data by time to match with labels
        const sortedData = data
          .sort((a: BalanceHistoryData, b: BalanceHistoryData) => new Date(a.time).getMonth() - new Date(b.time).getMonth())
          .map((entry: BalanceHistoryData) => entry.value);

        setChartData(sortedData.length > 0 ? sortedData : [0, 0, 0, 0, 0, 0]);
      } catch (error) {
        console.error('Error fetching balance history:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchBalanceHistory();
  }, [token]);

  const maxValueIndex = chartData.indexOf(Math.max(...chartData));

  const data = {
    labels: labels,
    datasets: [
      {
        label: 'Balance History',
        data: chartData,
        backgroundColor: chartData.map((_, index) =>
          index === maxValueIndex ? 'rgba(0, 204, 204, 0.8)' : 'rgba(0, 0, 0, 0.05)'
        ),
        hoverBackgroundColor: chartData.map((_, index) =>
          index === maxValueIndex ? 'rgba(0, 204, 204, 1)' : 'rgba(0, 204, 204, 0.5)'
        ),
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
    <div className="bg-white dark:bg-gray-600 p-4 pt-8 rounded-3xl shadow-md w-full max-w-[600px] h-[230px]">
      {loading ? <Shimmer /> : <Bar data={data} options={options} />}
    </div>
  );
};

export default BarChart;



// import React, { useState, useEffect } from 'react';
// import { Bar } from 'react-chartjs-2';
// import { Chart as ChartJS, BarElement, CategoryScale, LinearScale, Tooltip, Legend, TooltipItem } from 'chart.js';

// ChartJS.register(BarElement, CategoryScale, LinearScale, Tooltip, Legend);

// // Dummy Data
// const dummyBalanceHistory = [
//   { time: 'Jan', value: 1200 },
//   { time: 'Feb', value: 900 },
//   { time: 'Mar', value: 1400 },
//   { time: 'Apr', value: 800 },
//   { time: 'May', value: 1700 },
//   { time: 'Jun', value: 1500 },
// ];

// const BarChart: React.FC<{ token: string }> = () => {
//   const [chartData, setChartData] = useState<number[]>([]);
//   const [loading, setLoading] = useState<boolean>(true);
//   const [labels, setLabels] = useState<string[]>(dummyBalanceHistory.map((entry) => entry.time));

//   useEffect(() => {
//     const values = dummyBalanceHistory.map((entry) => entry.value);
//     setChartData(values);
//     setLoading(false);
//   }, []);

//   // Find the index of the maximum value
//   const maxValueIndex = chartData.indexOf(Math.max(...chartData));

//   const data = {
//     labels: labels,
//     datasets: [
//       {
//         label: 'Balance History',
//         data: chartData,
//         backgroundColor: chartData.map((_, index) =>
//           index === maxValueIndex ? 'rgba(0, 204, 204, 0.8)' : 'rgba(0, 0, 0, 0.05)'
//         ),
//         hoverBackgroundColor: chartData.map((_, index) =>
//           index === maxValueIndex ? 'rgba(0, 204, 204, 1)' : 'rgba(0, 204, 204, 0.5)'
//         ),
//         borderRadius: 10,
//         borderSkipped: false,
//         barPercentage: 0.6,
//       },
//     ],
//   };

//   const options = {
//     responsive: true,
//     maintainAspectRatio: false,
//     plugins: {
//       legend: {
//         display: false,
//       },
//       tooltip: {
//         callbacks: {
//           label: function (tooltipItem: TooltipItem<'bar'>) {
//             const value = tooltipItem.raw as number;
//             return `$${value.toLocaleString()}`;
//           },
//         },
//       },
//     },
//     scales: {
//       x: {
//         grid: {
//           display: false,
//         },
//         ticks: {
//           color: 'rgba(0, 0, 0, 0.5)',
//         },
//       },
//       y: {
//         beginAtZero: true,
//         display: false,
//       },
//     },
//   };

//   return (
//     <div className="bg-white dark:bg-gray-600 p-4 pt-8 rounded-3xl shadow-md w-full max-w-[600px] h-[230px]">
//       {loading ? <p>Loading...</p> : <Bar data={data} options={options} />}
//     </div>
//   );
// };

// export default BarChart;
