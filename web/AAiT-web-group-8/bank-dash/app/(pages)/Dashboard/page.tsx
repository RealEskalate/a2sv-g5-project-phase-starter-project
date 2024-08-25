
import ExpenseStatisticsChart from "@/components/dashboard/Charts/ExpenseStatisticsChart";
import WeeklyActivityChart from "@/components/dashboard/Charts/WeeklyActivityChart";
import CreditCard from "@/components/dashboard/CreditCard";
import QuickTransfer from "@/components/dashboard/QuickTransfer";
import RecentTransaction from "@/components/dashboard/RecentTransaction";
import LineGraph from "@/components/stats/linegraph";
import LightChip from "@/public/Chip_Card _Light.svg";
import DarkChip from "@/public/Chip_Card.svg";
import firstIcon from "@/public/Group 313.svg";
import secondIcon from "@/public/Group 314.svg";
import thirdIcon from "@/public/Group 315.svg";


const Dashboard = () => {
  return (
    <div className="container mx-auto p-4">
      <div className="flex flex-col lg:flex-row justify-between space-y-6 lg:space-y-0 lg:space-x-6">
        <div className="flex-1">
          <div className="flex justify-between items-center mb-4">
            <h2 className="text-lg font-semibold text-[#343C6A]">My Cards</h2>
            <a href="#" className="text-[#343C6A]">
              See All
            </a>
          </div>
          <div className="flex flex-row justify-start space-x-4 overflow-x-auto">
            <div className="flex space-x-4">
              <CreditCard
                balance="$5,756"
                cardHolder="Eddy Cusuma"
                validThru="12/22"
                cardNumber="3778 **** **** 1234"
                gradientFrom="#4C49ED"
                gradientTo="#0A06F4"
                chipImage={LightChip}
                textColor="text-white"
                bottomBackground="bg-gradient-to-br from-[rgba(76,73,237,0.8)] to-[rgba(76,73,237,1)]"
              />
              <CreditCard
                balance="$7,500"
                cardHolder="Jane Smith"
                validThru="11/23"
                cardNumber="3778 **** **** 1234"
                gradientFrom="white"
                gradientTo="white"
                textColor=""
                bottomBackground="border-t-2 bg-white"
                chipImage={DarkChip}
              />
            </div>
          </div>
        </div>
        <div>
          <h2 className="text-lg font-semibold mb-4 text-[#343C6A]">
            Recent Transactions
          </h2>
          <div className="bg-white p-4 rounded-3xl shadow">
            <div className="space-y-4">
              <RecentTransaction
                title="Deposit from my Card"
                date="28 January 2021"
                amount="-$850"
                type="expense"
                imageSrc={firstIcon}
              />
              <RecentTransaction
                title="Deposit Paypal"
                date="25 January 2021"
                amount="+$2,500"
                type="income"
                imageSrc={secondIcon}
              />
              <RecentTransaction
                title="Jemi Wilson"
                date="21 January 2021"
                amount="+$5,400"
                type="income"
                imageSrc={thirdIcon}
              />
            </div>
          </div>
        </div>
      </div>

      <div className="mt-10 grid grid-cols-1 lg:grid-cols-4 gap-6 h-72">
        <div className="col-span-3">
          <h2 className="text-lg font-semibold mb-4 text-[#343C6A]">
            Weekly Activity
          </h2>
          <div className="h-70">
            <WeeklyActivityChart />
          </div>
        </div>

        <div className="col-span-1">
          <h2 className="text-lg font-semibold mb-4 text-[#343C6A]">
            Expense Statistics
          </h2>
          <div className="relative bg-white p-6 rounded-2xl shadow h-72">
            <ExpenseStatisticsChart />
          </div>
        </div>
      </div>
      <div></div>
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8 px-4 mt-8">
        <div className="flex flex-col space-y-4 py-8 pb-10 min-h-full">
          <p className="font-bold text-[#343C6A]">Quick Transfer</p>
          <QuickTransfer />
        </div>
        <div className="flex flex-col py-6 min-h-full mt-8 lg:mt-0">
          <div className="space-y-2">
            <p className="font-bold text-xl text-[#343C6A] ml-4">
              Balance History
            </p>
            <LineGraph
              data={{
                labels: [
                  "Jul",
                  "Aug",
                  "Sep",
                  "Oct",
                  "Nov",
                  "Dec",
                  "Jan",
                  "Feb",
                  "Mar",
                  "Apr",
                  "May",
                  "Jun",
                ],
                values: [65, 59, 80, 81, 56, 55, 40, 81, 56, 55, 40, 60],
              }}
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
