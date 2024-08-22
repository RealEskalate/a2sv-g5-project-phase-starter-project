import React from 'react';


const CreditCard: React.FC<CardProps> = ({ balance, cardHolder, validThru, cardNumber }) => {
return (
    <div className="flex flex-col w-[350px] h-[235px] border-2 rounded-3xl bg-white text-gray-600">
        <div className="flex flex-col justify-between gap-8 rounded-t-3xl bg-white">
            <div className='flex justify-between mt-6 mx-6'>
                <div>
                    <p className="text-sm">Balance</p>
                    <p className="text-2xl font-bold">{balance}</p>
                </div>
                <img src="/Chip_Card_dark.png" alt="Chip Card" style={{width:'35px', height:'35px'}}/>
            </div>
            <div className='flex items-start space-x-8 mx-6 mb-4'>
                <div>
                    <p className="uppercase tracking-wider text-sm">Card Holder</p>
                    <p className="uppercase tracking-wider font-semibold text-sm">{cardHolder}</p>
                </div>
                <div>
                    <p className="uppercase tracking-wider text-sm">Valid Thru</p>
                    <p className="uppercase tracking-wider font-semibold text-sm">{validThru}</p>
                </div>
            </div>
        </div>
        <div className="flex justify-between items-center align-bottom border-t-2 p-6 py-4 rounded-b-3xl">
            <div className="text-xl tracking-widest">{cardNumber}</div>
            <img src="/Group17.png" alt="MasterCard" style={{height:'30px'}}/>
        </div>
    </div>
);
};

export default CreditCard;
