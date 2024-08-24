import QuickTransferImage from "./QuickTransferImage";

const QuickTransfer = () => {
  return (
    <div className="bg-white w-[445px] h-[276px] rounded-[25px] pt-[5px] pl-[25px]">
      <div className="flex items-center w-[394px] h-[127px] mt-[35px] gap-[28px] mb-[29px]">
        <QuickTransferImage
          name="Fuad Mo"
          title="Student"
          imgUrl="/images/profile_1.png"
          isActive={false}
        />
        <QuickTransferImage
          name="Etlak So"
          title="Broke"
          imgUrl="/images/profile_2.png"
          isActive={true}
        />
        <QuickTransferImage
          name="Redit Mu"
          title="Ind"
          imgUrl="/images/profile_3.png"
          isActive={false}
        />
        <div className="w-[50px] h-[50px] flex justify-center items-center rounded-full shadow-lg">
          <span className="text-slate-400">{`>`}</span>
        </div>
      </div>
      <div className="flex justify-between items-center w-[395px] h-[50px] mb-[35px]">
        <span className="text-[16px] text-slate-400">Write Amount</span>{" "}
        <div className="relative">
          <input
            type="text"
            placeholder="525.50"
            className="bg-[#EDF1F7]  w-[265px] h-[50px] rounded-[50px] focus:outline-none pl-5"
          />
          <div className="flex justify-center items-center rounded-[50px] bg-[#1814F3] text-white absolute bottom-[0] left-[140px] w-[125px] h-[50px] border">
            <span>Send</span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default QuickTransfer;
