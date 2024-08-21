import { useUser } from "@/contexts/UserContext";
import Image from "next/image";

interface itemProp {
  icon: string;
  name: string;
}

const ServiceList = ({ icon, name }: itemProp) => {
  const { isDarkMode } = useUser();
  return (
    <div
      className={`flex justify-between items-center p-3 md:p5 rounded-xl ${
        isDarkMode
          ? "bg-gray-800 text-gray-300"
          : "bg-white text-gray-900 border-2 border-gray-200"
      }`}
    >
      <div className="flex ml-1 gap-1">
        <div className="w-10 h-10">
          <Image
            src="/servicesIcons/saftey.svg"
            alt="Service Icon"
            width={30}
            height={30}
          />
        </div>
        <div className="px-2">
          <div className="mt-1 font-medium">{name}</div>
          <div
            className={`text-xs ${
              isDarkMode ? "text-gray-500" : "text-gray-600"
            }`}
          >
            It&apos;s a long established fact.
          </div>
        </div>
      </div>
      <div className="hidden md:flex justify-between gap-4 w-1/2">
        <div className="px-2">
          <div className="mt-1 font-medium">Business loans</div>
          <div
            className={`text-xs ${
              isDarkMode ? "text-gray-500" : "text-gray-600"
            }`}
          >
            It&apos;s a long established fact.
          </div>
        </div>
        <div className="px-2">
          <div className="mt-1 font-medium">Business loans</div>
          <div
            className={`text-xs ${
              isDarkMode ? "text-gray-500" : "text-gray-600"
            }`}
          >
            It&apos;s a long established fact.
          </div>
        </div>
        <div className="px-2">
          <div className="mt-1 font-medium">Business loans</div>
          <div
            className={`text-xs ${
              isDarkMode ? "text-gray-500" : "text-gray-600"
            }`}
          >
            It&apos;s a long established fact.
          </div>
        </div>
      </div>
      <div
        className={`text-sm font-bold lg:border-2 lg:px-6 lg:rounded-full lg:my-auto lg:py-1 ${
          isDarkMode
            ? "text-gray-300 border-gray-600 bg-gray-700"
            : "text-gray-900 border-gray-300 bg-white"
        }`}
      >
        View Detail
      </div>
    </div>
  );
};

export default ServiceList;
