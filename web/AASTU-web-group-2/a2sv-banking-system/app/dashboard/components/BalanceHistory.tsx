"use client";
import { TrendingUp } from "lucide-react";
import { Area, AreaChart, CartesianGrid, XAxis, YAxis } from "recharts";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { getSession } from "next-auth/react";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import Refresh from "@/app/api/auth/[...nextauth]/token/RefreshToken";
import { getBalanceHistory } from "@/lib/api/transactionController";



const chartConfig = {
  Balance: {
    label: "Balance",
    color: "hsl(var(--chart-1))",
  },
} satisfies ChartConfig;
type DataItem = {
  heading: string;
  text: string;
  headingStyle: string;
  dataStyle: string;
};

// eslint-disable-next-line react-hooks/rules-of-hooks


type Data = {
  access_token: string;
  data: string;
  refresh_token: string;
};
type SessionDataType = {
  user: Data;
};
export function BalanceHistory() {
  const [chartData, setChartData] = useState<{ month: string; balance: number }[]>([]);

  const [loading, setLoading] = useState(true);
  const [session, setSession] = useState<Data | null>(null);
  const router = useRouter();
  const [access_token, setAccess_token] = useState("");


  useEffect(() => {
    const fetchSession = async () => {
      const sessionData = (await getSession()) as SessionDataType | null;
      setAccess_token(await Refresh());

      if (sessionData && sessionData.user) {
        setSession(sessionData.user);
      } else {
        router.push(
          `./api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`
        );
      }
      setLoading(false);
    };

    fetchSession();
  }, [router]);
  useEffect(() => {
    const fetchBalanceHistory = async () => {
      if (!access_token) return;
      try {
        const balanceHistory = await getBalanceHistory(access_token);

        if (balanceHistory.success) {
          const formattedData = balanceHistory.data.map((item) => ({
            month: item.time, // Adjust the time to match your needs (e.g., month)
            balance: item.value,
          }));
          setChartData(formattedData);
        }
      } catch (error) {
        console.error("Error fetching balance history:", error);
      }
    };

    fetchBalanceHistory();

  }, [access_token]);
  return (
    <Card className="my-4 mx-4 rounded-3xl flex-grow md:w-[75%]">
      <CardHeader>
        <CardTitle className="text-[#343C6A] font-bold text-xl md:hidden">
          Balance History
        </CardTitle>
      </CardHeader>
      <CardContent>
        <div className="w-full">
          <ChartContainer config={chartConfig} className="md:h-48 md:w-full">
            <AreaChart
              width={5}
              height={300}
              data={chartData}
              className="aspect-square h-60 w-full max-w-[300px]"
            >
              <CartesianGrid
                strokeDasharray="3 3"
                stroke="rgba(0, 0, 0, 0.5)"
                vertical={true}
                horizontal={true}
              />
              <XAxis
                dataKey="month"
                tickLine={false}
                axisLine={false}
                tickMargin={8}
                tickFormatter={(value) => value.slice(0, 3)}
              />
              <YAxis
                tickLine={false}
                axisLine={false}
                tickMargin={8}
                interval={0}
                ticks={[0, 200, 400, 600, 800]}
                domain={[0, 800]}
              />
              <ChartTooltip
                cursor={false}
                content={<ChartTooltipContent indicator="line" />}
              />
              <Area
                dataKey="Balance"
                type="natural"
                fill="rgba(0, 0, 255, 0.2)"
                stroke="blue"
              />
            </AreaChart>
          </ChartContainer>
        </div>
      </CardContent>
    </Card>
  );
}
