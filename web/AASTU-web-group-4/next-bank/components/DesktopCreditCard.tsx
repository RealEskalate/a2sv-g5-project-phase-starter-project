import Image from 'next/image';

interface CreditCardProps {
  bgColor: string;
  textColor: string;
}

const DesktopCreditCard: React.FC<CreditCardProps> = ({ bgColor, textColor }) => {
  return (
    <div className={`${bgColor} w-[300px] h-[200px] md:w-[350px] md:h-[235px] rounded-xl relative`}>
      {/* Top Section */}
      <div className="flex justify-between w-[95%]">
        <div className="mt-1 ml-3 p-2">
          <span className={`text-[10px] md:text-[12px] ${textColor}`}>Balance</span>
          <span className={`block text-[18px] md:text-[20px] font-bold ${textColor}`}>$5,756</span>
        </div>
        <div className="mt-4 mr-2">
          <Image src="/icons/chip.png" width={30} height={30} alt="chip card" className="h-[25px] md:h-[29px]" />
        </div>
      </div>

      {/* Middle Section */}
      <div className="flex justify-between w-auto mt-2">
        <div className="ml-3 p-2">
          <span className={`text-[10px] md:text-[12px] ${textColor}`}>CARD HOLDER</span>
          <span className={`block text-[13px] md:text-[15px] font-bold ${textColor}`}>Tekola Chane</span>
        </div>

        <div className="mr-9 p-2">
          <span className={`text-[10px] md:text-[12px] ${textColor}`}>VALID THRU</span>
          <span className={`block text-[13px] md:text-[15px] font-bold ${textColor}`}>12/22</span>
        </div>
      </div>

      {/* Bottom Section */}
      <div className={`flex justify-between items-center absolute bottom-0 left-0 right-0 bg-gradient-to-b ${bgColor} rounded-b-xl p-3`}>
        <span className={`ml-2 text-[18px] md:text-[22px] ${textColor}`}>3778 **** **** 1234</span>
        <div className="mr-3">
          <Image src="/icons/masterCard.png" width={44} height={42} alt="master card icon" className="h-[30px] md:h-[42px]" />
        </div>
      </div>
    </div>
  );
};

export default DesktopCreditCard;

