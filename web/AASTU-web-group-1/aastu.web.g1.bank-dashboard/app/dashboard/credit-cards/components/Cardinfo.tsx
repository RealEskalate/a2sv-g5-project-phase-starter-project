import { useUser } from "@/contexts/UserContext";
import Image from "next/image";

const Cardinfo = () => {
  const { isDarkMode } = useUser();

  return (
    <div>
      <div
        className={`flex p-5 rounded-xl mb-2 mt-4 max-w-screen-sm justify-between min-w-[325px] ${
          isDarkMode ? "bg-gray-800" : "bg-white"
        }`}
      >
        <div className="flex-initial w-[2/12] m-3">
          <div
            className={`font-semibold py-1 px-2 rounded-lg text-sm w-[45px] min-w-[20px] ${
              isDarkMode
                ? "text-blue-300 bg-blue-900"
                : "text-blue-500 bg-opacity-25"
            }`}
          >
            <Image
              src={`/icons/Cardbill.svg`}
              alt={"Cards"}
              width={27}
              height={18}
            />
          </div>
        </div>
        <div className="flex-initial w-[4/12] m-3">
          <div>
            <h2 className={`${isDarkMode ? "text-gray-200" : "text-black"}`}>
              Card Type
            </h2>
            <p className={`${isDarkMode ? "text-gray-400" : "text-gray-500"}`}>
              Secondary
            </p>
          </div>
        </div>
        <div className="flex-initial w-[3/12] m-3">
          <div>
            <h2 className={`${isDarkMode ? "text-gray-200" : "text-black"}`}>
              Bank
            </h2>
            <p className={`${isDarkMode ? "text-gray-400" : "text-gray-500"}`}>
              DBL Bank
            </p>
          </div>
        </div>
        <div className="flex-initial w-[3/12] m-3">
          <p className={`${isDarkMode ? "text-blue-400" : "text-[#1814F3]"}`}>
            View Details
          </p>
        </div>
      </div>
    </div>
  );
};

export default Cardinfo;
