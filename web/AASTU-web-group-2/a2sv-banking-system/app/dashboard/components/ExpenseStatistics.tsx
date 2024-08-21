"use client";

import { TrendingUp } from "lucide-react";
import { LabelList, Pie, PieChart } from "recharts";
import { useState, useEffect } from "react";
import { getSession } from "next-auth/react";
import { useRouter } from "next/navigation";

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
import {
  getTransactionIncomes,
  getTransactionsExpenses,
} from "@/lib/api/transactionController";

const initialChartData = [
  { browser: "shopping", visitors: 0, fill: "var(--color-shopping)" },
  { browser: "transfer", visitors: 0, fill: "var(--color-transfer)" },
  { browser: "other", visitors: 0, fill: "var(--color-other)" },
];

const chartConfig = {
  visitors: {
    label: "Visitors",
  },
  shopping: {
    label: "Shopping",
    color: "hsl(var(--chart-1))",
  },
  transfer: {
    label: "Transfer",
    color: "hsl(var(--chart-2))",
  },
  other: {
    label: "Other",
    color: "hsl(var(--chart-3))",
  },
} satisfies ChartConfig;
type Data = {
  access_token: string;
  data: string;
  refresh_token: string;
};

type SessionDataType = {
  user: Data;
};
export function ExpenseStatistics() {
  const [chartData, setChartData] = useState(initialChartData);
  const [loading, setLoading] = useState(true);
  const [session, setSession] = useState<Data | null>(null);
  const router = useRouter();


  useEffect(() => {
    const fetchSession = async () => {
      const sessionData = (await getSession()) as SessionDataType | null;
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
    const fetchAndProcessExpenses = async () => {
      try {
        const sessionToken = session?.access_token;
        if (sessionToken) {
          const { data } = await getTransactionsExpenses(0, 1, sessionToken);
          
          const typeAmounts: { [key: string]: number } = {
            shopping: 0,
            transfer: 0,
            other: 0,
          };

          data.content.forEach((transaction: any) => {
            if (typeAmounts.hasOwnProperty(transaction.type)) {
              typeAmounts[transaction.type] += transaction.amount;
            } else {
              typeAmounts.other += transaction.amount; // If type is not predefined, add to "other"
            }
          });

          setChartData([
            { browser: "shopping", visitors: typeAmounts.shopping, fill: "var(--color-shopping)" },
            { browser: "transfer", visitors: typeAmounts.transfer, fill: "var(--color-transfer)" },
            { browser: "other", visitors: typeAmounts.other, fill: "var(--color-other)" },
          ]);

          setLoading(false);
        }
      } catch (error) {
        console.error("Error fetching expenses:", error);
        setLoading(false);
      }
    };

    fetchAndProcessExpenses();
  }, []);

  return (
    <Card className="mx-4 my-6 md:my-0 flex-grow rounded-3xl">
      <CardHeader className="items-left pb-0">
        <CardTitle className="text-[#343C6A] font-bold text-xl md:hidden">
          Expense Statistics
        </CardTitle>
      </CardHeader>
      <CardContent className="flex-1 pb-0">
        <div className="flex justify-center items-center">
          {loading ? (
            <p>Loading...</p>
          ) : (
            <ChartContainer 
              config={chartConfig}
              className="aspect-square h-72 w-full max-w-[300px]" // Ensure full width within a max limit
            >
              <PieChart>
                <ChartTooltip
                  content={<ChartTooltipContent nameKey="visitors" hideLabel />}
                />
                <Pie
                  data={chartData}
                  dataKey="visitors"
                  paddingAngle={5} // Adds margin between the slices
                >
                  <LabelList
                    dataKey="browser"
                    className="fill-background"
                    stroke="none"
                    fontSize={12}
                    formatter={(value: keyof typeof chartConfig) =>
                      chartConfig[value]?.label
                    }
                  />
                </Pie>
              </PieChart>
            </ChartContainer>
          )}
        </div>
      </CardContent>
    </Card>
  );
}
