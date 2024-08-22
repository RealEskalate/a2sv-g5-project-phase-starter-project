import { Bar, BarChart, ResponsiveContainer, XAxis } from "recharts";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { useState, useEffect } from "react";
import { getExpenses, getIncomes } from "@/services/transactionfetch";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";

// Define the chartConfig object
const chartConfig = {
  debit: {
    color: "#FF6384", 
  },
  credit: {
    color: "#36A2EB",
  },
};

interface Transaction {
  date: string;
  amount: number;
}

type DayOfWeek = "Sat" | "Sun" | "Mon" | "Tue" | "Wed" | "Thu" | "Fri";

export default function Component() {
  const [bottomMargin, setBottomMargin] = useState(90);
  const [barSize, setBarSize] = useState(20);
  const [chartData, setChartData] = useState<{ day: string; debit: number; credit: number }[]>([]);




  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const incomeData = await getIncomes(0, 5);
        const expenseData = await getExpenses(0, 5);

        const incomeArray = incomeData.data.content;
        const expenseArray = expenseData.data.content;

        console.log("Income Data:", incomeArray);
        console.log("Expense Data:", expenseArray);

        // Ensure that the arrays are present
        // if (!Array.isArray(incomeData.data) || !Array.isArray(expenseData.data)) {
        //   throw new Error('Invalid data structure');
        // }


       const dataByDay: Record<DayOfWeek, { debit: number; credit: number }> = {
  "Sat": { debit: 0, credit: 0 },
  "Sun": { debit: 0, credit: 0 },
  "Mon": { debit: 0, credit: 0 },
  "Tue": { debit: 0, credit: 0 },
  "Wed": { debit: 0, credit: 0 },
  "Thu": { debit: 0, credit: 0 },
  "Fri": { debit: 0, credit: 0 },
};

incomeArray.forEach((transaction: Transaction) => {
  const day = new Date(transaction.date).toLocaleDateString('en-US', { weekday: 'short' }) as DayOfWeek;
  console.log("date for income :", day)
  if (day in dataByDay) {
    dataByDay[day].credit += transaction.amount;
  }
});
console.log("aray after income update:",dataByDay)

expenseArray.forEach((transaction: Transaction) => {
  const days = new Date(transaction.date).toLocaleDateString('en-US', { weekday: 'short' }) as DayOfWeek;
  console.log("date:", days);
  
  if (days in dataByDay) {
    dataByDay[days].debit += transaction.amount;
  }
});

console.log("aray after expense update:",dataByDay)
    
      
        // Ensure `day` is typed as `DayOfWeek`
const newChartData = Object.keys(dataByDay).map((day) => {
  const typedDay = day as DayOfWeek; // Explicitly cast `day` to `DayOfWeek`
  return {
    day: typedDay,
    debit: dataByDay[typedDay].debit,
    credit: dataByDay[typedDay].credit,
  };
});

setChartData(newChartData);

      } catch (error) {
        console.error("Failed to fetch transactions", error);
      }
    };

    fetchTransactions();

    const handleResize = () => {
      const width = window.innerWidth;
      if (width < 768) { 
        setBottomMargin(20);
        setBarSize(10);
      } else if (width >= 768 && width < 1024) { 
        setBottomMargin(30);
        setBarSize(15);
      } else { 
        setBottomMargin(50);
        setBarSize(20);
      }
    };

    window.addEventListener('resize', handleResize);
    handleResize();

    return () => window.removeEventListener('resize', handleResize);
  }, []);

  const totalDebit = chartData.reduce((sum, data) => sum + data.debit, 0);
  const totalCredit = chartData.reduce((sum, data) => sum + data.credit, 0);

  return (
    <Card className="flex flex-col w-full h-full">
      <CardHeader className="flex justify-between">
        <div className="flex flex-row justify-between space-x-4">
          <div className="hidden md:flex text-sm font-normal">
            <span className="font-bold">${totalDebit}</span>&nbsp;Debited
            &&nbsp;
            <span className="font-bold"> ${totalCredit}</span>&nbsp;Credited in
            this Week
          </div>
          <div className="flex px-3 text-right">
            <span className="w-4 h-4 rounded-xl" style={{ backgroundColor: chartConfig.debit.color }}></span>
            <span className="text-sm font-normal pl-2 pr-4">Debit</span>
            <span className="w-4 h-4 rounded-xl" style={{ backgroundColor: chartConfig.credit.color }}></span>
            <span className="text-sm font-normal pl-2">Credit</span>
          </div>
        </div>
      </CardHeader>
      <CardContent className="flex-1">
        <ChartContainer config={chartConfig} className="p-0">
          <ResponsiveContainer width="100%" height="100%">
            <BarChart
              data={chartData}
              barCategoryGap="10%"
              barGap={5}
              barSize={barSize}
              margin={{ top: 10, right: 20, left: 20, bottom: bottomMargin }}
            >
              <XAxis dataKey="day" axisLine={true} tickLine={false} />
              <ChartTooltip cursor={false} content={<ChartTooltipContent indicator="dashed" />} />
              <Bar dataKey="debit" fill={chartConfig.debit.color} radius={[5, 5, 0, 0]} />
              <Bar dataKey="credit" fill={chartConfig.credit.color} radius={[5, 5, 0, 0]} />
            </BarChart>
          </ResponsiveContainer>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
