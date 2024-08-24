import {
  cardBackground,
  cardTextColor,
  cardLightTextColor,
  cardBottomBackground,
  chipImage,
  logoImage,
} from "./paymentCardStyles";

interface PaymentCardProps {
  isWhite: boolean;
}

const PaymentCard: React.FC<PaymentCardProps> = ({ isWhite }) => {
  return (
    <div
      className={`flex flex-col gap-[33px] w-[350px] h-[235px] ${cardBackground(
        isWhite
      )} ${cardTextColor(isWhite)} rounded-3xl`}
    >
      <div className="flex justify-between mt-[24px]">
        <div className="flex flex-col ml-6">
          <span className={`text-xs ${cardLightTextColor(isWhite)}`}>
            Balance
          </span>
          <span className="text-lg">$5,756</span>
        </div>
        <img
          src={chipImage(isWhite)}
          className="w-[35px] h-[35px] mr-6"
          alt="Card Chip"
        />
      </div>
      <div className="flex gap-16 mt-0">
        <div className="flex flex-col ml-6">
          <span className={`text-sm ${cardLightTextColor(isWhite)}`}>
            CARD HOLDER
          </span>
          <span>Eddy Tusuma</span>
        </div>
        <div className="flex flex-col">
          <span className={`text-sm ${cardLightTextColor(isWhite)}`}>
            VALID THRU
          </span>
          <span>12/22</span>
        </div>
      </div>
      <div
        className={`flex pt-[15px] justify-between ${cardBottomBackground(
          isWhite
        )} h-[70px] rounded-b-[25px] ${
          isWhite ? "border-t border-gray-200" : ""
        }`}
      >
        <span className="ml-6">3778 **** *** 1234</span>
        <img
          className="mr-6 w-[44px] h-[30px]"
          src={logoImage(isWhite)}
          alt="Logo"
        />
      </div>
    </div>
  );
};

export default PaymentCard;
