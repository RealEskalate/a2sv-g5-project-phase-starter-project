"use client";
import React, { useEffect, useState } from "react";
import { getSession } from "next-auth/react";
import { useRouter } from "next/navigation";
import Refresh from "../../api/auth/[...nextauth]/token/RefreshToken";
import Card1 from "../components/Card1";
import { getServerSession } from "next-auth";
import { options } from "../../api/auth/[...nextauth]/options";
import { TrendingUp } from "lucide-react";
import { CartesianGrid, Line, LineChart, XAxis, YAxis } from "recharts";
import { getRandomInvestementData } from "../back/Invest";

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/app/loans/components/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/app/loans/components/chart";
import { cp } from "fs";
const chartConfig = {
  value: {
    label: "value",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;
interface arr {
  time: string;
  value: number;
}
const token =
  "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJsc2FqZGxzanNuIiwiaWF0IjoxNzI0MTU1NzkzLCJleHAiOjE3MjQyNDIxOTN9.wi7oRgF81zMp1v8tPzRPmAj4GOLaYy4bV_TMVvtWmzg2mjrTThiruT_Fswcyu1eq";

interface info {
  totalInvestment: number;
  rateOfReturn: number;
  yearlyTotalInvestment: arr[];
  monthlyRevenue: arr[];
}

type Data = {
  access_token: string;
  data: string;
  refresh_token: string;
};

type SessionDataType = {
  user: Data;
};

export default function Monthly() {
  const [session, setSession] = useState<Data | null>(null);
  const [access_token, setAccess_token] = useState("");
  const router = useRouter();
  const [loading, setloading] = useState(true);
  const [Loading, setLoading] = useState(true);
  const [data, setdata] = useState<info>({
    totalInvestment: 1,
    rateOfReturn: 1,
    yearlyTotalInvestment: [],
    monthlyRevenue: [],
  });

  // Getting the session from the server and Access Token From Refresh
  useEffect(() => {
    const fetchSession = async () => {
      try {
        const sessionData = (await getSession()) as SessionDataType | null;
        setAccess_token(await Refresh());
        if (sessionData && sessionData.user) {
          setSession(sessionData.user);
        } else {
          router.push(
            `./api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`
          );
        }
      } catch (error) {
        console.error("Error fetching session:", error);
      } finally {
        setloading(false);
      }
    };

    fetchSession();
  }, [router]);

  // Combined fetching data to reduce multiple useEffect hooks
  useEffect(() => {
    const fetchData = async () => {
      if (!access_token) return;

      try {
        // Fetch data
        const d: info = await getRandomInvestementData(11, 2021, access_token);
        setdata(d);
      } catch (error) {
        console.error("Error fetching data:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [access_token]);

  if (loading || Loading)
    return (
      <div className="bg-white rounded-3xl border-none p-8 animate-shimmer">
        <div className="h-6 w-full bg-gray-300 rounded-md mb-4"></div>
        <div className="h-4 w-3/4 bg-gray-300 rounded-md mb-4"></div>
        <div className="h-4 w-1/2 bg-gray-300 rounded-md mb-4"></div>
        <div className="h-4 w-2/3 bg-gray-300 rounded-md"></div>
      </div>
    );
  // console.log(data);
  const { monthlyRevenue } = data;
  return (
    <Card className="bg-white rounded-3xl border-none ">
      <CardContent className="pt-8 pb-6 ">
        <ChartContainer config={chartConfig}>
          <LineChart
            accessibilityLayer
            data={monthlyRevenue}
            margin={{
              left: 12,
              right: 12,
            }}
          >
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="time"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              tickFormatter={(value) => value.slice(0, 7)}
            />
            <YAxis
              dataKey="value"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              // tickFormatter={(value) => value.slice(0, 3)}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent hideLabel={false} />}
            />
            <Line
              dataKey="value"
              type="natural"
              stroke="#16DBCC"
              strokeWidth={3}
              dot={false}
            />
          </LineChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
