'use client';

import React, { useEffect, useState } from 'react';
import {
  AreaChart,
  Area,
  XAxis,
  YAxis,
  Tooltip,
  CartesianGrid,
  ResponsiveContainer,
} from 'recharts';
import axios from 'axios';
import { useSession } from 'next-auth/react';

interface ExtendedUser {
  name?: string;
  email?: string;
  image?: string;
  accessToken?: string;
  refreshToken?: string;
}

const LineGraphComponent = () => {
  const { data: session } = useSession();
  const user = session?.user as ExtendedUser;

  const [chartData, setChartData] = useState<{ time: string; value: number }[]>([]);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      if (user?.accessToken) {
        try {
          const response = await axios.get('https://bank-dashboard-irbd.onrender.com/transactions/random-balance-history?monthsBeforeFirstTransaction=6', {
            headers: {
              Authorization: `Bearer ${user.accessToken}`,
            },
          });

          // Map the response data to match the expected structure for the chart
          const data = response.data.data.map((transaction: { time: string; value: number }) => {
            const date = new Date(transaction.time);
            const formattedTime = date.toLocaleString('en-US', { month: 'short' });
            return {
              time: formattedTime,
              value: transaction.value,
            };
          });

          setChartData(data);
          setError(null); // Clear any previous errors
        } catch (error) {
          if (error.response?.status === 401) {
            // Token might be expired, try to refresh it
            try {
              const refreshResponse = await axios.post('https://bank-dashboard-irbd.onrender.com/auth/refresh_token', {}, {
                headers: {
                  "Content-Type": "application/json",
                  Authorization: `Bearer ${user.refreshToken}`,
                },
              });

              const refreshedTokens = refreshResponse.data.data;
              const newAccessToken = refreshedTokens.access_token;

              // Retry the original request with the new access token
              const retryResponse = await axios.get('https://bank-dashboard-rsf1.onrender.com/transactions/random-balance-history?monthsBeforeFirstTransaction=6', {
                headers: {
                  Authorization: `Bearer ${newAccessToken}`,
                },
              });

              const retriedData = retryResponse.data.data.map((transaction: { time: string; value: number }) => {
                const date = new Date(transaction.time);
                const formattedTime = date.toLocaleString('en-US', { month: 'short' });
                return {
                  time: formattedTime,
                  value: transaction.value,
                };
              });

              setChartData(retriedData);
            } catch (refreshError) {
              console.error("Failed to refresh access token:", refreshError);
              setError("Failed to refresh access token. Please log in again.");
            }
          } else {
            console.error("Failed to fetch data:", error);
            setError("Failed to fetch data. Please check the console for more details.");
          }
        }
      }
    };

    fetchData();
  }, [user]);

  if (error) return <p>Error: {error}</p>;

  return (
    <div className="w-full h-64 bg-white dark:bg-gray-800 border dark:border-white">
      <ResponsiveContainer width="100%" height="100%">
        <AreaChart
          data={chartData}
          margin={{ top: 16, left: 4, right: 8, bottom: 4 }}
        >
          <CartesianGrid strokeDasharray="3 3" stroke="#666666" />
          <XAxis dataKey="time" tick={{ fill: '#666666' }} axisLine={{ stroke: '#666666' }} />
          <YAxis tick={{ fill: '#666666' }} axisLine={{ stroke: '#666666' }} />
          <Tooltip contentStyle={{ backgroundColor: '#333333', borderColor: '#666666' }} labelStyle={{ color: '#ffffff' }} />
          <Area
            dataKey="value"
            type="monotone"
            fill="#8884d8"
            stroke="#8884d8"
            fillOpacity={0.2}
          />
        </AreaChart>
      </ResponsiveContainer>
    </div>
  );
};

export default LineGraphComponent;
