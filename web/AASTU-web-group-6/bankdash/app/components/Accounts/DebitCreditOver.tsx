"use client";

import { TrendingUp } from "lucide-react";
import { Bar, BarChart, CartesianGrid, XAxis } from "recharts";
import { useState, useEffect } from "react";
import DebitCredit from "@/app/Services/api/Debitcredit";
import { useSession } from "next-auth/react";
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
import { useAppSelector } from "@/app/Redux/store/store";
const chartConfig = {
  debit: {
    label: "Debit",
    color: "hsl(var(--chart-1))",
  },
  credit: {
    label: "Credit",
    color: "hsl(var(--chart-2))",
  },
} satisfies ChartConfig;
interface chartData {
  day: string;
  debit: number;
  credit: number;
}
interface GroupedTransactions {
  [weekKey: string]: {
    income: { [dayKey: string]: number };
    expense: { [dayKey: string]: number };
  };
}

function getDayOfWeek(dateString: string): string {
  const date = new Date(dateString);
  const daysOfWeek = [
    "Sunday",
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
  ];
  return daysOfWeek[date.getDay()];
}
export function DebitCreditOver() {
  const [data, setData] = useState<chartData[]>([]);
  const [totalIncome, setTotalIncome] = useState(0);
  const [totalExpense, setTotalExpense] = useState(0);
  const [currentWeek, setCurrentWeek] = useState(0);
  const [transactions, setTransactions] = useState<GroupedTransactions | null>(
    null
  );
  const [dateRange, setDateRange] = useState("");
  const {data:session} = useSession();
  const accessToken =  session?.accessToken as string;
  console.log(accessToken , "accessTokenDebit")
  const income = useAppSelector((state) => state.transactions.income )
  const expense = useAppSelector((state) => state.transactions.expense )
  useEffect(() => {
    const fetchData = async () => {
      // console.log(ac)
      try {
        const ans: GroupedTransactions | undefined = await DebitCredit(accessToken , income , expense);
        if (ans) {
          setTransactions(ans);
          processWeekData(ans, currentWeek);
        } else {
          console.error("No data returned from DebitCredit.");
        }
      } catch (error) {
        console.error("Error fetching data", error);
      }
    };

    fetchData();
  }, []);

  const processWeekData = (
    transactions: GroupedTransactions,
    weekIndex: number
  ) => {
    const weekKeys = Object.keys(transactions);

    if (weekIndex < 0 || weekIndex >= weekKeys.length) {
      setCurrentWeek(weekKeys.length - 1);
      console.error("Invalid week index.");
      return;
    }

    const selectedWeek = weekKeys[weekIndex];
    const startDate = new Date(selectedWeek);
    const endDate = new Date(startDate);
    endDate.setDate(startDate.getDate() + 6);

    const options: Intl.DateTimeFormatOptions = {
      month: "short",
      day: "numeric",
    };
    const formattedDateRange = `${startDate.toLocaleDateString(
      undefined,
      options
    )} - ${endDate.toLocaleDateString(undefined, options)}`;

    setDateRange(formattedDateRange);

    const chartData: chartData[] = [];
    let incomeSum = 0;
    let expenseSum = 0;

    const dailyData: { [key: string]: { debit: number; credit: number } } = {};

    Object.keys(transactions[selectedWeek].income).forEach((dayKey) => {
      incomeSum += transactions[selectedWeek].income[dayKey];
      const dayName = getDayOfWeek(dayKey);
      if (!dailyData[dayName]) {
        dailyData[dayName] = { debit: 0, credit: 0 };
      }
      dailyData[dayName].credit += transactions[selectedWeek].income[dayKey];
    });

    Object.keys(transactions[selectedWeek].expense).forEach((dayKey) => {
      expenseSum += transactions[selectedWeek].expense[dayKey];
      const dayName = getDayOfWeek(dayKey);
      if (!dailyData[dayName]) {
        dailyData[dayName] = { debit: 0, credit: 0 };
      }
      dailyData[dayName].debit += transactions[selectedWeek].expense[dayKey];
    });

    const daysOfWeek = [
      "Sunday",
      "Monday",
      "Tuesday",
      "Wednesday",
      "Thursday",
      "Friday",
      "Saturday",
    ];

    daysOfWeek.forEach((day) => {
      if (dailyData[day]) {
        chartData.push({
          day: day,
          debit: dailyData[day].debit,
          credit: dailyData[day].credit,
        });
      } else {
        chartData.push({
          day: day,
          debit: 0,
          credit: 0,
        });
      }
    });

    setData(chartData);
    setTotalExpense(expenseSum);
    setTotalIncome(incomeSum);
  };

  const nextWeek = () => {
    setCurrentWeek((prev) => {
      const newWeek = prev + 1;
      processWeekData(transactions as GroupedTransactions, newWeek);
      return newWeek;
    });
  };

  const prevWeek = () => {
    setCurrentWeek((prev) => {
      const newWeek = prev > 0 ? prev - 1 : 0;
      processWeekData(transactions as GroupedTransactions, newWeek);
      return newWeek;
    });
  };

  useEffect(() => {
    if (transactions) {
      processWeekData(transactions, currentWeek);
    }
  }, [currentWeek, transactions]);
  return (
    <Card className="rounded-3xl shadow-lg dark:bg-[#232328]  ">
      <CardHeader>
        <div className="flex justify-end lg:justify-between ">
          <CardTitle className="hidden gap-2 lg:block lg:text-[12px] xl:text-base text-base font-normal font-inter text-[#718EBF] dark:text-gray-400 ">
            <span className="font-medium text-black dark:text-gray-300">
              ${totalExpense}
            </span>{" "}
            Debited &{" "}
            <span className="font-medium text-black dark:text-gray-300">
              ${totalIncome}
            </span>{" "}
            Credited in this Week
            <p className="pt-3 font-medium text-black dark:text-gray-300">
              Week of {dateRange}
            </p>
          </CardTitle>

          <div className="flex gap-5 ">
            <div className="flex items-center gap-2">
              <div className="border border-[#4C78FF] w-[15px] h-[15px] rounded-sm bg-[#4C78FF]"></div>
              <p className="font-inter font-normal text-base text-[#718EBF] dark:text-gray-300">
                Debit
              </p>
            </div>
            <div className="flex items-center gap-2">
              <div className="border border-[#FCAA0B] w-[15px] h-[15px] rounded-sm bg-[#FCAA0B]"></div>
              <p className="font-inter font-normal text-base text-[#718EBF] dark:text-gray-300">
                Credit
              </p>
            </div>
          </div>
        </div>
      </CardHeader>
      <CardContent>
        <ChartContainer config={chartConfig} className="lg:h-[350px] w-[100%]">
          <BarChart accessibilityLayer data={data}>
            <CartesianGrid vertical={false} className="h-[50%] lg:h-[70%]" />
            <XAxis
              dataKey="day"
              tickLine={false}
              tickMargin={10}
              axisLine={false}
              fontSize={10}
              tickFormatter={(value) => value.slice(0, 3)}
            />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dashed" />}
            />
            <Bar dataKey="debit" fill="#1A16F3" radius={10} />
            <Bar dataKey="credit" fill="#FCAA0B" radius={10} />
          </BarChart>
        </ChartContainer>
      </CardContent>
      <CardFooter className="flex justify-end md:pr-6">
        <div className="flex gap-6 ">
          <button onClick={nextWeek} className="btn-prev">
            <Prev />
          </button>
          <button onClick={prevWeek} className="btn-next">
            <Next />
          </button>
        </div>
      </CardFooter>
    </Card>
  );
}
function Prev() {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="1.5"
      stroke="currentColor"
      className="size-9"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="m11.25 9-3 3m0 0 3 3m-3-3h7.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
      />
    </svg>
  );
}
function Next() {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 24 24"
      stroke-width="1.5"
      stroke="currentColor"
      className="size-9"
    >
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        d="m12.75 15 3-3m0 0-3-3m3 3h-7.5M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
      />
    </svg>
  );
}
