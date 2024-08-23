'use client';

import React, { useEffect, useState } from 'react';
import { BarChart, Bar, XAxis, YAxis, Tooltip, CartesianGrid, ResponsiveContainer, Cell } from 'recharts';
import axios from 'axios';
import { useSession } from 'next-auth/react';

interface ExtendedUser {
  name?: string;
  email?: string;
  image?: string;
  accessToken?: string;
}

const BarChartComponent = () => {
  const { data: session } = useSession();
  const user = session?.user as ExtendedUser;

  const [chartData, setChartData] = useState([]);

  useEffect(() => {
    if (user?.accessToken) {
      axios
        .get('https://bank-dashboard-1tst.onrender.com/transactions?page=0', {
          headers: {
            'Authorization': `Bearer ${user.accessToken}`,
          },
        })
        .then((response) => {
          const data = response.data.map((transaction) => {
            return {
              month: new Date(transaction.date).toLocaleString('default', { month: 'long' }),
              amount: transaction.amount,
            };
          });

          // Grouping data by month and calculating expenses per month
          const groupedData = data.reduce((acc, curr) => {
            const month = curr.month;
            if (!acc[month]) {
              acc[month] = { month, amount: 0 };
            }
            if (curr.amount < 0) {
              acc[month].amount += curr.amount;
            }
            return acc;
          }, {});

          const formattedData = Object.values(groupedData);
          setChartData(formattedData);
        })
        .catch((error) => {
          console.error('Error fetching chart data:', error);
        });
    }
  }, [user?.accessToken]);

  return (
    <div className="w-full h-64">
      <ResponsiveContainer width="100%" height="100%">
        <BarChart data={chartData}>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis dataKey="month" />
          <YAxis />
          <Tooltip />
          <Bar dataKey="amount" fill="#8884d8" onMouseOver={() => console.log('Hovered!')}>
            {chartData.map((entry, index) => (
              <Cell
                key={`cell-${index}`}
                fill={index % 2 === 0 ? '#82ca9d' : '#8884d8'}
              />
            ))}
          </Bar>
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
};

export default BarChartComponent;
