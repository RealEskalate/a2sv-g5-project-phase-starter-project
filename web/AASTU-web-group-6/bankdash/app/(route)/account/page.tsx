import Card from '../../components/Accounts/account'
import LastTransList from '@/app/components/Accounts/LastTransPa';
import { DebitCreditOver } from "../../components/Accounts/DebitCreditOver";
import InvoiceCard from "../../components/Accounts/InvoiceCard";
import VisaCard from '@/app/components/Card/VisaCard';
export default function Home() {
  return (
    <>
      <div className="flex flex-col lg:flex-row gap-7">
        <div className="flex lg:w-[45%] gap-7">
          <Card
            title="My Balance"
            amount="$12,750"
            color="#FFF5D9"
            icon="/assets/money-tag 1.svg"
            width="w-[45%]"
          />
          <Card
            title="Income"
            amount="$5,600"
            color="#E7EDFF"
            icon="/assets/expense.svg"
            width="w-[45%]"
          />
        </div>
        <div className="flex  lg:w-[45%] gap-7">
          <Card
            title="Expense"
            amount="$3,460"
            color="#FFE0EB"
            icon="/assets/income.svg"
            width="w-[45%]"
          />
          <Card
            title="Total Saving"
            amount="$7,920"
            color="#DCFAF8"
            icon="/assets/saving.svg"
            width="w-[45%]"
          />
        </div>
      </div>

      <div className="flex  flex-col lg:flex-row my-5 justify-between">
        <div className="lg:w-[65%]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
            Last Transaction
          </p>
          <div className=" bg-white border rounded-3xl p-3 shadow-lg border-gray-300">
            <LastTransList/>
          </div>
        </div>
        <div className="lg:w-[30%] lg:h-[250px]">
          <div className="flex justify-between">
            <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
              My Card
            </p>
            <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
              See All
            </p>
          </div>
          <VisaCard isBlack={false} isFade={true} isSimGray={false}/>
        </div>
      </div>
      <div className="flex flex-col lg:flex-row justify-between my-5">
        <div className="lg:w-[65%] lb:h-[364px]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
            Debit & Credit Overview
          </p>
          <DebitCreditOver />
        </div>
        <div className="lg:w-[30%]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5 ">
            Invoices Sent
          </p>
          <div className="border border-solid rounded-3xl p-9  bg-white shadow-lg border-gray-300">
            <InvoiceCard
              title="Apple Store"
              date="5h ago"
              amount="$450"
              icon="/assets/apple.svg"
              color="#DCFAF8"
            />
            <InvoiceCard
              title="Michael"
              date="2 days ago"
              amount="$160"
              icon="/assets/userr.svg"
              color="#FFF5D9"
            />
            <InvoiceCard
              title="Playstation"
              date="5 days ago"
              amount="$1085"
              icon="/assets/Group.svg"
              color="#E7EDFF"
            />
            <InvoiceCard
              title="William"
              date="10 days ago"
              amount="$90"
              icon="/assets/userr.svg"
              color="#FFE0EB"
            />
          </div>
        </div>
      </div>
    </>
  );
}
