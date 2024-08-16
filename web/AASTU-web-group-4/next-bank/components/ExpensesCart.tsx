import { Bar } from 'react-chartjs-2';
import { Chart as ChartJS, CategoryScale, LinearScale, BarElement, Tooltip, Legend } from 'chart.js';

ChartJS.register(CategoryScale, LinearScale, BarElement, Tooltip, Legend);

const ExpensesChart = () => {
  const data = {
    labels: ['January', 'February', 'March', 'April', 'May', 'June'],
    datasets: [
      {
        label: 'Monthly Expenses',
        data: [500, 800, 400, 700, 900, 600],
        backgroundColor: 'rgba(237, 240, 247, 1)', // Default gray
        borderRadius: 12, // Rounded corners on both top and bottom
        hoverBackgroundColor: 'rgba(22, 219, 204, 1)', // Lemon green on hover
      },
    ],
  };

  const options = {
    scales: {
      x: {
        grid: {
          display: false, // Hide grid lines
        },
        ticks: {
          color: 'rgba(75, 85, 99, 1)', // Color for the labels on the x-axis
        },
        border: {
          display: false, // Hide x-axis line
        },
      },
      y: {
        grid: {
          display: false, // Hide grid lines
        },
        ticks: {
          display: false, // Hide ticks on y-axis
        },
        border: {
          display: false, // Hide y-axis line
        },
      },
    },
    plugins: {
      legend: {
        display: false, // Hide legend
      },
      tooltip: {
        callbacks: {
          label: function (context: any) {
            return `$${context.raw}`;
          },
        },
      },
    },
  };

  return (
    <div className="w-full h-48 overflow-hidden">
      <Bar data={data} options={options} />
    </div>
  );
};

export default ExpensesChart;
