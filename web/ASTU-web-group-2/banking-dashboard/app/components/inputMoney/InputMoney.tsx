import { Inter } from "next/font/google";
const inter = Inter({ subsets: ["latin"] });

const InputMoney = () => {
  return (
    <div className="flex items-center justify-between">
      <div className={`text-[#718EBF] ${inter.className}`}>Write Amount</div>
      <div className="flex bg-[#EDF1F7] rounded-3xl">
        <input
          type="number"
          placeholder="525.5"
          className="rounded-3xl p-3 w-24 bg-[#EDF1F7] border-none outline-none"
        />
        <button className="flex bg-[#1814F3] py-3 px-4 rounded-3xl gap-2">
          <div className="text-white">Send</div>
          <img src="/assets/inputMoney/telegram.svg" alt="telegram" />
        </button>
      </div>
    </div>
  );
};

export default InputMoney;
