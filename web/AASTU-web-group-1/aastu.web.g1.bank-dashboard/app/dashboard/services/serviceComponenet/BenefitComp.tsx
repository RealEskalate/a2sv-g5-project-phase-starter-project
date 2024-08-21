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
      className={`flex p-5  rounded-2xl gap-2  ${
        isDarkMode
          ? "bg-gray-600 border-gray-600 text-gray-300"
          : "bg-white border-gray-200 text-gray-900"
      }`}
    >
      <div className="flex ml-3 mr-1 w-16 h-16 bg-blue-50 items-center justify-center rounded-full">
        <Image src={items.icon} alt="Benefit Icon" width={30} height={30} />
      </div>
      <div className="flex flex-col px-2 flex-grow justify-center">
        <div className="text-lg">{items.name}</div>
        <div
          className={`text-xs w-[130px] ${
            isDarkMode ? "text-gray-400" : "text-gray-600"
          }`}
        >
          Unlimited protection
        </div>
      </div>
    </div>
  );
};

export default BenefitComp;
