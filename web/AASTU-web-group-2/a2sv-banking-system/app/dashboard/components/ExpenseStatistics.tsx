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
import Refresh from "@/app/api/auth/[...nextauth]/token/RefreshToken";
import { IconType } from "react-icons";

const initialChartData = [
  { browser: "shopping", amount: 0, fill: "var(--color-shopping)" },
  { browser: "transfer", amount: 0, fill: "var(--color-transfer)" },
  { browser: "deposit", amount: 0, fill: "var(--color-deposit)" },
  {browser: "service", amount: 0, fill: "var(--color-service)"}
];

const chartConfig = {
  Amount: {
    label: "Amount",
  },
  shopping: {
    label: "Shopping",
    color: "hsl(var(--chart-1))",
  },
  transfer: {
    label: "Transfer",
    color: "hsl(var(--chart-2))",
  },
  deposit: {
    label: "Deposit",
    color: "hsl(var(--chart-3))",
  },
  service: {
    label: "Service",
    color: "hsl(var(--chart-4))",
  },
} satisfies ChartConfig;

type DataItem = {
  heading: string;
  text: string;
  headingStyle: string;
  dataStyle: string;
};

type Column = {
  icon: IconType;
  iconStyle: string;
  data: DataItem[];
};

type Data = {
  access_token: string;
  data: string;
  refresh_token: string;
};

type SessionDataType = {
  user: Data;
};
const ShimmerEffect = () => (
  <div className="h-64 w-64 max-w-[300px] bg-gray-200 animate-pulse rounded-full"></div>
);
export function ExpenseStatistics() {
  const [chartData, setChartData] = useState(initialChartData);
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
    const fetchAndProcessExpenses = async () => {
      try {
        if (access_token) {
          const { data } = await getTransactionsExpenses(0, 1000, access_token);
          console.log("worked", data)
          const typeAmounts: { [key: string]: number } = {
            shopping: 0,
            transfer: 0,
            deposit: 0,
            service: 0
          };

          data.content.forEach((transaction: any) => {
            if (typeAmounts.hasOwnProperty(transaction.type)) {
              typeAmounts[transaction.type] += transaction.amount;
            } else {
              typeAmounts.other += transaction.amount; // If type is not predefined, add to "other"
            }
          });

          setChartData([
            { browser: "shopping", amount: typeAmounts.shopping, fill: "var(--color-shopping)" },
            { browser: "transfer", amount: typeAmounts.transfer, fill: "var(--color-transfer)" },
            { browser: "deposit", amount: typeAmounts.deposit, fill: "var(--color-deposit)" },
            { browser: "service", amount: typeAmounts.service, fill: "var(--color-service)" },
          ]);

          setLoading(false);
        }
      } catch (error) {
        console.error("Error fetching expenses:", error);
        setLoading(false);
      }
    };

    fetchAndProcessExpenses();
  });

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
            <ShimmerEffect/>
          ) : (
            <ChartContainer 
              config={chartConfig}
              className="aspect-square h-72 w-full max-w-[300px]" // Ensure full width within a max limit
            >
              <PieChart>
                <ChartTooltip
                  content={<ChartTooltipContent nameKey="amount" hideLabel />}
                />
                <Pie
                  data={chartData}
                  dataKey="amount"
                  paddingAngle={0} // Adds margin between the slices
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
