import CreditCard from "../_components/Credit_Card";
import InfoCard from "./components/InfoCard";
import LastTransaction from "./components/LastTransaction";
import { ChartWeekly } from "@/components/ui/BarchartWeekly";
import Invoices from "./components/Invoices";

const Accounts = () => {
  return (
    <div className="p-5 md:pr-20">
      <div className="grid grid-cols-2 md:grid-cols-4 gap-4 md:flex-row p-3">
        <div className="bg-white p-4 rounded-3xl">
          <InfoCard
            title="My balance"
            amount={12750}
            image="/icons/money.svg"
            color="bg-yellow"
          />
        </div>
        <div className="bg-white p-4 rounded-3xl">
          <InfoCard
            title="Income"
            amount={5600}
            image="/icons/handmoney.svg"
            color="bg-blue"
          />
        </div>
        <div className="bg-white p-4 rounded-3xl">
          <InfoCard
            title="Expense"
            amount={3460}
            image="/icons/001-medical.svg"
            color="bg-pink"
          />
        </div>
        <div className="bg-white p-4 rounded-3xl items-start">
          <InfoCard
            title="Total Saving"
            amount={7920}
            image="/icons/003-saving.svg"
            color="bg-green"
          />
        </div>
      </div>
      <div className="md:flex md:gap-12">
        <div className="w-[70%]">
          <h1 className="text-xl">Last Transactions</h1>
          <div className=" ">
            <LastTransaction
              image="/icons/Bell.svg"
              alttext="bell"
              description="Spotify Subscription"
              transaction={-150}
              colorimg="bg-green"
              date="25 Jan 2021"
              type="Shopping"
              account="1234 ****"
              status="Pending"
            />
            <LastTransaction
              image="/icons/tools.svg"
              alttext="bell"
              description="Mobile Service"
              transaction={-340}
              colorimg="bg-blue"
              date="25 Jan 2021"
              type="Service"
              account="1234 ****"
              status="Completed"
            />
            <LastTransaction
              image="/icons/user.svg"
              alttext="settings"
              description="Emilly Wilson"
              transaction={780}
              colorimg="bg-pink"
              date="25 Jan 2021"
              type="Transfer"
              account="1234 ****"
              status="Completed"
            />
          </div>
        </div>
        <div className="md:w-[30%]">
        <div className="flex justify-between font-inter text-[16px] font-semibold mb-4">
          <h4>My Cards</h4>
          <h4>See All</h4>
        </div>
          <div className=" mb-4">
            <CreditCard
              isBlue={true}
              balance={5894}
              creditNumber="3778*** ****1234"
              name="Eddy Cusuma"
              textColor="text-white"
            />
          </div>
        </div>
      </div>
      <div className="md:flex gap gap-6 mb-5">
        <div>
          <h1 className="text-xl mb-4"> Debit & Credit overview </h1>
          <div className=" mb-4">
            <ChartWeekly />
          </div>
        </div>
        <div>
          <div>
            <h1 className="text-xl mb-4">Invoices Sent</h1>
            <div className="bg-white justify-between">
              <Invoices
                image="/icons/apple.svg"
                title="Apple Store"
                date="5h ago"
                expense={450}
                color="bg-green"
              />
              <Invoices
                image="/icons/useryello.svg"
                title="Michael"
                date="2 days ago"
                expense={450}
                color="bg-yellow"
              />
              <Invoices
                image="/icons/playstation.svg"
                title="Apple Store"
                date="2 days ago"
                expense={1085}
                color="bg-blue"
              />
              <Invoices
                image="/icons/user.svg"
                title="William"
                date="10 days ago"
                expense={90}
                color="bg-pink"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Accounts;
