"use client";
import React from "react";
import "chart.js/auto";
import { ChartData } from "chart.js";
import { Line } from "react-chartjs-2";

import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from "chart.js";
import { useGetInvestmentItemsQuery } from "@/lib/redux/slices/investmentSlice";
import GraphSkeletons from "../AllSkeletons/chartSkeleton/graphSkeletons";

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);

const YearlyTotalInvestment = () => {
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

  const Datavalues = rawdata?.data?.yearlyTotalInvestment;

  Datavalues?.forEach((element) => {
    Labels.push(element.time);
    values.push(element.value);
  });
  Labels.reverse();
  values.reverse();

  const Alldata: ChartData<"line"> = {
    labels: Labels,
    datasets: [
      {
        data: values,
        borderColor: "#FCAA0B",
        borderWidth: 2,
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
    },
  };

  return (
    <div className="w-full md:w-1/2">
      <h1 className="text-[#333B69] text-20px py-2 font-semibold">
        Yearly Total Investment
      </h1>
      <div className="bg-white p-6 rounded-3xl">
        <Line data={Alldata} options={options} className="w-full" />
      </div>
    </div>
  );
};

export default YearlyTotalInvestment;
