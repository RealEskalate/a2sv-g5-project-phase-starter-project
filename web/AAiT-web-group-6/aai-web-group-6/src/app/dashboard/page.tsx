import PaymentCard from "@/components/home/PaymentCard";
import RecentTransactionList from "@/components/home/RecentTransactionList";
import QuickTransfer from "@/components/home/QuickTransfer";

function page() {
  return (
    <div className=" flex flex-col items-center">
      <div className="flex gap-[30px] mt-[24px]">
        <div className="w-[730px] h-[282px] flex flex-col justify-between">
          <div className="flex justify-between">
            <span className="leading-[26.63px] text-[22px] font-[600] text-[#343C6A]">
              My Cards
            </span>
            <span className="leading-[20.57px] text-[17px] font-[600] text-[#343C6A]">
              See All
            </span>
          </div>
          <div className="flex gap-[30px]">
            <PaymentCard isWhite={false} />
            <PaymentCard isWhite={true} />
          </div>
        </div>
        <div className="flex flex-col justify-between w-[350px] h-[282px]">
          <div className="flex justify-start">
            <span className="leading-[26.63px] text-[22px] font-[600] text-[#343C6A]">
              Recent Transaction
            </span>
          </div>

          <div>
            <RecentTransactionList />
          </div>
        </div>
      </div>
      <div className="flex gap-[30px] border">
        <div className="w-[730px] h-[367px] border"></div>
        <div className="w-[350px] h-[367px] border"></div>
      </div>
      <div className="flex gap-[30px] mb-[39px]">
        <div className="flex flex-col justify-between w-[445px]  h-[323px]">
          <span className="leading-[26.63px] text-[22px] font-[600] text-[#343C6A]">
            Quick Transfer
          </span>
          <QuickTransfer />
        </div>
        <div className="w-[635px] h-[323px] border"></div>
      </div>
    </div>
  );
}

export default page;
