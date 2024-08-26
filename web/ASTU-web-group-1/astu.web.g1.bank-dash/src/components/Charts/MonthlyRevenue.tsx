'use client';
import React from 'react';
import 'chart.js/auto';
import { ChartData } from 'chart.js';
import { Line } from 'react-chartjs-2';

import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';
import { useGetInvestmentItemsQuery } from '@/lib/redux/slices/investmentSlice';
import GraphSkeletons from '../AllSkeletons/chartSkeleton/graphSkeletons';

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

const MonthlyRevenue = () => {
  // Define the chart data with type annotations
  const { data: rawdata, isLoading } = useGetInvestmentItemsQuery({
    years: 10,
    months: 5,
  });

  if (isLoading) {
    return <GraphSkeletons />;
  }
  const Labels: string[] = [];
  const values: number[] = [];

  const Datavalues = rawdata?.data?.monthlyRevenue;

  Datavalues?.forEach((element) => {
    Labels.push(element.time);
    values.push(element.value);
  });
  Labels.reverse();
  values.reverse();

  const data: ChartData<'line'> = {
    labels: Labels,
    datasets: [
      {
        data: values,
        borderColor: '#16DBCC',
        borderWidth: 2,
        tension: 0.2,
      },
    ],
  };

  const options = {
    responsive: true,
    plugins: {
      legend: {
        display: false,
      },
      title: {
        display: false,
      },
      datalabels: {
        display: false, // This will hide the tooltip
      },
    },
  };

  return (
    <div className='w-full md:w-1/2'>
      <h1 className='text-[#333B69] text-20px py-2 font-semibold'>Monthly Revenue</h1>
      <div className='bg-white p-6 rounded-3xl'>
        <Line data={data} options={options} className='w-full' />
      </div>
    </div>
  );
};

export default MonthlyRevenue;
