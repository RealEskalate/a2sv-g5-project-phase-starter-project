import { useUser } from "@/contexts/UserContext";
import { SafteyIcon } from "../serviceIcons/icons";
import Image from "next/image";

interface itemsProp {
  icon: string;
  name: string;
}

interface Benefitprop {
  items: itemsProp;
}
const BenefitComp = ({ items }: Benefitprop) => {
  console.log(items.icon);

  return (
    <div className="flex p-5 border-2 rounded-2xl bg-white ">
      <div className="ml-3 mr-1 w-10 h-10">
      <Image src={items.icon} alt=""  width={30} height={30}/>
      </div>
      <div className="flex flex-col px-2 flex-grow justify-center ">
        <div className="text-lg">{items.name}</div>
        <div className="text-xs text-[#718EBF] w-[130px]">
          Unlimited protection
        </div>{" "}
      </div>
    </div>
  );
};

export default BenefitComp;
