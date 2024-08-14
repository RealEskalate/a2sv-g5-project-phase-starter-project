import AccountInformation from "@/components/AccountInformation/AccountInformation";
import Image from "next/image";

export default function Home() {
  return (
    <div className="w-full grid gap-4 grid-cols-2 md:grid-cols-4 ">
      <AccountInformation
        image="/assets/balance.png"
        name="My Balance"
        balance="12700"
        color="bg-[#FFF5D9]"
      />

      <AccountInformation
        image="/assets/income.png"
        name="Income"
        balance="5,600"
        color="bg-[#E7EDFF]"
      />
      <AccountInformation
        image="/assets/expense.png"
        name="Expense"
        balance="3,460"
        color="bg-[#FFE0EB]"
      />
      <AccountInformation
        image="/assets/saving.png"
        name="Total Saving"
        balance="7,920"
        color="bg-[#DCFAF8]"
      />
    </div>
  );
}
