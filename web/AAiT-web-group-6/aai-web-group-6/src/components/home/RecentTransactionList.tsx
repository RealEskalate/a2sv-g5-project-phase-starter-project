import RecentTransaction from "./RecentTransaction";

const RecentTransactionList: React.FC = () => {
  return (
    <div className="p-[25px] flex flex-col gap-1 bg-white rounded-[25px]">
      <RecentTransaction
        title="Deposit from my card"
        date="28 January 2021"
        transcAmount="850"
        transcType="1"
        icon="/images/card.png"
      />

      <RecentTransaction
        title="Deposit Paypal"
        date="28 January 2021"
        transcAmount="2500"
        transcType="2"
        icon="/images/paypal.png"
      />

      <RecentTransaction
        title="Jemil Willson"
        date="28 January 2021"
        transcAmount="5400"
        transcType="3"
        icon="/images/will.png"
      />
    </div>
  );
};

export default RecentTransactionList;
