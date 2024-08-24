interface RecentTransactionProps {
  title: string;
  icon: string;
  date: string;
  transcType: string;
  transcAmount: string;
}

const RecentTransaction: React.FC<RecentTransactionProps> = ({
  title,
  icon,
  date,
  transcType,
  transcAmount,
}) => {
  const styling = transcType == "1" ? `text-[#FF4B4A]` : `text-[#41D4A8]`;
  const val = transcType == "1" ? `-$${transcAmount}` : `+$${transcAmount}`;
  return (
    <div className="flex justify-between items-center w-[301px] h-[55px] bg-white">
      <img src={icon} className="w-[55px] h-[55px]" alt="" />
      <div className="flex flex-col items-start w-[167px] h-[44px]">
        <span>{title}</span>
        <span className="text-slate-400">{date}</span>
      </div>
      <span className={styling}>{val}</span>
    </div>
  );
};

export default RecentTransaction;
