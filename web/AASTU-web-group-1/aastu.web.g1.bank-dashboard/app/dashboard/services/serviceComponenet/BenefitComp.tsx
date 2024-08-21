import { useUser } from "@/contexts/UserContext";
import Image from "next/image";

interface itemsProp {
  icon: string;
  name: string;
}

interface Benefitprop {
  items: itemsProp;
}

const BenefitComp = ({ items }: Benefitprop) => {
  const { isDarkMode } = useUser();
  return (
    <div
      className={`flex p-5 border-2 rounded-2xl ${
        isDarkMode
          ? "bg-gray-800 border-gray-600 text-gray-300"
          : "bg-white border-gray-200 text-gray-900"
      }`}
    >
      <div className="ml-3 mr-1 w-10 h-10">
        <Image src={items.icon} alt="Benefit Icon" width={30} height={30} />
      </div>
      <div className="flex flex-col px-2 flex-grow justify-center">
        <div className="text-lg">{items.name}</div>
        <div
          className={`text-xs w-[130px] ${
            isDarkMode ? "text-gray-500" : "text-gray-600"
          }`}
        >
          Unlimited protection
        </div>
      </div>
    </div>
  );
};

export default BenefitComp;
