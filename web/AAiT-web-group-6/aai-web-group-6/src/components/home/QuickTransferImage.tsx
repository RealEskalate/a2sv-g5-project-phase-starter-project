interface QuickTransferImageProps {
  name: string;
  imgUrl: string;
  title: string;
  isActive: boolean;
}

const QuickTransferImage: React.FC<QuickTransferImageProps> = ({
  name,
  imgUrl,
  title,
  isActive,
}) => {
  const activeUser = isActive
    ? `text-slate-950 font-[500] text-[16px] text-[#232323]`
    : `text-slate-800 font-[300] text-[16px] text-[#232323]`;
  return (
    <div className="flex flex-col items-center gap-[15px] w-[85px] h-[127px]">
      <img src={imgUrl} alt="" className="w-[70px] h-[70px] rounded-full" />
      <div className="flex flex-col  items-center w-[85px] h-[42px]">
        <span className={activeUser}>{name}</span>
        <span
          className={isActive ? `text-[#718EBF] font-[500]` : `text-slate-300`}
        >
          {title}
        </span>
      </div>
    </div>
  );
};

export default QuickTransferImage;
