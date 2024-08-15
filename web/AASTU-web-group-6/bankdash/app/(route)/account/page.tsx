import Card from '../../components/Accounts/account'
import Last_trans from "../../components/Accounts/Last_trans";
import { DebitCreditOver } from "../../components/Accounts/DebitCreditOver";
import InvoiceCard from "../../components/Accounts/InvoiceCard";
import VisaCard from '@/app/components/Card/VisaCard';
export default function Home() {
  return (
    <>
      <div className="flex gap-7">
        <Card
          title="My Balance"
          amount="$12,750"
          color="#FFF5D9"
          icon="/assets/money-tag 1.svg"
          width=' w-[22%]'
        />
        <Card
          title="Income"
          amount="$5,600"
          color="#E7EDFF"
          icon="/assets/expense.svg"
          width='w-[22%]'
        />
        <Card
          title="Expense"
          amount="$3,460"
          color="#FFE0EB"
          icon="/assets/income.svg"
          width='w-[22%]'
        />
        <Card
          title="Total Saving"
          amount="$7,920"
          color="#DCFAF8"
          icon="/assets/saving.svg"
          width='w-[22%]'
        />
      </div>

      <div className="flex my-5 justify-between">
        <div className="w-[65%]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
            Last Transaction
          </p>
          <div className=" bg-white border rounded-3xl p-3">
            <Last_trans
              title="Spotify Subscription"
              date="25 Jan 2021"
              type="Shopping"
              account_no="1234 ****"
              status="Pending"
              amount="-$150"
              color="#DCFAF8"
              icon="/assets/renew.svg"
            />
            <Last_trans
              title="Mobile Service"
              date="25 Jan 2021"
              type="Service"
              account_no="1234 ****"
              status="Completed"
              amount="-$340"
              color="#E7EDFF"
              icon="/assets/settings.svg"
            />
            <Last_trans
              title="Emilly Wilson"
              date="25 Jan 2021"
              type="Transfer"
              account_no="1234 ****"
              status="Completed"
              amount="+$780"
              color="#FFE0EB"
              icon="/assets/user.svg"
            />
          </div>
        </div>
        <div className='w-[30%]'>
          <div className='flex justify-between'>
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
           My Card
          </p>
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">See All</p>
          </div>
          <VisaCard isBlack={false}/>
       
        </div>
        
      </div>
      <div className="flex justify-between my-5">
        <div className="w-[65%] h-[364px]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
            Debit & Credit Overview
          </p>
          <DebitCreditOver />
        </div>
        <div className="w-[30%]">
          <p className="font-inter font-semibold text-[22px] text-[#333B69] mb-5">
            Invoices Sent
          </p>
          <div className="border border-solid rounded-3xl p-9">
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
              icon="/assets/user.svg"
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
              icon="/assets/user.svg"
              color="#FFE0EB"
            />
          </div>
        </div>
      </div>
    </>
  );
}
