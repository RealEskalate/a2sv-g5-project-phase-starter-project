import { SafteyIcon } from "../serviceIcons/icons";
const BenefitComp = () => {
  return (
    <div className="flex p-5 border-2 rounded-2xl bg-white ">
      <div className="ml-3 mr-1">
        <SafteyIcon />
      </div>
      <div className="flex flex-col px-2 flex-grow justify-center ">
        <div className="text-lg">Life insurance</div>
        <div className="text-xs text-[#718EBF] w-[130px]">
          Unlimited protection
        </div>{" "}
      </div>
    </div>
  );
};

export default BenefitComp;
