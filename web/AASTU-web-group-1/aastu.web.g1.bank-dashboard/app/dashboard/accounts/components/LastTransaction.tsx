import { useUser } from "@/contexts/UserContext";
import Image from "next/image";

interface Prop {
  image: string;
  alttext: string;
  description: string;
  transaction: number;
  type: string;
  account: string;
  status: string;
  colorimg: string;
  date: string;
}

const LastTransaction = ({
  image,
  alttext,
  description,
  transaction,
  colorimg,
  date,
  type,
  account,
  status,
}: Prop) => {
  const { isDarkMode } = useUser();

  // Determine color and background based on dark mode
  const containerClass = isDarkMode
    ? "bg-gray-800 text-gray-100"
    : "bg-white text-gray-800";
  const colorClass = isDarkMode
    ? `${colorimg} bg-opacity-40`
    : `${colorimg} bg-opacity-25`;
  const textColor = isDarkMode ? "text-gray-400" : "text-[#718EBF]";

  // Format the transaction value with a "+" or "-" sign
  const formattedTransaction =
    transaction < 0 ? (
      <span className="text-red-500">-${Math.abs(transaction)}</span>
    ) : (
      <span className="text-green-500">+${transaction}</span>
    );

  return (
    <div className={`rounded-xl mb-2 mt-4 min-w-[325px] ${containerClass}`}>
      <div className="md:grid md:grid-cols-6 flex grid-cols-2 items-center justify-between">
        <div className="flex flex-initial col-span-2 m-3">
          <div
            className={`${colorClass} font-semibold py-2 px-2 rounded-lg text-sm w-[45px]`}
          >
            <Image src={image} alt={alttext} width={27} height={18} />
          </div>
          <div className="flex flex-col ml-3">
            <div>{description}</div>
            <div className={`text-sm ${textColor}`}>{date}</div>
          </div>
        </div>
        <div className="hidden md:block flex-initial w-[2/12]">{type}</div>
        <div className="hidden md:block flex-initial w-[2/12]">{account}</div>
        <div className="hidden md:block flex-initial w-[2/12]">{status}</div>
        <div className="flex-initial w-[4/12] m-3">{formattedTransaction}</div>
      </div>
    </div>
  );
};

export default LastTransaction;
