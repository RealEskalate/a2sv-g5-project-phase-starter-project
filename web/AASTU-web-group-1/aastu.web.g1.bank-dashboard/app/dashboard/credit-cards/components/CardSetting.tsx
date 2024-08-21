import { useUser } from "@/contexts/UserContext";
import Image from "next/image";

interface Props {
  image: string;
  color: string;
  title: string;
  description: string;
}

const CardSetting = ({ image, color, title, description }: Props) => {
  const { isDarkMode } = useUser();

  return (
    <div
      className={`flex mb-3 rounded-xl ${
        isDarkMode ? "bg-gray-800" : "bg-white"
      }`}
    >
      <div
        className={`flex-initial w-[5/12] m-3 text-[16px] ${
          isDarkMode ? "text-white" : "text-black"
        }`}
      >
        <div
          className={`${color} bg-opacity-25 font-semibold py-1 px-2 rounded-lg text-sm w-[45px]`}
        >
          <Image
            src={image}
            alt={title}
            width={20}
            height={20}
            className="mx-auto"
          />
        </div>
      </div>
      <div
        className={`flex-initial w-[7/12] m-3 ${
          isDarkMode ? "text-gray-300" : "text-gray-700"
        }`}
      >
        <div>
          <h1>{title}</h1>
          <p
            className={`text-[#718EBF] ${
              isDarkMode ? "text-blue-400" : "text-[#718EBF]"
            }`}
          >
            {description}
          </p>
        </div>
      </div>
    </div>
  );
};

export default CardSetting;
