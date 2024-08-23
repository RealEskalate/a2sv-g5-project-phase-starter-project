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
}

const LineGraphComponent = () => {
  const { data: session } = useSession();
  const user = session?.user as ExtendedUser;

  const [chartData, setChartData] = useState([]);

  useEffect(() => {
    if (user?.accessToken) {
      axios
        .get('https://bank-dashboard-o9tl.onrender.com/transactions/balance-history', {
          headers: {
            Authorization: `Bearer ${user.accessToken}`,
          },
        })
        .then((response) => {
          // Assuming the API response has a data array with time and value fields
          const data = response.data.map((transaction: { time: string; value: number }) => ({
            time: transaction.time,
            value: transaction.value,
          }));
          setChartData(data);
        })
        .catch((error) => {
          console.error('Error fetching chart data:', error);
        });
    }
  }, [user?.accessToken]);

  return (
    <div className="w-full h-64">
      <ResponsiveContainer width="100%" height="100%">
        <AreaChart
          data={chartData}
          margin={{ top: 16, left: 4, right: 8, bottom: 4 }}
        >
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="time" />
          <YAxis />
          <Tooltip />
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
