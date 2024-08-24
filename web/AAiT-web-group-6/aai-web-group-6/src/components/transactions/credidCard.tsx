import React from 'react'
import { CreditCardicon } from '../dashboard-layout/icons';
import { colors } from '@mui/material';

type Props = {
    balance: Number;
    cardHolder: string;
    validTHRU: string;
    Code: string;
    isWhite: boolean;
}

import {
    cardBackground,
    cardTextColor,
    cardLightTextColor,
    cardBottomBackground,
    chipImage,
    logoImage,
} from "@/components/transactions/paymentCard";
import Image from 'next/image';



const PaymentCard: React.FC<Props> = ({ isWhite, Code, balance, cardHolder, validTHRU }) => {
    return (
        <div
            className={`flex flex-col gap-[33px] md:w-[231px]  w-[265px] xl:w-[350px] flex-shrink-0 h-[235px] ${cardBackground(
                isWhite
            )} ${cardTextColor(isWhite)} rounded-3xl`}
        >
            <div className="flex justify-between mt-[24px]">
                <div className="flex flex-col ml-6">
                    <span className={`text-xs ${cardLightTextColor(isWhite)}`}>
                        Balance
                    </span>
                    <span className="text-lg">{balance.toString()}</span>
                </div>
                <Image
                    src={chipImage(isWhite)}
                    className="w-[35px] h-[35px] mr-6"
                    alt="Card Chip"
                    width={35}
                    height={35}
                />
            </div>
            <div className="flex gap-16 mt-0">
                <div className="flex flex-col ml-6">
                    <span className={`text-sm ${cardLightTextColor(isWhite)}`}>
                        CARD HOLDER
                    </span>
                    <span>{cardHolder}</span>
                </div>
                <div className="flex flex-col">
                    <span className={`text-sm ${cardLightTextColor(isWhite)}`}>
                        VALID THRU
                    </span>
                    <span>{validTHRU}</span>
                </div>
            </div>
            <div
                className={`flex pt-[15px] justify-between ${cardBottomBackground(
                    isWhite
                )} h-[70px] rounded-b-[25px] ${isWhite ? "border-t border-gray-200" : ""
                    }`}
            >
                <span className="ml-6">{Code}</span>
                <Image
                    className="mr-6 w-[44px] h-[30px]"
                    src={logoImage(isWhite)}
                    alt="Logo"
                    width={44}
                    height={30}
                />
            </div>
        </div>
    );
};

export default PaymentCard;